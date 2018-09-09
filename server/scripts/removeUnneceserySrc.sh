#!/bin/sh

echo "Removing unnecessery src from cd /Applications/GitMaker-Server/repos/$1"

cd /Applications/GitMaker-Server/repos

rm $1

echo "Unnecessery code removed!"

exit 0
