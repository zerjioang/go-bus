#!/bin/bash

cd "$(dirname "$0")"

# move to project root dir from ./scripts to ./
cd ..

echo "Formatting source code with Go fmt"

#get all packages excluding vendors
filelist=$(go list ./... | grep -vendor)
for file in $filelist
do
	echo "Formatting file $file"
	go fmt $file
done

echo "Code formatting done!"