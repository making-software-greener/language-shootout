for i in `seq ${1}`
do
   echo $i
   time cat ../${2}.dat | (java BubbleSort >/dev/null)
   echo ""
done
