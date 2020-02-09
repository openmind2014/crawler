package main

import (
	"crawler/distributed/config"
	"crawler/distributed/persist/cilent"
	"crawler/distributed/worker/client"
	"crawler/engine"
	"crawler/schduler"
	"crawler/zhenai/parser"
	"fmt"
)

func main() {
	itemChan, err := cilent.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	processor, err := client.CreateProcessor()
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:        &schduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	//e.Run(engine.Request{
	//	Url:    "https://www.zhenai.com/zhenghun",
	//	Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	//})

	e.Run(engine.Request{
		Url:    "https://www.zhenai.com/zhenghun/shanghai",
		Parser: engine.NewFuncParser(parser.ParseCity, config.ParseCity),
	})
}
