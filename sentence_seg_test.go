package senseg

import (
	"testing"
)

func TestReNewLine(t *testing.T) {
	cases := []string{"case 1, ____**** hello this is now ====", "haha\nthis is cool"}
	for _, s := range cases {
		ret := ReNewLine.FindAllStringSubmatch(s, -1)
		t.Log(ret)
	}
}
