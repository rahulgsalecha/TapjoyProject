package main

import (
	"reflect"
	"testing"
)

func Test_checkExternalSubstringsValidity(t *testing.T) {
	resultBools1 := []bool{true, true, true}
	resultBools2 := []bool{true, false, true}
	resultBools3 := []bool{false, false, false}

	type args struct {
		results []bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test_checkExternalSubstringsValidity_true_case_1",
			args: args{
				results: resultBools1,
			},
			want: true,
		},
		{
			name: "Test_checkExternalSubstringsValidity_true_case_2",
			args: args{
				results: resultBools2,
			},
			want: true,
		},
		{
			name: "Test_checkExternalSubstringsValidity_false_case_1",
			args: args{
				results: resultBools3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkExternalSubstringsValidity(tt.args.results); got != tt.want {
				t.Errorf("checkExternalSubstringsValidity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkInternalSubstringsValidity(t *testing.T) {
	resultBools1 := []bool{true, true, true}
	resultBools2 := []bool{true, false, true}
	resultBools3 := []bool{false, false, false}

	type args struct {
		results []bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test_checkInternalSubstringsValidity_true_case_1",
			args: args{
				results: resultBools1,
			},
			want: true,
		},
		{
			name: "Test_checkInternalSubstringsValidity_false_case_1",
			args: args{
				results: resultBools2,
			},
			want: false,
		},
		{
			name: "Test_checkInternalSubstringsValidity_false_case_2",
			args: args{
				results: resultBools3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInternalSubstringsValidity(tt.args.results); got != tt.want {
				t.Errorf("checkInternalSubstringsValidity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidSubstring(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test_isValidSubstring_case1_true",
			args: args{
				str: "rttr",
			},
			want: true,
		},
		{
			name: "Test_isValidSubstring_case2_false",
			args: args{
				str: "rmtr",
			},
			want: false,
		},
		{
			name: "Test_isValidSubstring_case3_false",
			args: args{
				str: "bbbb",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidSubstring(tt.args.str); got != tt.want {
				t.Errorf("isValidSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lineValidator(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test_lineValidator_case1_true",
			args: args{
				line: "rttr[mnop]qrst",
			},
			want: true,
		},
		{
			name: "Test_lineValidator_case2_false",
			args: args{
				line: "efgh[baab]ommo",
			},
			want: false,
		},
		{
			name: "Test_lineValidator_case3_false",
			args: args{
				line: "bbbb[qwer]ereq",
			},
			want: false,
		},
		{
			name: "Test_lineValidator_case4_true",
			args: args{
				line: "irttrj[asdfgh]zxcvbns",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lineValidator(tt.args.line); got != tt.want {
				t.Errorf("lineValidator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitLineToSubstrings(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test_splitLineToSubstrings",
			args: args{
				line: "irttrj[asdfgh]zxcvbns",
			},
			want: []string{"irttrj", "asdfgh", "zxcvbns"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitLineToSubstrings(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitLineToSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countAndPrintValidLines(t *testing.T) {
	type args struct {
		inputFile string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test_countAndPrintValidLines_case1_validFile",
			args: args{
				inputFile: "../data/sample.txt",
			},
			wantErr: false,
		},
		{
			name: "Test_countAndPrintValidLines_case2_validFile",
			args: args{
				inputFile: "../data/input.txt",
			},
			wantErr: false,
		},
		{
			name: "Test_countAndPrintValidLines_case3_InvalidFile",
			args: args{
				inputFile: "../data/input1.txt",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := countAndPrintValidLines(tt.args.inputFile); (err != nil) != tt.wantErr {
				t.Errorf("countAndPrintValidLines() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}