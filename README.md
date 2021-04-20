# verbose
[![CI](https://github.com/nwtgck/verbose/actions/workflows/ci.yml/badge.svg)](https://github.com/nwtgck/verbose/actions/workflows/ci.yml)

CLI which makes your input verbose, flooding buffer

## Installation

Get executable binaries from [GitHub Releases](https://github.com/nwtgck/verbose/releases)

## Usage

Here is examples to use `verbose`.

```bash
printf "hello" | verbose
# => hhhhhhhhhheeeeeeeeeelllllllllllllllllllloooooooooo
```

Specify the number how verbose.

```bash
printf "hello" | verbose -n 50
# => hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhheeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeelllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllloooooooooooooooooooooooooooooooooooooooooooooooooo
```

Encode & decode.

```bash
printf "hello" | verbose | verbose -d
# => hello
```

Encode & decode specifying how much amplified.

```bash
printf "hello" | verbose -n 100 | verbose -n 100 -d
# => hello
```
