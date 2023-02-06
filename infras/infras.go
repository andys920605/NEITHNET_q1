package infras

import (
	"context"
	"fmt"
	"log"
	"q1/infras/configs"
	infras_itf "q1/infras/interface"
	"q1/models/commons"
	"q1/utils"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// region members
var (
	cfgTemp  *configs.Config
	viperTmp *viper.Viper
) // Austin 20220525
// endregion

// Options
type Options struct {
	// context
	Ctx context.Context
	// system information
	Info *commons.SystemInfo
	// fixed config data
	Config *configs.Config //`wire:"-"`
	// logger
	Logger infras_itf.IApiLogger
}

// region public methods
// new config
func ProvideConfig() *configs.Config {
	buffer, err := configs.LoadConfig(utils.GetConfigPath())
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	cfgTemp, err = configs.ParseConfig(buffer)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}
	viperTmp = buffer // for backup
	return cfgTemp
}

// on config change -- call back
// 監看檔案變化
func (opts Options) OnConfigChange() {
	viperTmp.WatchConfig() // 監控配定及重新取得設定值
	viperTmp.OnConfigChange(func(e fsnotify.Event) {
		if e.Op == fsnotify.Write {
			opts.Logger.InitLogger()
			if temp, err := configs.ParseConfig(viperTmp); err != nil {
				log.Fatalf("ParseConfig: %v", err)
			} else {
				cfgTemp = temp
			}
			fmt.Println("Config file changed:", e.Name)
		}
	})
}

// get new config
// 取新變化的值
func (opts Options) GetConfig() *configs.Config {
	return cfgTemp
}

// endregion
