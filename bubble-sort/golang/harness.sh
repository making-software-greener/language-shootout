for i in `seq ${1}`
do
   echo $i
   time cat ../${2}.dat | (./bubble >/dev/null)
   echo ""
done
