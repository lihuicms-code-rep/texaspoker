#!/usr/bin/env bash
ps aux|grep texaspoker|grep -v grep|awk '{print $1}'|xargs kill -s 9
rm -fr ./gameLog/ ./backLog/
echo "stop texaspoker success"