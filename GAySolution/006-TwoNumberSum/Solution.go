LeetCode это задача №1 — Two Sum
Time: O(n)
Space: O(n)

package main
import "sort"

func (_ Solution) TwoNumberSum(array []int, target int) []int {
	var num2 int
targeth := make(map[int]int)
for i:= 0; i<len(array);i++{
	num2 = target - array[i]
if _, ok := targeth[num2]; ok {
        result := []int{num2, array[i]}
sort.Ints(result)
return result
}
targeth[array[i]] = array[i]
}
return []int{}
}
