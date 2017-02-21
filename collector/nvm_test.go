package collector

import (
	"testing"
	"os"
)

func TestGetLogPage(t *testing.T) {
	f, err := os.Open("/dev/nvme0n1")
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	lids := []uint8{2, 221, 202, 197, 193, 194}
	for i := 0;i < len(lids);i ++ {
		buf := make([]byte, 256)
		err = GetLogPage(f.Fd(), lids[i], buf)
		if err != nil {
			t.Logf("lid: %d error : %s", lids[i], err.Error())
		} else {
			t.Logf("lid: %d bytes: %v", lids[i], buf)
		}

	}
}
