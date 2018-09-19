package utils

import (
	"fmt"
	"go.uber.org/zap"
	"strconv"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	AccessToken string
	ChatIDs     []int64

	subscribers map[string]*Subscriber
)

type Subscriber struct {
	Name       string
	NotifyChan chan int
}

func init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	loadConfig()
	subscribers = make(map[string]*Subscriber)

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config ReloadConfig!", e.Name)
		Logger.Info("reload config", zap.String("config_path", e.Name))
		ReloadConfig()

		// 通知所有 subscribers
		publish()
	})
}

func Subscribe(s *Subscriber) {
	subscribers[s.Name] = s
}

func UnSubscribe(name string) {
	delete(subscribers, name)
}

func publish() {
	for _, s := range subscribers {
		s.NotifyChan <- 1
	}
}

func ReloadConfig() {
	loadConfig()
}

func loadConfig() {
	AccessToken = viper.GetString("telegram.access_token")
	ChatIDs = *parseChatIDs()
}

func parseChatIDs() *[]int64 {
	var chatIDs []int64
	chatIDText := viper.GetStringSlice("telegram.chat_ids")

	for _, v := range chatIDText {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			Logger.Error(err.Error())
		}
		chatIDs = append(chatIDs, i)
	}

	return &chatIDs
}
