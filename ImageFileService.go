package ImageFileMicroservices

func RunService() {
	workToDo := make(chan string)
	finishedWorkMap := make(map[string]bool, 1024)
	go startProcessor(workToDo, &finishedWorkMap)
	setupWebInterface(workToDo, &finishedWorkMap)
}
