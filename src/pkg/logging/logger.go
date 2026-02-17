package logging

import "github.com/Stevesadr/golang-backend-project/config"

type Logger interface{
	Init()

	Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Debugf(template string, arg interface{})

	Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Infof(template string, arg interface{})

	Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Warnf(template string, arg interface{})

	Error(err error,cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Errorf(template string, arg interface{})	

	Fatal(err error,cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Fatalf(template string, arg interface{})
}

func NewLogger(cfg config.Config) Logger {
	return nil
}