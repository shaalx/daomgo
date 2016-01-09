#!/bin/sh
for i in `ls | grep '[0-9]'`;do
echo $i
rm -rf $i
done
