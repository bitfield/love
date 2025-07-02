#!/bin/bash

set -e
for TOPIC in */; do
  for LISTING in ${TOPIC}*/; do
    # Update deps:
    # cd $LISTING && go get -t -u && cd -
    # cd $LISTING && go fmt *.go && cd -
    ./run_testscript.sh $LISTING
  done
done
