package main

import (
	"bus-booking/util"
)

func main() {
	util.Init()
	router := Route()
	_ = router.Run(util.Port)
}
