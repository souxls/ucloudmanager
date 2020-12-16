package log

import (
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	log = logrus.New()
	// 设置别名
	Debugln = log.Debugln
	Infoln  = log.Infoln
	Warnln  = log.Warningln
	Errorln = log.Errorln
	Fatalln = log.Fatalln
	Panicln = log.Panicln
	Debugf  = log.Debugf
	Infof   = log.Infof
	Warnf   = log.Warningf
	Errorf  = log.Errorf
	Fatalf  = log.Fatalf
	Panicf  = log.Panicf
)

// Init 日志模块初始化
func Init() {

	log.Formatter = &logrus.TextFormatter{
		TimestampFormat:        "2006-01-02 15:04:05.000",
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			file = f.File[strings.LastIndex(f.File, "/")+1:] + ":" + strconv.Itoa(f.Line)
			function = f.Function[strings.LastIndex(f.Function, ".")+1:]
			return function, file
		},
	}

	log.SetReportCaller(true)
	log.SetLevel(5)
	log.SetNoLock()
	log.SetOutput(os.Stdout)

}
