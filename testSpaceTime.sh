#这个脚本用于测试时空值的计算。


echo "起metb"
echo "等待2分钟"
echo "上传文件"

echo "开始每隔1分钟测试一次"

second=0
testtime=1
theoretical=0
size=4096
cycle=5
while true
do
    if [ `expr $second % $cycle` -eq 0 ]
    then
        theoretical=`expr $size \* $second`
        printf "第%d次测试结果\n理论时空值：%d\n实际时空值：XXX\n" $testtime $theoretical
        testtime=`expr $testtime + 1`
    fi
    second=`expr $second + 1`
    sleep 1
    
done



