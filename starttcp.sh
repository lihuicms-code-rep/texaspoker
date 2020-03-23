#!/usr/bin/env bash
if [ -d backLog ];
then
   echo "backLog dir exists"
else
   mkdir ./backLog/
fi

./texaspoker.exe -config ./res/ -st 1 2>&1 > ./backLog/tcpserver.log &
echo "start texaspoker tcp by daemon mode success"