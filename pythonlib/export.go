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
    output, err := prismaid.RunReview(goInput)
    if err != nil {
        // Handle the error, return an error message prefixed with "error:"
        errorMsg := "error:" + err.Error()
        return C.CString(errorMsg)
    }

    // Return the output as a C string
    return C.CString(output)
}

//export FreeCString
func FreeCString(str *C.char) {
    C.free(unsafe.Pointer(str))
}

func main() {}
