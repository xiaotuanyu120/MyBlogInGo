package config

import (
	"path"
	"path/filepath"

	"github.com/spf13/viper"
	"github.com/xiaotuanyu120/MyBlogInGo/utils"
)

type Config struct {
	BaseDir     string
	SrcDir      string `mapstructure:"srcDir"`
	DstDir      string `mapstructure:"dstDir"`
	ChromaStyle string `mapstructure:"chromaStyle"`
}

type viperConfParam struct {
	Name      string
	Extension string
	Path      string
}

/*
MBConfig
global app configuration variable
*/
var MBConfig = new(Config)

/*
viperConf
load customized configuration file
or
load default configuration file [[`projectdir`/conf]]
*/
func viperConf(conf viperConfParam) (*Config, error) {
	projectDir := filepath.Join(utils.ExecuteDir(), "..")

	if conf.Name == "" {
		conf.Name = "blog"
	}

	if conf.Extension == "" {
		conf.Extension = "yaml"
	}

	if conf.Path == "" {
		conf.Path = path.Join(projectDir, "conf")
	}

	// load viper config
	v := viper.New()
	v.SetConfigName(conf.Name)
	v.SetConfigType(conf.Extension)
	v.AddConfigPath(conf.Path)

	// load app config
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	v.Set("BaseDir", projectDir)

	// unmarshal app config
	if err := v.Unmarshal(MBConfig); err != nil {
		return nil, err
	}

	return MBConfig, nil
}

/*
GetConfig
return loaded configuration global variable MBConfig
*/
func GetConfig() (*Config, error) {
	// return it directly if config is already loaded
	if MBConfig.BaseDir != "" {
		return MBConfig, nil
	}

	var vcp = viperConfParam{}
	return viperConf(vcp)
}
