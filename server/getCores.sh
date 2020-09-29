#!/bin/sh

number=`lscpu | grep NUMA | head -n 1 | awk '{print $3}'` && lscpu | grep NUMA | tail -n $number | sed "s/node//g" | awk '{print $2" "$4}'