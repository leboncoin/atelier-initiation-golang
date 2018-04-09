#!/usr/bin/env bash

TARGET_DIR=$1

GO_VERSION=go1.10.1

if [ -z "$TARGET_DIR" ]; then
    echo "target directory is missing"
    echo ""
    echo "  $0 /media/myusername/myusbkeyname"
    echo ""
    echo "creates a directory in the target directory, with a clone of the git repo"
    echo "plus the installation binaries for $GO_VERSION"
    exit 1
fi

if [[ -e "$TARGET_DIR/atelier-go" ]]; then
    read -p "$TARGET_DIR/atelier-go already exists. Are you sure you want to delete & recreate it? " -n 1 -r
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        exit 1
    fi

    rm -rf $TARGET_DIR/atelier-go
fi

mkdir -p $TARGET_DIR/atelier-go
cd $TARGET_DIR/atelier-go
rm -rf atelier-initiation-golang
git clone https://github.com/leboncoin/atelier-initiation-golang.git
mkdir $TARGET_DIR/atelier-go/install
cd $TARGET_DIR/atelier-go/install

wget -O $GO_VERSION.darwin-amd64.pkg https://dl.google.com/go/$GO_VERSION.darwin-amd64.pkg
wget -O $GO_VERSION.linux-amd64.tar.gz https://dl.google.com/go/$GO_VERSION.linux-amd64.tar.gz
wget -O $GO_VERSION.windows-amd64.msi https://dl.google.com/go/$GO_VERSION.windows-amd64.msi
