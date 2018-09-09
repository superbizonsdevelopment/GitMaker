#!/bin/sh

mainfile=$1
echo "Running on Building go projects!"
echo "Main file: $mainfile"

cd /Applications/GitMaker-Server/apps/$2

go build $mainfile

echo "Go Projects Built"

exit 0
