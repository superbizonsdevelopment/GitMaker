#!/bin/sh

gitProjectToDownload=$1

echo "Downloading src from: $gitProjectToDownload"

cd /Applications/GitMaker-Server/repos

git clone $gitProjectToDownload

echo "Src Downloaded!!!"

exit 0
