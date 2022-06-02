package xlog_test

import (
	"ddd/domain/xlog"
	"ddd/infrastructure/logfmt"
	"testing"
)

func TestLog(t *testing.T) {
	xlog.Log = new(logfmt.Log)
	xlog.Debug("a\n")
}
