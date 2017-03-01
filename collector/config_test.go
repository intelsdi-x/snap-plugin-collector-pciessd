package collector

import (
	"testing"
)

func TestReadConfig(t *testing.T) {
	ret, err := ReadConfig("../conf/config.json")
	if err != nil {
		t.Fatalf("%v", err)
	}
	t.Logf("%v", ret)
}
