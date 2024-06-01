package windows

import (
	"unicode/utf16"
	"unsafe"
)

func GoWString(s *uint16) string {
	if s == nil {
		return ""
	}
	p := (*[1<<30 - 1]uint16)(unsafe.Pointer(s))
	sz := 0
	for p[sz] != 0 {
		sz++
	}
	return string(utf16.Decode(p[:sz:sz]))
}
