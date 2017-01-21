package adler32

const (
	modAdler = 65521
)

func Checksum(data []byte) uint32 {
	var (
		a uint32 = 1
		b uint32 = 0
	)
	for _, d := range data {
		a = (a + uint32(d)) % modAdler
		b = (b + a) % modAdler
	}
	return (b << 16) | a
}
