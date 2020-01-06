#!/bin/bash

FILES=$(find examples/correctInput -name "*.txt" | cut -d "/" -f 3- | tr "\n" " ")
IFS=$' ' read -r -a FILETAB <<< "$FILES"
for elem in "${FILETAB[@]}"
do
	echo ---------------------
	PROGRAM=$(make tests correctInput/$elem | grep solution)
	SOLUTION=$(cat examples/correctInput/$elem | grep solution)
	echo "$elem"
	echo $PROGRAM
	echo $SOLUTION
	if [ "$SOLUTION" = "$PROGRAM" ]; then
		echo OK
	else
		echo NOPE
	fi;
done

# RESULT=

#echo "$RESULT"

