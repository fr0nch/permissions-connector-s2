#!/bin/bash
# build.sh - For Linux builds

set -ex

mkdir -p $PREFIX/bin

cp bin/libpermissions-connector-s2.so $PREFIX/bin/
cp permissions-connector-s2.pplugin $PREFIX/
cp permissions-connector-s2.yaml $PREFIX/

chmod 755 $PREFIX/bin/libpermissions-connector-s2.so
chmod 644 $PREFIX/permissions-connector-s2.pplugin $PREFIX/permissions-connector-s2.yaml
