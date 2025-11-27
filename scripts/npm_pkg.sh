#!/usr/bin/env bash

find ./npm/platforms -name "package.json" -execdir sh -c "$*" ";"
