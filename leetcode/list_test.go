package leetcode

import "testing"

func checkListLen(t *testing.T, l *ListNode, length int) bool {
	if n := l.Len(); n != length {
		t.Errorf("l.Len() = %d, want %d", n, length)
		return false
	}
	return true
}

func TestList(t *testing.T) {
	nums := []int{}
	l := NewList(nums)
	if !checkListLen(t, l, len(nums)) {
		return
	}

	nums = append(nums, 1, 2, 3)
	l = NewList(nums)
	if !checkListLen(t, l, len(nums)) {
		return
	}
}
