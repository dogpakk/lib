#!/bin/bash

cd slice
go test
cd ..

cd str
go test
cd ..

cd financial
go test
cd ..
