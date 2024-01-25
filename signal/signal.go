// Package signal contains signals used to communicate to the gops agents.
package signal

const (
	// StackTrace represents a command to print stack trace.
	StackTrace = byte(0x1)

	// GC runs the garbage collector.
	GC = byte(0x2)

	// MemStats reports memory stats.
	MemStats = byte(0x3)

	// Version prints the Go version.
	Version = byte(0x4)

	// HeapProfile starts `go tool pprof` with the current memory profile.
	HeapProfile = byte(0x5)

	// CPUProfile starts `go tool pprof` with the current CPU profile
	CPUProfile = byte(0x6)

	// Stats returns Go runtime statistics such as number of goroutines, GOMAXPROCS, and NumCPU.
	Stats = byte(0x7)

	// Trace starts the Go execution tracer, waits 5 seconds and launches the trace tool.
	Trace = byte(0x8)

	// BinaryDump returns running binary file.
	BinaryDump = byte(0x9)

	// SetGCPercent sets the garbage collection target percentage.
	SetGCPercent = byte(0x10)

	// analyze On-CPU as well as Off-CPU (e.g. I/O) time together. https://github.com/felixge/fgprof
	CPUAllProfile = byte(0x11)

	// dotnet dump
	DotnetDump = byte(0x12)
	// dotnet trace
	DotnetTrace = byte(0x13)
	// dotnet counter
	DotnetCounter = byte(0x14)
)
