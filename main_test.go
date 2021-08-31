package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"training_tests/mocks"
)

func Test_simple_sum(t *testing.T) {
	ar := new(mocks.GetArgs)

	ar.On("GetA").Return(2)
	ar.On("GetB").Return(4)

	res:= Sum(ar)
	assert.Equal(t, 6, res)
}
func Test_sum(t *testing.T) {

	ar := new(mocks.GetArgs)
	ar.On("GetA").Return(2)
	ar.On("GetB" ).Return(4)

	ar2 := new(mocks.GetArgs)
	ar2.On("GetA").Return(2)
	ar2.On("GetB").Return(-2)

	type args struct {
		arguments GetArgs
	}

	tt := []struct {
		name string
		args args
		want int
	}{
		{
			name: "The first",
			args: args{
				arguments: ar,
			},
			want: 6,
		},
		{
			name: "The second",
			args: args{
				arguments: ar2,
			},
			want: 0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := Sum(tc.args.arguments)
			if got != tc.want {
				t.Errorf("got = %v, wanted = %v", got, tc.want)
				return
			}
			assert.Equal(t, tc.want, got)
		})
	}
}