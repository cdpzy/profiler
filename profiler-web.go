package profiler

import (
	"fmt"
	"strings"
)
func info_off_html() ([]byte, error) {
	return []byte{
		0x3c, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0x0a, 0x09, 0x3c, 0x68, 0x65, 0x61,
		0x64, 0x3e, 0x0a, 0x09, 0x09, 0x3c, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e,
		0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x20, 0x74, 0x72, 0x61, 0x63, 0x6b,
		0x69, 0x6e, 0x67, 0x20, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64,
		0x3c, 0x2f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e, 0x0a, 0x09, 0x3c, 0x2f,
		0x68, 0x65, 0x61, 0x64, 0x3e, 0x0a, 0x09, 0x3c, 0x62, 0x6f, 0x64, 0x79,
		0x3e, 0x0a, 0x09, 0x09, 0x3c, 0x70, 0x3e, 0x4d, 0x65, 0x6d, 0x6f, 0x72,
		0x79, 0x20, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x20,
		0x69, 0x73, 0x20, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x6f, 0x66,
		0x66, 0x2e, 0x3c, 0x2f, 0x70, 0x3e, 0x0a, 0x09, 0x09, 0x3c, 0x70, 0x3e,
		0x3c, 0x61, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x2f, 0x70, 0x72,
		0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x74,
		0x22, 0x3e, 0x54, 0x75, 0x72, 0x6e, 0x20, 0x69, 0x74, 0x20, 0x6f, 0x6e,
		0x3c, 0x2f, 0x61, 0x3e, 0x3c, 0x2f, 0x70, 0x3e, 0x0a, 0x09, 0x3c, 0x2f,
		0x62, 0x6f, 0x64, 0x79, 0x3e, 0x0a, 0x3c, 0x2f, 0x68, 0x74, 0x6d, 0x6c,
		0x3e, 0x0a,
	}, nil
}

func info_html() ([]byte, error) {
	return []byte{
		0x3c, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0x0a, 0x20, 0x20, 0x3c, 0x68, 0x65,
		0x61, 0x64, 0x3e, 0x0a, 0x09, 0x3c, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
		0x20, 0x73, 0x72, 0x63, 0x3d, 0x22, 0x2f, 0x2f, 0x61, 0x6a, 0x61, 0x78,
		0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
		0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6a, 0x61, 0x78, 0x2f, 0x6c, 0x69, 0x62,
		0x73, 0x2f, 0x6a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x31, 0x2e, 0x31,
		0x31, 0x2e, 0x31, 0x2f, 0x6a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6d,
		0x69, 0x6e, 0x2e, 0x6a, 0x73, 0x22, 0x3e, 0x3c, 0x2f, 0x73, 0x63, 0x72,
		0x69, 0x70, 0x74, 0x3e, 0x0a, 0x09, 0x3c, 0x21, 0x2d, 0x2d, 0x0a, 0x09,
		0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x20, 0x61, 0x6e, 0x20, 0x47,
		0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x20, 0x63, 0x68, 0x61, 0x72, 0x74, 0x20,
		0x73, 0x68, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x20, 0x61, 0x20, 0x73, 0x65,
		0x72, 0x76, 0x69, 0x63, 0x65, 0x27, 0x73, 0x20, 0x6d, 0x65, 0x6d, 0x6f,
		0x72, 0x79, 0x20, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x0a, 0x09, 0x2d,
		0x2d, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x73, 0x63, 0x72, 0x69,
		0x70, 0x74, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x74, 0x65, 0x78,
		0x74, 0x2f, 0x6a, 0x61, 0x76, 0x61, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
		0x22, 0x20, 0x73, 0x72, 0x63, 0x3d, 0x22, 0x68, 0x74, 0x74, 0x70, 0x73,
		0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
		0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x73, 0x61, 0x70, 0x69, 0x22,
		0x3e, 0x3c, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x3e, 0x0a, 0x20,
		0x20, 0x20, 0x20, 0x3c, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x20, 0x74,
		0x79, 0x70, 0x65, 0x3d, 0x22, 0x74, 0x65, 0x78, 0x74, 0x2f, 0x6a, 0x61,
		0x76, 0x61, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x22, 0x3e, 0x0a, 0x09,
		0x09, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x61, 0x64,
		0x28, 0x22, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74,
		0x69, 0x6f, 0x6e, 0x22, 0x2c, 0x20, 0x22, 0x31, 0x22, 0x2c, 0x20, 0x7b,
		0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x3a, 0x5b, 0x22, 0x63,
		0x6f, 0x72, 0x65, 0x63, 0x68, 0x61, 0x72, 0x74, 0x22, 0x5d, 0x7d, 0x29,
		0x3b, 0x0a, 0x0a, 0x09, 0x09, 0x2f, 0x2f, 0x20, 0x46, 0x65, 0x74, 0x63,
		0x68, 0x20, 0x74, 0x68, 0x65, 0x20, 0x48, 0x65, 0x61, 0x70, 0x20, 0x4d,
		0x65, 0x6d, 0x6f, 0x72, 0x79, 0x20, 0x53, 0x74, 0x61, 0x74, 0x73, 0x20,
		0x4a, 0x53, 0x4f, 0x4e, 0x20, 0x65, 0x76, 0x65, 0x72, 0x79, 0x20, 0x31,
		0x20, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x2c, 0x20, 0x61, 0x6e, 0x64,
		0x20, 0x64, 0x72, 0x61, 0x77, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x65,
		0x20, 0x63, 0x68, 0x61, 0x72, 0x74, 0x0a, 0x09, 0x09, 0x66, 0x75, 0x6e,
		0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x66, 0x65, 0x74, 0x63, 0x68, 0x48,
		0x65, 0x61, 0x70, 0x4d, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x73, 0x28,
		0x29, 0x20, 0x7b, 0x0a, 0x09, 0x09, 0x20, 0x20, 0x20, 0x20, 0x24, 0x2e,
		0x61, 0x6a, 0x61, 0x78, 0x28, 0x7b, 0x0a, 0x09, 0x09, 0x09, 0x09, 0x74,
		0x79, 0x70, 0x65, 0x3a, 0x20, 0x27, 0x47, 0x45, 0x54, 0x27, 0x2c, 0x0a,
		0x09, 0x09, 0x09, 0x09, 0x75, 0x72, 0x6c, 0x3a, 0x20, 0x27, 0x69, 0x6e,
		0x66, 0x6f, 0x27, 0x2c, 0x0a, 0x09, 0x09, 0x09, 0x09, 0x64, 0x61, 0x74,
		0x61, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x20, 0x27, 0x6a, 0x73, 0x6f, 0x6e,
		0x27, 0x2c, 0x0a, 0x09, 0x09, 0x09, 0x09, 0x73, 0x75, 0x63, 0x63, 0x65,
		0x73, 0x73, 0x3a, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
		0x20, 0x28, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x49,
		0x6e, 0x66, 0x6f, 0x29, 0x20, 0x7b, 0x0a, 0x09, 0x09, 0x09, 0x09, 0x09,
		0x2f, 0x2f, 0x20, 0x64, 0x72, 0x61, 0x77, 0x20, 0x74, 0x68, 0x65, 0x20,
		0x63, 0x68, 0x61, 0x72, 0x74, 0x0a, 0x09, 0x09, 0x09, 0x09, 0x09, 0x64,
		0x72, 0x61, 0x77, 0x48, 0x65, 0x61, 0x70, 0x43, 0x68, 0x61, 0x72, 0x74,
		0x28, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x6e,
		0x66, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x29,
		0x3b, 0x0a, 0x0a, 0x09, 0x09, 0x09, 0x09, 0x09, 0x2f, 0x2f, 0x20, 0x64,
		0x72, 0x61, 0x77, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x61, 0x62, 0x6c,
		0x65, 0x20, 0x6f, 0x66, 0x20, 0x65, 0x78, 0x74, 0x72, 0x61, 0x20, 0x69,
		0x6e, 0x66, 0x6f, 0x0a, 0x09, 0x09, 0x09, 0x09, 0x09, 0x76, 0x61, 0x72,
		0x20, 0x65, 0x78, 0x74, 0x72, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x20, 0x3d,
		0x20, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x6e,
		0x66, 0x6f, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x53, 0x65, 0x72, 0x76,
		0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x0a, 0x09, 0x09, 0x09, 0x09,
		0x09, 0x69, 0x66, 0x28, 0x65, 0x78, 0x74, 0x72, 0x61, 0x49, 0x6e, 0x66,
		0x6f, 0x20, 0x21, 0x3d, 0x20, 0x6e, 0x75, 0x6c, 0x6c, 0x29, 0x20, 0x7b,
		0x0a, 0x09, 0x09, 0x09, 0x09, 0x09, 0x09, 0x24, 0x28, 0x27, 0x23, 0x65,
		0x78, 0x74, 0x72, 0x61, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x27, 0x29, 0x2e,
		0x68, 0x74, 0x6d, 0x6c, 0x28, 0x27, 0x27, 0x29, 0x0a, 0x09, 0x09, 0x09,
		0x09, 0x09, 0x09, 0x24, 0x28, 0x27, 0x23, 0x65, 0x78, 0x74, 0x72, 0x61,
		0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x27, 0x29, 0x2e, 0x61, 0x70, 0x70, 0x65,
		0x6e, 0x64, 0x28, 0x27, 0x3c, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x63,
		0x65, 0x6c, 0x6c, 0x73, 0x70, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x3d, 0x22,
		0x31, 0x30, 0x22, 0x20, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3d, 0x22, 0x62,
		0x6f, 0x72, 0x64, 0x65, 0x72, 0x3a, 0x31, 0x70, 0x78, 0x20, 0x73, 0x6f,
		0x6c, 0x69, 0x64, 0x20, 0x23, 0x64, 0x64, 0x64, 0x3b, 0x22, 0x20, 0x63,
		0x65, 0x6c, 0x6c, 0x73, 0x70, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x3d, 0x22,
		0x30, 0x22, 0x20, 0x63, 0x65, 0x6c, 0x6c, 0x70, 0x61, 0x64, 0x64, 0x69,
		0x6e, 0x67, 0x3d, 0x22, 0x32, 0x22, 0x20, 0x77, 0x69, 0x64, 0x74, 0x68,
		0x3d, 0x22, 0x35, 0x30, 0x30, 0x22, 0x3e, 0x27, 0x29, 0x0a, 0x09, 0x09,
		0x09, 0x09, 0x09, 0x09, 0x66, 0x6f, 0x72, 0x20, 0x28, 0x76, 0x61, 0x72,
		0x20, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x4e, 0x61, 0x6d,
		0x65, 0x20, 0x69, 0x6e, 0x20, 0x65, 0x78, 0x74, 0x72, 0x61, 0x49, 0x6e,
		0x66, 0x6f, 0x29, 0x20, 0x7b, 0x0a, 0x09, 0x09, 0x09, 0x09, 0x09, 0x09,
		0x09, 0x69, 0x66, 0x20, 0x28, 0x65, 0x78, 0x74, 0x72, 0x61, 0x49, 0x6e,
		0x66, 0x6f, 0x2e, 0x68, 0x61, 0x73, 0x4f, 0x77, 0x6e, 0x50, 0x72, 0x6f,
		0x70, 0x65, 0x72, 0x74, 0x79, 0x28, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
		0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x29, 0x29, 0x20, 0x7b, 0x0a, 0x09,
		0x09, 0x09, 0x09, 0x09, 0x09, 0x09, 0x09, 0x24, 0x28, 0x27, 0x23, 0x65,
		0x78, 0x74, 0x72, 0x61, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x27, 0x29, 0x2e,
		0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x28, 0x27, 0x3c, 0x74, 0x72, 0x3e,
		0x3c, 0x74, 0x64, 0x20, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3d, 0x22, 0x70,
		0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x20, 0x31, 0x30, 0x70, 0x78,
		0x3b, 0x20, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3a, 0x20, 0x32, 0x35, 0x30,
		0x70, 0x78, 0x3b, 0x20, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2d, 0x6c,
		0x65, 0x66, 0x74, 0x3a, 0x31, 0x70, 0x78, 0x20, 0x73, 0x6f, 0x6c, 0x69,
		0x64, 0x20, 0x23, 0x64, 0x64, 0x64, 0x3b, 0x20, 0x62, 0x6f, 0x72, 0x64,
		0x65, 0x72, 0x2d, 0x72, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x31, 0x70, 0x78,
		0x20, 0x73, 0x6f, 0x6c, 0x69, 0x64, 0x20, 0x23, 0x64, 0x64, 0x64, 0x3b,
		0x20, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2d, 0x62, 0x6f, 0x74, 0x74,
		0x6f, 0x6d, 0x3a, 0x31, 0x70, 0x78, 0x20, 0x73, 0x6f, 0x6c, 0x69, 0x64,
		0x20, 0x23, 0x64, 0x64, 0x64, 0x3b, 0x22, 0x3e, 0x27, 0x20, 0x2b, 0x20,
		0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65,
		0x20, 0x2b, 0x20, 0x27, 0x3c, 0x2f, 0x74, 0x64, 0x3e, 0x3c, 0x74, 0x64,
		0x20, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3d, 0x22, 0x70, 0x61, 0x64, 0x64,
		0x69, 0x6e, 0x67, 0x3a, 0x20, 0x31, 0x30, 0x70, 0x78, 0x3b, 0x20, 0x77,
		0x69, 0x64, 0x74, 0x68, 0x3a, 0x20, 0x32, 0x35, 0x30, 0x70, 0x78, 0x3b,
		0x20, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2d, 0x72, 0x69, 0x67, 0x68,
		0x74, 0x3a, 0x31, 0x70, 0x78, 0x20, 0x73, 0x6f, 0x6c, 0x69, 0x64, 0x20,
		0x23, 0x64, 0x64, 0x64, 0x3b, 0x20, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72,
		0x2d, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x3a, 0x31, 0x70, 0x78, 0x20,
		0x73, 0x6f, 0x6c, 0x69, 0x64, 0x20, 0x23, 0x64, 0x64, 0x64, 0x3b, 0x22,
		0x3e, 0x27, 0x20, 0x2b, 0x20, 0x65, 0x78, 0x74, 0x72, 0x61, 0x49, 0x6e,
		0x66, 0x6f, 0x5b, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x4e,
		0x61, 0x6d, 0x65, 0x5d, 0x20, 0x2b, 0x20, 0x27, 0x3c, 0x2f, 0x74, 0x64,
		0x3e, 0x3c, 0x2f, 0x74, 0x72, 0x3e, 0x27, 0x29, 0x0a, 0x09, 0x09, 0x09,
		0x09, 0x09, 0x09, 0x09, 0x7d, 0x0a, 0x09, 0x09, 0x09, 0x09, 0x09, 0x09,
		0x7d, 0x0a, 0x09, 0x09, 0x09, 0x09, 0x09, 0x09, 0x24, 0x28, 0x27, 0x23,
		0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x27, 0x29,
		0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x28, 0x27, 0x3c, 0x2f, 0x74,
		0x61, 0x62, 0x6c, 0x65, 0x3e, 0x27, 0x29, 0x0a, 0x09, 0x09, 0x09, 0x09,
		0x09, 0x7d, 0x0a, 0x09, 0x09, 0x09, 0x09, 0x7d, 0x2c, 0x0a, 0x09, 0x09,
		0x09, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x20,
		0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x28, 0x64, 0x61,
		0x74, 0x61, 0x29, 0x20, 0x7b, 0x0a, 0x09, 0x09, 0x09, 0x09, 0x09, 0x2f,
		0x2f, 0x20, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x20, 0x74,
		0x68, 0x65, 0x20, 0x6e, 0x65, 0x78, 0x74, 0x20, 0x70, 0x6f, 0x6c, 0x6c,
		0x0a, 0x09, 0x09, 0x09, 0x09, 0x09, 0x73, 0x65, 0x74, 0x54, 0x69, 0x6d,
		0x65, 0x6f, 0x75, 0x74, 0x28, 0x66, 0x65, 0x74, 0x63, 0x68, 0x48, 0x65,
		0x61, 0x70, 0x4d, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x73, 0x2c, 0x20,
		0x31, 0x30, 0x30, 0x30, 0x29, 0x3b, 0x0a, 0x09, 0x09, 0x09, 0x09, 0x7d,
		0x0a, 0x09, 0x09, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x29, 0x3b, 0x0a, 0x09,
		0x09, 0x7d, 0x0a, 0x0a, 0x09, 0x09, 0x2f, 0x2f, 0x20, 0x66, 0x65, 0x74,
		0x63, 0x68, 0x20, 0x74, 0x68, 0x65, 0x20, 0x68, 0x65, 0x61, 0x70, 0x20,
		0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x20, 0x73, 0x74, 0x61, 0x74, 0x73,
		0x20, 0x69, 0x6e, 0x20, 0x31, 0x20, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64,
		0x0a, 0x09, 0x09, 0x73, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75,
		0x74, 0x28, 0x66, 0x65, 0x74, 0x63, 0x68, 0x48, 0x65, 0x61, 0x70, 0x4d,
		0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x73, 0x2c, 0x20, 0x31, 0x30, 0x30,
		0x30, 0x29, 0x3b, 0x0a, 0x0a, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69,
		0x6f, 0x6e, 0x20, 0x64, 0x72, 0x61, 0x77, 0x48, 0x65, 0x61, 0x70, 0x43,
		0x68, 0x61, 0x72, 0x74, 0x28, 0x68, 0x65, 0x61, 0x70, 0x44, 0x61, 0x74,
		0x61, 0x29, 0x20, 0x7b, 0x0a, 0x09, 0x09, 0x2f, 0x2f, 0x20, 0x62, 0x75,
		0x69, 0x6c, 0x64, 0x20, 0x74, 0x68, 0x65, 0x20, 0x63, 0x68, 0x61, 0x72,
		0x74, 0x20, 0x64, 0x61, 0x74, 0x61, 0x0a, 0x09, 0x09, 0x63, 0x68, 0x61,
		0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x20, 0x3d, 0x20, 0x5b, 0x5d, 0x3b,
		0x0a, 0x09, 0x09, 0x76, 0x61, 0x72, 0x20, 0x6e, 0x75, 0x6d, 0x53, 0x74,
		0x61, 0x74, 0x73, 0x20, 0x3d, 0x20, 0x68, 0x65, 0x61, 0x70, 0x44, 0x61,
		0x74, 0x61, 0x2e, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x3b, 0x0a, 0x09,
		0x09, 0x63, 0x68, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x5b, 0x63,
		0x68, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x6c, 0x65, 0x6e,
		0x67, 0x74, 0x68, 0x5d, 0x20, 0x3d, 0x20, 0x5b, 0x27, 0x54, 0x69, 0x6d,
		0x65, 0x27, 0x2c, 0x20, 0x27, 0x53, 0x79, 0x73, 0x20, 0x28, 0x4d, 0x42,
		0x29, 0x27, 0x2c, 0x20, 0x27, 0x48, 0x65, 0x61, 0x70, 0x53, 0x79, 0x73,
		0x20, 0x28, 0x4d, 0x42, 0x29, 0x27, 0x2c, 0x20, 0x27, 0x48, 0x65, 0x61,
		0x70, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x20, 0x28, 0x4d, 0x42, 0x29, 0x27,
		0x2c, 0x20, 0x27, 0x48, 0x65, 0x61, 0x70, 0x49, 0x64, 0x6c, 0x65, 0x20,
		0x28, 0x4d, 0x42, 0x29, 0x27, 0x2c, 0x20, 0x27, 0x48, 0x65, 0x61, 0x70,
		0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x20, 0x28, 0x4d, 0x42,
		0x29, 0x27, 0x5d, 0x3b, 0x0a, 0x09, 0x09, 0x66, 0x6f, 0x72, 0x20, 0x28,
		0x76, 0x61, 0x72, 0x20, 0x69, 0x20, 0x3d, 0x20, 0x30, 0x3b, 0x20, 0x69,
		0x20, 0x3c, 0x20, 0x6e, 0x75, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x73, 0x3b,
		0x20, 0x69, 0x2b, 0x2b, 0x29, 0x20, 0x7b, 0x0a, 0x09, 0x09, 0x09, 0x63,
		0x68, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x5b, 0x63, 0x68, 0x61,
		0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x6c, 0x65, 0x6e, 0x67, 0x74,
		0x68, 0x5d, 0x20, 0x3d, 0x20, 0x5b, 0x68, 0x65, 0x61, 0x70, 0x44, 0x61,
		0x74, 0x61, 0x5b, 0x69, 0x5d, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73,
		0x41, 0x67, 0x6f, 0x2f, 0x31, 0x30, 0x30, 0x30, 0x2c, 0x20, 0x68, 0x65,
		0x61, 0x70, 0x44, 0x61, 0x74, 0x61, 0x5b, 0x69, 0x5d, 0x2e, 0x53, 0x79,
		0x73, 0x4b, 0x62, 0x2f, 0x31, 0x30, 0x30, 0x30, 0x2e, 0x30, 0x2c, 0x20,
		0x68, 0x65, 0x61, 0x70, 0x44, 0x61, 0x74, 0x61, 0x5b, 0x69, 0x5d, 0x2e,
		0x48, 0x65, 0x61, 0x70, 0x53, 0x79, 0x73, 0x4b, 0x62, 0x2f, 0x31, 0x30,
		0x30, 0x30, 0x2e, 0x30, 0x2c, 0x20, 0x68, 0x65, 0x61, 0x70, 0x44, 0x61,
		0x74, 0x61, 0x5b, 0x69, 0x5d, 0x2e, 0x48, 0x65, 0x61, 0x70, 0x41, 0x6c,
		0x6c, 0x6f, 0x63, 0x4b, 0x62, 0x2f, 0x31, 0x30, 0x30, 0x30, 0x2e, 0x30,
		0x2c, 0x20, 0x68, 0x65, 0x61, 0x70, 0x44, 0x61, 0x74, 0x61, 0x5b, 0x69,
		0x5d, 0x2e, 0x48, 0x65, 0x61, 0x70, 0x49, 0x64, 0x6c, 0x65, 0x4b, 0x62,
		0x2f, 0x31, 0x30, 0x30, 0x30, 0x2e, 0x30, 0x2c, 0x20, 0x68, 0x65, 0x61,
		0x70, 0x44, 0x61, 0x74, 0x61, 0x5b, 0x69, 0x5d, 0x2e, 0x48, 0x65, 0x61,
		0x70, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x4b, 0x62, 0x2f,
		0x31, 0x30, 0x30, 0x30, 0x2e, 0x30, 0x5d, 0x3b, 0x0a, 0x09, 0x09, 0x7d,
		0x0a, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x76, 0x61,
		0x72, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x3d, 0x20, 0x67, 0x6f, 0x6f,
		0x67, 0x6c, 0x65, 0x2e, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a,
		0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x61, 0x72, 0x72, 0x61, 0x79, 0x54,
		0x6f, 0x44, 0x61, 0x74, 0x61, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x28, 0x63,
		0x68, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x29, 0x3b, 0x0a, 0x0a,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x76, 0x61, 0x72, 0x20,
		0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x3d, 0x20, 0x7b, 0x0a,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x74, 0x69,
		0x74, 0x6c, 0x65, 0x3a, 0x20, 0x27, 0x48, 0x65, 0x61, 0x70, 0x20, 0x4d,
		0x65, 0x6d, 0x6f, 0x72, 0x79, 0x20, 0x55, 0x73, 0x61, 0x67, 0x65, 0x27,
		0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x63, 0x75, 0x72, 0x76, 0x65, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x20, 0x27,
		0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x27, 0x2c, 0x0a, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x68, 0x41, 0x78,
		0x69, 0x73, 0x3a, 0x20, 0x7b, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3a, 0x20,
		0x27, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x20, 0x66, 0x72, 0x6f,
		0x6d, 0x20, 0x6e, 0x6f, 0x77, 0x27, 0x2c, 0x20, 0x76, 0x69, 0x65, 0x77,
		0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x3a, 0x20, 0x7b, 0x6d, 0x69, 0x6e,
		0x3a, 0x20, 0x2d, 0x36, 0x30, 0x2c, 0x20, 0x6d, 0x61, 0x78, 0x3a, 0x30,
		0x7d, 0x7d, 0x2c, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x76, 0x41, 0x78, 0x69, 0x73, 0x3a, 0x20, 0x7b, 0x74, 0x69,
		0x74, 0x6c, 0x65, 0x3a, 0x20, 0x27, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
		0x20, 0x55, 0x73, 0x61, 0x67, 0x65, 0x27, 0x2c, 0x20, 0x66, 0x6f, 0x72,
		0x6d, 0x61, 0x74, 0x3a, 0x20, 0x22, 0x23, 0x2c, 0x23, 0x23, 0x23, 0x4d,
		0x42, 0x22, 0x2c, 0x20, 0x76, 0x69, 0x65, 0x77, 0x57, 0x69, 0x6e, 0x64,
		0x6f, 0x77, 0x3a, 0x20, 0x7b, 0x6d, 0x69, 0x6e, 0x3a, 0x20, 0x30, 0x7d,
		0x7d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x3b,
		0x0a, 0x0a, 0x09, 0x09, 0x76, 0x61, 0x72, 0x20, 0x63, 0x68, 0x61, 0x72,
		0x74, 0x20, 0x3d, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x67, 0x6f, 0x6f, 0x67,
		0x6c, 0x65, 0x2e, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61,
		0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x43, 0x68, 0x61,
		0x72, 0x74, 0x28, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
		0x67, 0x65, 0x74, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79,
		0x49, 0x64, 0x28, 0x27, 0x63, 0x68, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x69,
		0x76, 0x27, 0x29, 0x29, 0x3b, 0x0a, 0x09, 0x09, 0x63, 0x68, 0x61, 0x72,
		0x74, 0x2e, 0x64, 0x72, 0x61, 0x77, 0x28, 0x64, 0x61, 0x74, 0x61, 0x2c,
		0x20, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x29, 0x3b, 0x0a, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c,
		0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x3e, 0x0a, 0x20, 0x20, 0x3c,
		0x2f, 0x68, 0x65, 0x61, 0x64, 0x3e, 0x0a, 0x20, 0x20, 0x3c, 0x62, 0x6f,
		0x64, 0x79, 0x20, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x3d, 0x22, 0x63, 0x65,
		0x6e, 0x74, 0x65, 0x72, 0x22, 0x20, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3d,
		0x22, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79,
		0x3a, 0x20, 0x41, 0x72, 0x69, 0x61, 0x6c, 0x22, 0x3e, 0x0a, 0x09, 0x3c,
		0x64, 0x69, 0x76, 0x20, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3d, 0x22, 0x77,
		0x69, 0x64, 0x74, 0x68, 0x3d, 0x31, 0x30, 0x30, 0x25, 0x3b, 0x20, 0x74,
		0x65, 0x78, 0x74, 0x2d, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x3a, 0x20, 0x72,
		0x69, 0x67, 0x68, 0x74, 0x22, 0x3e, 0x0a, 0x09, 0x09, 0x3c, 0x61, 0x20,
		0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69,
		0x6c, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f, 0x70, 0x22, 0x3e, 0x53, 0x74,
		0x6f, 0x70, 0x20, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x69, 0x6e, 0x67,
		0x3c, 0x2f, 0x61, 0x3e, 0x0a, 0x09, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e,
		0x0a, 0x09, 0x3c, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x09, 0x20, 0x20, 0x20,
		0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x63, 0x68,
		0x61, 0x72, 0x74, 0x5f, 0x64, 0x69, 0x76, 0x22, 0x20, 0x73, 0x74, 0x79,
		0x6c, 0x65, 0x3d, 0x22, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3a, 0x20, 0x31,
		0x30, 0x30, 0x30, 0x70, 0x78, 0x3b, 0x20, 0x68, 0x65, 0x69, 0x67, 0x68,
		0x74, 0x3a, 0x20, 0x36, 0x30, 0x30, 0x70, 0x78, 0x3b, 0x20, 0x6d, 0x61,
		0x72, 0x67, 0x69, 0x6e, 0x3a, 0x20, 0x30, 0x20, 0x61, 0x75, 0x74, 0x6f,
		0x3b, 0x22, 0x3e, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x09, 0x3c,
		0x2f, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x0a, 0x09, 0x3c, 0x64, 0x69, 0x76,
		0x3e, 0x0a, 0x09, 0x09, 0x3c, 0x68, 0x33, 0x20, 0x73, 0x74, 0x79, 0x6c,
		0x65, 0x3d, 0x22, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x73, 0x69, 0x7a, 0x65,
		0x3a, 0x20, 0x31, 0x35, 0x70, 0x78, 0x3b, 0x66, 0x6f, 0x6e, 0x74, 0x2d,
		0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x62, 0x6f, 0x6c, 0x64,
		0x3b, 0x22, 0x3e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x20, 0x53, 0x65, 0x72,
		0x76, 0x69, 0x63, 0x65, 0x20, 0x49, 0x6e, 0x66, 0x6f, 0x3c, 0x2f, 0x68,
		0x33, 0x3e, 0x0a, 0x09, 0x09, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x69, 0x64,
		0x3d, 0x22, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
		0x22, 0x20, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3d, 0x22, 0x77, 0x69, 0x64,
		0x74, 0x68, 0x3a, 0x35, 0x30, 0x30, 0x70, 0x78, 0x3b, 0x20, 0x6d, 0x61,
		0x72, 0x67, 0x69, 0x6e, 0x3a, 0x20, 0x30, 0x20, 0x61, 0x75, 0x74, 0x6f,
		0x3b, 0x22, 0x3e, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x09, 0x3c,
		0x2f, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x20, 0x20, 0x3c, 0x2f, 0x62, 0x6f,
		0x64, 0x79, 0x3e, 0x0a, 0x3c, 0x2f, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0x0a,
	}, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"info-off.html": info_off_html,
	"info.html": info_html,
}
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"info-off.html": &_bintree_t{info_off_html, map[string]*_bintree_t{
	}},
	"info.html": &_bintree_t{info_html, map[string]*_bintree_t{
	}},
}}