/*
@Time : 2020/7/13 12:00 下午
@Author : lucbine
@File : config.go
*/
package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type configWrap struct {
	ConfigPath string //配置文件逻辑
	Env        string //环境
	Vipers     map[string]*viper.Viper
}

const ConfigFileType = "toml"

var cw *configWrap

//初始化 配置文件
func InitConfig(env string) error {
	cw = &configWrap{
		ConfigPath: getBasePath(env),
		Env:        env,
	}

	//读取配置文件
	configFileInfos, err := ioutil.ReadDir(cw.ConfigPath)
	if err != nil {
		return err
	}
	cw.Vipers = make(map[string]*viper.Viper, len(configFileInfos))

	for _, value := range configFileInfos {
		v := viper.New()
		v.SetConfigFile(filepath.Join(cw.ConfigPath, value.Name()))
		v.SetConfigType(ConfigFileType)

		if err := v.ReadInConfig(); err != nil {
			fmt.Println(value.Name())
			return err
		}
		fileName := strings.SplitN(value.Name(), ".", 2)[0]
		fmt.Println("using config file :", v.ConfigFileUsed())
		cw.Vipers[fileName] = v
	}
	return nil
}

//获得路径
func getBasePath(env string) string {
	var (
		file     string
		err      error
		basePath string
	)

	if file, err = exec.LookPath(os.Args[0]); err != nil {
		panic(err)
	}

	if basePath, err = filepath.Abs(file); err != nil {
		panic(err)
	}
	configPath := path.Dir(path.Dir(basePath)) + "/config/" + env
	fmt.Println("config dir: ", configPath)
	return configPath
}

//获得当前的环境
func Env() string {
	return cw.Env
}

//配置文件解析
func Unmarshal(fileName string, rawVal interface{}) error {
	return cw.Vipers[fileName].Unmarshal(rawVal)
}
