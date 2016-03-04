# Simple Interactive Wrapper of CLI with sub-command with Go

Nowdays, we have lots of CLI tools with many sub-commands, eg. hdfs, etcdctl... 
When we use these tools continuously， we may input main-command again by again.

This little wrapper will help you to reduce that useless input.

## Usage

hdfs is the CLI.

```bash
# install test cli(hdfs), just for example
go get github.com/colinmarc/hdfs/cmd/hdfs
# config to connect to your hadoop
export HADOOP_NAMENODE="wxnn:8020"
```

```bash
go build -o interactive
./interactive hdfs
```

1. `!quit` will quit the current interactive session
-  `!cmd` will execute `cmd` in shell, eg. `!ls` will list files.
-  `cmd` will execute `cmd` will the main-cmd, eg. `ls /` will run `hdfs ls /`
and list files in the hdfs.

```
./interactive hdfs
> ls -l /
drwxr-xr-x hbase       hbase  0 Feb 23 14:15 hbase
drwxrwxrwx  hdfs  supergroup  0 Feb 23 14:23 tmp
drwxrwxrwx  hive        hive  0 Feb 23 17:00 user
> !ls
interactive
main.go
README.md
> !quit
```
