package godynlib

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include <stdlib.h>
import "C"
import "unsafe"
import "errors"

const (
	RTLD_LAZY     = C.RTLD_LAZY
	RTLD_NOW      = C.RTLD_NOW
	RTLD_GLOBAL   = C.RTLD_GLOBAL
	RTLD_LOCAL    = C.RTLD_LOCAL
	RTLD_NODELETE = C.RTLD_NODELETE
	RTLD_NOLOAD   = C.RTLD_NOLOAD
	RTLD_DEEPBIND = C.RTLD_DEEPBIND
)

func dlopen(filename string, flag int) (uintptr, error) {
	Cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(Cfilename))
	Cflag := C.int(flag)

	Chandle, _ := C.dlopen(Cfilename, Cflag)
	if Chandle == nil {
		// error happened
		CErrString := C.dlerror()
		return 0, errors.New(C.GoString(CErrString))
	} else {
		return uintptr(Chandle), nil
	}
}

func dlsym(handle uintptr, symbol string) (uintptr, error) {
	Csymbol := C.CString(symbol)
	defer C.free(unsafe.Pointer(Csymbol))
	Chandle := unsafe.Pointer(handle)

	// First clean preview error
	_, _ = C.dlerror()

	// Then call dlsym
	CSymbolHandle, _ := C.dlsym(Chandle, Csymbol)

	// Test error now
	CErrString, _ := C.dlerror()

	if CErrString == nil {
		return uintptr(CSymbolHandle), nil
	} else {
		return 0, errors.New(C.GoString(CErrString))
	}
}

func dlclose(handle uintptr) error {
	Chandle := unsafe.Pointer(handle)
	errno, _ := C.dlclose(Chandle)
	if errno != 0 {
		CErrString, _ := C.dlerror()
		return errors.New(C.GoString(CErrString))
	} else {
		return nil
	}
}
