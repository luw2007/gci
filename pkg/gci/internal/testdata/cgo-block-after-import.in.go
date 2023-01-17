package main

import (
	"fmt"

	"github.com/luw2007/gci"
	g "github.com/golang"
)

// #cgo CFLAGS: -DPNG_DEBUG=1
// #cgo amd64 386 CFLAGS: -DX86=1
// #cgo LDFLAGS: -lpng
// #include <png.h>
import "C"
