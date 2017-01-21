#!/bin/sh

go build -o dev_server/dev_server dev_server/main.go
./dev_server/dev_server &
pid=$!
make install
JOBCAN_CODE=a JOBCAN_GROUP_ID=1 OHACAN_DEV_SERVER="http://localhost:8080" ohacan in
sleep 1
JOBCAN_CODE=a JOBCAN_GROUP_ID=1 OHACAN_DEV_SERVER="http://localhost:8080" ohacan out
kill $pid
rm dev_server/dev_server
