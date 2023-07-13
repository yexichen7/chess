package api

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"server/dao"
	"server/tool"
	"testing"
)

func Test_router(t *testing.T) {

	go func() {
		//先连接数据库
		dao.RUNDB()
		//初始化logger
		tool.InitLogger()
		defer tool.Logger.Sync()
		//接受机器数据的服务
		sev := NewServer("0.0.0.0", 8084)
		//广播ws
		go WsBroadcast()
		//启动tcp服务
		go sev.Start()
		//启动引擎
		RUNENGINE()
	}()

	http.Get("0.0.0.0:8084")

	//tests := []struct {
	//	name     string
	//	Usermail string
	//	expect   string
	//}{
	//	{name: "good case", Usermail: "2771872337@qq.com", expect: "1"},
	//	{"bad case", "", "we need a UserId"},
	//}
	//r := gin.Default()
	//r.POST("/getwincount", getwincount)
	//for _, tt := range tests {
	//	param, _ := json.Marshal(&tt.Usermail)
	//	params := bytes.NewReader(param)
	//	t.Run(tt.name, func(t *testing.T) {
	//		// httptest 这个包是为了mock一个HTTP请求
	//		// 参数分别是请求方法，请求URL，请求Body
	//		// Body只能使用io.Reader
	//		req := httptest.NewRequest(
	//			"GET",          // 请求方法
	//			"/getwincount", // 请求URL
	//			params,         // 请求参数
	//		)
	//
	//		// mock一个响应记录器
	//		w := httptest.NewRecorder()
	//
	//		// 让server端处理mock请求并记录返回的响应内容
	//		r.ServeHTTP(w, req)
	//
	//		// 校验状态码是否符合预期
	//		assert.Equal(t, http.StatusOK, w.Code)
	//
	//		// 解析并检验响应内容是否复合预期
	//		var resp map[string]string
	//		err := json.Unmarshal([]byte(w.Body.String()), &resp)
	//		assert.Nil(t, err)
	//		assert.Equal(t, tt.expect, resp["msg"])
	//	})
	//}
}

func Test_login(t *testing.T) {
	tests := []struct {
		name   string
		mail   string
		expect string
	}{
		{name: "good case", mail: "2771872337@qq.com", expect: ""},
		{name: "good case", mail: "2743212332@qq.com", expect: ""},
		{name: "bad case", mail: "277187qq.com", expect: ""},
	}
	r := gin.Default()
	r.POST("/login", login)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// httptest 这个包是为了mock一个HTTP请求
			// 参数分别是请求方法，请求URL，请求Body
			// Body只能使用io.Reader

			mail, _ := json.Marshal(&tt.mail)
			reader := bytes.NewReader(mail)

			req := httptest.NewRequest(
				"POST",   // 请求方法
				"/login", // 请求URL
				reader,   // 请求参数
			)

			// mock一个响应记录器
			w := httptest.NewRecorder()

			// 让server端处理mock请求并记录返回的响应内容
			r.ServeHTTP(w, req)

			// 校验状态码是否符合预期
			assert.Equal(t, http.StatusOK, w.Code)

			// 解析并检验响应内容是否复合预期
			var resp map[string]string
			err := json.Unmarshal([]byte(w.Body.String()), &resp)
			if err != nil {
			}
		})
	}
}
