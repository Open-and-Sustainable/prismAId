// File: pythonlib/export.go
package main

/*
#include <stdlib.h>
*/
import "C"

import (
    "github.com/open-and-sustainable/prismaid"
    "unsafe"
)

//export RunReviewPython
func RunReviewPython(input *C.char) *C.char {
    // Convert C string to Go string
    goInput := C.GoString(input)

    // Call your Go function with the input
    err := prismaid.RunReview(goInput)
    if err != nil {
        // Return the error message as a C string
        return C.CString(err.Error())
    }

    // No error, return NULL
    return nil
}

//export FreeCString
func FreeCString(str *C.char) {
    C.free(unsafe.Pointer(str))
}

func main() {}
