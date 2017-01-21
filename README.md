# Ohacan

A CLI tool and Golang Client Library for jobcan(http://jobcan.ne.jp).
Now it support clock in/out only.

## Setup

```
go get github.com/swdyh/ohacan/ohacan
```

TODO: download binary

set environment variables.

```
JOBCAN_GROUP_ID=111
JOBCAN_CODE=xxx
```

JOBCAN_CODE is access url param. ex) http://jobcan.jp/m?code=xxx
JOBCAN_GROUPID is select button value in clock in page.

I recommend to use envcahin(https://github.com/sorah/envchain).

## Usage

Clock in

```
ohacan in
```

or

```
ohacan oha
```


Clock out

```
ohacan out
```

or

```
ohacan otsu
```

## Lisence

APACHE LICENSE, VERSION 2.0 
http://www.apache.org/licenses/
