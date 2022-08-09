# A CLI

A starter for self installing, self completing CLI using the amazing `opts` library.

You ask why's it amazing;
- **No Flag Parsing**, reflection based `struct` stuffing (field attributes turn flags into `args`). 
- Did I mentioned the binaries are **self completers**, add a sub-command or field, recompile and walla it's in the tab completion.

## Badges

[![Release](https://img.shields.io/github/release/wxio/acli.svg?style=for-the-badge)](https://github.com/wxio/acli/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=for-the-badge)](/LICENSE.md)
[![Powered By: Opts CLI Library](https://img.shields.io/badge/powered%20by-opts_cli-green.svg?style=for-the-badge)](https://github.com/jpillora/opts)
[![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=for-the-badge)](https://github.com/goreleaser)

## Quick Start

**Add a new subcommand**
``` bash
git clone https://github.com/wxio/acli.git mycli
cd mycli
go build
./acli cli new_sub_command --entire-reg --parent testing sample >> register_sample.go
go build
./acli testing sample
# edit internal/testing/sample/sample.go
go build
./acli testing sample
# repeat
```

**Rename CLI**

``` bash
./acli cli rename freddo mycli
rm acli
go build
./mycli
```

**Install it as a self-completer**
``` bash
# this assumes $HOME/go/bin is in you path
go install
mycli --install
# from zsh, differs for bash and fish
source ~/.zshrc
# tab completion now active
./mycli cl<tab> new<tab>  --entire-reg --par<tab> testing sample2 >> register_sample2.go
go install
mycli te<tab> sa<tab>
```

## Detailed Usage

## `acli cli md_tmpl`
<!--tmpl,code=plain:acli cli md_tmpl --help -->
``` plain 

  Usage: acli cli md_tmpl [options] <filename>

  Simple markdown templating using shell commands. see https://github.com/jpillora/md-tmpl
  General usage is `acli cli md_tmpl README.md`.
  Works nicely with `acli --gen-docs >> README.md`

  Options:
  --working-dir, -w  default .
  --preview, -p
  --help, -h         display help

```
<!--/tmpl-->

## `acli cli new_sub_command`
<!--tmpl,code=plain:acli cli new_sub_command --help -->
``` plain 

  Usage: acli cli new_sub_command [options] [name] [name] ...

  Creates starter code for new subcommands & print sample registration code for the main package.

  Starter code is written to  in `internal/parent/subcommand/subcommand.go`

  Use --parent to nest the subcommands.

  Sample code is written to stdOut, all other output is on stdErr, this allows for the following
  `acli cli new_sub_command subcmd2 >> main.go`
  or
  `acli cli new_sub_command --entire-reg --parent topsy/turvy top1 top2 > reg02.go`

  Options:
  --debug, -d
  --parent, -p       path of parent commands eg foo/bar
  --org, -o          default wxio
  --project          default acli
  --module-path, -m  the parent path of the internal src directory (default `pwd`)
  --overwrite
  --entire-reg, -e   (only valid with --parent) print to standard output an entire go file to
                     register the subcommand. If false only the func init is printed.
  --help, -h         display help

```
<!--/tmpl-->

## `acli cli rename`
<!--tmpl,code=plain:acli cli rename --help -->
``` plain 

  Usage: acli cli rename [options] [to] [to] ...

  Does a search and replace on the name 'acli'* in all the source file.
  Respects .gitignore.

  * to is either 'org name' or just 'name' eg 'freddo frog' or 'frog'

  Options:
  --debug, -d
  --from-org, -o     default wxio
  --from-name, -n    default acli
  --module-path, -m  the parent path of the internal src directory (default `pwd`)
  --help, -h         display help

```
<!--/tmpl-->

