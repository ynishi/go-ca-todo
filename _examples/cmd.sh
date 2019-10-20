#!/bin/sh

BASEDIR=$(dirname "$0")
PRJDIR=$BASEDIR/../

sh $PRJDIR/build.sh

$PRJDIR/cmd/todo/todo

$PRJDIR/cmd/todo/todo -sub list

TITLE=title1 TAG=tag1 DUE="2020-01-01T00:00:00+00:00" $PRJDIR/cmd/todo/todo -sub add