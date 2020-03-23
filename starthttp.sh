#!/usr/bin/env bash
if [ -d backLog ];
then
   echo "backLog dir exists"
else
   mkdir ./backLog/
fi

./texaspoker.exe -config ./res/ -st 2 2>&1 > ./backLog/httpserver.log &
echo "start texaspoker http by daemon mode success"