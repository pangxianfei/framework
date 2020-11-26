package graceful

func ShutDown(quietly bool) {
	logInfo(quietly, "tmaic is shutting down")
	closeQueue(quietly)
	closeCache(quietly)
	closeDB(quietly)
	closeMonitor(quietly)
	logInfo(quietly, "tmaic is shut down")
}
