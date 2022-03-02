package common

//常见的能用语言描述的操作
//算法思路：先将算法用可以写出的程序的语言描述出来
/*
1.从前往后遍历[iter = start(); iter != end(); iter++](结束时iter指向end() end()表示容器中最后一个元素的后一个元素)
2.从后往前遍历
3.在某个顺序容器中查找某个数的位置
4.排序
5.往链表尾部插入一个元素
6.最大值(n>=2)
7.最小值(m>=2)
8.map[对数组所有元素做一个映射]
9.reduce[对数组所有元素做一个聚集操作]
10.filter[根据条件对数组进行一次过滤]
11.count-if
12.swap
13.判断哈希表是否存在某个key并取得该key
if val, ok := hashTable[target-value]; ok {}
14.处理两个标准数的运算结果，并得出是否溢出或进位
sum ,carry = sum%10,carry/10
15.得出数组中两个元素间的元素个数
j-i+1
16.使某个变量始终为两个数组中的最小长度
func f(nums1,nums2){
	len1 = len(nums1)
	len2 = len(nums2)
	if(len1 > len2){
		f(nums2,nums1)
	}
}
17.初始化二维数组
row,col
arr := make([][]int,row)
for i < row{
	arr[i] = make([]int,col)
	i++
}
*/
