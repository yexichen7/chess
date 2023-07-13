package service

import (
	"reflect"
	"server/model"
	"testing"
)

func TestAddWinCount(t *testing.T) {
	type args struct {
		userid int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"1",
			args{1},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddWinCount(tt.args.userid); (err != nil) != tt.wantErr {
				t.Errorf("AddWinCount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUserInfo(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    model.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"good",
			args{"1"},
			model.User{1, "1", "2771872337@qq.com"},
			false,
		},
		{
			"bad",
			args{""},
			model.User{1, "1", "2771872337@qq.com"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserInfo(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUserNameAndMailRight(t *testing.T) {
	type args struct {
		Userrname string
		Userrmail string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"good",
			args{"1", "2771872337@qq.com"},
			true,
			false,
		},
		{
			"bad",
			args{"", "2771872337"},
			false,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsUserNameAndMailRight(tt.args.Userrname, tt.args.Userrmail)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsUserNameAndMailRight() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsUserNameAndMailRight() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUserNameExist(t *testing.T) {
	type args struct {
		Username string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"good",
			args{"1"},
			true,
			false,
		},
		{
			"bad",
			args{""},
			false,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsUserNameExist(tt.args.Username)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsUserNameExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsUserNameExist() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUser(t *testing.T) {
	type args struct {
		User model.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"good",
			args{model.User{1, "1", "2771872337@qq.com"}},
			false,
		},
		{
			"good",
			args{model.User{-2, "", "2771872q.com"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NewUser(tt.args.User); (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSearchWinCount(t *testing.T) {
	type args struct {
		userid int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"good",
			args{1},
			1,
		},
		{
			"good",
			args{-1},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchWinCount(tt.args.userid); got != tt.want {
				t.Errorf("SearchWinCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
