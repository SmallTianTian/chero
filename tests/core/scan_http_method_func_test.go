package core

import (
	"testing"

	"github.com/SmallTianTian/chero/core"
)

func TestScanHttpMethodFunc(t *testing.T) {
	exceptMethod := []string{
		"GetU0", "PostU1", "PutU2", "PatchU3", "DeleteU4", "OptionsU5",
		"CommentGet", "CommentPost", "CommentPut", "CommentPatch", "CommentDelete", "CommentOptions",
		"GetChild",
	}
	result := core.ScanHttpMethodFunc(".")

	emM := make(map[string]struct{}, len(exceptMethod))
	for _, v := range exceptMethod {
		emM[v] = struct{}{}
	}

	for _, v := range result {
		if _, ok := emM[v.F.Name.Name]; ok {
			delete(emM, v.F.Name.Name)
		} else {
			t.Errorf("Method(%s) not in exceptMethod.", v.F.Name.Name)
		}
	}
	for k := range emM {
		t.Errorf("Method(%s) not be scan.", k)
	}
}
