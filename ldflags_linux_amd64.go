//go:build linux && amd64

package asposepdf

//#cgo LDFLAGS: -L${SRCDIR}/lib/ -lAsposePDFforGo_linux_amd64
//#cgo LDFLAGS: -Wl,-rpath,$ORIGIN/
//#cgo LDFLAGS: -Wl,-rpath,$ORIGIN/lib/
//#cgo LDFLAGS: -Wl,-rpath,$ORIGIN/asposepdf/lib/
//#cgo LDFLAGS: -Wl,-rpath,$ORIGIN/../asposepdf/lib/
//#cgo LDFLAGS: -Wl,-rpath,${SRCDIR}/
//#cgo LDFLAGS: -Wl,-rpath,${SRCDIR}/lib/
import "C"
