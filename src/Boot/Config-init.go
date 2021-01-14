package Boot

import (
	"fmt"
	"github.com/micro/go-micro/v2/config"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
	"os"
	"time"
)

type GlobalConfig struct {
	Config *struct{
		Address string
		Path 	string
		Port 	uint64
	}
	Service *struct{
		Namespace 	string
		Name 		string
	}
	Data *DataConfig
}

type DataConfig struct {
	Mysql *MySqlConfig
	Redis *RedisConfig
}

var JConfig *GlobalConfig
var nacosClient config_client.IConfigClient

//初始化主配置文件(本地配置)和初始化nacos链接
func InitConfig() {
	configFile := "app.yaml"
	err := config.LoadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}
	JConfig = &GlobalConfig{Data:new(DataConfig)}
	err = config.Get("100txy").Scan(JConfig)
	if err != nil {
		log.Fatal(err)
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:JConfig.Config.Address,
			ContextPath:JConfig.Config.Path,
			Port:JConfig.Config.Port,
		},
	}
	nacosClient, err = clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
	})
	//开始加重数据相关配置
	//JConfig.Data.Mysql = new(MySqlConfig)
	listenNacos("100txy-sysconfig","100txy_GROUP",JConfig.Data)
}

func listenNacos(dataid string, group string, model interface{}){
	err := nacosClient.ListenConfig(vo.ConfigParam{
		DataId:   dataid,
		Group:    group,
		Content:  "",
		OnChange: func(namespace, group, dataId, data string) {
			time.Sleep(time.Second*1)
			//shouldReload := reload
			//if !model.IsLoad() {
			//	shouldReload = false //如果model没有被加载过，则不需要做重载
			//}
			cacheFile := fmt.Sprintf("./runtime/configcache/%s-%s.yaml",group,dataid)
			file,err := os.OpenFile(cacheFile,os.O_RDWR|os.O_CREATE|os.O_TRUNC,0666)
			if err != nil {
				log.Println(err)
				return
			}
			defer file.Close()
			_, err = file.WriteString(data)
			if err != nil {
				log.Println(err)
				return
			}
			err = config.LoadFile(cacheFile)
			if err != nil {
				log.Println(err)
				return
			}
			err = config.Scan(model)
			if err != nil {
				log.Println(err)
				return
			}
			//if shouldReload { //重载关键代码
			//	err := model.Reload()
			//	if err != nil {
			//		logger.Error(err)
			//		return
			//	}else{
			//		logger.Info(dataid,"重载完成")
			//	}
			//}
		},
	})
	if err != nil {
		log.Println("listen config error,dataid:",dataid,err)
	}
}