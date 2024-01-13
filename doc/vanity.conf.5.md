---
title: VANITY.CONF
header: Vanity Manual
footer: 1.0.0
date: Jan 20, 2022
section: 5
---
# NAME
**vanity.conf** – vanity configuration file.

# DESCRIPTION
The vanity(1) command line interface generates static websites,
managing Go vanity imports as configured in the **vanity.yaml** file.

# FORMAT
## Top level object keys:
**domain-name** *string* (required)
: The domain where the static website will be hosted.

**imports** *array\<import\>*
: The list of vanity imports.

## The import object format:
**vcs** *string* (required)
: The vcs is one of "bzr", "fossil", "git", "hg", "svn".

**repo-root** *string* (required)
: The repo-root is the root of the version control system containing a
  scheme and not containing a .vcs qualifier.

**import-prefix** *string* (required)
: The vanity import name.

# EXAMPLES
Minimal vanity configuration file:

```yaml
---
domain-name: "go.gearno.de"
imports: []
```

A bit more complete configuration file:

```yaml
---
domain-name: "go.gearno.de"
imports:
  - vcs: "git"
    repo-root: "github.com/gearnode/privatebin"
    import-prefix: "privatebin"
  - vcs: "git"
    repo-root: "github.com/gearnode/vanity"
    import-prefix: "vanity"
```

# AUTHORS
Bryan Frimin.
