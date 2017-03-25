# brainfuck-go
Brainf*ck written in Golang

# Install
```bash
go get -u github.com/nownabe/brainfuck-go
```

# Usage
```bash
# Brainfuck program source code
$ cat helloworld.bf
+++++++++[>++++++++>+++++++++++>+++++<<<-]>.>++.+++++++..+++.>-.------------.<++++++++.--------.+++.------.--------.>+.

# Run brainfuck program
$ brainfuck-go helloworld.bf
Input: +++++++++[>++++++++>+++++++++++>+++++<<<-]>.>++.+++++++..+++.>-.------------.<++++++++.--------.+++.------.--------.>+.


Hello, world!

(END)

# Suppress output
$ brainfuck-go helloworld.bf 2> /dev/null
Hello, world!
```

# Your Custom Brainfuck
See [examples](https://github.com/nownabe/brainfuck-go/tree/master/examples).

You can execute original programming language like brainfuck with JSON config file.

This is a config JSON for a standard brainfuck.

```json
{
  "next": ">",
  "prev": "<",
  "inc": "+",
  "dec": "-",
  "read": ",",
  "write": ".",
  "open": "[",
  "close": "]",
  "whitespaces": " \t\r\n"
}
```

When you want PPAP programming language, you can get it with a following JSON:

```json
{
  "next": "apple",
  "prev": "pinapple",
  "inc": "I",
  "dec": "have",
  "read": "pen",
  "write": "Uh",
  "open": "[",
  "close": "]",
  "whitespaces": " \t\r\n"
}
```

You can execute your PPAP program:

```bash
# PPAP program source code
$ cat helloworld.ppap
IIIIIIIII[appleIIIIIIIIappleIIIIIIIIIIIappleIIIIIpinapplepinapplepinapplehave]appleUhappleIIUhIIIIIIIUhUhIIIUhapplehaveUhhavehavehavehavehavehavehavehavehavehavehavehaveUhpinappleIIIIIIIIUhhavehavehavehavehavehavehavehaveUhIIIUhhavehavehavehavehavehaveUhhavehavehavehavehavehavehavehaveUhappleIUh

# Execute your PPAP program
$ brainfuck-go -conf ppap.json helloworld.ppap
Input: IIIIIIIII[appleIIIIIIIIappleIIIIIIIIIIIappleIIIIIpinapplepinapplepinapplehave]appleUhappleIIUhIIIIIIIUhUhIIIUhapplehaveUhhavehavehavehavehavehavehavehavehavehavehavehaveUhpinappleIIIIIIIIUhhavehavehavehavehavehavehavehaveUhIIIUhhavehavehavehavehavehaveUhhavehavehavehavehavehavehavehaveUhappleIUh


Hello, world!

(END)
```

## Convert Source Code
You can also convert brainfuck source code to another source of your original programming language.

Install converter:

```bash
go get -u github.com/nownabe/brainfuck-go/converter
```

And then you can convert:

```bash
$ converter -conf ppap.json helloworld.bf
IIIIIIIII[appleIIIIIIIIappleIIIIIIIIIIIappleIIIIIpinapplepinapplepinapplehave]appleUhappleIIUhIIIIIIIUhUhIIIUhapplehaveUhhavehavehavehavehavehavehavehavehavehavehavehaveUhpinappleIIIIIIIIUhhavehavehavehavehavehavehavehaveUhIIIUhhavehavehavehavehavehaveUhhavehavehavehavehavehavehavehaveUhappleIUh
```
