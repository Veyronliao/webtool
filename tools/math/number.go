package math

import "sort"

func FindMiddleNumber(a, b, c float64) float64 {
	//创建切片保存三个数字
	number := []float64{a, b, c}
	//对切片进行排序
	sort.Float64s(number)
	//返回切片中间的数
	return number[1]
}
func minDifference(arr []int) (int, int, int) {
	if len(arr) < 2 {
		return 0, 0, -1 // 如果数组元素不足两个，则返回-1表示错误或无解
	}
	// 对数组进行排序
	sort.Ints(arr)
	minDiff := arr[1] - arr[0]      // 初始化最小差值为第一个元素的差值
	first, second := arr[0], arr[1] // 初始化最小差值对应的两个数
	// 遍历数组，查找最小差值
	for i := 1; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]
		if diff < minDiff {
			minDiff = diff
			first, second = arr[i], arr[i+1]
		}
	}
	return first, second, minDiff
}
