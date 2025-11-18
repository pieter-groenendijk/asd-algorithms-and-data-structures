#!/bin/bash

# Config
DIR="$(dirname "$(realpath "$0")")"
SOURCE_DIR="${DIR}/source"
BUILD_DIR="${DIR}/generated/build"
PACKAGE_DIR="${DIR}/generated/package"

# Prepare
sh "${DIR}/build.sh"

rm -r "$PACKAGE_DIR"
mkdir "$PACKAGE_DIR"

# Execution
cd "${BUILD_DIR}"

zip -r "${PACKAGE_DIR}/package.zip" ./
