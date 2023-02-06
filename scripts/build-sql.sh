#!/bin/bash

loadArray () {
  file=$1
  quad=$2
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
   NAME=$name QUADRANT=$quad URL=$url CSVTABLE="ATS2023" SAMPSZ=244 envsubst <ats.sql
  if [ "$iter" != "${#arr[@]}" ] ; then
     echo "Union"
  fi
  let iter++
  done
}

readControl() {
  loadArray language.txt LANGUAGE
  echo "Union"
  loadArray webframe.txt "WEB FRAMEWORKS"
  echo "Union"
  loadArray otherframe.txt "OTHER FRAMEWORKS"
  echo "Union"
  loadArray database.txt "DATABASE"
}
readControl