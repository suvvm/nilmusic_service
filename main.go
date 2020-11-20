package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"suvvm.work/nilmusic_service/config"
	"suvvm.work/nilmusic_service/dal/db"
)

var (
	dbConfig = "./conf/db_config.yml"
)

// InitConfig 初始化配置信息
func InitConfig() {
	str, err := os.Getwd() // 获取相对路径
	if err != nil {
		panic(fmt.Sprintf("filepath failed, err=%v", err))
	}
	filename, err := filepath.Abs(filepath.Join(str, dbConfig)) // 获取db配置文件路径
	if err != nil {
		panic(fmt.Sprintf("filepath failed, err=%v", err))
	}
	conf := config.Init(filename)                    // 读取db配置文件
	if err = db.InitDB(&conf.DBConfig); err != nil { // 初始化db链接
		panic(fmt.Sprintf("init db conn err=%v", err))
	}
}

func main() {
	InitConfig() // 初始化服务端配置信息
	r := gin.Default()
	register(r)                     // 注册路由
	if err := r.Run(); err != nil { // 启动api层服务端
		log.Println("run api service fail")
	}
}
