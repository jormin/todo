### About

Todo Manager

### Install

```
go install
```

### Usage

```shell
NAME:
   todo - A simple tool to manage your todo list

USAGE:
   todo [global options] command [command options] [arguments...]

VERSION:
   v1.0.0

DESCRIPTION:
   A simple tool to manage your todo list

COMMANDS:
   add      add todo
   edit     edit todo
   ls       show config fund list
   rm       remove todo
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

#### ls

```shell
todo ls [command options]
```

- -a: show all todo list (default: false)
- -d: show todo list of one date (default: today)

##### show all todo list

```shell
➜  todo git:(master) ✗ todo ls -a
#     ID                       Date                Content     Level     Status         UpdateTime
1     c3en56g6n88j8kosta40     20210701(today)     aaa         Low       Incomplete     2021-07-01 15:45:30
2     c3emlm86n88invcs8s40     20210630            aaa         Low       Incomplete     2021-07-01 15:12:25
3     c3emlko6n88innjbuf30     20210701(today)     123         High      Completed      2021-07-01 15:12:19
```

##### show todo list of today

```shell
➜  todo git:(master) ✗ todo ls   
#     ID                       Date         Content     Level     Status         UpdateTime
1     c3en56g6n88j8kosta40     20210701     aaa         Low       Incomplete     2021-07-01 15:45:30
2     c3emlko6n88innjbuf30     20210701     123         High      Completed      2021-07-01 15:12:19
```

##### show todo list of one date

```shell
➜  todo git:(master) ✗ todo ls -d 20210630
#     ID                       Date         Content     Level     Status         UpdateTime
1     c3emlm86n88invcs8s40     20210630     aaa         Low       Incomplete     2021-07-01 15:12:25
```

#### add

```shell
todo add [command options] [content: todo content]
```

- -d: the date of todo (default: today)
- -l: the level of todo, optional values are 1-3, 1 is the minimum level and 3 is the maximum level (default: 1)
- -s: the status of todo, optional values are 0-1, 0 means incomplete and 1 means completed (default: 0)

```shell
➜  todo git:(master) ✗ todo add -d 20210630 -l 3 -s 0 "test for add"
c3eog906n88kgk167ht0
```

#### edit

```shell
todo edit [command options] [id: unique todo id]
```

- -c: the content of todo
- -d: the date of todo (default: today)
- -l: the level of todo, optional values are 1-3, 1 is the minimum level and 3 is the maximum level (default: 1)
- -s: the status of todo, optional values are 0-1, 0 means incomplete and 1 means completed (default: 0)

```shell
➜  todo git:(master) ✗ todo edit -c "test for edit" -l 3 -s 0 c3en56g6n88j8kosta40
edit todo c3en56g6n88j8kosta40 success
```

#### rm

- -a: show all todo list (default: false)

```shell
todo rm [command options] [id1: unique todo id] [id2] ... [idn]
```

##### remove todo by id

```shell
➜  todo git:(master) ✗ todo rm c3emlko6n88innjbuf30
remove todo c3emlko6n88innjbuf30 success
```

##### remove all todos

```shell
➜  todo git:(master) ✗ todo rm -a
remove all todos success
```