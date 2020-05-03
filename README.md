# Mydb kv & sql

Supports:

- status    () Display server status.
- keys      () List 10 keys.
- set       () set key value.
- get       () get key.

## Installation


``` shell
git clone https://github.com/yeziyitao/mydb.git
make
```

## Quickstart

``` shell
yezi$ bin/Server 
start server...listen: 12358


Import keys:

``` shell
make bench
```

yezi$ bin/Client
Welcome to the mydb v0.0.2
Type 'help;' or '\h' for help.
connected
mydb v0.0.2> help
List of all mydb v0.0.2 commands:

exit      () Exit mydb.
help      (\h) Display this help.
status    () Display server status.
keys      () List 10 keys.
set       () set key value.
get       () get key.

not implement
mydb v0.0.2> \h
List of all mydb v0.0.2 commands:

exit      () Exit mydb.
help      (\h) Display this help.
status    () Display server status.
keys      () List 10 keys.
set       () set key value.
get       () get key.

not implement
mydb v0.0.2> status
[status]
Count:101
GetNum:1
SetNum:201
DelNum:0

mydb v0.0.2> keys
key68=68
key76=76
key82=82
key58=58
key60=60
key83=83
key90=90
key39=39
key42=42
key52=52

mydb v0.0.2> set a 1
set a 1 success
a set ok
mydb v0.0.2> get a
get a success
1
mydb v0.0.2> get aa
get aa success
aa not exist
mydb v0.0.2> get key1
get key1 success
1
mydb v0.0.2> exit
```
