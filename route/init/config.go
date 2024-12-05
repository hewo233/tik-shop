package init

import (
	"github.com/hewo/tik-shop/route/config"
	"github.com/hewo/tik-shop/shared/consts"
	"github.com/spf13/viper"
	"log"
)

func InitConfig() {
	v := viper.New()
	v.SetConfigFile(consts.ApiConfigPath)

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Successfully read config info from %s : %v", v.ConfigFileUsed(), v.AllSettings())

	// 设置全局变量
	if err := v.Unmarshal(&config.GlobalServerConfig); err != nil {
		log.Fatal(err)
	}
	log.Printf("Successfully set config info to GlobalServerConfig : %v", config.GlobalServerConfig)

}
