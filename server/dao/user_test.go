package dao

import (
	"reflect"
	"server/model"
	"testing"
)

func TestInsertUser(t *testing.T) {
	type args struct {
		user model.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"1",
			args{model.User{1, "1", "2771872337@qq.com"}},
			false,
		},
		{
			"1",
			args{model.User{1, "1", "2771872337"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InsertUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("InsertUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSelectUserByUserName(t *testing.T) {
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
			"1",
			args{"1"},
			model.User{6, "1", "2771872337@qq.com"},
			false,
		},
		{
			"1",
			args{"1"},
			model.User{6, "1", "277187233"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SelectUserByUserName(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectUserByUserName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectUserByUserName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectWinCount(t *testing.T) {
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
			"1",
			args{1},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectWinCount(tt.args.userid); got != tt.want {
				t.Errorf("SelectWinCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateWinCount(t *testing.T) {
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
			if err := UpdateWinCount(tt.args.userid); (err != nil) != tt.wantErr {
				t.Errorf("UpdateWinCount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
