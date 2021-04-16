package os

import (
	"errors"
	"os"
	"strconv"
)

// SetListeningPort is to set listening port
func SetListeningPort(port string) error { // 设置服务监听端口
	if len(port) == 0 {
		return errors.New("port is nil")
	}
	value, err := strconv.Atoi(port)
	if err != nil {
		return err
	}
	if value < 1 || value > 65535 { // port 越界
		return errors.New("out of port")
	}
	os.Setenv("PORT", port) // 设置端口
	return nil
}
