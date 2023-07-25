package easy

import (
	"grind75/leetcode/easy"
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s			string
		expected 	bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
	}

	for _, test := range tests {
		output := easy.IsPalindrome(test.s)

		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("MaxProfit(%s) actual: %t - expected: %t",
				test.s, output, test.expected)
		}
	}
}
