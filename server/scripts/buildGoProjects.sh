#!/bin/sh

mainfile=$1
echo "Running on Building go projects!"
echo "Main file: $mainfile"

go build $mainfile

echo "Go Projects Builded"

exit 0
