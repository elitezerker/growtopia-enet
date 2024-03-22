package enet

// #cgo windows CFLAGS: -Ienet/include/
// #cgo windows LDFLAGS: -lWs2_32 -lWinmm
// #include <enet/enet.h>
import "C"
import "fmt"

func Initialize() int {
	return int(C.enet_initialize())
}

func Deinitialize() {
	C.enet_deinitialize()
}

func LinkedVersion() string {
	var version = uint32(C.enet_linked_version())
	major := uint8(version >> 16)
	minor := uint8(version >> 8)
	patch := uint8(version)
	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}
