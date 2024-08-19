package biz

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is user not found.
)

// GreeterRepo is a Greater repo.
type MapRepo interface {
	Save(context.Context, *Greeter) (*Greeter, error)
	Update(context.Context, *Greeter) (*Greeter, error)
	FindByID(context.Context, int64) (*Greeter, error)
	ListByHello(context.Context, string) ([]*Greeter, error)
	ListAll(context.Context) ([]*Greeter, error)
}

// GreeterUsecase is a Greeter usecase.
type MapUsecase struct {
	repo MapRepo
	log  *log.Helper
}

type DirectionDrivingResp struct {
	Status string `json:"status"`
	Info   string `json:"info"`
	Count  string `json:"count"`
	Route  struct {
		Origin      string `json:"origin"`
		Destination string `json:"destination"`
		TaxiCost    string `json:"taxi_cost"`
		Paths       []struct {
			Distance string `json:"distance"`
			Duration string `json:"duration"`
		}
	} `json:"route"`
}

// NewGreeterUsecase new a Greeter usecase.
func NewMapUsecase(repo MapRepo, logger log.Logger) *MapUsecase {
	return &MapUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (msbiz *MapUsecase) GetDrivingInfo(origin, destination string) (string, string, error) {
	key := "da7032dd11069b885a4273fa68ec101f"
	api := "https://restapi.amap.com/v3/direction/driving"
	params := fmt.Sprintf("key=%s&origin=%s&destination=%s&extensions=base&output=JSON", key, origin, destination)
	url := fmt.Sprintf("%s?%s", api, params)
	resp, err := http.Get(url)
	if err != nil {
		return "", "", err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("http request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	ddResp := &DirectionDrivingResp{}
	err = json.Unmarshal(body, ddResp)
	if err != nil {
		return "", "", err
	}

	if ddResp.Status != "1" {
		return "", "", fmt.Errorf("ddResp.Status != 1, %s", ddResp.Info)
	}

	fmt.Println("花费：" + ddResp.Route.TaxiCost + "元")

	return ddResp.Route.Paths[0].Distance, ddResp.Route.Paths[0].Duration, nil
}
