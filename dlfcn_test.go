package godynlib

import "testing"

func TestDlfcn(t *testing.T) {
	handle, err := dlopen("./libtest.so", RTLD_NOW)
	if handle == 0 {
		t.Fatalf("dlopen return NULL: %v", err)
	}

	symbolHandle, err := dlsym(handle, "test")
	if symbolHandle == 0 {
		t.Fatalf("dlsym return NULL: %v", err)
	}

	err = dlclose(handle)
	if err != nil {
		t.Fatalf("dlclose failed: %v", err)
	}
}
