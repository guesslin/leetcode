#!/bin/bash
declare -a arr=()
read -a line < "file.txt"
COLS=${#line[@]}
index=0
while read -a line ; do
	for (( COUNTER=0; COUNTER<${#line[@]}; COUNTER++ )); do
		arr[$index]=${line[$COUNTER]}
		((index++))
	done
done < "file.txt"

for (( ROW = 0; ROW < COLS; ROW++ )); do
	tmpstr=""
	for (( COUNTER = ROW; COUNTER < ${#arr[@]}; COUNTER += COLS )); do
		tmpstr+=${arr[$COUNTER]}
		tmpstr+=" "
	done
	printf "%s\n" "$(echo -e "${tmpstr}" | sed -e 's/[[:space:]]*$//')"
done
