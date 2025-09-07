package gormslog

import (
	"context"
	"log/slog"
	"time"

	"gorm.io/gorm/logger"
)

func NoopLogger(_ context.Context, _ string, _ ...any) {}

type SlogLogFunc = func(context.Context, string, ...any)

type SlogLogger struct {
	log              *slog.Logger
	level            logger.LogLevel
	slowSQLThreshold time.Duration
	info             SlogLogFunc
	warn             SlogLogFunc
	error            SlogLogFunc
}

func New(log *slog.Logger, level logger.LogLevel, slowSQLThreshold time.Duration) *SlogLogger {
	l := &SlogLogger{
		log:              log,
		level:            level,
		slowSQLThreshold: slowSQLThreshold,
		info:             NoopLogger,
		warn:             NoopLogger,
		error:            NoopLogger,
	}
	if l.level >= logger.Info {
		l.info = l.log.InfoContext
	}
	if l.level >= logger.Warn {
		l.warn = l.log.WarnContext
	}
	if l.level >= logger.Error {
		l.error = l.log.ErrorContext
	}
	return l
}

func (l *SlogLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.level = level
	return &newLogger
}

func (l *SlogLogger) Info(ctx context.Context, msg string, data ...any) {
	l.info(ctx, msg, data)
}

func (l *SlogLogger) Warn(ctx context.Context, msg string, data ...any) {
	l.warn(ctx, msg, data)
}

func (l *SlogLogger) Error(ctx context.Context, msg string, data ...any) {
	l.error(ctx, msg, data)
}

func (l *SlogLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.level <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	switch {
	case err != nil && l.level >= logger.Error:
		l.log.ErrorContext(ctx, "SQL error",
			"err", err, "elapsed", elapsed, "rows", rows, "sql", sql)
	case elapsed > l.slowSQLThreshold && l.slowSQLThreshold != 0 && l.level >= logger.Warn:
		l.log.WarnContext(ctx, "Slow SQL",
			"elapsed", elapsed, "rows", rows, "sql", sql)
	case l.level == logger.Info:
		l.log.InfoContext(ctx, "SQL",
			"elapsed", elapsed, "rows", rows, "sql", sql)
	}
}
