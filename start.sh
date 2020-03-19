#!/usr/bin/env bash
mkdir ./backLog/
./texaspoker.exe -config ./res/ -st 2 2>&1 > ./backLog/texaspoker.log &
echo "start texaspoker by daemon mode success"