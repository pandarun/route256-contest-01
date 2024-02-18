#!/bin/bash

# cat scripts from 1 to 75 and compare with i.a result

#cd ../test
for i in {1..11}
do
    echo "Test $i"
    cat $i | ./main > out
    diff out $i.a
done