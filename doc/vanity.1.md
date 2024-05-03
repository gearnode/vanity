---
title: VANITY
header: Vanity Manual
footer: 1.0.0
date: Jan 20, 2022
section: 1
---
# NAME
**vanity** - generate static websites for Golang vanity import paths.

# SYNOPSIS
**vanity** [-help] [-version] [-output=\<dirname\>] [-cfg=\<filename\>]

# DESCRIPTION

A minimalist, open source command line interface to manage Golang
vanity import as a static website.

# OPTIONS
**-help**
: Show help message.

**-output** \<dirname\>
: the path where generated file will be written (default "dist").

**-cfg** \<filename\>
: The path of the configuration file (default "vanity.yaml").

**-flat-file-urls**
: Generates .html files named after their respective paths, avoiding
  directory-style URLs to prevent unwanted URL redirection, such as
  the addition of a trailing slash that implies a directory structure.

**-version**
: Prints the vanity cli version.

# EXIT STATUS
The **vanity** utility exits 0 on success, and >0 if an error occurs.

# EXAMPLES
Generate static website:

    $ vanity -cfg /usr/local/etc/vanity.yaml -output /usr/local/www/example.org

# SEE ALSO
**vanity.conf**(5)

# AUTHORS
Bryan Frimin.
