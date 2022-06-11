package main

import (
	"testing"
)

func TestWriteIndex(t *testing.T) {
	index, err := CreateOrLoadIndex()
	if err != nil {
		t.Errorf("CreateOrLoadIndex Failure: %v", err)
		return
	}
	err = WriteIndex(index)
	if err != nil {
		t.Errorf("WriteIndex Failure: %v", err)
		return
	}
	t.Log(index.DocCount())
}

func TestWriteBatchIndex(t *testing.T) {
	index, err := CreateOrLoadIndex()
	if err != nil {
		t.Errorf("CreateOrLoadIndex Failure: %v", err)
		return
	}
	err = WriteBatchIndex(index)
	if err != nil {
		t.Errorf("WriteIndex Failure: %v", err)
		return
	}
	t.Log(index.DocCount())
}
