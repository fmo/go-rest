package logger

type Logger interface {
    Debut(msg string, args ...any)
    Info(msg string, args ...any)
    Warn(msg string, args ...any)
}
