package trace

import (
	"runtime"
	"strings"
)

func caller(position int) (string, int, string) {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(position, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	name := strings.Split(f.Name(), "/")
	if strings.Contains(name[len(name)-1], "logger.") {
		runtime.Callers(4, pc)
		f := runtime.FuncForPC(pc[0])
		file, line = f.FileLine(pc[0])
		name = strings.Split(f.Name(), "/")
	}
	return file, line, name[len(name)-1]
}

func CallerTrace() (string, int, string) {
	return caller(3)
}

func CallerTraceWithPosition(position int) (string, int, string) {
	return caller(position)
}
