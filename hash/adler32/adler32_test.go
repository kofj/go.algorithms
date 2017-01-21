package adler32

import (
	"hash/adler32"

	"testing"
)

func TestChecksum(t *testing.T) {
	var tests = []struct {
		data []byte
	}{
		{[]byte("Wikipedia")},
		{[]byte("abc")},
		{[]byte("data")},
		{[]byte("中文字符串")},
		{[]byte("This is a sentence.")},
	}

	for _, test := range tests {
		checksum := Checksum(test.data)
		expect := adler32.Checksum(test.data)
		t.Logf("data: %s, checksum: %d, expect: %d\n", test.data, checksum, expect)
		if checksum != expect {
			t.Fail()
		}
	}
}
