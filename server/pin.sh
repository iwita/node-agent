#!/bin/sh
POD_NAME=$1
CORES=$2


#echo "POD name: $POD_NAME"
#echo "Cores to pin: $CORES"
CONTAINER_ID=$(docker ps | awk '{printf $1; for (i=3;i<=NF;i++) printf FS$i; print NL}' | awk -v pn=${POD_NAME} '/POD/{next} $0 ~ pn {print $1}')

#echo "Found Container: $CONTAINER_ID"
docker update --cpuset-cpus ${CORES} ${CONTAINER_ID}

if [ $? = 0 ]
then
  echo "Pin ok"
  return 0
else
  echo "Pin NOT ok"
  return 1
fi
