//go:build darwin && amd64

package asposepdf

//#cgo LDFLAGS: -L${SRCDIR}/lib/ -lAsposePDFforGo_darwin_amd64
//#cgo LDFLAGS: -Wl,-rpath,./
//#cgo LDFLAGS: -Wl,-rpath,../
//#cgo LDFLAGS: -Wl,-rpath,./lib/
//#cgo LDFLAGS: -Wl,-rpath,../lib/
//#cgo LDFLAGS: -Wl,-rpath,./asposepdf/lib/
//#cgo LDFLAGS: -Wl,-rpath,../asposepdf/lib/
//#cgo LDFLAGS: -Wl,-rpath,${SRCDIR}/
//#cgo LDFLAGS: -Wl,-rpath,${SRCDIR}/lib/
import "C"
