package main

/*
#include "wrapper.h"
*/
import "C"

func main() {
	var foo C.SourceType
	_ = foo
}
