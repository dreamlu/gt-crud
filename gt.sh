#!/usr/bin/env bash
git pull
./devMode.sh prod

cd docker
./pushAll.sh
cd ..
./devMode.sh dev