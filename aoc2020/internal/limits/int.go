package limits

// Max and min values of int and uint on the compiling system.
const (
	UintMax = ^uint(0)
	UintMin = uint(0)

	IntMax = int(UintMax >> 1) 
	IntMin = -IntMax - 1
)
