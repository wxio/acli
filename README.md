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

``` bash
clone
build
./acli --install
# new shell
./acli <tab><tab>
```

