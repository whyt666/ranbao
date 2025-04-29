package middleware

import (
	"fmt"
	"log"

	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// FormLogger 打印表单数据的中间件
func FormLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 继续处理请求
		c.Next()

		//start := time.Now()

		//解析前先初始化日志数据
		logData := make(map[string]interface{})

		// 1. 解析表单（兼容不同Content-Type）
		contentType := c.ContentType()
		switch {
		case strings.HasPrefix(contentType, "multipart/form-data"):
			// 处理文件上传，内存限制设为32MB
			if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
				log.Printf("解析multipart表单错误: %v", err)
			}
		default:
			// 普通表单解析
			if err := c.Request.ParseForm(); err != nil {
				log.Printf("解析表单错误: %v", err)
			}
		}

		// 2. 收集URL查询参数
		queryParams := make(map[string]string)
		for k, v := range c.Request.URL.Query() {
			queryParams[k] = strings.Join(v, ",")
		}
		if len(queryParams) > 0 {
			logData["query"] = queryParams
		}

		// 3. 收集POST表单字段
		formData := make(map[string]interface{})
		fmt.Println(c.Request.Form)
		if c.Request.Form != nil {
			for k, v := range c.Request.Form {
				if len(v) == 1 {
					formData[k] = v[0]
				} else {
					formData[k] = v
				}
			}
		}
		if len(formData) > 0 {
			logData["form"] = formData
		}

		// 4. 收集上传文件信息
		if c.Request.MultipartForm != nil && c.Request.MultipartForm.File != nil {
			fileInfo := make(map[string][]string)
			for key, files := range c.Request.MultipartForm.File {
				var names []string
				for _, f := range files {
					names = append(names, fmt.Sprintf("%s (size: %d)", f.Filename, f.Size))
				}
				fileInfo[key] = names
			}
			if len(fileInfo) > 0 {
				logData["files"] = fileInfo
			}
		}

		// 记录日志
		log.Printf("[FormLog] %v | %s | %s | %s | Data: %+v",
			time.Now().Format("2006/01/02 - 15:04:05"),
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL.Path,
			logData,
		)

		// 记录处理耗时（可选）
		//log.Printf("[FormLog] Completed in %v", time.Since(start))
	}
}
