package generate_ids

import "sync"

var mx sync.Mutex
var markIDGenerate = 0

func GetMarkID() int {
	mx.Lock()
	defer mx.Unlock()

	markIDGenerate++
	return markIDGenerate
}
