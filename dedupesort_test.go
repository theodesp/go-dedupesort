package go_dedupesort_test

import (
	. "github.com/smartystreets/goconvey/convey"
	utils "github.com/theodesp/go-dedupesort"
	"testing"
)

func TestSpec(t *testing.T) {
	Convey("Given an array of integers in random order", t, func() {
		var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 42, 7586, -5467984, 7586}

		Convey("it sorts the array in a ascending order", func() {
			So(utils.DedupeSort(ints[0:]), ShouldResemble,
				[]int{-5467984, -784, 0, 42, 59, 74, 238, 905, 959, 7586, 9845})
		})
	})

	Convey("Given an array of integers that contain duplicates", t, func() {
		var ints = [...]int{11, 22, 11, 0, 0, 2, 2, 1, 1, 7}

		Convey("it removes any duplicates", func() {
			So(utils.DedupeSort(ints[0:]), ShouldResemble, []int{0, 1, 2, 7, 11, 22})
		})
	})

}
