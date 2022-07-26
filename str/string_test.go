package str

import (
	"gin_xrom_framework/pkg/assert"
	"testing"
)

func TestBlurEmail(t *testing.T) {
	type Case struct {
		email  string
		expect string
	}
	cases := []Case{
		{"1313831783@qq.com", "13****3@qq.com"},
		{"belonk@126.com", "be****k@126.com"},
	}
	for _, c := range cases {
		dst := BlurEmail(c.email)
		if dst != c.expect {
			t.Errorf("test failed, expect: %v, but found: %v", c.expect, dst)
		}
	}
}

func TestEndsWith(t *testing.T) {
	assert.True(EndsWith("", ""))
	assert.True(EndsWith("a", ""))
	assert.True(!EndsWith("", "a"))

	str := "aaabb123b"
	assert.True(EndsWith(str, "b"))
	assert.True(EndsWith(str, "3b"))
	assert.True(EndsWith(str, "23b"))
	assert.True(EndsWith(str, "123b"))
	assert.True(!EndsWith(str, "a"))

	assert.True(StartsWith("", ""))
	assert.True(StartsWith("a", ""))
	assert.True(!StartsWith("", "a"))

	assert.True(StartsWith(str, "a"))
	assert.True(StartsWith(str, "aa"))
	assert.True(StartsWith(str, "aaa"))
	assert.True(StartsWith(str, "aaab"))
	assert.True(!StartsWith(str, "aaab1"))
	assert.True(!StartsWith(str, "1aaab1"))
}
