package util_test

import (
	"fmt"
	"github.com/foxcapades/golour/v1/internal/util"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestClampF32(t *testing.T) {
	Convey("ClampF32", t, func() {
		Convey("Given a value less than the minimum", func() {
			Convey("returns the minimum value", func() {
				So(util.ClampF32(10, 12, 13), ShouldEqual, 12)
			})
		})

		Convey("Given a value greater than the maximum", func() {
			Convey("returns the minimum value", func() {
				So(util.ClampF32(15, 12, 13), ShouldEqual, 13)
			})
		})

		Convey("Given a value between the minimum and maximum", func() {
			Convey("returns the minimum value", func() {
				So(util.ClampF32(12, 11, 13), ShouldEqual, 12)
			})
		})
	})
}

func TestPercentStringSize(t *testing.T) {
	Convey("PercentStringSize", t, func() {
		tests := [...]struct{
			i float32
			o int
		} {
			{0, 2},      // 0%
			{1, 4},      // 100%
			{1.1, 4},    // 100%
			{0.1, 3},    // 10%
			{0.01, 2},   // 1%
			{0.001, 4},  // 0.1%
			{0.0001, 2}, // 0%
		}

		for _, test := range tests {
			Convey(fmt.Sprintf("(%f) -> %d", test.i, test.o), func() {
				So(util.PercentStringSize(test.i), ShouldEqual, test.o)
			})
		}
	})
}

func TestTruncateF32(t *testing.T) {
	Convey("TruncateF32", t, func() {
		tests := [...]struct{
			i float32
			p int
			o float32
		}{
			{0.0001, 3, 0},
			{0.1234, 3, 0.123},
		}

		for _, test := range tests {
			Convey(fmt.Sprintf("(%f, %d) -> %f", test.i, test.p, test.o), func() {
				So(util.TruncateF32(test.i, test.p), ShouldEqual, test.o)
			})
		}
	})
}

func TestPrecisionPercentToBytes(t *testing.T) {
	Convey("PercentStringSize", t, func() {
		tests := [...]struct{
			i float32
			o string
		} {
			{0, "0"},
			{1, "100"},
			{1.1, "100"},
			{0.1, "10"},
			{0.01, "1"},
			{0.15, "15"},
			{0.001, "0.1"},
			{0.0001, "0"},
		}

		for _, test := range tests {
			Convey(fmt.Sprintf("(%f, []byte) -> %d", test.i, len(test.o)), func() {
				buf := make([]byte, len(test.o))
				So(util.PrecisionPercentToBytes(test.i, buf), ShouldEqual, len(test.o))
				So(string(buf), ShouldEqual, test.o)
			})
		}
	})
}
