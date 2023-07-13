package tool

import "testing"

func TestFloat64ToFloat2(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		args args
		want float64
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
			if got := Float64ToFloat2(tt.args.f); got != tt.want {
				t.Errorf("Float64ToFloat2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetInterfaceToInt(t *testing.T) {
	type args struct {
		t1 interface{}
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
			if got := GetInterfaceToInt(tt.args.t1); got != tt.want {
				t.Errorf("GetInterfaceToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringTOInt(t *testing.T) {
	type args struct {
		number string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"1",
			args{"1"},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringTOInt(tt.args.number); got != tt.want {
				t.Errorf("StringTOInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToFloat(t *testing.T) {
	type args struct {
		number string
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		// TODO: Add test cases.
		{
			"1",
			args{"1.2"},
			1.2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToFloat(tt.args.number); got != tt.want {
				t.Errorf("StringToFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
