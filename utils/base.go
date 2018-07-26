package utils

import (
	"log"
)

func Info(info ...string) {
	for _, i := range info {
		log.Println("=>Info :", i)
	}
}

func PanicRecover(errs ...error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("=>Error :", r)
		}
	}()
	if len(errs) != 0 {
		panic(errs[0].Error())
	} else {
		panic("=>Error not passed")
	}
}
