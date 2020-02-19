package global

import (
	"gowallet/config"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	// "github.com/weekface/mgorus"
	// "github.com/rogierlommers/logrus-redis-hook"
	// "github.com/Abramovic/logrus_influxdb"
	// "github.com/bshuster-repo/logrus-logstash-hook"
)

type DefaultFieldHook struct {
}

func (hook *DefaultFieldHook) Fire(entry *log.Entry) error {
	entry.Data["appName"] = "tapi"
	return nil
}

func (hook *DefaultFieldHook) Levels() []log.Level {
	return log.AllLevels
}

type XMLFormater struct {
}

func (f *XMLFormater) Format(entry *log.Entry) ([]byte, error) {
	return nil, nil
}

var (
	// 工程配置
	ProjectConfig *config.JSONConfig
)

// 配置日志文件
func configLocalFilesystemLogger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {
	baseLogPath := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d_%H%M",
		// rotatelogs.WithLinkName(baseLogPath), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge), // 文件最大保存时间
		// rotatelogs.WithRotationCount(365),  // 最多存365个文件
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)

	if err != nil {
		log.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer, // 为不同级别设置不同的输出目的
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	log.AddHook(lfHook)

	// testHook := &DefaultFieldHook{}
	// log.AddHook(testHook)

	// hooker, err := mgorus.NewHooker("localhost:27017", "db", "collection")
	// if err == nil {
	// 	log.AddHook(hooker)
	// }

	log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{})
	// log.SetFormatter(&XMLFormater{})

	if ProjectConfig.Servers.Debug == 0 {
		log.SetLevel(log.ErrorLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

// 加载全局配置
func LoadGlobalConfig() error {
	var err error
	ProjectConfig, err = config.LoadJSONConfig("conf/conf.json")
	if err != nil {
		return err
	}

	configLocalFilesystemLogger("log", "wallet", time.Hour*24*365, time.Hour*24)

	return nil
}
