#!/bin/bash
#

./biu

if [[ $? -eq 1 ]]; then
  echo ok
else
  echo error
  exit 1
fi
