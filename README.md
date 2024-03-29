# Vanity

This repository offers a CLI tool to easily create static websites for
Golang vanity import paths.

## Abstract

For the past decade, I've used NGINX for vanity Go imports in personal
projects. Recently, setting up Go vanity imports at work made me
realize that using NGINX for this is too complex and adds unnecessary
maintenance. So, I created a tool that simplifies this process. Now,
anyone can easily set up vanity imports on static websites like GitHub
Pages or S3 with Cloudfront

## Install

### From source

```
git clone https://github.com/gearnode/vanity.git
cd vanity
make
make install
```

## Usage

You can generate the static website with:

	vanity -cfg example.org.yaml -output /usr/local/www/example.org

## Build

You can build the command line interface with:

	make build

## Documentation

The [handbook](doc/handbook.md) contains informations about various
aspects of the command line interface.

You can also use the standard Go documentation tool to read code
documentation, for example:

	go doc -all go.gearno.de/vanity
	
## Contact

If you find a bug or have any question, feel free to open a Github
issue or to contact me [by email](mailto:bryan@frimin.fr).

Please note that I do not currently review or accept any contribution.

## Licence

Released under the ISC license.

Copyright (c) 2023 Bryan Frimin <bryan@frimin.fr>.

Permission to use, copy, modify, and/or distribute this software for
any purpose with or without fee is hereby granted, provided that the
above copyright notice and this permission notice appear in all
copies.

THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL
WARRANTIES WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE
AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL
DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR
PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER
TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR
PERFORMANCE OF THIS SOFTWARE.
