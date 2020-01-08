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
		echo "\033[32mOK\033[0m"
	else
		echo "\033[31mNOPE\033[0m"
	fi;
done

# RESULT=

#echo "$RESULT"

