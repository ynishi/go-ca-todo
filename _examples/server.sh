#!/bin/sh

BASEDIR=$(dirname "$0")
PRJDIR=$BASEDIR/../

sh $PRJDIR/build.sh

$PRJDIR/web/todo/todo &
TODO_PID=$!

curl 'localhost:8080/'
curl 'localhost:8080/add?title=ti2&tag=ta1&due=2020-01-01T00%3A00%3A00%2B00%3A00'
curl 'localhost:8080/'

kill $TODO_PID