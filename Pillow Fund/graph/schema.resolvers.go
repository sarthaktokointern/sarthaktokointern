package graph

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/sarthak/Pillow/graph/model"
	"gopkg.in/gcfg.v1"
)

type config struct {
	Config struct {
		XWaitTill  int
		XRateLimit int
	}
}

var (
	RequestCnt = 0
	cfg        config
	FinishTime time.Time
	StartTime  time.Time
	Response   *model.RequestData
	Status     = 429
	Message    = "Too Many Requests!!"
	XWaitTill  = ""
)

// RequestCounter is the resolver for the RequestCounter field.
func (r *queryResolver) RequestCounter(context.Context) (*model.RequestData, error) {
	err := gcfg.ReadFileInto(&cfg, "/Users/sarthaksrivastava/GolandProjects/sarthaktokointern/Pillow Fund/config.ini")

	if err != nil {
		return &model.RequestData{}, fmt.Errorf("failed to parse config file %v", err)
	}

	return UpdateRequestCounter()
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

func UpdateRequestCounter() (*model.RequestData, error) {
	var wg sync.WaitGroup
	var m sync.Mutex
	wg.Add(1)
	go func() {
		defer wg.Done()
		m.Lock()
		if RequestCnt == cfg.Config.XRateLimit {
			if FinishTime.IsZero() {

				StartTime = time.Now()

				FinishTime = StartTime.Add(time.Duration(cfg.Config.XWaitTill) * time.Second)

				XWaitTill = FinishTime.String()

			} else if time.Now().After(FinishTime) {
				RequestCnt = 1
				StartTime = time.Time{}
				FinishTime = time.Time{}
				Response = &model.RequestData{
					Data:       &model.Data{NumberOfRequests: &RequestCnt},
					Extensions: nil,
				}

			} else {
				Response = &model.RequestData{
					Extensions: &model.Extensions{
						Status:     &Status,
						Message:    &Message,
						XWaitTill:  &XWaitTill,
						XRateLimit: &cfg.Config.XRateLimit,
					},
					Data: nil,
				}
			}

		} else {

			RequestCnt++
			Response = &model.RequestData{
				Data:       &model.Data{NumberOfRequests: &RequestCnt},
				Extensions: nil,
			}

		}
		m.Unlock()
	}()

	return Response, nil
}
