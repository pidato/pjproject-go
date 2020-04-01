// +build linux

package pjsua2

/*
#cgo CXXFLAGS: -DPJ_AUTOCONF=1 -O3 -Wno-delete-non-virtual-dtor -Wno-unused-function
#cgo CXXFLAGS: -Wall -fPIC -fno-strict-aliasing -Wno-maybe-uninitialized
#cgo CXXFLAGS: -DPJ_IS_BIG_ENDIAN=0 -DPJ_IS_LITTLE_ENDIAN=1
#cgo CXXFLAGS: -DPJMEDIA_USE_OLD_FFMPEG=1
#cgo CXXFLAGS: -I./include/linux
#cgo LDFLAGS: -ldl -luuid  -lm -lrt -lpthread
#cgo LDFLAGS: -L./linux
#cgo LDFLAGS: -lpjproject-2.10
*/
import "C"
