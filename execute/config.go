package execute

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

const (
	ConfigFilePath = "config.json"
)

type Config struct {
	GethPath 	string
	Port 		int
	DataDir		string
	NetworkId 	int
	Mine		bool
	MineThread	int
	CoinBase	string
	GenesisFile string
	Rpc			bool
	RpcAddr		string
	RpcPort		int
	GethLogPath string
}


/*
func (Config* cf) loadConfig() bool{
	return true
}

func (Config* cf) saveConfig()bool{
	return true
}
*/


func Init(){
	var config Config
	viper.SetConfigFile("/home/yuki/dev/workplace/eth-connect/config/config.json")
	if err := viper.ReadInConfig(); err != nil{
		log.Fatal(err)
	}

	viper.Unmarshal(&config)
	j,_ :=json.Marshal(config)
	fmt.Print(string(j))
}
