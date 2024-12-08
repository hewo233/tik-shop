package init

import (
	"github.com/hewo/tik-shop/rpc/user-service/config"
	"github.com/hewo/tik-shop/shared/consts"
	"github.com/spf13/viper"
	"log"
)

func InitConfig() {
	v := viper.New()
	v.SetConfigFile(consts.UserConfigPath)
	if err := v.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Successfully read config info from %s : %v", v.ConfigFileUsed(), v.AllSettings())

	if err := v.Unmarshal(&config.GlobalUserServerConfig); err != nil {
		log.Fatal(err)
	}
	// TODO: Delete after testing
	log.Printf("Successfully set config info to GlobalUserServerConfig : %v", config.GlobalUserServerConfig)
}
