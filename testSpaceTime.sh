#这个脚本用于测试时空值的计算。


echo "起metb"
metb init -n 7
metb for-each mefs config Test true --json
metb start
metb run 1 mefs config Role keeper
metb run 2 mefs config Role provider
metb run 3 mefs config Role provider
metb run 4 mefs config Role provider
metb run 5 mefs config Role provider
metb run 6 mefs config Role provider

metb restart
metb connect 1 2
metb connect 1 3
metb connect 1 4
metb connect 1 5
metb connect 1 6
metb connect 2 3
metb connect 2 4
metb connect 2 5
metb connect 2 6
metb connect 3 4
metb connect 3 5
metb connect 3 6
metb connect 4 5
metb connect 4 6 
metb connect 0 1
metb connect 0 2
metb connect 0 3
metb connect 0 4
metb connect 0 5
metb connect 0 6   #user与所有节点相连#keeper需要与所有节点相连，provider不需要
echo "完成！"


