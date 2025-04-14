package util

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/obutora/open-api-mini-layerd-architecture-template/internal/config"
)

// Logger はアプリケーションのロガーを表します
type Logger struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	DebugLogger *log.Logger
}

// NewLogger は新しいLoggerインスタンスを作成します
func NewLogger(cfg *config.Config) (*Logger, error) {
	var infoWriter, errorWriter, debugWriter io.Writer

	// ログファイルが設定されている場合はファイルに出力
	if cfg.Log.File != "" {
		// ディレクトリが存在しない場合は作成
		dir := filepath.Dir(cfg.Log.File)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, err
		}

		// ログファイルを開く
		file, err := os.OpenFile(cfg.Log.File, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}

		// 標準出力とファイルの両方に出力
		infoWriter = io.MultiWriter(os.Stdout, file)
		errorWriter = io.MultiWriter(os.Stderr, file)
		debugWriter = io.MultiWriter(os.Stdout, file)
	} else {
		// ファイルが設定されていない場合は標準出力のみ
		infoWriter = os.Stdout
		errorWriter = os.Stderr
		debugWriter = os.Stdout
	}

	// ログレベルに応じてデバッグログを無効化
	if cfg.Log.Level != "debug" {
		debugWriter = io.Discard
	}

	return &Logger{
		InfoLogger:  log.New(infoWriter, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		ErrorLogger: log.New(errorWriter, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		DebugLogger: log.New(debugWriter, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
	}, nil
}

// Info はINFOレベルのログを出力します
func (l *Logger) Info(format string, v ...interface{}) {
	l.InfoLogger.Printf(format, v...)
}

// Error はERRORレベルのログを出力します
func (l *Logger) Error(format string, v ...interface{}) {
	l.ErrorLogger.Printf(format, v...)
}

// Debug はDEBUGレベルのログを出力します
func (l *Logger) Debug(format string, v ...interface{}) {
	l.DebugLogger.Printf(format, v...)
}

// DefaultLogger はデフォルトのロガーを返します
func DefaultLogger() *Logger {
	return &Logger{
		InfoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		ErrorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		DebugLogger: log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
