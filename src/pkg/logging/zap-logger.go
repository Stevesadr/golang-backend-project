package logging

import (
	"github.com/Stevesadr/golang-backend-project/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/lumberjack.v2"
)

type zapLogger struct{
	cfg *config.Config
	logger *zap.SugaredLogger
}

var zapLevelMap map[string]zapcore.Level = map[string]zapcore.Level{
	"debug" : zap.DebugLevel,
	"info" : zap.InfoLevel,
	"warn" : zap.WarnLevel,
	"error" : zap.ErrorLevel,
	"fatal" : zap.FatalLevel,
}

func (l *zapLogger) GetLogLevel() zapcore.Level{
	level, exist := zapLevelMap[l.cfg.Logger.Level]
	if !exist{
		return zapcore.DebugLevel
	}
	return level
}

func newZapLogger(cfg *config.Config) *zapLogger {
	logger := &zapLogger{cfg:cfg}
	logger.Init()
	return logger
}

func (l *zapLogger)Init(){
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename: l.cfg.Logger.FilePath,
		MaxSize: 1,
		MaxAge: 5,
		MaxBackups: 10,
		Compress: true,
	})

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		w,
		l.GetLogLevel(),
	)

	logger := zap.New(core, zap.AddCaller(),zap.AddCallerSkip(1), zap.AddStacktrace(zap.ErrorLevel)).Sugar()
	logger = logger.With("AppName", "MyApp", "LoggerName", "ZapLog")
	l.logger = logger
}

func prepareLogKey(cat Category, sub SubCategory, extra map[ExtraKey]interface{}) []interface{}{
	if extra == nil {
		extra = make(map[ExtraKey]interface{}, 0)
	}
	extra["Category"] = cat
	extra["SubCategory"] = sub
	params := mapToZapParams(extra)
	return params
}


func(l *zapLogger)Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKey(cat, sub, extra)
	l.logger.Debugw(msg, params...)
}
func(l *zapLogger)Debugf(template string, arg ...interface{}){
	l.logger.Debugf(template, arg)
}

func(l *zapLogger)Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKey(cat, sub, extra)
	l.logger.Infow(msg, params...)
}
func(l *zapLogger)Infof(template string, arg ...interface{}){
	l.logger.Infof(template, arg)
}

func(l *zapLogger)Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKey(cat, sub, extra)
	l.logger.Warnw(msg, params...)
}
func(l *zapLogger)Warnf(template string, arg ...interface{}){
	l.logger.Warnf(template, arg)
}

func(l *zapLogger)Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKey(cat, sub, extra)
	l.logger.Errorw(msg, params...)
}
func(l *zapLogger)Errorf(template string, arg ...interface{}){
	l.logger.Errorf(template, arg)
}

func(l *zapLogger)Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKey(cat, sub, extra)
	l.logger.Fatalw(msg, params...)
}
func(l *zapLogger)Fatalf(template string, arg ...interface{}){
	l.logger.Fatalf(template, arg)
}