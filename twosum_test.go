package leetcode

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {

	nums := [3]int{3, 2, 4}
	target := 6
	fmt.Println(TwoSum(nums[:], target))

}
