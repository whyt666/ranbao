package ini

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var Conf *ServerConf

type ServerConf struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	MailBySender Mail   `mapstructure:"mail_by_sender"`
	RedisConfig  Redis  `mapstructure:"redis"`
	MysqlConfig  Mysql  `mapstructure:"mysql"`
}

type Mail struct {
	QQMail string `mapstructure:"qq_mail"`
}

type Mysql struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

type Redis struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

func InitConfig(filePath string) {
	viper.SetConfigFile(filePath)

	err := viper.ReadInConfig() // 读取配置信息
	if err != nil {
		// 读取配置信息失败
		zap.S().Panicf("viper.ReadInConfig failed, err:%v\n", err)
		return
	}

	// 把读取到的配置信息反序列化到 Conf 变量中
	if err := viper.Unmarshal(&Conf); err != nil {
		zap.S().Panicf("viper.Unmarshal failed, err:%v\n", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err := viper.Unmarshal(&Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
	})
	zap.S().Info("config初始化成功")
	fmt.Println("ok")

	return
}
