package logger

import (
	"os"
	"path/filepath"

	"github.com/arfanxn/welding/internal/infrastructure/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

// NewLoggerFromConfig creates a new logger that writes to both console and file
// config contains the configuration including the log file path
func NewLoggerFromConfig(cfg *config.Config) (*Logger, error) {
	// Create log directory if it doesn't exist
	logDir := filepath.Dir(cfg.LogFilepath)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, err
	}

	// Create or open the log file
	logFile, err := os.OpenFile(
		cfg.LogFilepath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		return nil, err
	}

	// Configure encoder
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Create file write syncer
	fileSyncer := zapcore.AddSync(logFile)

	// Create core for file logging
	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(fileSyncer),
		zap.InfoLevel,
	)

	// Create core for console logging
	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.AddSync(os.Stdout),
		zap.InfoLevel,
	)

	// Combine cores
	core := zapcore.NewTee(
		fileCore,
		consoleCore,
	)

	// Create the logger with a core that writes to both file and console.
	// The logger is configured with:
	// - core: The combined logging core that handles both file and console output
	// - zap.AddCaller(): Automatically adds the calling function's file and line number to each log entry
	// - zap.AddStacktrace(zapcore.ErrorLevel): Adds stack traces for error-level logs and above
	lgr := &Logger{
		Logger: zap.New(
			core,
			zap.AddCaller(),
			zap.AddStacktrace(zapcore.ErrorLevel),
		),
	}

	// Ensure logs are written on program exit
	defer lgr.Sync()

	return lgr, nil
}
