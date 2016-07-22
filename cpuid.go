package cpuid

func cpuid(info *[4]uint32, ax uint32)

func Cpuid(info *[4]uint32, ax uint32) {
	cpuid(info, ax)
	return
}
