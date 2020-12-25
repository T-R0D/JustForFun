package limits

// Max and min values of int and uint on the compiling system.
const (
	Uint64Max = ^uint64(0)
	Uint64Min = uint64(0)
	
	Int64Max = int64(Uint64Max >> 1) 
	Int64Min = -Int64Max - 1
)
