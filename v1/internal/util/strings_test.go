package util_test

import (
	"fmt"
	"github.com/foxcapades/golour/v1/internal/util"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)


func TestPercentBytes(t *testing.T) {
	Convey("PercentBytes", t, func() {
		tests := [...]struct{
			i uint8
			o string
		} {
			{0, "0"},
			{2, "0.02"},
			{4, "0.04"},
			{8, "0.08"},
			{16, "0.16"},
			{32, "0.32"},
			{64, "0.64"},
			{100, "1"},
		}

		for _, test := range tests {
			Convey(fmt.Sprintf("(%d, []byte) -> %d", test.i, len(test.o)), func() {
				buf := make([]byte, len(test.o))
				So(util.PercentBytes(test.i, buf), ShouldEqual, len(test.o))
				So(string(buf), ShouldEqual, test.o)
			})
		}
	})
}