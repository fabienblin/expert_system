#!/bin/bash

FICHIERS=$(ls -l examples/correctInput/other/test_outputs | cut -d ":" -f 2- | cut -d " " -f 2 | tr "\n" " ")
IFS=$' ' read -r -a FILETAB <<< "$FICHIERS"
for elem in "${FILETAB[@]}"
do
  if [ "$elem" != "${FILETAB[0]}" ]; then
    if [ "$elem" != "${FILETAB[1]}" ]; then
      if [ "$elem" != "${FILETAB[2]}" ]; then
        $(mv "examples/correctInput/other/test_outputs/$elem" "examples/correctInput/other/test_outputs/$elem.txt")
        $(cat "examples/correctInput/other/test_outputs/$elem.txt")
      fi;
    fi;
  fi;
done