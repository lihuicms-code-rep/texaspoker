#!/usr/bin/env bash
rm -fr ./gameLog/
ps aux|grep texaspoker|grep -v grep|awk '{print $1}'|xargs kill -s 9
echo "stop texaspoker success"