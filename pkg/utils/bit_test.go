package utils

import "testing"

func TestBitsToByte(t *testing.T) {
	type args struct {
		bits []int
	}
	tests := []struct {
		name    string
		args    args
		want    byte
		wantErr bool
	}{
		{"[0 0 0 0 0 0 0 0] --> 0", args{[]int{0, 0, 0, 0, 0, 0, 0, 0}}, 0, false},
		{"[1 0 0 0 0 0 0 0] --> 0", args{[]int{1, 0, 0, 0, 0, 0, 0, 0}}, 128, false},
		{"[1 1 1 1 1 1 1 1] --> 0", args{[]int{1, 1, 1, 1, 1, 1, 1, 1}}, 255, false},
		{"[0 0 0 0 0 0 0] --> 0", args{[]int{0, 0, 0, 0, 0, 0, 0}}, 0, false},
		{"[1 0 0 0 0 0 0] --> 0", args{[]int{1, 0, 0, 0, 0, 0, 0}}, 128, false},
		{"[1 1 1 1 1 1 1] --> 0", args{[]int{1, 1, 1, 1, 1, 1, 1}}, 254, false},
		{"[0 0 0 0 0 0 0 0 0] --> 0", args{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0}}, 0, true},
		{"[1 0 0 0 0 0 0 0 0] --> 0", args{[]int{1, 0, 0, 0, 0, 0, 0, 0, 0}}, 128, true},
		{"[1 1 1 1 1 1 1 1 0] --> 0", args{[]int{1, 1, 1, 1, 1, 1, 1, 1, 0}}, 255, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BitsToByte(tt.args.bits)
			if (err != nil) != tt.wantErr {
				t.Errorf("BitsToByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BitsToByte() = %v, want %v", got, tt.want)
			}
		})
	}
}
