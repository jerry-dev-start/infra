package logs

import (
	"fmt"
	"time"
)

// 定义 Banner 字符串
const banner = `
   ______   _____  _   __
  / ____/  /  _/ |/ /  / /
 / / __    / / |   /  / / 
/ /_/ /  _/ / /   |  /_/  
\____/  /___//_/|_| (_)   
                          
 Gin Framework - Custom Base Starter
 Version: v1.0.0
`

func PrintBanner() {
	// 使用颜色转义字符可以让 Banner 看起来更像 SpringBoot (绿色)
	// \033[32m 表示绿色，\033[0m 表示重置颜色
	fmt.Printf("\033[32m%s\033[0m", banner)
	fmt.Printf(" [Started at: %s]\n\n", time.Now().Format("2006-01-02 15:04:05"))
}
