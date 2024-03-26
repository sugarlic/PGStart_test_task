#!/bin/bash

A="10"
B="5"
C=`expr $A + $B`
printf "A=10 B=5 C=expr \$A + \$B C=%d \n" "$C"

# пример цикла по i
I=0
while [ $I -lt 15 ]
do
    printf "0x%02x " "$I"
    I=`expr $I + 1`
done
echo