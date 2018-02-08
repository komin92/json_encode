package main

import "testing"

func TestSeqToJSON(t *testing.T) {
	data := []byte{0x31, 0xa, 0x32, 0xa, 0x33, 0xa, 0x34, 0xa, 0x35, 0xa, 0x36, 0xa, 0x37, 0xa, 0x38, 0xa, 0x39, 0xa, 0x31, 0x30, 0xa}
	JSON := string(ConvertSliceJSON(ConvertInputToLines(string(data))))
	expectedOutput := "[\"1\",\"2\",\"3\",\"4\",\"5\",\"6\",\"7\",\"8\",\"9\",\"10\"]"
	if expectedOutput != JSON {
		t.Fail()
	}
}
