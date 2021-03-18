package runner

import (
	"code/gen/util/conf"
	"code/gen/util/font"
	"code/gen/util/logger"
)

func Runner() {
	font.GetZhFont()
	font.GetTemp()
	conf.Init()
	logger.InitLog()
}
