package font

import (
	"code/gen/util/common"
	"code/gen/util/logger"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	zhFontUrl      = "http://f.idmiss.com/file/font/WenQuanWeiMiHei.ttf"
	tempMapperUrl  = "http://images.idmiss.com/golang/temp/mapper.go.tpl"
	tempApiUrl     = "http://images.idmiss.com/golang/temp/api.go.tpl"
	tempModelUrl   = "http://images.idmiss.com/golang/temp/model.go.tpl"
	tempServiceUrl = "http://images.idmiss.com/golang/temp/service.go.tpl"
	tempRouterUrl  = "http://images.idmiss.com/golang/temp/router.go.tpl"
)

func GetZhFont() {
	logger.Log.WithFields(logrus.Fields{"data": ""}).Info("检测字体是否存在...")
	if _, err := common.DownloadFileForUrl(zhFontUrl, "resource/font/", "WenQuanWeiMiHei.ttf"); err != nil {
		logger.Log.WithFields(logrus.Fields{"err": err}).Error("字体下载失败...")

	}
	logger.Log.Info("开始配置字体...")
	os.Setenv("FYNE_FONT", "resource/font/WenQuanWeiMiHei.ttf")
	logger.Log.Info("配置字体成功")
}
func GetTemp() {

	logger.Log.WithFields(logrus.Fields{"data": ""}).Info("检测Mapper模板是否存在...")
	if _, err := common.DownloadFileForUrl(tempMapperUrl, "resource/temp/", "mapper.go.tpl"); err != nil {
		logger.Log.WithFields(logrus.Fields{"err": err}).Error("下载mapper模板失败")
	}
	logger.Log.WithFields(logrus.Fields{"data": ""}).Info("检测api模板是否存在...")
	if _, err := common.DownloadFileForUrl(tempApiUrl, "resource/temp/", "api.go.tpl"); err != nil {
		logger.Log.WithFields(logrus.Fields{"err": err}).Error("下载api模板失败")
	}
	logger.Log.WithFields(logrus.Fields{"data": ""}).Info("检测Model模板是否存在...")
	if _, err := common.DownloadFileForUrl(tempModelUrl, "resource/temp/", "model.go.tpl"); err != nil {
		logger.Log.WithFields(logrus.Fields{"err": err}).Error("下载Model模板失败")
	}
	logger.Log.WithFields(logrus.Fields{"data": ""}).Info("检测service模板是否存在...")
	if _, err := common.DownloadFileForUrl(tempServiceUrl, "resource/temp/", "service.go.tpl"); err != nil {
		logger.Log.WithFields(logrus.Fields{"err": err}).Error("下载service模板失败")
	}
	logger.Log.WithFields(logrus.Fields{"data": ""}).Info("检测router模板是否存在...")
	if _, err := common.DownloadFileForUrl(tempRouterUrl, "resource/temp/", "router.go.tpl"); err != nil {
		logger.Log.WithFields(logrus.Fields{"err": err}).Error("下载tempApiUrl模板失败")
	}

}
