package cpuid

func cpuid(info *[4]uint32, ax uint32) {
	return
}

func Cpuid(info *[4]uint32, ax uint32) {
	cpuid(info, 0)
	return
}
