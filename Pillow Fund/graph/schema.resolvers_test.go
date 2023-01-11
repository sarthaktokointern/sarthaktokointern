package graph

import (
	"context"
	"github.com/sarthak/Pillow/graph/model"
	"testing"
)

func TestRequestCounter(t *testing.T) {
	type TestCase struct {
		want *model.RequestData
		Err  error
	}

	TestCases := []TestCase{
		{
			want: nil,
			Err:  nil,
		},
	}
	resolver := queryResolver{}
	for _, test := range TestCases {

		got, Er := resolver.RequestCounter(context.Background())
		if Er != nil {
			t.Errorf("got the following error %v", Er)
		} else {
			if got != test.want {
				t.Errorf("got: %v and want %v", got, test.want)
			}
		}

	}

}
