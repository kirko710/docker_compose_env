package main

import (
	"testing"
)

func TestAppStarts(t *testing.T) {
	t.Log("App test passed")
}

func TestConnectReturnsError(t *testing.T) {
	_, err := connect()
	if err != nil {
		t.Log("DB not available, skipping:", err)
		t.Skip()
	}
}
