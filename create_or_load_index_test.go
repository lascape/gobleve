package main

import (
	"testing"
)

func TestCreateOrLoadIndex(t *testing.T) {
	index, err := CreateOrLoadIndex()
	if err != nil {
		t.Error("CreateOrLoadIndex Failure:", err)
		return
	}
	t.Log("CreateOrLoadIndex Success:", index.Name())
}
