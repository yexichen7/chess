package api

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_checkroom(t *testing.T) {
	tests := []struct {
		name   string
		roomId string
		expect string
	}{
		{name: "good case", roomId: "1", expect: ""},
		{name: "good case", roomId: "2", expect: ""},
		{},
	}
	r := gin.Default()
	r.POST("/checkroom", checkroom)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// httptest 这个包是为了mock一个HTTP请求
			// 参数分别是请求方法，请求URL，请求Body
			// Body只能使用io.Reader

			mail, _ := json.Marshal(&tt.roomId)
			reader := bytes.NewReader(mail)

			req := httptest.NewRequest(
				"POST",       // 请求方法
				"/checkroom", // 请求URL
				reader,       // 请求参数
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

func Test_enterroom(t *testing.T) {
	tests := []struct {
		name   string
		roomId string
		expect string
	}{
		{name: "good case", roomId: "1", expect: ""},
		{name: "good case", roomId: "2", expect: ""},
		{},
	}
	r := gin.Default()
	r.POST("/enterroom", enterroom)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// httptest 这个包是为了mock一个HTTP请求
			// 参数分别是请求方法，请求URL，请求Body
			// Body只能使用io.Reader

			mail, _ := json.Marshal(&tt.roomId)
			reader := bytes.NewReader(mail)

			req := httptest.NewRequest(
				"POST",       // 请求方法
				"/enterroom", // 请求URL
				reader,       // 请求参数
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
