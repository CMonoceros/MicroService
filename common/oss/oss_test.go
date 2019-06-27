package oss

import "testing"

func TestParseObjectName(t *testing.T) {
	res, err := ParseObjectName(BUCKET_BRICK_SET, "https://snowbrick.oss-cn-shanghai.aliyuncs.com/set/10264-1/0d1bfc9ac8e89395fb01d1efa5b703b5.jpg")
	if err != nil {
		t.Fatalf("TestParseObjectName: ParseObjectName failed:%v", err)
	}
	if res != "set/10264-1/0d1bfc9ac8e89395fb01d1efa5b703b5.jpg" {
		t.Fatalf("TestParseObjectName: ParseObjectName wrong result")
	}
}
