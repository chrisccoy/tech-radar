#!/bin/bash

loadArray () {
  file=$1
  quad=$2
  high=$3
  mid=$4
  sqlfile=$5

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
   NAME=$name QUADRANT=$quad URL=$url CSVTABLE="ATS2023" SAMPSZ=244 HIGH=$high MID=$mid envsubst <$sqlfile
  if [ "$iter" != "${#arr[@]}" ] ; then
     echo "Union"
  fi
  let iter++
  done
}

buildLang() {
  loadArray language.txt LANGUAGE ".1" ".05" "ats.sql"
  echo "Union"
  loadArray webframe.txt "WEB FRAMEWORKS" ".1" ".05" "ats.sql"
  echo "Union"
  loadArray otherframe.txt "OTHER FRAMEWORKS" ".1" ".05" "ats.sql"
  echo "Union"
  loadArray database.txt "DATABASE" ".1" ".05" "ats.sql"
}
buildCICD() {
    loadArray ide.txt "DEV IDE" ".25" ".1" "ats.sql"
    echo "Union"
    loadArray revcs.txt "REVISION CONTROL" ".25" ".1" "ats-personalvprofessional.sql"
    echo "Union"
    loadArray opsys.txt "OPERATING SYSTEMS" ".25" ".1" "ats-personalvprofessional.sql"
    echo "Union"
    loadArray cicd.txt "CI/CD TOOLS" ".25" ".1" "ats.sql"
}
buildLang > lang.sql
buildCICD > cicd.sql
