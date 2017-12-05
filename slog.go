package saclogsend

import (
	"fmt"
	"log/syslog"
)

// Slog is a logger instance. All methods support Printf string replacement.
type Slog struct {
	log *syslog.Writer
}

// New creates a new logger and connects it to the remote host
func New(tag, remoteHost string) (logger *Slog, err error) {
	logger = &Slog{}
	l, err := syslog.Dial("udp", remoteHost, syslog.LOG_LOCAL0, tag)
	logger.log = l
	return logger, err
}

// Err writes an error message to the syslog
func (slog *Slog) Err(message string, args ...interface{}) {
	smsg := fmt.Sprintf(message, args...)
	fmt.Println(smsg)
	if slog.log != nil {
		slog.log.Err(smsg)
	}
}

// Info writes an informational message to the syslog
func (slog *Slog) Info(message string, args ...interface{}) {
	smsg := fmt.Sprintf(message, args...)
	fmt.Println(smsg)
	if slog.log != nil {
		slog.log.Info(smsg)
	}
}
