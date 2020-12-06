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

# GOOS=linux GOARCH=amd64 go build -vxn -o conc ./cmd
GOOS=windows GOARCH=amd64 go build -vxn -o conc.exe ./cmd
