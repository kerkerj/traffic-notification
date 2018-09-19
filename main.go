package main

import (
	"fmt"
	"sync"
	"time"
	
	"github.com/kerkerj/traffic-notification/source"
	"github.com/kerkerj/traffic-notification/utils"
	"go.uber.org/zap"

)

var wg sync.WaitGroup

func main() {
	utils.Logger.Info("Start ")
	
	sub := Subscribe()
	go configUpdated(sub)
	
	src := source.New()

	matchPattern := false
	for {
		// fetch data
		utils.Logger.Info("Fetch..")
		
		src.FetchXML(source.IncidentDataURL).ParseXML()

		if matchPattern {
			// send message to subscribed channel
		}

		utils.Logger.Info("get token", zap.String("token", utils.AccessToken), zap.Int64s("chat_ids", utils.ChatIDs))

		time.Sleep(10 * time.Second)
	}

	wg.Wait()
}

func Subscribe() *utils.Subscriber {
	dataChan := make(chan int, 1)
	sub := &utils.Subscriber{
		Name:       "traffic",
		NotifyChan: dataChan,
	}
	utils.Subscribe(sub)
	return sub
}

func configUpdated(s *utils.Subscriber) {
	for i := range s.NotifyChan {
		utils.Logger.Info(fmt.Sprintf("reload config: %v", i==1))
	}
}
