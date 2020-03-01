package execute

import (
	"github.com/spf13/viper"
	"log"
)

type Connector struct {
	checker	 	*ProcessChecker
}

func (conn *Connector)LoadConfig(configPath string) *Config {
	if configPath == ""{
		configPath = "/home/yuki/dev/workplace/eth-connect/src/gethProcess/config.json"
	}
	viper.SetConfigFile(configPath)
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil{
		log.Fatal(err)
		return nil
	}

	return &Config{
		GethPath:   viper.GetString("GethPath"),
		Port:       viper.GetInt("Port"),
		DataDir:    viper.GetString("DataDir"),
		NetworkId:  viper.GetInt("NetworkId"),
		Mine:       viper.GetBool("Mine"),
		MineThread: viper.GetInt("MineThread"),
		CoinBase:   viper.GetString("CoinBase"),
		GenesisFile:viper.GetString("GenesisFile"),
		Rpc:		viper.GetBool("Rpc"),
		RpcAddr:	viper.GetString("RpcAddr"),
		RpcPort:	viper.GetInt("RpcPort"),
		GethLogPath:viper.GetString("GethLogPath"),
	}

}

func (conn *Connector)Start(configPath string){
	config := conn.LoadConfig(configPath)
	if config == nil{
		log.Fatal("config load unsuccessful")
		return
	}
	conn.checker = NewProcessChecker(config)
	conn.checker.Execute()
}

