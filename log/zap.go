package log

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

//不同类型logger
var (
	Console *zap.SugaredLogger //控制台
	Logic   *zap.SugaredLogger //游戏逻辑
)

const (
	consolePath = "./gameLog/console/"
	logicPath   = "./gameLog/logic/"
)

//初始化所有ZapLogger
func InitZapLogger() {
	t := time.Now()
	timeStr := fmt.Sprintf("%04d-%02d-%02d", t.Year(), t.Month(), t.Day())
	cName := fmt.Sprintf("server_%s.log", timeStr)
	lName := fmt.Sprintf("game_%s.log", timeStr)

	Console = initSingleZapLogger(consolePath, cName, 0)
	Logic = initSingleZapLogger(logicPath, lName, 1)
}

//初始化具体logger
//日志目录,日志名,编码种类(0:普通, 1:JSON)
func initSingleZapLogger(path, filename string, encoderType int) *zap.SugaredLogger {
	//输出目录是否存在
	if _, err := os.Stat(path); err != nil {
		os.MkdirAll(path, 0777)
	}

	writer := getLogWriter(path, filename)
	var enc zapcore.Encoder
	switch encoderType {
	case 0:
		enc = getNormalEncoder()
	case 1:
		enc = getJSONEncoder()
	default:
		enc = getNormalEncoder()
	}

	zCore := zapcore.NewCore(enc, writer, zapcore.InfoLevel) //第三个参数表示要>=该级别才会写入
	zLogger := zap.New(zCore, zap.AddCaller())               //添加具体调用函数
	return zLogger.Sugar()
}

//Writer
func getLogWriter(path, filename string) zapcore.WriteSyncer {
	fn := fmt.Sprintf("%s%s", path, filename)
	fmt.Println("getLogWriter fn", fn)
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fn,   //文件路径
		MaxSize:    10,   //进行切割之前,日志最大大小MB
		MaxBackups: 5,    //旧文件保留个数
		MaxAge:     30,   //旧文件最大保留天数
		Compress:   true, //是否压缩/归档旧文件
	}

	return zapcore.AddSync(lumberJackLogger)
}

//encoder:normal
func getNormalEncoder() zapcore.Encoder {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderCfg)
}

//encoder:json
func getJSONEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}
