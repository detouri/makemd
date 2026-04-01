#!/bin/bash

git clone https://github.com/detouri/makemd.git
cd makemd
make build
./dist/makemd version
