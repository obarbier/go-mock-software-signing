package log_utils

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"io"
	"log"
	"os"
	"strings"
)

// TODO(obarbier): 	configuration hierarchy should be file->env
//					currently only have environment

type LogConfig struct {
	LogLevel      string `envconfig:"log_level" split_word:"true" required:"true"`
	LogFileOutput bool   `envconfig:"log_file_output" split_word:"true" default:"false"`
	LogFile       string `envconfig:"log_file" split_word:"true" default:"core.log"`
	LogOutput     string
}

// logUtils -.
type logUtils struct {
	logger *zerolog.Logger
}

var l *logUtils

func init() {
	// Trying to do a singleton pattern. never have more than 1 l
	if l == nil {
		// TODO(obarbier): add some file vs local
		l = newLog()
	}
}

func newLog() *logUtils {
	var nl zerolog.Level
	var s LogConfig
	err := envconfig.Process("myapp", &s)
	if err != nil {
		// TODO(obarbier): maybe panic
		log.Fatal(err.Error())
	}
	// FIXME(obarbier): not populating the data
	switch strings.ToLower(s.LogLevel) {
	case "trace":
		nl = zerolog.TraceLevel
	case "error":
		nl = zerolog.ErrorLevel
	case "warn":
		nl = zerolog.WarnLevel
	case "info":
		nl = zerolog.InfoLevel
	case "debug":
		nl = zerolog.DebugLevel
	default:
		nl = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(nl)

	var stout io.Writer
	switch strings.ToLower(s.LogOutput) {
	default:
		stout = os.Stdout
	}
	// TODO(obarbier):	Understand CallerWithSkipFrameCount
	//					skipFrameCount := 3
	//					zerolog.New(stout).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()

	var writers []io.Writer
	writers = append(writers, zerolog.ConsoleWriter{Out: stout})
	if s.LogFileOutput {
		file, err := os.OpenFile(s.LogFile, os.O_RDWR|os.O_CREATE, 0777)
		if err != nil {
			log.Fatalf("Error: %s", err)
			os.Exit(1)
		}
		// FIXME(obarbier): not writing to file
		writers = append(writers, file)
	}
	logger := zerolog.New(io.MultiWriter(writers...)).With().Timestamp().Logger()

	return &logUtils{
		logger: &logger,
	}
}

// Tace -.
func Trace(message interface{}, args ...interface{}) {
	// FIXME(obarbier): Fields is not working
	l.logger.Trace().Fields(args).Msg(message.(string))
}

// Debug -.
func Debug(message interface{}, args ...interface{}) {
	l.logger.Debug().Fields(args).Msg(message.(string))
}

// Info -.
func Info(message string, args ...interface{}) {
	l.logger.Info().Fields(args).Msg(message)
}

// Warn -.
func Warn(message string, args ...interface{}) {
	l.logger.Warn().Fields(args).Msg(message)
}

// Error -.
func Error(message interface{}, args ...interface{}) {
	if l.logger.GetLevel() == zerolog.DebugLevel {
		l.logger.Debug().Fields(args).Msg(message.(string))
	}
	l.logger.Error().Fields(args).Msg(message.(string))
}

// Fatal -.
func Fatal(message interface{}, args ...interface{}) {
	l.logger.Fatal().Fields(args).Msg(message.(string))
	os.Exit(1)
}

func LogAny() func(string, ...interface{}) {
	return func(s string, i ...interface{}) {
		l.logger.Debug().Fields(i).Msg(fmt.Sprintf(s, i...))
	}
}
