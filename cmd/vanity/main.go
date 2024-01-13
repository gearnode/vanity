// Copyright (c) 2023 Bryan Frimin <bryan@frimin.fr>.
//
// Permission to use, copy, modify, and/or distribute this software
// for any purpose with or without fee is hereby granted, provided
// that the above copyright notice and this permission notice appear
// in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL
// WARRANTIES WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE
// AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT, INDIRECT, OR
// CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM LOSS
// OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT,
// NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF OR IN
// CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

package main // import "go.gearno.de/vanity/cmd/vanity"

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path"

	"go.gearno.de/vanity"
	"sigs.k8s.io/yaml"
)

const (
	version = "0.0.1"
)

type (
	Cfg struct {
		DomainName string      `json:"domain-name"`
		Imports    []ImportCfg `json:"imports"`
	}

	ImportCfg struct {
		VCS          string `json:"vcs"`
		RepoRoot     string `json:"repo-root"`
		ImportPrefix string `json:"import-prefix"`
	}
)

func info(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "info: "+format+"\n", args...)
}

func fail(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "error: "+format+"\n", args...)
	os.Exit(1)
}

func loadCfgFile(filename string) (*Cfg, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %w", err)
	}

	y, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("cannot read file: %w", err)
	}

	j, err := yaml.YAMLToJSON(y)
	if err != nil {
		return nil, fmt.Errorf("cannot convert yaml to json: %w", err)
	}

	cfg := Cfg{}
	err = json.Unmarshal(j, &cfg)
	if err != nil {
		return nil, fmt.Errorf("cannot json unmarshal: %w", err)
	}

	return &cfg, nil
}

func main() {
	cfgPath := flag.String("cfg", "", `the path of the configuration file (default "./vanity.yaml")`)
	outputPath := flag.String("output", "", `the path where generated file will be written (default "./dist")`)
	showHelp := flag.Bool("help", false, "shows this help message")
	showVersion := flag.Bool("version", false, "prints the vanity cli version")

	flag.Parse()

	if *showHelp {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *showVersion {
		fmt.Printf("vanity cli version %s\n", version)
		os.Exit(0)
	}

	cwd, err := os.Getwd()
	if err != nil {
		fail("cannot get current working directory: %v", err)
	}

	if *cfgPath == "" {
		*cfgPath = path.Join(cwd, "vanity.yaml")
	}

	if *outputPath == "" {
		*outputPath = path.Join(cwd, "dist")
	}

	info("loading configuration")
	cfg, err := loadCfgFile(*cfgPath)
	if err != nil {
		fail("cannot load configuration: %v", err)
	}

	prefixes := map[string]struct{}{}
	for _, importCfg := range cfg.Imports {
		if importCfg.VCS == "" {
			fail("the 'vcs' configuration field cannot be left blank")
		}

		if importCfg.ImportPrefix == "" {
			fail("the 'import-prefix' configuration field cannot be left blank")
		}

		if importCfg.RepoRoot == "" {
			fail("the 'repo-root' configuration field cannot be left blank")
		}

		_, ok := prefixes[importCfg.ImportPrefix]
		if ok {
			fail("duplicated import prefix: %s", importCfg.ImportPrefix)
		}
		prefixes[importCfg.ImportPrefix] = struct{}{}
	}

	for _, importCfg := range cfg.Imports {
		vanityImport := vanity.NewImport(cfg.DomainName, importCfg.VCS, importCfg.RepoRoot, importCfg.ImportPrefix)
		info("generating %s", vanityImport.ImportRoot())

		dirname := path.Join(*outputPath, importCfg.ImportPrefix)
		err := os.MkdirAll(dirname, os.ModePerm)
		if err != nil {
			fail("cannot create %q directory: %v", dirname, err)
		}

		filename := path.Join(dirname, "index.html")
		os.WriteFile(filename, []byte(vanityImport.HTMLPage()), 0644)
	}
}
