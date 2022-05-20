package stack

import (
	"reflect"
	"testing"
)

func TestNewStack(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want *Stack
	}{
		{"Empty Stack", args{size: 0}, &Stack{size: 0, arr: make([]any, 0), top: 0}},
		{"Stack with 10", args{size: 10}, &Stack{size: 10, arr: make([]any, 10), top: 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStack(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	type fields struct {
		arr  []any
		size int
		top  int
	}
	type args struct {
		e any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Push No Error",
			fields: fields{
				arr: make([]any, 5),
				size: 5,
				top : 0,
			},
			args: args{e: 1},
			wantErr: false,
		},

		{
			name: "Push Error",
			fields: fields{
				arr: []any{1,2,3,4,5},
				size: 5,
				top : 5,
			},
			args: args{e: 1},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				arr:  tt.fields.arr,
				size: tt.fields.size,
				top:  tt.fields.top,
			}
			err := s.Push(tt.args.e);
			if  (err != nil) != tt.wantErr {
				t.Errorf("Stack.Push() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	type fields struct {
		arr  []any
		size int
		top  int
	}
	tests := []struct {
		name    string
		fields  fields
		want    any
		wantErr bool
	}{
		{
			name: "Pop Error",
			fields: fields{
				arr: make([]any, 5),
				size: 5,
				top : 0,
			},
			want: nil,
			wantErr: true,
		},

		{
			name: "Pop No Error",
			fields: fields{
				arr: []any{1,2,3,4,5},
				size: 5,
				top : 5,
			},
			want: 5,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				arr:  tt.fields.arr,
				size: tt.fields.size,
				top:  tt.fields.top,
			}
			got, err := s.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("Stack.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}