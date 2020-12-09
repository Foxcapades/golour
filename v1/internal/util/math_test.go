package util_test

import (
	"fmt"
	"testing"

	. "github.com/foxcapades/golour/v1/internal/util"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMinU8(t *testing.T) {
	Convey("MinU8(uint8, uint8) uint8", t, func() {
		So(MinU8(1, 2), ShouldEqual, 1)
		So(MinU8(1, 1), ShouldEqual, 1)
		So(MinU8(2, 1), ShouldEqual, 1)
	})
}

func TestMinI(t *testing.T) {
	Convey("MinI(int, int) int", t, func() {
		So(MinI(1, 2), ShouldEqual, 1)
		So(MinI(1, 1), ShouldEqual, 1)
		So(MinI(2, 1), ShouldEqual, 1)
	})
}

func TestMinF64(t *testing.T) {
	Convey("MinF64(...float64) float64", t, func() {
		So(MinF64(1, 2, 3), ShouldEqual, 1)
		So(MinF64(2, 1, 3), ShouldEqual, 1)
		So(MinF64(2, 3, 1), ShouldEqual, 1)

		So(MinF64(0.1, 0.2, 0.3), ShouldEqual, 0.1)
		So(MinF64(0.2, 0.1, 0.3), ShouldEqual, 0.1)
		So(MinF64(0.2, 0.3, 0.1), ShouldEqual, 0.1)
	})
}

func TestMaxF64(t *testing.T) {
	Convey("MaxF64(...float64) float64", t, func() {
		So(MaxF64(1, 2, 3), ShouldEqual, 3)
		So(MaxF64(2, 1, 3), ShouldEqual, 3)
		So(MaxF64(2, 3, 1), ShouldEqual, 3)

		So(MaxF64(0.1, 0.2, 0.3), ShouldEqual, 0.3)
		So(MaxF64(0.2, 0.1, 0.3), ShouldEqual, 0.3)
		So(MaxF64(0.2, 0.3, 0.1), ShouldEqual, 0.3)
	})
}

func TestU8Pow(t *testing.T) {
	Convey("U8Pow(uint8, uint8) uint", t, func() {
		tests := [...]struct {
			val, pow uint8
			out      uint
		}{
			{2, 0, 1},
			{2, 1, 2},
			{2, 2, 4},
			{2, 3, 8},
			{2, 4, 16},
			{2, 5, 32},
			{2, 6, 64},
			{2, 7, 128},
			{2, 8, 256},
			{2, 9, 512},
			{2, 10, 1_024},
			{2, 11, 2_048},
			{2, 12, 4_096},
			{2, 13, 8_192},
			{2, 14, 16_384},
			{2, 15, 32_768},
			{2, 16, 65_536},
			{10, 0, 1},
			{10, 1, 10},
			{10, 2, 100},
			{10, 3, 1_000},
			{10, 4, 10_000},
			{10, 5, 100_000},
			{10, 6, 1_000_000},
			{10, 7, 10_000_000},
			{10, 8, 100_000_000},
		}
		for i := range tests {
			test := &tests[i]
			Convey(fmt.Sprintf("%d ^ %d = %d", test.val, test.pow, test.out), func() {
				So(U8Pow(test.val, test.pow), ShouldEqual, test.out)
			})
		}

	})
}

func TestIPow(t *testing.T) {
	Convey("IPow(uint8, uint8) uint", t, func() {
		tests := [...]struct {
			val, pow, out int
		}{
			{2, 0, 1},
			{2, 1, 2},
			{2, 2, 4},
			{2, 3, 8},
			{2, 4, 16},
			{2, 5, 32},
			{2, 6, 64},
			{2, 7, 128},
			{2, 8, 256},
			{2, 9, 512},
			{2, 10, 1_024},
			{2, 11, 2_048},
			{2, 12, 4_096},
			{2, 13, 8_192},
			{2, 14, 16_384},
			{2, 15, 32_768},
			{2, 16, 65_536},
			{10, 0, 1},
			{10, 1, 10},
			{10, 2, 100},
			{10, 3, 1_000},
			{10, 4, 10_000},
			{10, 5, 100_000},
			{10, 6, 1_000_000},
			{10, 7, 10_000_000},
			{10, 8, 100_000_000},
		}
		for i := range tests {
			test := &tests[i]
			Convey(fmt.Sprintf("%d ^ %d = %d", test.val, test.pow, test.out), func() {
				So(IPow(test.val, test.pow), ShouldEqual, test.out)
			})
		}

	})
}

func TestBluntRound(t *testing.T) {
	Convey("BluntRound(float64) float64", t, func() {
		So(BluntRound(1.05), ShouldEqual, 1.1)
		So(BluntRound(1.045), ShouldEqual, 1)
	})
}

func ExampleBluntRound() {
	fmt.Printf("%.2f\n", BluntRound(1.005))
	fmt.Printf("%.2f\n", BluntRound(1.05))
	fmt.Printf("%.2f\n", BluntRound(1.5))

	// Output:
	// 1.00
	// 1.10
	// 1.50
}

func TestBluntRoundByte(t *testing.T) {
	Convey("BluntRoundByte64(float64) uint8", t, func() {
		So(BluntRoundByte64(1.0), ShouldEqual, 1)
		So(BluntRoundByte64(1.1), ShouldEqual, 1)
		So(BluntRoundByte64(1.2), ShouldEqual, 1)
		So(BluntRoundByte64(1.3), ShouldEqual, 1)
		So(BluntRoundByte64(1.4), ShouldEqual, 1)
		So(BluntRoundByte64(1.5), ShouldEqual, 2)
		So(BluntRoundByte64(1.6), ShouldEqual, 2)
		So(BluntRoundByte64(1.7), ShouldEqual, 2)
		So(BluntRoundByte64(1.8), ShouldEqual, 2)
		So(BluntRoundByte64(1.9), ShouldEqual, 2)
	})
}

func ExampleBluntRoundByte() {
	fmt.Println(BluntRoundByte64(1.4))
	fmt.Println(BluntRoundByte64(1.5))

	// Output:
	// 1
	// 2
}
