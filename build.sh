#!/bin/bash

# Config
PROJECT_TITLE="algorithms-and-data-structures-pieter-groenendijk"
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

go build -o "$BUILD_DIR/$PROJECT_TITLE" -v "./" 

cp -rv "." "$BUILD_DIR/source"


## Document
### Config


### Execution
cd "$DOCUMENT_DIR"

pandoc \
    -f markdown+raw_html \
    --output="${BUILD_DIR}/${PROJECT_TITLE}.pdf" \
    --pdf-engine=pdflatex \
    --citeproc \
    --toc \
    -M colorlinks \
    "${DOCUMENT_DIR}/readme.md"

cp -rv "." "$BUILD_DIR/document"
