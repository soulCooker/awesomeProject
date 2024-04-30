# 二分查找

查找存在的下标或者插入位置
 （k）序列，mid = k,start=k,end=k+1, 如果小于mid，则end=k,start=k，循环结束，插入的位置是k；如果大于mid，start=k+1,end=k+1,插入的位置是k+1，都是start
(k, k+1)序列，mid=k+1，如果小于mid，则回到情况一;如果大于mid，start=k+2，end=k+1,插入的位置是k+2，也是start

综上所述，可搜索空间结束时，start都是待插入的位置。
