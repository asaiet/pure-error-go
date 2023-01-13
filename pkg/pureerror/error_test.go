package pureerror

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

const (
	CodeExample = "example"
)

var generalErr = errors.New("this is general error")

func TestNew(t *testing.T) {
	type args struct {
		code string
		err  error
	}
	tests := []struct {
		name string
		args args
		want PureError
	}{
		{
			name: "normal",
			args: args{
				code: CodeExample,
				err:  nil,
			},
			want: &pureError{
				code:    CodeExample,
				msg:     "",
				wrapped: nil,
			},
		},
		{
			name: "wrapped error",
			args: args{
				code: CodeExample,
				err:  generalErr,
			},
			want: &pureError{
				code:    CodeExample,
				msg:     "",
				wrapped: generalErr,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.code, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pureError_Code(t *testing.T) {
	tests := []struct {
		name string
		err  *pureError
		want string
	}{
		{
			name: "normal",
			err: &pureError{
				code:    CodeExample,
				msg:     "",
				wrapped: generalErr,
			},
			want: CodeExample,
		},
		{
			name: "nil error",
			err:  nil,
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Code(); got != tt.want {
				t.Errorf("Code() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pureError_Error(t *testing.T) {
	tests := []struct {
		name string
		err  *pureError
		want string
	}{
		{
			name: "normal",
			err: &pureError{
				code:    CodeExample,
				msg:     "",
				wrapped: nil,
			},
			want: "example",
		},
		{
			name: "wrapped error",
			err: &pureError{
				code:    CodeExample,
				msg:     "",
				wrapped: generalErr,
			},
			want: "example: this is general error",
		},
		{
			name: "with message",
			err: &pureError{
				code:    CodeExample,
				msg:     "reason",
				wrapped: generalErr,
			},
			want: "example: reason: this is general error",
		},
		{
			name: "nil error",
			err:  nil,
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pureError_Is(t *testing.T) {
	var nilPureError *pureError
	var nilAnotherPureError *pureError

	tests := []struct {
		name   string
		err    error
		target error
		want   bool
	}{
		{
			name:   "normal",
			err:    New(CodeExample, nil),
			target: New(CodeExample, nil),
			want:   true,
		},
		{
			name:   "wrap general error",
			err:    New(CodeExample, generalErr),
			target: generalErr,
			want:   true,
		},
		{
			name:   "wrapped by general error",
			err:    fmt.Errorf("test: %w", New(CodeExample, nil)),
			target: New(CodeExample, nil),
			want:   true,
		},
		{
			name:   "nil error",
			err:    nilPureError,
			target: nilAnotherPureError,
			want:   true,
		},
		{
			name:   "not same error",
			err:    New(CodeExample, nil),
			target: generalErr,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := errors.Is(tt.err, tt.target); tt.want != got {
				t.Errorf("errors.Is(err, target) = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pureError_Unwrap(t *testing.T) {
	tests := []struct {
		name string
		err  *pureError
		want error
	}{
		{
			name: "normal",
			err: &pureError{
				code:    CodeExample,
				msg:     "human fault",
				wrapped: generalErr,
			},
			want: generalErr,
		},
		{
			name: "wrap nil",
			err: &pureError{
				code:    CodeExample,
				msg:     "human fault",
				wrapped: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Unwrap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unwrap() error = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pureError_Why(t *testing.T) {
	err := New(CodeExample, nil).Why("reason")
	want := &pureError{
		code:    CodeExample,
		msg:     "reason",
		wrapped: nil,
	}

	if !reflect.DeepEqual(err, want) {
		t.Errorf("error = %v, want %v", err, want)
	}
}
