#！/usr/bin/env bash

# $GOOS         $GOARCH
# darwin         386
# darwin         amd64
# freebsd     386
# freebsd     amd64
# linux         386
# linux         amd64
# linux         arm     incomplete
# windows     386     incomplete

rm -rf ./history
rm concatenate.tar.gz
GOOS=linux GOARCH=amd64 go build -v -o conc ./cmd
GOOS=windows GOARCH=amd64 go build -v -o conc.exe ./cmd
cd ..
tar -zcf concatenate.tar.gz concatenate
cd -
