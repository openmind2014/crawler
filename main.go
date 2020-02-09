package main

import (
	"crawler/engine"
	"crawler/persist"
	"crawler/schduler"
	"crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		//Scheduler: &schduler.SimpleScheduler{},
		Scheduler:   &schduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
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
