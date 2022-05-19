#!/usr/bin/env bash

echo "run api ..."
nohup ./api 2>&1 &
sleep 1
git rm -r --cached .