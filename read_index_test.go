package main

import (
	"testing"
)

func TestReadIndex(t *testing.T) {
	index, err := CreateOrLoadIndex()
	if err != nil {
		t.Errorf("CreateOrLoadIndex Failure: %v", err)
		return
	}
	result, err := ReadIndex(index)
	if err != nil {
		t.Errorf("ReadIndex Failure: %v", err)
		return
	}
	t.Log(result.String())
}

func TestReadHighlightIndex(t *testing.T) {
	index, err := CreateOrLoadIndex()
	if err != nil {
		t.Errorf("CreateOrLoadIndex Failure: %v", err)
		return
	}
	result, err := ReadHighlightIndex(index)
	if err != nil {
		t.Errorf("ReadIndex Failure: %v", err)
		return
	}
	t.Log(result.String())
}
