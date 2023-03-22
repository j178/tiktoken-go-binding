#!/usr/bin/env bash

set -ex

cargo -C tiktoken-cffi build --release -Z unstable-options --out-dir .
go test -v
