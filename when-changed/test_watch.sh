#!/usr/bin/env bash

when-changed -v -r -1 -s ./ "go test -v -test.run $1"