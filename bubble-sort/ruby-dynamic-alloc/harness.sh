for i in `seq ${1}`
do
   echo $i
   time cat ../${2}.dat | (ruby bubble.rb >/dev/null)
   echo ""
done
