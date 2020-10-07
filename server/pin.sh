#!/bin/sh
POD_NAME=$1
CORES=$2

#echo $CORES
CONTAINER_ID=$(docker ps | awk '{printf $1; for (i=3;i<=NF;i++) printf FS$i; print NL}' | grep relaxed | awk '{print $1}')

#echo "Found Container: $CONTAINER_ID"
docker update --cpuset-cpus ${CORES} ${CONTAINER_ID}

if [ $? = 0 ]
then
  echo "Pin ok"
else
  echo "Pin NOT ok"
fi
