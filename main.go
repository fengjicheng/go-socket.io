package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const contetx = "spreadjs"

func main() {
	// 设置文件服务器路由
	http.HandleFunc(fmt.Sprintf("/%s/", contetx), enableCORS(streamFileHandler))
	// 设置用户数据API路由
	http.HandleFunc("/api/users", enableCORS(usersHandler))

	// 启动服务器
	port := "80"
	fmt.Printf("服务器运行在端口 %s 上...\n", port)
	fmt.Println(fmt.Sprintf("访问 http://localhost/%s/**", contetx))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

// enableCORS 中间件添加CORS头
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 允许所有域名进行跨域调用
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// 允许的请求方法
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		// 允许的请求头
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// 处理预检请求
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// 继续处理请求
		next(w, r)
	}
}

// streamFileHandler 处理文件流请求
func streamFileHandler(w http.ResponseWriter, r *http.Request) {
	// 获取请求的文件路径
	filePath := filepath.Join("./spread", r.URL.Path[len("/spreadjs/"):])

	// 检查文件是否存在
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		http.Error(w, "文件不存在", http.StatusNotFound)
		return
	}

	// 检查是否是目录
	if fileInfo.IsDir() {
		http.Error(w, "请求的是目录，不是文件", http.StatusBadRequest)
		return
	}

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "无法打开文件", http.StatusInternalServerError)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("关闭文件时出错: %v", err)
		}
	}(file)

	// 设置响应头
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filepath.Base(filePath)))
	w.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

	// 使用io.Copy将文件内容流式传输到响应
	_, err = io.Copy(w, file)
	if err != nil {
		log.Printf("文件传输错误: %v", err)
		http.Error(w, "文件传输过程中发生错误", http.StatusInternalServerError)
		return
	}
}

// User 表示用户数据结构
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	CreatedAt string `json:"created_at"`
}

// usersHandler 处理用户列表请求
func usersHandler(w http.ResponseWriter, r *http.Request) {
	// 设置响应头
	w.Header().Set("Content-Type", "application/json")

	// 返回用户列表
	json.NewEncoder(w).Encode(generateMockUsers(20))
}

// generateMockUsers 生成模拟用户数据
func generateMockUsers(count int) []User {
	rand.Seed(time.Now().UnixNano())

	names := []string{
		"张三", "李四", "王五", "赵六", "钱七", "孙八", "周九", "吴十",
		"小明", "小红", "小刚", "小丽", "小强", "小美", "小亮", "小华",
	}

	domains := []string{"gmail.com", "yahoo.com", "hotmail.com", "example.com"}

	var result []User

	for i := 1; i <= count; i++ {
		name := names[rand.Intn(len(names))]
		email := fmt.Sprintf("%s%d@%s", name, i, domains[rand.Intn(len(domains))])
		age := rand.Intn(40) + 18 // 18-57岁

		// 创建时间在过去1年内
		createdAgo := time.Duration(rand.Intn(365*24)) * time.Hour
		createdAt := time.Now().Add(-createdAgo).Format(time.RFC3339)

		result = append(result, User{
			ID:        i,
			Name:      name,
			Email:     email,
			Age:       age,
			CreatedAt: createdAt,
		})
	}

	return result
}
