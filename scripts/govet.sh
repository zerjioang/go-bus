#!/bin/bash

cd "$(dirname "$0")"

# move to project root dir from ./scripts to ./
cd ..

echo "Checking source code with Go vet"

#get all files excluding vendors
filelist=$(go list ./... | grep -vendor)
for file in ${filelist}
do
	echo "static analysis of package $file"
	go vet $@ ${file}
done

echo "Code checking done!"
