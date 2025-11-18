#!/bin/bash

# Config
DIR="$(dirname "$(realpath "$0")")"
SOURCE_DIR="${DIR}/source"
DOCUMENT_DIR="${DIR}/document"
BUILD_DIR="${DIR}/generated/build"


# Prepare
rm -r "$BUILD_DIR"
mkdir "$BUILD_DIR"


## Source
### Config


### Execution
cd "$SOURCE_DIR"

go build -o "$BUILD_DIR" -v "./" 

cp -rv "." "$BUILD_DIR/source"


## Document
### Config
### Execution
cd "$DOCUMENT_DIR"
