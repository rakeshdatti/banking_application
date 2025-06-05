package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)


var log *zap.Logger
func init() {
	var err error
	// log,err = zap.NewProduction(zap.AddCallerSkip(1))
	
	//creating our own configuration  to the key in info 
	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey="timestamp"
	encoderConfig.StacktraceKey=""
	encoderConfig.EncodeTime=zapcore.ISO8601TimeEncoder
	config.EncoderConfig=encoderConfig

	log,err =config.Build(zap.AddCallerSkip(1))
	if err!=nil{
		panic(err)
	}
}


func Info(message string ,fileds ...zap.Field){
	log.Info(message,fileds...)
}

func Debug(message string ,fileds ...zap.Field){
	log.Debug(message,fileds...)
}

func Error(message string ,fileds ...zap.Field){
	log.Error(message,fileds...)
}
