package logging

import (
	"os"

	"github.com/Stevesadr/golang-backend-project/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var zeroLevelMap map[string]zerolog.Level = map[string]zerolog.Level{
	"debug" : zerolog.DebugLevel,
	"info" : zerolog.InfoLevel,
	"warn" : zerolog.WarnLevel,
	"error" : zerolog.ErrorLevel,
	"fatal" : zerolog.FatalLevel,
}

type zeroLogger struct{
	cfg config.Config
	logger *zerolog.Logger
}

func newZeroLogger(cfg *config.Config) *zeroLogger {
	logger := &zeroLogger{cfg: *cfg}
	logger.Init()
	return logger
}

func(l *zeroLogger)GetLogLevel() zerolog.Level{
	level, exist := zeroLevelMap[l.cfg.Logger.Level]
	if !exist{
		return zerolog.DebugLevel
	}
	return level
}

func(l *zeroLogger)Init(){
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	file, err := os.OpenFile(l.cfg.Logger.FilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil{
		panic("could not open file")
	}

	var logger = zerolog.New(file).With().Timestamp().Str("AppName", "MyApp").Str("LoggerName", "MyApp").Logger() 

	zerolog.SetGlobalLevel(l.GetLogLevel())
	l.logger = &logger
}

func(l *zeroLogger)Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}){
	l.logger.Debug().Str("Category", string(cat)).Str("SubCategory", string(sub)).Fields(extra).Msg(msg)
}
func(l *zeroLogger)Debugf(template string, arg ...interface{}){
	l.logger.Debug().Msgf(template, arg...)
}

func(l *zeroLogger)Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}){
	l.logger.Info().Str("Category", string(cat)).Str("SubCategory", string(sub)).Fields(extra).Msg(msg)
}
func(l *zeroLogger)Infof(template string, arg ...interface{}){
	l.logger.Info().Msgf(template, arg...)
}

func(l *zeroLogger)Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}){
	l.logger.Warn().Str("Category", string(cat)).Str("SubCategory", string(sub)).Fields(extra).Msg(msg)
}
func(l *zeroLogger)Warnf(template string, arg ...interface{}){
	l.logger.Warn().Msgf(template, arg...)
}

func(l *zeroLogger)Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}){
	l.logger.Error().Str("Category", string(cat)).Str("SubCategory", string(sub)).Fields(extra).Msg(msg)
}
func(l *zeroLogger)Errorf(template string, arg ...interface{}){
	l.logger.Error().Msgf(template, arg...)
}

func(l *zeroLogger)Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}){
	l.logger.Fatal().Str("Category", string(cat)).Str("SubCategory", string(sub)).Fields(extra).Msg(msg)
}
func(l *zeroLogger)Fatalf(template string, arg ...interface{}){
	l.logger.Fatal().Msgf(template, arg...)
}