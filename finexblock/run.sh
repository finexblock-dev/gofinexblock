#!/bin/bash

directories=("secure" "safety" "database" "interceptor")

for dir in ${directories[@]}
do
    touch $dir/$dir.md
done