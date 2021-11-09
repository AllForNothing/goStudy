package split_string

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("abc", "b")
	want := []string{"a", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:%v but got %v", want, got)
	}
}


func twoSum(nums []int, target int) []int {
	ret :=make([]int,2,2)
	map1 := make(map[int]int,len(nums))
	for index,v := range nums {
		if i, ok := map1[target-v]; ok{
			ret[0] =index
			ret[1] =i
			return ret
		}else{
			map1[v] = index
		}
	}
	return ret
}