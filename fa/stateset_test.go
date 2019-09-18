package fa

import "testing"

func TestStateSet_Equal(t *testing.T) {
	var ss StateSet = make(map[int]bool)
	if !ss.Equal(ss) {
		t.Fatalf("ss.Equal() expected %t. got %t", true, false)
	}

	var s1 StateSet = map[int]bool{
		1: true,
		2: true,
	}
	var s2 StateSet = map[int]bool{
		2: true,
		1: true,
	}
	if !s2.Equal(s1) || !s1.Equal(s2) {
		t.Fatalf("ss.Equal() expected %t. got %t", true, false)
	}

	var s3 StateSet = map[int]bool{
		1: true,
		2: true,
		3: true,
	}
	if s1.Equal(s3) || s3.Equal(s1) || s2.Equal(s3) || s3.Equal(s2) {
		t.Fatalf("ss.Equal() expected %t. got %t", false, true)
	}
}

func TestStateSet_Add(t *testing.T) {
	var ss StateSet = make(map[int]bool)
	ss.Add(NewSet(1))
	if !ss.Equal(NewSet(1)) {
		t.Fatalf("ss.Add() expected %v. got %v", NewSet(1), ss)
	}

	ss.Add(NewSet(2, 3))
	if !ss.Equal(NewSet(1, 2, 3)) {
		t.Fatalf("ss.Add() expected %v. got %v", NewSet(1, 2, 3), ss)
	}

	ss.Add(NewSet(2, 3, 4))
	if !ss.Equal(NewSet(1, 2, 3, 4)) {
		t.Fatalf("ss.Add() expected %v. got %v", NewSet(1, 2, 3, 4), ss)
	}
}

func TestStateSet_LE(t *testing.T) {
	var ss StateSet = make(map[int]bool)
	if !ss.LE(ss) {
		t.Fatalf("ss.LE() expected %t. got %t", true, false)
	}

	var s1 StateSet = map[int]bool{
		1: true,
		2: true,
	}
	var s2 StateSet = map[int]bool{
		2: true,
		1: true,
	}
	if !s2.LE(s1) || !s1.LE(s2) {
		t.Fatalf("ss.LE() expected %t. got %t", true, false)
	}

	var s3 StateSet = map[int]bool{
		1: true,
		2: true,
		3: true,
	}
	if !s1.LE(s3) || s3.LE(s1) || !s2.LE(s3) || s3.LE(s2) {
		t.Fatalf("ss.LE() expected %t. got %t", false, true)
	}
}
