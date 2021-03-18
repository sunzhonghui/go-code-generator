package conf

import (
	"code/gen/util/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConf struct {
	UrlName      string `json:"urlName"`
	IP           string `json:"ip"`
	DatabaseName string `json:"databaseName"`
	Port         string `json:"port"`
	UserName     string `json:"userName"`
	Password     string `json:"password"`
}

type ProjectConf struct {
	Name       string `json:"name"`       //项目名称
	Abbr       string `json:"abbr"`       //项目缩写
	ModName    string `json:"modName"`    //mod 名称
	RouterName string `json:"routerName"` //路由前缀
}

var Database = &DatabaseConf{}
var Project = &ProjectConf{}
var DB *gorm.DB

func Init() {

	//viper.SetConfigType("yaml") // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath("resource") // 设置读取路径：就是在此路径下搜索配置文件。
	//viper.AddConfigPath("$HOME/.appname")  // 多次调用以添加多个搜索路径
	viper.SetConfigName("conf") // 设置被读取文件的全名，包括扩展名。
	//viper.SetConfigName("server") // 设置被读取文件的名字： 这个方法 和 SetConfigFile实际上仅使用一个就够了
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // 读取配置文件： 这一步将配置文件变成了 Go语言的配置文件对象包含了 map，string 等对象。
	if err != nil {
		logger.Log.WithFields(logrus.Fields{"err": err}).Warn("未找到配置文件，创建配置文件")
		viper.SetDefault("database.url", "localhost")
		viper.SetDefault("database.ip", "192.168.1.1")
		viper.SetDefault("database.port", "3306")
		viper.SetDefault("database.username", "root")
		viper.SetDefault("database.password", "root")
		viper.SetDefault("database.database", "")

		viper.SetDefault("project.name", "Go代码生成器v0.0.1")
		viper.SetDefault("project.abbr", "code-gen")
		viper.SetDefault("project.mod", "code/gen")
		viper.SetDefault("project.router", "api/test")
		//项目默认配置
		viper.WriteConfigAs("resource/conf.yaml")
	}

	ResetData()

	// 控制台输出： map[first:panda last:8z] 99 panda [Coding Movie Swimming]

}

func ResetData() {
	Database.UrlName = viper.GetString("database.url")
	Database.IP = viper.GetString("database.ip")
	Database.Port = viper.GetString("database.port")
	Database.DatabaseName = viper.GetString("database.database")
	Database.UserName = viper.GetString("database.username")
	Database.Password = viper.GetString("database.password")

	Project.Name = viper.GetString("project.name")
	Project.Abbr = viper.GetString("project.abbr")
	Project.ModName = viper.GetString("project.mod")
	Project.RouterName = viper.GetString("project.router")
	logger.Log.WithFields(logrus.Fields{"data": Database, "data2": Project}).Info("读取本地配置")
}

func (d *DatabaseConf) Save() {
	logger.Log.WithFields(logrus.Fields{"data": d}).Info("保存数据库配置到文件")

	viper.Set("database.url", d.UrlName)
	viper.Set("database.ip", d.IP)
	viper.Set("database.port", d.Port)
	viper.Set("database.username", d.UserName)
	viper.Set("database.password", d.Password)
	viper.Set("database.database", d.DatabaseName)
	err := viper.WriteConfigAs("resource/conf.yaml")
	if err != nil {
		logger.Log.WithFields(logrus.Fields{"data": d, "err": err}).Error("保存数据库配置到文件失败")
	}
}

func (d *ProjectConf) Save() {
	logger.Log.WithFields(logrus.Fields{"data": d}).Info("保存项目配置到文件")

	viper.Set("project.name", Project.Name)
	viper.Set("project.abbr", Project.Abbr)
	viper.Set("project.mod", Project.ModName)
	viper.Set("project.router", Project.RouterName)
	err := viper.WriteConfigAs("resource/conf.yaml")
	if err != nil {
		logger.Log.WithFields(logrus.Fields{"data": d, "err": err}).Error("保存数据库配置到文件失败")
	}
}
func (d *DatabaseConf) GetDB() (err error) {
	mysqlUrl := d.UserName + ":" + d.Password + "@(" + d.IP + ":" + d.Port + ")/" + d.DatabaseName + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(mysqlUrl), &gorm.Config{})
	if err == nil {
		logger.Log.WithFields(logrus.Fields{"data": mysqlUrl}).Info("数据库连接地址")
		DB = db
	}
	return
}
