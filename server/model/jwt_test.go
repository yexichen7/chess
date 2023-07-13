package model

import (
	"github.com/dgrijalva/jwt-go"
	"reflect"
	"testing"
)

func TestGenToken(t *testing.T) {
	type args struct {
		user User
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"1",
			args{User{1, "1", "2771872337@qq.com"}},
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MSwiZXhwIjoxNjkxNzE3MTc4LCJpc3MiOiJjaGVzcyJ9.J-BbrUecMmUhySQVF51IcevvMabGIzU0tqIU8xhT9Dk",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenToken(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseToken(t *testing.T) {
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		args    args
		want    *MyClaims
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "1",
			args:    args{"1"},
			want:    &MyClaims{1, jwt.StandardClaims{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseToken(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
