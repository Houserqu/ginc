package ginc

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Logger = logrus.New() // logrus 实例

func InitLogger() {
	Logger.SetReportCaller(true)       // 日志记录文件命名
	Logger.SetLevel(logrus.DebugLevel) // 日志输出级别
	Logger.Out = ioutil.Discard        // 禁止 logrus 的输出

	var writer io.Writer // 输出位置

	if viper.GetString("log.out") == "file" {
		// 配置文件输入 hook
		logFolderPath, err := filepath.Abs(viper.GetString("log.path"))
		if err != nil {
			log.Fatalf(err.Error())
		}

		// 如果日志目录不存在则创建
		_, err = os.Stat(logFolderPath)
		if err != nil {
			if os.IsNotExist(err) {
				if err = os.Mkdir(logFolderPath, os.ModePerm); err != nil {
					log.Fatalf("create log file err: %v", errors.WithStack(err))
				}

			} else {
				log.Fatal(err.Error())
			}
		}

		logPath := path.Join(logFolderPath, "default.log")

		writer, err = rotatelogs.New(
			logPath+".%Y-%m-%d.log",
			rotatelogs.WithLinkName(logPath),          // 生成软链，指向最新日志文件
			rotatelogs.WithMaxAge(14*24*time.Hour),    // 文件最大保存时间
			rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
		)

		if err != nil {
			log.Fatalf("config log file system logger error. %v", errors.WithStack(err))
		}
	} else {
		// 默认输出到控制台
		writer = os.Stdout
	}

	// 不同级别输出
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     viper.GetBool("log.pretty"),
	})

	Logger.AddHook(lfHook)
}

func Log(c *gin.Context) *logrus.Entry {
	value, _ := c.Get("reqId")
	return Logger.WithFields(logrus.Fields{"type": "DEFAULT", "reqId": value})
}

func LogResty(resp *resty.Response) {
	Logger.WithField("type", "REQUEST").Info(resp.Request.URL, resp.StatusCode(), string(resp.Body()))
}
