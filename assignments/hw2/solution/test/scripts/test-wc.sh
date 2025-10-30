#!/bin/bash
go run cmd/wordcount/main.go master sequential test/data/pg-*.txt
sort -n -k2 mrtmp.wcseq | tail -10 | diff - test/scripts/test-wc.out > diff.out
if [ -s diff.out ]
then
echo "Failed test. Output should be as in test-wc.out. Your output differs as follows (from diff.out):" > /dev/stderr
  cat diff.out
else
  echo "Passed test" > /dev/stderr
fi

