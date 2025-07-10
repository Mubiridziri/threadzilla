package utils

import (
	"os"
	"testing"
)

func TestGetEnvStr_Exists(t *testing.T) {
	os.Setenv("TEST_KEY", "test_value")
	t.Cleanup(func() { os.Unsetenv("TEST_KEY") })
	if val := GetEnvStr("TEST_KEY", "default"); val != "test_value" {
		t.Errorf("expected 'test_value', got %q", val)
	}
}

func TestGetEnvStr_NotExists(t *testing.T) {
	os.Unsetenv("TEST_KEY_NOT_SET")
	if val := GetEnvStr("TEST_KEY_NOT_SET", "default"); val != "default" {
		t.Errorf("expected 'default', got %q", val)
	}
}
