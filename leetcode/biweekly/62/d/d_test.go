// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	examples := [][]string{
		{
			`[2,-1,2]`, `3`, 
			`1`,
		},
		{
			`[0,0,0]`, `1`, 
			`2`,
		},
		{
			`[22,4,-25,-20,-15,15,-16,7,19,-10,0,-13,-14]`, `-33`, 
			`4`,
		},
		
	}
	targetCaseNum := 1
	if err := testutil.RunLeetCodeFuncWithExamples(t, waysToPartition, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-62/problems/maximum-number-of-ways-to-partition-an-array/