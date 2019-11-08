# bcrypt-tool

Command line tool for generating bcrypt passwords.

## Build
```sh
$ make
```

## Use
check version
```sh
$ ./bcrypt-tool -version
```
generate bcrypt password
```sh
$ ./bcrypt-tool mypassword
```
generate bcrypt password from stdin
```sh
$ echo "mypassword" | ./bcrypt-tool
```
