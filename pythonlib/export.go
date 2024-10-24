// File: pythonlib/export.go
package main

/*
#include <stdlib.h>
*/
import "C"

import (
    "github.com/open-and-sustainable/prismaid" // Adjust the import path to match your module name
)

//export RunReviewPython
func RunReviewPython() {
    prismaid.RunReview()
}

func main() {}
