package main

import (
	"os"
	"testing"
)

func TestFileStore_PutValue(t *testing.T) {
	testStoreDir := "./test-dir"
	os.Mkdir(testStoreDir, 0755)

	t.Run("storer_positive_1", func(t *testing.T) {
		fs := NewFileStore(testStoreDir)
		if err := fs.PutValue("test1", "testvalue1"); err != nil {
			t.Errorf("FileStore.PutValue() error = %v", err)
		}
		value, err := fs.GetValue("test1")
		if err != nil {
			t.Errorf("FileStore.PutValue() error = %v", err)
		}
		if value != "testvalue1" {
			t.Errorf("values are not equal")
		}
	})
	t.Run("storer_negative_1", func(t *testing.T) {
		fs := NewFileStore(testStoreDir)
		if err := fs.PutValue("test2", "testvalue2"); err != nil {
			t.Errorf("FileStore.PutValue() error = %v", err)
		}
		os.WriteFile(testStoreDir+"/test2", []byte("wrongtestvalue2"), 0755)
		value, err := fs.GetValue("test2")
		if err != nil {
			t.Errorf("FileStore.PutValue() error = %v", err)
		}
		if value == "testvalue2" {
			t.Errorf("values are equal, but they should not")
		}
	})
	t.Run("storer_negative_2", func(t *testing.T) {
		fs := NewFileStore("testbaddir")
		if err := fs.PutValue("test2", "testvalue2"); err == nil {
			t.Errorf("shouldn't put file to the directory which is not exist")
		}
	})
	t.Run("storer_negative_3", func(t *testing.T) {
		fs := NewFileStore(testStoreDir)
		_, err := fs.GetValue("badtest3")
		if err == nil {
			t.Errorf("shouldn't get from the file twhich is not exist")
		}
	})

}
