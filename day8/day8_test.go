package main

import "testing"

func TestIsVisibleFromLeft(t *testing.T) {

	treeMap := &TreeMap{
		treeMap: [][]int{{3, 0, 2, 7, 3}},
	}

	t.Run("visible", func(t *testing.T) {

		result := treeMap.isVisibleFromLeft(3, 0)
		if !result {
			t.Fatalf("tree marked invisible")
		}
	})

	t.Run("invisible", func(t *testing.T) {

		result := treeMap.isVisibleFromLeft(2, 0)
		if result {
			t.Fatalf("tree marked visible")
		}
	})

	t.Run("equal height", func(t *testing.T) {
		treeMap2 := &TreeMap{
			treeMap: [][]int{{3, 3, 3, 7, 3}},
		}

		result := treeMap2.isVisibleFromLeft(2, 0)
		if result {
			t.Fatalf("tree marked visible")
		}

	})

}

func TestIsVisibleFromRight(t *testing.T) {

	treeMap := &TreeMap{
		treeMap: [][]int{{3, 9, 2, 7, 3}},
	}

	t.Run("visible", func(t *testing.T) {

		result := treeMap.isVisibleFromRight(1, 0)
		if !result {
			t.Fatalf("tree marked invisible")
		}
	})

	t.Run("invisible", func(t *testing.T) {

		result := treeMap.isVisibleFromRight(2, 0)
		if result {
			t.Fatalf("tree marked visible")
		}
	})
}

func TestIsVisibleFromTop(t *testing.T) {

	treeMap := &TreeMap{
		treeMap: [][]int{
			{3, 9, 2, 7, 3},
			{3, 9, 2, 7, 3},
			{3, 9, 9, 6, 3},
		},
	}

	t.Run("visible", func(t *testing.T) {

		result := treeMap.isVisibleFromTop(2, 2)
		if !result {
			t.Fatalf("tree marked invisible")
		}
	})

	t.Run("invisible", func(t *testing.T) {

		result := treeMap.isVisibleFromTop(3, 2)
		if result {
			t.Fatalf("tree marked visible")
		}
	})
}

func TestIsVisibleFromBottom(t *testing.T) {

	treeMap := &TreeMap{
		treeMap: [][]int{
			{3, 9, 2, 7, 3},
			{3, 9, 2, 7, 3},
			{3, 9, 9, 6, 3},
		},
	}

	t.Run("visible", func(t *testing.T) {

		result := treeMap.isVisibleFromBottom(1, 3)
		if !result {
			t.Fatalf("tree marked invisible")
		}
	})

	t.Run("invisible", func(t *testing.T) {

		result := treeMap.isVisibleFromBottom(2, 1)
		if result {
			t.Fatalf("tree marked visible")
		}
	})
}
