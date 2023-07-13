package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendMail(t *testing.T) {
	type args struct {
		mails []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"good",
			args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendMail(tt.args.mails); err != nil {
				t.Errorf("InsertUser() error = %v, wantErr %v", err, false)
			}
		})
	}
}

func Test_authHandler(t *testing.T) {
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
	r.POST("/auth", authHandler)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// httptest 这个包是为了mock一个HTTP请求
			// 参数分别是请求方法，请求URL，请求Body
			// Body只能使用io.Reader

			req := httptest.NewRequest(
				"POST",  // 请求方法
				"/auth", // 请求URL
				nil,     // 请求参数
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
