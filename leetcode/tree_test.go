package leetcode

import "testing"

type TestData struct {
	nums []any
	want string
}

func TestTree(t *testing.T) {
	testDatas := []TestData{
		{
			nums: []any{},
			want: "[]",
		},
		{
			nums: []any{3},
			want: "[3]",
		},
		{
			nums: []any{3, 9, 20},
			want: "[3, 9, 20]",
		},
		{
			nums: []any{3, 9, 20, nil, nil, 15, 7},
			want: "[3, 9, 20, 15, 7]",
		},
	}
	for _, testData := range testDatas {
		root := NewTree(testData.nums)
		str := root.outputString()
		if str != testData.want {
			t.Fatalf("str = %v, want %v", str, testData.want)
		}
	}
}
