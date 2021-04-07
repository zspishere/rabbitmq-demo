package global

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

var (
	RabbitmqUrl string
)

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(RootDir + "/config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	RabbitmqUrl = viper.GetString("rabbitmq-url")
	log.Println("RabbitMQ Url", RabbitmqUrl)

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		viper.ReadInConfig()
		RabbitmqUrl = viper.GetString("rabbitmq-url")
	})
}
