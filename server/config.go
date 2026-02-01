package server

import (
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	DBConn     string `mapstructure:"DB_CONN"`
}

func LoadConfig() *Config {
	// viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	var config Config
	bindSystemEnv(config)

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("unable to load config: %v", err)
	}

	// config := Config{
	// 	DBConn:     viper.GetString("DB_CONN"),
	// 	ServerPort: viper.GetString("SERVER_PORT"),
	// }

	return &config
}

func bindSystemEnv(iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			bindSystemEnv(v.Interface(), append(parts, tv)...)
		default:
			viper.BindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}
