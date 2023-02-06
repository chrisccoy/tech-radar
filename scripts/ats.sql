select '$NAME' as "name",
 CASE
    WHEN EXISTS (select count(*) from csv."$CSVTABLE" where "$NAME" like '%NEXT%' or "$NAME" like '%PAST%' or "$NAME" like '%Both%' having count(*)::float/$SAMPSZ > .1) THEN 'HIGH'
    WHEN EXISTS (select count(*) from csv."$CSVTABLE" where "$NAME" like '%NEXT%' or "$NAME" like '%PAST%' or "$NAME" like '%Both%' having count(*)::float/$SAMPSZ between .05 and .1) THEN 'MID'
    WHEN EXISTS (select count(*) from csv."$CSVTABLE" where "$NAME" like '%NEXT%' or "$NAME" like '%PAST%' or "$NAME" like '%Both%' having count(*) > 0) THEN 'LOW'
    ELSE 'UNUSED'
 END as "ring",
 '$QUADRANT' as "quadrant",
'FALSE' as "isNew",
'$URL' as "url",
CASE
   WHEN EXISTS (select 1 as foo where (select count(*) from csv."$CSVTABLE" where "$NAME" like '%NEXT%') > (select count(*) from csv."$CSVTABLE" where "$NAME" like '%PAST%') ) THEN 1
   WHEN EXISTS (select 1 as foo where ((select count(*) from csv."$CSVTABLE" where "$NAME" like '%NEXT%') + 
                (select count(*) from csv."$CSVTABLE" where "$NAME" like '%Both%')) < (select count(*) from csv."$CSVTABLE" where "$NAME" like '%PAST%') ) THEN -1   ELSE 0
END as "Direction"
