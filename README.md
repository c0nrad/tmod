# tmod

Modify a file's timestamp. Make it look like you modified a file in the past or the future.

## Usage
```bash
$ ./tmod -h
Usage of ./tmod:
  -file="": file to modify date
  -offset=-300: number of seconds to offset
```

## Example

```
$ touch demo

$ ls -la
total 3064
drwxr-xr-x  5 c0nrad  staff      170 Jul 17 22:15 .
drwxr-xr-x  8 c0nrad  staff      272 Jul 17 22:15 ..
-rw-r--r--  1 c0nrad  staff        0 Jul 17 22:15 demo
-rwxr-xr-x  1 c0nrad  staff  1562860 Jul 17 22:15 tmod

$ ./tmod -h
Usage of ./tmod:
  -file="": file to modify date
  -offset=-300: number of seconds to offset

$ ./tmod --file "demo" --offset "-300"

$ ls -la
drwxr-xr-x  5 c0nrad  staff      170 Jul 17 22:15 .
drwxr-xr-x  8 c0nrad  staff      272 Jul 17 22:15 ..
-rw-r--r--  1 c0nrad  staff        0 Jul 17 22:10 demo
-rwxr-xr-x  1 c0nrad  staff  1562860 Jul 17 22:15 tmod
```

## Contact

c0nrad@c0nrad.io