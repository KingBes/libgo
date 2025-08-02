package main

/*
#include <stdlib.h>
#include <string.h>
*/
import "C"
import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// 进度条结构体
type ProgressBar struct {
	total   int64   // 总字节数
	current int64   // 当前已下载字节数
	percent float64 // 当前百分比
}

func NewProgressBar(total int64) *ProgressBar {
	return &ProgressBar{total: total}
}

// 更新进度并显示
func (p *ProgressBar) Update(read int64) {
	p.current += read
	p.percent = float64(p.current) / float64(p.total) * 100

	// 构建进度条
	barWidth := 50
	filled := int(float64(barWidth) * p.percent / 100)
	empty := barWidth - filled
	bar := "[" + strings.Repeat("=", filled) + strings.Repeat(" ", empty) + "]"

	// 格式化显示
	fmt.Printf("\r%s %.2f%% %s/%s ",
		bar,
		p.percent,
		formatBytes(p.current),
		formatBytes(p.total))
}

//export download_file
func download_file(c_url, c_filepath *C.char) *C.char {
	url := C.GoString(c_url)
	filepath := C.GoString(c_filepath)
	// 获取文件大小
	resp, err := http.Head(url)
	if err != nil {
		return C.CString(fmt.Sprintf("HEAD请求失败: %v", err))
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return C.CString(fmt.Sprintf("服务器返回错误状态: %d", resp.StatusCode))
	}

	size, err := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
	if err != nil {
		return C.CString(fmt.Sprintf("无法获取文件大小: %v", err))
	}

	// 创建本地文件
	file, err := os.Create(filepath)
	if err != nil {
		return C.CString(fmt.Sprintf("文件创建失败: %v", err))
	}
	defer file.Close()

	// 开始下载
	resp, err = http.Get(url)
	if err != nil {
		return C.CString(fmt.Sprintf("下载失败: %v", err))
	}
	defer resp.Body.Close()

	progressBar := NewProgressBar(size)
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	done := make(chan bool)
	go func() {
		// 显示初始进度条
		progressBar.Update(0)
		for {
			select {
			case <-ticker.C:
				progressBar.Update(0) // 仅触发显示刷新
			case <-done:
				return
			}
		}
	}()

	// 使用 TeeReader 同时写入文件并更新进度
	reader := io.TeeReader(resp.Body, &ProgressWriter{progressBar})
	_, err = io.Copy(file, reader)
	close(done) // 停止进度刷新

	// 最终完成显示
	fmt.Println() // 换行结束进度条
	return C.CString("")
}

// 自定义 Writer 用于更新进度
type ProgressWriter struct {
	pb *ProgressBar
}

func (pw *ProgressWriter) Write(p []byte) (int, error) {
	n := len(p)
	pw.pb.Update(int64(n))
	return n, nil
}

// 格式化字节显示
func formatBytes(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}
