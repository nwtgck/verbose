# verbose
[![CircleCI](https://circleci.com/gh/nwtgck/verbose.svg?style=shield)](https://circleci.com/gh/nwtgck/verbose)

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
