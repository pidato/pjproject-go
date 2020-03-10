// +build darwin2

package pjsua2

/*
#cgo CXXFLAGS: -I/usr/local/include -I/usr/include -g -O2 -Wno-delete-non-virtual-dtor -DPJ_AUTOCONF=1 -O2 -DPJ_IS_BIG_ENDIAN=0 -DPJ_IS_LITTLE_ENDIAN=1
#cgo CXXFLAGS: -I/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/include/c++/v1
#cgo LDFLAGS: -L/usr/lib -L/usr/local/lib
#cgo LDFLAGS: -lstdc++
#cgo LDFLAGS: -lssl
#cgo LDFLAGS: -lcrypto
#cgo LDFLAGS: -lfvad
#cgo LDFLAGS: -lopus
#cgo LDFLAGS: -lopencore-amrnb
#cgo LDFLAGS: -lpj-x86_64-apple-darwin19.3.0
#cgo LDFLAGS: -lpjsip-ua-x86_64-apple-darwin19.3.0
#cgo LDFLAGS: -lpjsip-simple-x86_64-apple-darwin19.3.0
#cgo LDFLAGS: -lpjsip-x86_64-apple-darwin19.3.0
#cgo LDFLAGS: -lpjlib-util-x86_64-apple-darwin19.3.0
#cgo LDFLAGS: -lpjnath-x86_64-apple-darwin19.3.0
#cgo LDFLAGS: -lpjmedia-x86_64-apple-darwin19.3.0
#cgo LDFLAGS: -lpjmedia-codec-x86_64-apple-darwin19.3.0
#cgo LDFLAGS: -lpjmedia-videodev-x86_64-apple-darwin19.3.0
#cgo LDFLAGS: -lpjmedia-audiodev-x86_64-apple-darwin19.3.0
#cgo LDFLAGS: -lsrtp-x86_64-apple-darwin19.3.0
#cgo LDFLAGS: -lresample-x86_64-apple-darwin19.3.0
#cgo LDFLAGS: -lgsmcodec-x86_64-apple-darwin19.3.0
#cgo LDFLAGS: -lspeex-x86_64-apple-darwin19.3.0
#cgo LDFLAGS: -lilbccodec-x86_64-apple-darwin19.3.0
#cgo LDFLAGS: -lg7221codec-x86_64-apple-darwin19.3.0
#cgo LDFLAGS: -lwebrtc-x86_64-apple-darwin19.3.0
#cgo LDFLAGS: -lpjsua-x86_64-apple-darwin19.3.0
#cgo LDFLAGS: -lpjsua2-x86_64-apple-darwin19.3.0
#cgo LDFLAGS: -framework CoreAudio -framework CoreServices -framework AudioUnit -framework AudioToolbox -framework Foundation -framework AppKit -framework AVFoundation -framework CoreGraphics -framework QuartzCore -framework CoreVideo -framework CoreMedia -framework VideoToolbox -framework Security
*/
import "C"
