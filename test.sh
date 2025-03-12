#!/bin/bash
#

./biu

if [[ $? -eq 3 ]]; then
  echo ok
else
  echo error
  exit 1
fi
