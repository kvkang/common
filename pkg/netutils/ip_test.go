package netutils

import (
	"net"
	"reflect"
	"testing"
)

func TestIPv4MaskEnd(t *testing.T) {
	type args struct {
		ip   net.IP
		mask net.IPMask
	}
	tests := []struct {
		name string
		args args
		want net.IP
	}{
		{"0.0.0.0/0", args{net.ParseIP("0.0.0.0"), net.CIDRMask(0, 32)}, net.ParseIP("255.255.255.255")},
		{"0.0.0.0/1", args{net.ParseIP("0.0.0.0"), net.CIDRMask(1, 32)}, net.ParseIP("127.255.255.255")},
		{"0.0.0.0/2", args{net.ParseIP("0.0.0.0"), net.CIDRMask(2, 32)}, net.ParseIP("63.255.255.255")},
		{"0.0.0.0/3", args{net.ParseIP("0.0.0.0"), net.CIDRMask(3, 32)}, net.ParseIP("31.255.255.255")},
		{"0.0.0.0/4", args{net.ParseIP("0.0.0.0"), net.CIDRMask(4, 32)}, net.ParseIP("15.255.255.255")},
		{"0.0.0.0/5", args{net.ParseIP("0.0.0.0"), net.CIDRMask(5, 32)}, net.ParseIP("7.255.255.255")},
		{"0.0.0.0/6", args{net.ParseIP("0.0.0.0"), net.CIDRMask(6, 32)}, net.ParseIP("3.255.255.255")},
		{"0.0.0.0/7", args{net.ParseIP("0.0.0.0"), net.CIDRMask(7, 32)}, net.ParseIP("1.255.255.255")},
		{"0.0.0.0/8", args{net.ParseIP("0.0.0.0"), net.CIDRMask(8, 32)}, net.ParseIP("0.255.255.255")},
		{"0.0.0.0/9", args{net.ParseIP("0.0.0.0"), net.CIDRMask(9, 32)}, net.ParseIP("0.127.255.255")},
		{"0.0.0.0/10", args{net.ParseIP("0.0.0.0"), net.CIDRMask(10, 32)}, net.ParseIP("0.63.255.255")},
		{"0.0.0.0/11", args{net.ParseIP("0.0.0.0"), net.CIDRMask(11, 32)}, net.ParseIP("0.31.255.255")},
		{"0.0.0.0/12", args{net.ParseIP("0.0.0.0"), net.CIDRMask(12, 32)}, net.ParseIP("0.15.255.255")},
		{"0.0.0.0/13", args{net.ParseIP("0.0.0.0"), net.CIDRMask(13, 32)}, net.ParseIP("0.7.255.255")},
		{"0.0.0.0/14", args{net.ParseIP("0.0.0.0"), net.CIDRMask(14, 32)}, net.ParseIP("0.3.255.255")},
		{"0.0.0.0/15", args{net.ParseIP("0.0.0.0"), net.CIDRMask(15, 32)}, net.ParseIP("0.1.255.255")},
		{"0.0.0.0/16", args{net.ParseIP("0.0.0.0"), net.CIDRMask(16, 32)}, net.ParseIP("0.0.255.255")},
		{"0.0.0.0/17", args{net.ParseIP("0.0.0.0"), net.CIDRMask(17, 32)}, net.ParseIP("0.0.127.255")},
		{"0.0.0.0/18", args{net.ParseIP("0.0.0.0"), net.CIDRMask(18, 32)}, net.ParseIP("0.0.63.255")},
		{"0.0.0.0/19", args{net.ParseIP("0.0.0.0"), net.CIDRMask(19, 32)}, net.ParseIP("0.0.31.255")},
		{"0.0.0.0/20", args{net.ParseIP("0.0.0.0"), net.CIDRMask(20, 32)}, net.ParseIP("0.0.15.255")},
		{"0.0.0.0/21", args{net.ParseIP("0.0.0.0"), net.CIDRMask(21, 32)}, net.ParseIP("0.0.7.255")},
		{"0.0.0.0/22", args{net.ParseIP("0.0.0.0"), net.CIDRMask(22, 32)}, net.ParseIP("0.0.3.255")},
		{"0.0.0.0/23", args{net.ParseIP("0.0.0.0"), net.CIDRMask(23, 32)}, net.ParseIP("0.0.1.255")},
		{"0.0.0.0/24", args{net.ParseIP("0.0.0.0"), net.CIDRMask(24, 32)}, net.ParseIP("0.0.0.255")},
		{"0.0.0.0/25", args{net.ParseIP("0.0.0.0"), net.CIDRMask(25, 32)}, net.ParseIP("0.0.0.127")},
		{"0.0.0.0/26", args{net.ParseIP("0.0.0.0"), net.CIDRMask(26, 32)}, net.ParseIP("0.0.0.63")},
		{"0.0.0.0/27", args{net.ParseIP("0.0.0.0"), net.CIDRMask(27, 32)}, net.ParseIP("0.0.0.31")},
		{"0.0.0.0/28", args{net.ParseIP("0.0.0.0"), net.CIDRMask(28, 32)}, net.ParseIP("0.0.0.15")},
		{"0.0.0.0/29", args{net.ParseIP("0.0.0.0"), net.CIDRMask(29, 32)}, net.ParseIP("0.0.0.7")},
		{"0.0.0.0/30", args{net.ParseIP("0.0.0.0"), net.CIDRMask(30, 32)}, net.ParseIP("0.0.0.3")},
		{"0.0.0.0/31", args{net.ParseIP("0.0.0.0"), net.CIDRMask(31, 32)}, net.ParseIP("0.0.0.1")},
		{"0.0.0.0/32", args{net.ParseIP("0.0.0.0"), net.CIDRMask(32, 32)}, net.ParseIP("0.0.0.0")},

		//		{"0.0.0.0/17", args{net.ParseIP("0.0.0.0"), net.CIDRMask(17, 32)}, net.ParseIP("0.1.255.255.255")},
		//		{"0.0.0.0/18", args{net.ParseIP("0.0.0.0"), net.CIDRMask(18, 32)}, net.ParseIP("0.0.255.255.255")},
		//		{"0.0.0.0/19", args{net.ParseIP("0.0.0.0"), net.CIDRMask(19, 32)}, net.ParseIP("0.0.255.255.255")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IPv4MaskEnd(tt.args.ip, tt.args.mask); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IPv4MaskEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIPv4Add(t *testing.T) {
	type args struct {
		ip  net.IP
		add int
	}
	tests := []struct {
		name    string
		args    args
		wantGot net.IP
		wantErr bool
	}{
		{"0.0.0.0 + 1", args{net.ParseIP("0.0.0.0"), 1}, net.ParseIP("0.0.0.1"), false},
		{"0.0.0.255 + 1", args{net.ParseIP("0.0.0.255"), 1}, net.ParseIP("0.0.1.0"), false},
		{"255.255.255.255 + 1", args{net.ParseIP("255.255.255.255"), 1}, net.ParseIP("0.0.0.0"), true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGot, err := IPv4Add(tt.args.ip, tt.args.add)
			if (err != nil) != tt.wantErr {
				t.Errorf("IPv4Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotGot, tt.wantGot) {
				t.Errorf("IPv4Add() = %v, want %v", gotGot, tt.wantGot)
			}
		})
	}
}
