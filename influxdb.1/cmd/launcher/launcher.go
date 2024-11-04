package launcher

const (
	// DiskStore stores all REST resources to disk in boltdb and sqlite.
	DiskStore = "disk"
	// BoltStore also stores all REST resources to disk in boltdb and sqlite. Kept for backwards-compatibility.
	BoltStore = "bolt"
	// MemoryStore stores all REST resources in memory (useful for testing).
	MemoryStore = "memory"

	// LogTracing enables tracing via zap logs
	LogTracing = "log"
	// JaegerTracing enables tracing via the Jaeger client library
	JaegerTracing = "jaeger"
)