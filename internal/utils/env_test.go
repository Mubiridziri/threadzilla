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

func TestGetEnvBool(t *testing.T) {
	t.Setenv("TEST_BOOL_TRUE", "true")
	if !GetEnvBool("TEST_BOOL_TRUE", false) {
		t.Errorf("expected true for 'true'")
	}

	t.Setenv("TEST_BOOL_FALSE", "false")
	if GetEnvBool("TEST_BOOL_FALSE", true) {
		t.Errorf("expected false for 'false'")
	}

	t.Setenv("TEST_BOOL_ONE", "1")
	if !GetEnvBool("TEST_BOOL_ONE", false) {
		t.Errorf("expected true for '1'")
	}

	t.Setenv("TEST_BOOL_ZERO", "0")
	if GetEnvBool("TEST_BOOL_ZERO", true) {
		t.Errorf("expected false for '0'")
	}

	t.Setenv("TEST_BOOL_INVALID", "notabool")
	if !GetEnvBool("TEST_BOOL_INVALID", true) {
		t.Errorf("expected default true for invalid value")
	}
	if GetEnvBool("TEST_BOOL_INVALID", false) {
		t.Errorf("expected default false for invalid value")
	}

	os.Unsetenv("TEST_BOOL_MISSING")
	if !GetEnvBool("TEST_BOOL_MISSING", true) {
		t.Errorf("expected default true for missing var")
	}
	if GetEnvBool("TEST_BOOL_MISSING", false) {
		t.Errorf("expected default false for missing var")
	}
}
