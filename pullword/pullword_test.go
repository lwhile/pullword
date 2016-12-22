package pullword

import (
	"testing"
)

func TestPullword(t *testing.T) {
	request := NewRequest("马化腾是马云最大的威胁", false, true)
	res := request.ResM()
	if res == "" {
		t.Errorf("NULL")
	}
}
