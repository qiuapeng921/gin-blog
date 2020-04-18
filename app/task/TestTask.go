package task

import "log"

func TestOneTask() {
	for i := 0; i < 100; i++ {
		log.Println(i)
	}
}