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

go build -o conc ./cmd
