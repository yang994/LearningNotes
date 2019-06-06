#这个脚本用于测试时空值的计算。


echo "起metb"
metb init -n 11
metb for-each mefs config Test true --json
metb start
metb run 1 mefs config Role keeper
metb run 2 mefs config Role keeper
metb run 3 mefs config Role keeper
metb run 4 mefs config Role keeper
metb run 5 mefs config Role provider
metb run 6 mefs config Role provider
metb run 7 mefs config Role provider
metb run 8 mefs config Role provider
metb run 9 mefs config Role provider
metb run 10 mefs config Role provider
metb restart
metb connect 1 2
metb connect 1 3
metb connect 1 4
metb connect 1 5
metb connect 1 6
metb connect 1 7
metb connect 1 8
metb connect 1 9
metb connect 1 10
metb connect 2 3
metb connect 2 4
metb connect 2 5
metb connect 2 6
metb connect 2 7
metb connect 2 8
metb connect 2 9
metb connect 2 10
metb connect 3 4
metb connect 3 5
metb connect 3 6
metb connect 3 7
metb connect 3 8
metb connect 3 9
metb connect 3 10
metb connect 4 5
metb connect 4 6
metb connect 4 7
metb connect 4 8
metb connect 4 9
metb connect 4 10   #keeper需要与所有节点相连，provider不需要
metb connect 0 1
metb connect 0 2
metb connect 0 3
metb connect 0 4
metb connect 0 5
metb connect 0 6
metb connect 0 7
metb connect 0 8
metb connect 0 9
metb connect 0 10   #user与所有节点相连
echo "完成！"


:>>EOF
echo "等待2分钟"
echo "上传文件"
echo "开始每隔1分钟测试一次"
metb shell 2
mefs test stpay
exit
echo "结束！"
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
EOF



