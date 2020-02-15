PROCESS=`ps -ef|grep gorm_generate|grep -v grep|grep -v PPID|awk '{ print $2}'`
for i in $PROCESS
do
  echo "Kill the gorm_generate process [ $i ]"
  kill -9 $i
done
echo "stop ok"