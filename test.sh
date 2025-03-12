#!/bin/bash
#

./biu

if [[ $? -eq 32 ]]; then
  echo ok
else
  echo error
  exit 1
fi
