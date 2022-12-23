package main

import "testing"

func TestCountVisibleFromLeft(t *testing.T) {

	treeMap := &TreeMap{
		treeMap: [][]int{{3, 0, 2, 7, 3}},
	}

	t.Run("all visible", func(t *testing.T) {

		count := treeMap.countVisibleFromLeft(3, 0)
		want := 3
		if count != want {
			t.Fatalf("incorrect count visible. want: %v, got: %v", want, count)
		}
	})

	t.Run("edge tree", func(t *testing.T) {
		count := treeMap.countVisibleFromLeft(0, 0)
		want := 0
		if count != want {
			t.Fatalf("incorrect count visible. want: %v, got: %v", want, count)
		}
	})

	t.Run("equal height", func(t *testing.T) {
		treeMap2 := &TreeMap{
			treeMap: [][]int{{3, 7, 3, 7, 3}},
		}

		count := treeMap2.countVisibleFromLeft(3, 0)
		want := 2
		if count != want {
			t.Fatalf("incorrect count visible. want: %v, got: %v", want, count)
		}
	})

}

func TestCountVisibleFromRight(t *testing.T) {

	treeMap := &TreeMap{
		treeMap: [][]int{{3, 9, 2, 7, 3}},
	}

	t.Run("all visible", func(t *testing.T) {

		count := treeMap.countVisibleFromRight(1, 0)
		want := 3
		if count != want {
			t.Fatalf("incorrect count visible. want: %v, got: %v", want, count)
		}
	})

	t.Run("invisible", func(t *testing.T) {

		count := treeMap.countVisibleFromRight(2, 0)
		want := 1
		if count != want {
			t.Fatalf("incorrect count visible. want: %v, got: %v", want, count)
		}
	})

	t.Run("edge", func(t *testing.T) {

		count := treeMap.countVisibleFromRight(4, 0)
		want := 0
		if count != want {
			t.Fatalf("incorrect count visible. want: %v, got: %v", want, count)
		}
	})
}

func TestCountVisibleFromTop(t *testing.T) {

	treeMap := &TreeMap{
		treeMap: [][]int{
			{3, 0, 3, 7, 3},
			{2, 5, 5, 1, 2},
			{3, 9, 9, 6, 3},
		},
	}

	t.Run("example", func(t *testing.T) {

		count := treeMap.countVisibleFromTop(2, 1)
		want := 1
		if count != want {
			t.Fatalf("incorrect count visible. want: %v, got: %v", want, count)
		}
	})

	t.Run("edge", func(t *testing.T) {

		count := treeMap.countVisibleFromTop(2, 0)
		want := 0
		if count != want {
			t.Fatalf("incorrect count visible. want: %v, got: %v", want, count)
		}
	})

	t.Run("all visible", func(t *testing.T) {

		count := treeMap.countVisibleFromTop(2, 2)
		want := 2
		if count != want {
			t.Fatalf("incorrect count visible. want: %v, got: %v", want, count)
		}
	})
}

func TestCountIsVisibleFromBottom(t *testing.T) {

	treeMap := &TreeMap{
		treeMap: [][]int{
			{3, 9, 2, 7, 3},
			{3, 9, 2, 7, 3},
			{3, 9, 9, 6, 3},
		},
	}

	t.Run("edge", func(t *testing.T) {

		count := treeMap.countVisibleFromBottom(1, 3)
		want := 0
		if count != want {
			t.Fatalf("incorrect count visible. want: %v, got: %v", want, count)
		}
	})

	t.Run("blocked", func(t *testing.T) {

		count := treeMap.countVisibleFromBottom(2, 1)
		want := 1
		if count != want {
			t.Fatalf("incorrect count visible. want: %v, got: %v", want, count)
		}
	})
}
