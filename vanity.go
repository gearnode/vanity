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

package vanity

import (
	"fmt"
	"path"
)

type (
	Import struct {
		DomainName   string
		VCS          string
		RepoRoot     string
		ImportPrefix string
	}
)

func NewImport(domain, vcs, root, prefix string) Import {
	return Import{domain, vcs, root, prefix}
}

func (i Import) ImportRoot() string {
	return path.Join(i.DomainName, i.ImportPrefix)
}

func (i Import) GoImportMeta() string {
	return fmt.Sprintf(`<meta name="go-import" content="%s %s %s"/>`, i.VCS, i.RepoRoot, i.ImportRoot())
}

func (i Import) GoSourceMeta() string {
	return fmt.Sprintf(`<meta name="go-source" content="%s %s %s %s"/>`, i.ImportRoot(), "_", "_", "_")
}

func (i Import) HTMLPage() string {
	return fmt.Sprintf(
		`<!DOCTYPE><html lang="en"><head><title>%s</title><meta charset="utf-8"/><meta http-equiv="refresh" content="10; url=https://pkg.go.dev/%s">%s%s</head><body><p>Redirecting to docs at <a href="https://pkg.go.dev/%s">%s</a>...</p></body></html>`,
		i.ImportRoot(),
		i.ImportRoot(),
		i.GoImportMeta(),
		i.GoSourceMeta(),
		i.ImportRoot(),
		i.ImportRoot(),
	)
}
