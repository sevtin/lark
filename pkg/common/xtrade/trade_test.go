package xtrade

import (
	"fmt"
	"testing"
)

func TestGenerateSn(t *testing.T) {
	sn := GenerateSn("LK")
	fmt.Println(sn)
}
