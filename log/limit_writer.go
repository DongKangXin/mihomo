package log

import (
	"os"

	"github.com/metacubex/gvisor/pkg/sync"
)

// 定义 LimitedWriter（放在 main.go 中）
type LimitedWriter struct {
	logFile     *os.File
	logPath     string
	maxSize     int64
	currentSize int64
	autoClear   bool
	mu          sync.Mutex
}

func (lw *LimitedWriter) Write(p []byte) (n int, err error) {
	lw.mu.Lock()
	defer lw.mu.Unlock()

	if lw.currentSize+int64(len(p)) > lw.maxSize {
		if lw.autoClear {
			lw.logFile.Close()
			os.Truncate(lw.logPath, 0)

			file, _ := os.OpenFile(lw.logPath, os.O_WRONLY|os.O_APPEND, 0644)
			lw.logFile = file
			lw.currentSize = 0

			msg := "[WARNING] 日志文件超长已清空\n"
			lw.logFile.WriteString(msg)
			lw.currentSize = int64(len(msg))
		}
	}

	n, err = lw.logFile.Write(p)
	if err == nil {
		lw.currentSize += int64(n)
	}

	return n, err
}
