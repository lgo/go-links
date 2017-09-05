package backend

import (
	"errors"
	"strings"
	"testing"
)

func TestDictBackendKeyStorage(t *testing.T) {
	backend := Dict{}
	backend.Start()

	if err := backend.Store("key", "val"); err != nil {
		t.Errorf("Store(\"key\") got unexpected error: %s", err)
	}

	if val, err := backend.Get("key"); err != nil {
		t.Errorf("Get(\"key\") got unexpected error: %s", err)
	} else if val != "val" {
		t.Errorf("Get(\"key\") expected:%q actual:%q", "val", val)
	}

	if val, err := backend.Get("notkey"); err == nil {
		t.Errorf("Get(\"notkey\") expected error:%q instead got value:%q", errors.New("no key found"), val)
	} else if !strings.Contains(err.Error(), "no key found") {
		t.Errorf("Get(\"notkey\") expected error:%q actual:%q", "no key found", err.Error())
	}

	if err := backend.Store("key", "newval"); err != nil {
		t.Errorf("Store(\"key\") unexpected expected error:%s", err)
	}

	if err := backend.Store("key", "newval"); err != nil {
		t.Errorf("Store(\"key\") unexpected expected error:%s", err)
	}

	if val, err := backend.Get("key"); err != nil {
		t.Errorf("Get(\"key\") got unexpected error: %s", err)
	} else if val != "newval" {
		t.Errorf("Get(\"key\") expected:%q actual:%q", "newval", val)
	}

	if val := backend.Delete("key"); val != true {
		t.Errorf("Delete(\"key\") expected success")
	}

	if val := backend.Delete("key"); val != false {
		t.Errorf("Delete(\"key\") expected failure")
	}

	if val := backend.Delete("notkey"); val != false {
		t.Errorf("Delete(\"key\") expected failure")
	}
}
