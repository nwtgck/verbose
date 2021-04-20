# verbose
[![CI](https://github.com/nwtgck/verbose/actions/workflows/ci.yml/badge.svg)](https://github.com/nwtgck/verbose/actions/workflows/ci.yml)

Make your input verbose 🙈  
The main purpose of `verbose` is flooding buffer!

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

Decode verbose input

```bash
printf "hello" | verbose | verbose -d
# => hello
```
