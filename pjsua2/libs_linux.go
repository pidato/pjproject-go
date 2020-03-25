// +build linux

package pjsua2

/*
#cgo CXXFLAGS: -DPJ_AUTOCONF=1 -O2 -DPJ_IS_BIG_ENDIAN=0 -Wno-delete-non-virtual-dtor -Wunused-function
#cgo CXXFLAGS: -Wall -fPIC
#cgo CXXFLAGS: -DPJ_IS_LITTLE_ENDIAN=1
#cgo CXXFLAGS: -DPJMEDIA_USE_OLD_FFMPEG=1
#cgo CXXFLAGS: -I./include/linux
#cgo LDFLAGS: -ldl -luuid -lm -lrt -lpthread -lasound
#cgo LDFLAGS: -L./linux
#cgo LDFLAGS: -lpjproject-2.10
*/
import "C"
