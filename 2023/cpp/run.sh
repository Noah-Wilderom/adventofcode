#!/bin/bash

DAY=$1

if [ -z "$DAY" ]; then
	echo "Specify the day. For example: ./run.sh 01"
	exit 1
fi

echo "Compiling day $DAY"

cd build 

cmake -DDAY_DIRECTORY=$DAY ../

make

chmod +x day-$DAY

./day-$DAY

cd ../
