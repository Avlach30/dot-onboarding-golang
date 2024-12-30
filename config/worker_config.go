package config

var (
	// WORKER
	MaxWorkerQueue        = Get("MAX_WORKER_QUEUE", "3")
	MaxParalelWorkerQueue = Get("MAX_PARALEL_WORKER_QUEUE", "3")
	MaxTriesQueue         = Get("MAX_TRIES", "3")
)
