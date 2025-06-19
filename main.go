package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

const context = ""

func main() {
	// 设置文件服务器路由
	//http.HandleFunc(fmt.Sprintf("/%s/", context), enableCORS(streamFileHandler))
	//// 设置用户数据API路由
	//http.HandleFunc("/api/users", enableCORS(usersHandler))

	http.HandleFunc("/license", enableCORS(licenseHandler))
	http.HandleFunc("/license/", enableCORS(licenseHandler))

	// 创建 HTTP 服务器
	httpServer := &http.Server{
		Addr:    ":80",
		Handler: nil, // 使用默认多路复用器
	}

	// 创建 HTTPS 服务器
	httpsServer := &http.Server{
		Addr:    ":443",
		Handler: nil, // 使用默认多路复用器
	}

	// 启动 HTTP 服务器（生产环境可改为重定向到 HTTPS）
	go func() {
		log.Println("HTTP 服务器启动，监听端口:80")
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("启动 HTTP 服务器失败: %v", err)
		}
	}()

	// 启动 HTTPS 服务器
	go func() {
		log.Println("HTTPS 服务器启动，监听端口:443")
		if err := httpsServer.ListenAndServeTLS("cert.pem", "key.pem"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("启动 HTTPS 服务器失败: %v", err)
		}
	}()

	// 优雅关闭
	waitForShutdown(httpServer, httpsServer)

	// 启动服务器
	//port := "80"
	//fmt.Printf("服务器运行在端口 %s 上...\n", port)
	//pid := os.Getpid()
	//fmt.Printf("当前进程的ID是: %d\n", pid)
	//fmt.Println(fmt.Sprintf("访问 http://localhost/%s/", context))
	//log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func waitForShutdown(servers ...*http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// 等待中断信号
	<-interruptChan

	// 优雅关闭所有服务器
	for _, s := range servers {
		log.Printf("正在关闭服务器: %s", s.Addr)
		if err := s.Shutdown(nil); err != nil {
			log.Printf("关闭服务器失败: %v", err)
		}
	}

	log.Println("所有服务器已优雅关闭")
}

type licenseCapacity struct {
	Used int `json:"used"`
	Max  int `json:"max"`
}

var capacity = &licenseCapacity{
	Used: 0,
	Max:  100,
}

type EncryptLic struct {
	EncryptInfo string `json:"encrypt_info"`
}

// RequestInfo 封装HTTP请求的关键信息
type RequestInfo struct {
	IP            string `json:"ip"`
	Domain        string `json:"domain"`
	AppName       string `json:"appName"`
	ReportID      string `json:"reportID"`
	ReportVersion string `json:"reportVersion"`
	Certificate   string `json:"certificate"`
}

var (
	licStr = `{"VERSION": "11.0",
  "MACADDRESS": "fa:16:3e:29:5a:82",
  "SERIALNUMBER": null,
  "REGTOOL_SERIAL": "",
  "MAX_NUMBER": "1",
  "DEADLINE": 4904159797000,
  "LISTEN_PORT": null,
  "APPNAME": "report",
  "APPCONTENT": "",
  "UUID": "84434b73942118e89e4b65261435907aae1be1e16f12d2e70b15a42232ec9a33a3b4076736ed576a",
  "FS_USER": "0",
  "MOBILE_FS_USER": "0",
  "PROJECTNAME": "全椒县智慧家庭医生综合管理系统",
  "COMPANYNAME": "安徽晶奇网络科技有限公司",
  "CONCURRENCY": "0",
  "FUNCTION": "101155137800229",
  "PLUGIN": [],
  "KEY": "",
  "REPORTLETSCOUNT": "",
  "MUTICONNECTION": "1",
  "OFFICIAL": "true",
  "TYPE": "FILE",
  "LASTTIME": "0",
  "SESSIONID": "",
  "MAX_NODE": null,
  "REVISION": null,
  "DATABASE_TYPE": [""],
  "REGION": "",
  "CREATE_TIME": null,
  "END_USER": "全椒县智慧家庭医生综合管理系统"
}`
)

func getLic(custom bool) EncryptLic {
	// 读取原始授权文件内容
	buf := bytes.Buffer{}
	data, _ := os.ReadFile("FanRuan.lic")
	// 自定义授权
	if custom {
		data, _ = os.ReadFile("FanRuan3.lic")
	}
	buf.Write(data)
	// 转换为 Base64 编码
	encryptLic := EncryptLic{
		EncryptInfo: base64.StdEncoding.EncodeToString(buf.Bytes()),
	}
	return encryptLic
}

func licenseHandler(writer http.ResponseWriter, request *http.Request) {
	var (
		statusCode   = http.StatusOK
		info         RequestInfo
		responseBody []byte
	)

	// 获取请求路径
	path := request.URL.Path
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&info); err == nil {
		log.Printf("licenseHandler: %v", info)
	}
	// 处理/license路径
	if path == "/license" {
		encryptLic := getLic(true)
		responseBody, _ = json.Marshal(encryptLic)
		if capacity.Used < capacity.Max {
			capacity.Used++
		}
		log.Printf("max: %d, used: %d", capacity.Max, capacity.Used)
	} else if strings.HasPrefix(path, "/license/") {
		// 处理/license/下的子路径
		subPath := path[len("/license/"):]
		switch subPath {
		case "capacity":
			responseBody, _ = json.Marshal(capacity)
		case "deactivate":
			if capacity.Used > 0 {
				capacity.Used--
			}
		case "information":
			statusCode = http.StatusNotFound
		default:
			statusCode = http.StatusInternalServerError
		}
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	if responseBody != nil {
		_, _ = writer.Write(responseBody)
	}
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

type Result struct {
	Success bool   `json:"success"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    []any  `json:"data,omitempty"`
	Total   int    `json:"total,omitempty"`
	Summary any    `json:"summary,omitempty"`
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
func generateMockUsers(count int) Result {
	result := Result{
		Success: true,
		Total:   count,
	}
	rand.Seed(time.Now().UnixNano())

	names := []string{
		"张三", "李四", "王五", "赵六", "钱七", "孙八", "周九", "吴十",
		"小明", "小红", "小刚", "小丽", "小强", "小美", "小亮", "小华",
	}

	domains := []string{"gmail.com", "yahoo.com", "hotmail.com", "example.com"}

	var users []any

	for i := 1; i <= count; i++ {
		name := names[rand.Intn(len(names))]
		email := fmt.Sprintf("%s%d@%s", name, i, domains[rand.Intn(len(domains))])
		age := rand.Intn(40) + 18 // 18-57岁

		// 创建时间在过去1年内
		createdAgo := time.Duration(rand.Intn(365*24)) * time.Hour
		createdAt := time.Now().Add(-createdAgo).Format(time.RFC3339)

		users = append(users, User{
			ID:        i,
			Name:      name,
			Email:     email,
			Age:       age,
			CreatedAt: createdAt,
		})
	}

	result.Data = users

	return result
}
