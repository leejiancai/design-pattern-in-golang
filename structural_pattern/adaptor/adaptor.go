package adaptor

import (
	"fmt"
	"strings"
)

type OldLogger struct {
	name string
}

func (logger *OldLogger) Info(msg string) string {
	return fmt.Sprintf(msg)
}

type NewLoggerInterface interface {
	Info(msg ...string) string
}

type NewLogger struct {
	oldLogger *OldLogger
}

func (logger *NewLogger) Info(msg ...string) string {
	var sb strings.Builder
	for _, i := range msg {
		sb.WriteString(i)
	}
	return logger.oldLogger.Info(sb.String())
}
