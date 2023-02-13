#!/bin/bash

loadArray () {
  file=$1
  quad=$2
  high=$3
  mid=$4

  arr=()
  while IFS= read -r line; do
   arr+=("$line")
  done <$file  
  iter=1
  for i in "${arr[@]}"
  do
   name="$(echo $i| cut -d "|" -f1)"
   url="$(echo $i| cut -d "|" -f2)"
   if [ "$url" ==  "$name" ] ; then
     url="#"
   fi
   NAME=$name QUADRANT=$quad URL=$url CSVTABLE="ATS2023" SAMPSZ=244 HIGH=$high MID=$mid envsubst <ats.sql
  if [ "$iter" != "${#arr[@]}" ] ; then
     echo "Union"
  fi
  let iter++
  done
}

readControl() {
  loadArray ide.txt "DEV IDE"
  echo "Union"
  loadArray revcs.txt "REVISION CONTROL"
  echo "Union"
  loadArray opsys.txt "OPERATING SYSTEMS"
  echo "Union"
  loadArray cicd.txt "CI/CD TOOLS"
}
readControl