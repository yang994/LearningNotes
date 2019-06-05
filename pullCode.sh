#!/bin/bash

code_path="$GOPATH/src/github.com/ipfs/go-ipfs"
cd $code_path

git checkout private
git stash
git checkout mem
git pull
git checkout private
git rebase mem
git stash pop

echo "完成！"