package config

import (
	"fmt"
	"github.com/helix-drop/AnonyBBS/models"
	"github.com/spf13/viper"
	"os"
	"path"
)

func Init() {
	// 读取配置文件
	dir, _ := os.Getwd()
	configPath := path.Join(dir, "config")
	configName := "config"
	viper.SetConfigName(configName) // 指定配置文件的文件名称(不需要指定配置文件的扩展名)
	viper.AddConfigPath(configPath) // 设置配置文件的搜索目录
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("配置文件读取失败：%s", err))
	}

	// 连接mysql数据库
	addr := viper.GetString("mysql.address")
	port := viper.GetString("mysql.port")
	user := viper.GetString("mysql.user")
	pwd := viper.GetString("mysql.password")
	db := viper.GetString("mysql.database")
	// "${user}:${pwd}@tcp(${addr}:${port})/${db}?charset=utf8&parseTime=True&loc=Local"
	connectStr := user + ":" + pwd + "@tcp(" + addr + ":" + port + ")/" + db + "?charset=utf8mb4&parseTime=True&loc=Local"
	if err := models.ConnectMysql(connectStr); err != nil {
		panic(err)
	}
}
