#!/bin/sh

[[ -f "$PWD/package.json" ]]
jq .scripts "$PWD/package.json"
