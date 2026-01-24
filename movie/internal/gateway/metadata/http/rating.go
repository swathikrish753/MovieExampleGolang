package http

import (
	"context"
	"fmt"
	"net/http"

	"encoding/json"

	"movieexample.com/movie/internal/gateway"
	model "movieexample.com/rating/pkg"
)

type RGateway struct {
	addr string
}

func NewRgate(addr string) *RGateway {
	return &RGateway{addr}
}

func (g *RGateway) GetAggregatedRating(ctx context.Context, recordID model.RecordID, recordType model.RecordType) (float64, error) {
	req, err := http.NewRequest(http.MethodGet, g.addr+"/rating", nil)
	if err != nil {
		return 0, err
	}
	req = req.WithContext(ctx)
	values := req.URL.Query()
	values.Add("id", string(recordID))
	values.Add("type", string(recordType))
	req.URL.RawQuery = values.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return 0, gateway.ErrNotFound
	}
	if resp.StatusCode/100 != 2 {
		return 0, fmt.Errorf("non-2xx response %v", resp)
	}
	var v float64
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return 0, err
	}
	return v, nil

}
