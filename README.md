# verbose

Make your input verbose 🙈  
The main purpose of `verbose` is flooding buffer!

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
