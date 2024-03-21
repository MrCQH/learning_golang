package byte_size

type ByteSize float64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota) //
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)
