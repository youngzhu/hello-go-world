package quicksort

import "testing"

func TestQuickSort1(t *testing.T) {
	values := []int{5, 4, 2, 3, 1}
	QuickSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5 {
		t.Error("QuickSort() failed. Got", values, "Expected 1 2 3 4 5")
	}

}

func TestQuickSort2(t *testing.T) {
	values := []int{0, 4, 0, 3, 1}
	QuickSort(values)
	if values[0] != 0 || values[1] != 0 || values[2] != 1 || values[3] != 3 || values[4] != 4 {
		t.Error("QuickSort() failed. Got", values, "Expected 0 0 1 3 4")
	}

}

func TestQuickSort3(t *testing.T) {
	values := []int{5}
	QuickSort(values)
	if values[0] != 5 {
		t.Error("QuickSort() failed. Got", values, "Expected 5")
	}

}
