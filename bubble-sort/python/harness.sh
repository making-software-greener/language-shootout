for i in `seq ${1}`
do
   echo $i
   time cat ../${2}.dat | (python3 -u bubble.py >/dev/null)
   echo ""
done
