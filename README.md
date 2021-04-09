# Erlang String Parser

For those occasions where Erlang is just printing a list of numbers
when you really want to see the ASCII equivalent, use **erlnum**.
Erlnum prints a hexdump of the Erlang list.

## Usage

    erlnum file.txt
    cat file.txt | erlnum

## Example

```
$ cat input.txt
72,101,108,108,111,10
87,111,114,108,100,10

$ cat input.txt | erlnum
00000000  48 65 6c 6c 6f 0a 57 6f 72 6c 64 0a    |Hello.World.|
```

## Install

Checkout the code and simply:

    go install
    