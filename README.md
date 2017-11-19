vlookup
====

`vlookup` is command which behaves like vlookup function of Excel.

# Usage
```bash
Usage: vlookup [value](list file) [table](csv file) index_number
  -v	print version
```

# Example

```bash
$ cat example/value.csv
111
222
333

$ cat example/table.csv
111,hoge
333,hello

$ vlookup example/value.csv example/table.csv 2
111,hoge
222,nil
333,hello
```

# Install
```bash
$ go get -u github.com/himetani/vlookup
```
