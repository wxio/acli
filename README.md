# A CLI

A starter for self installing, self completing CLI using the amazing `opts` library.

You ask why's it amazing;
- **No Flag Parsing**, reflection based `struct` stuffing (field attributes turn flags into `args`). 
- Did I mentioned the binaries are **self completers**, add a sub-command or field, recompile and walla it's in the tab completion.

## Badges

[![Release](https://img.shields.io/github/release/wxio/acli.svg?style=for-the-badge)](https://github.com/gmwxio/acli/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=for-the-badge)](/LICENSE.md)
[![Powered By: Opts CLI Library](https://img.shields.io/badge/powered%20by-opts_cli-green.svg?style=for-the-badge)](https://github.com/jpillora/opts)
[![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=for-the-badge)](https://github.com/goreleaser)

## Quick Start

[Terminal Recording](./docs/acli_sample.svg) similar to below commands.

``` bash
git clone https://github.com/gmwxio/acli.git mycli
cd mycli
go build
./acli cli rename wxio/acli freddo/mycli
go build
./mycli --install
# from zsh, differs for bash and fish
source ~/.zshrc
# tab completion now active
./mycli cl<tab> new<tab> --par<tab> testing sample >> main.go
# edit main.go
go build
./mycli te<tab> sa<tab>
```

### Using `cli new_sub_command`
*without tab completion*

```
acli cli new_sub_command --parent go_tour/module1 step1
```

<!-- instructions for main -->
Modify `main.go`