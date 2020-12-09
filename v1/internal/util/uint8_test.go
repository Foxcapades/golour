package util_test

import (
	"fmt"
	"github.com/foxcapades/golour/v1/internal/util"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDigitToU8(t *testing.T) {
	Convey("DigitToU8", t, func() {
		tests := [...]struct{
			i byte
			o uint8
			e error
		} {
			{i: '0', o: 0},
			{i: '1', o: 1},
			{i: '2', o: 2},
			{i: '3', o: 3},
			{i: '4', o: 4},
			{i: '5', o: 5},
			{i: '6', o: 6},
			{i: '7', o: 7},
			{i: '8', o: 8},
			{i: '9', o: 9},
			{i: 'A', o: 10},
			{i: 'a', o: 10},
			{i: 'B', o: 11},
			{i: 'b', o: 11},
			{i: 'C', o: 12},
			{i: 'c', o: 12},
			{i: 'D', o: 13},
			{i: 'd', o: 13},
			{i: 'E', o: 14},
			{i: 'e', o: 14},
			{i: 'F', o: 15},
			{i: 'f', o: 15},
			{i: 'g', o: 0, e: util.ErrInvalidUint8Fmt},
		}

		for _, test := range tests {
			Convey(fmt.Sprintf("(%c) -> (%d, %s)", test.i, test.o, test.e), func() {
				a, b := util.DigitToU8(test.i)

				So(b, ShouldEqual, test.e)
				So(a, ShouldEqual, test.o)
			})
		}
	})
}