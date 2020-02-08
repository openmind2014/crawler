package main

import (
	"crawler/engine"
	"crawler/schduler"
	"crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		//Scheduler: &schduler.SimpleScheduler{},
		Scheduler:   &schduler.QueuedScheduler{},
		WorkerCount: 100,
	}

	//e.Run(engine.Request{
	//	Url:        "https://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	e.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})
}
