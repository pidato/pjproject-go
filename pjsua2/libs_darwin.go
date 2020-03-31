// +build darwin

package pjsua2

/*
#cgo CXXFLAGS: -DPJ_AUTOCONF=1 -O3 -DPJ_IS_BIG_ENDIAN=0 -Wno-delete-non-virtual-dtor -Wno-unused-function
#cgo CXXFLAGS: -Wall -fPIC
#cgo CXXFLAGS: -DPJ_IS_LITTLE_ENDIAN=1
#cgo CXXFLAGS: -DPJMEDIA_USE_OLD_FFMPEG=1
#cgo CXXFLAGS: -I./include/mac
#cgo LDFLAGS: -ldl -lm -lpthread
#cgo LDFLAGS: -framework CoreAudio -framework CoreServices -framework AudioUnit -framework AudioToolbox -framework Foundation -framework AppKit -framework AVFoundation -framework CoreGraphics -framework QuartzCore -framework CoreVideo -framework CoreMedia -framework VideoToolbox -framework Security
#cgo LDFLAGS: -L./mac
#cgo LDFLAGS: -lpjproject-2.10
*/
import "C"
