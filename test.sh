#!/bin/bash

DIR="$(dirname "$(realpath "$0")")"
SOURCE_DIR="${DIR}/source"

cd "$SOURCE_DIR"

go test -v --cover 

cd "$DIR"
