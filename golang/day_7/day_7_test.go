package main

import "testing"

func TestTotalDirs(t *testing.T) {

	rootDir := &Dir{
		size: 10000,
	}

	subDir := &Dir{
		parent: rootDir,
		size:   90000,
	}

	subDir2 := &Dir{
		parent: rootDir,
		size:   500000,
	}

	rootDir.subDirs = append(rootDir.subDirs, subDir)
	rootDir.subDirs = append(rootDir.subDirs, subDir2)

	want := 100000

	total := totalDirs(0, *rootDir)

	if total != want {
		t.Fatalf("wanted : %v, got: %v", want, total)
	}
}
