package file_manager

import (
	"encoding/base64"
	"os"
	"testing"
)

func TestFileManager_SaveBase64File(t *testing.T) {
	fm := FileManager{}
	filepath := "testfile.txt"
	defer os.Remove(filepath)

	content := "Hello, world!"
	encoded := base64.StdEncoding.EncodeToString([]byte(content))

	err := fm.SaveBase64File(filepath, encoded)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	data, err := os.ReadFile(filepath)
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}
	if string(data) != content {
		t.Errorf("expected file content %q, got %q", content, string(data))
	}
}

func TestFileManager_SaveBase64File_InvalidBase64(t *testing.T) {
	fm := FileManager{}
	filepath := "testfile_invalid.txt"
	defer os.Remove(filepath)

	invalid := "not_base64!@#"
	err := fm.SaveBase64File(filepath, invalid)
	if err == nil {
		t.Error("expected error for invalid base64, got nil")
	}
}

func TestFileManager_DeleteFile(t *testing.T) {
	fm := FileManager{}
	filepath := "testfile_delete.txt"
	content := []byte("delete me")
	if err := os.WriteFile(filepath, content, 0644); err != nil {
		t.Fatalf("setup failed: %v", err)
	}

	err := fm.DeleteFile(filepath)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if _, err := os.Stat(filepath); !os.IsNotExist(err) {
		t.Errorf("expected file to be deleted, but it exists")
	}
}

func TestFileManager_DeleteFile_NonExistent(t *testing.T) {
	fm := FileManager{}
	filepath := "does_not_exist.txt"

	err := fm.DeleteFile(filepath)
	if err == nil {
		t.Error("expected error when deleting non-existent file, got nil")
	}
}
