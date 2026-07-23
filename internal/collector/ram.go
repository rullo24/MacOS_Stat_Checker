package collector

type RamStats struct {
	AppMemoryInUseMb        int
	WiredMemoryInUseMb      int
	CompressedMemoryInUseMb int
	FreeMemoryMb            int
}
