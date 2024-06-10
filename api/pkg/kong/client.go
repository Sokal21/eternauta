package kong

import (
	"context"
	"encoding/json"
	"eternauta/pkg"
	"fmt"
)

type KongClient struct {
	client pkg.IClient
}

type CreateServiceRequest struct {
	Name           string `json:"name"`
	Protocol       string `json:"protocol"`
	Path           string `json:"path"`
	ReadTimeout    int    `json:"read_timeout"`
	Retries        int    `json:"retries"`
	Host           string `json:"host"`
	ConnectTimeout int    `json:"connect_timeout"`
	WriteTimeout   int    `json:"write_timeout"`
	Port           int    `json:"port"`
	// Tags null
	// Ca_certificates null
	// Client_certificate null
}

type CreateServiceRespose struct {
	Port           int    `json:"port"`
	Created_at     int    `json:"created_at"`
	Updated_at     int    `json:"updated_at"`
	WriteTimeout   int    `json:"write_timeout"`
	ConnectTimeout int    `json:"connect_timeout"`
	ReadTimeout    int    `json:"read_timeout"`
	Enabled        bool   `json:"enabled"`
	Protocol       string `json:"protocol"`
	Id             string `json:"id"`
	Name           string `json:"name"`
	Path           string `json:"path"`
	Retries        int    `json:"retries"`
	Host           string `json:"host"`
	// client_certificate null
	// tls_verify null
	// ca_certificates null
	// tls_verify_depth null
	// tags null
}

type Service struct {
	Id string `json:"id"`
}

type CreateRouteRequest struct {
	Service                 Service  `json:"service"`
	Name                    string   `json:"name"`
	Hosts                   []string `json:"hosts"`
	RegexPriority           int      `json:"regex_priority"`
	PathHandling            string   `json:"path_handling"`
	StripPath               bool     `json:"strip_path"`
	PreserveHost            bool     `json:"preserve_host"`
	HttpsRedirectStatusCode int      `json:"https_redirect_status_code"`
	Protocols               []string `json:"protocols"`
	Tags                    []string `json:"tags"`
	RequestBuffering        bool     `json:"request_buffering"`
	ResponseBuffering       bool     `json:"response_buffering"`
	// paths                      null
	// snis                       null
	// methods                    null
	// headers                    null
	// sources                    null
	// destinations               null
}

type CreateRouteResponse struct {
	Id string `json:"id"`
}

var _ IKongClient = &KongClient{}

func NewKongClient(client pkg.IClient) *KongClient {
	return &KongClient{
		client: client,
	}
}

func (c *KongClient) CreateService(ctx context.Context, requestBody CreateServiceRequest) (*CreateServiceRespose, error) {
	// if requestBody.Account != nil && requestBody.Alias != nil {
	// 	return nil, fmt.Errorf("get_alias: quering for account and for alias at the same time its not allowed")
	// }
	path := "/services"
	url := c.client.ResolvePath(path)

	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("create_service: failed encoding request: %w", err)
	}

	res, err := c.client.Post(ctx, url, body)
	if err != nil {
		return nil, fmt.Errorf("create_service: failed performing request: %w", err)
	}

	if !res.Successful {
		return nil, fmt.Errorf("create_service: failed creating service status code %d, body %s", res.StatusCode, string(res.Body))
	}

	decodedResponse := &CreateServiceRespose{}
	if err := json.Unmarshal(res.Body, decodedResponse); err != nil {
		return nil, fmt.Errorf("create_service: failed decoding response: %w, body: %s", err, string(res.Body))
	}
	return decodedResponse, nil
}

func (c *KongClient) CreateRoute(ctx context.Context, requestBody CreateRouteRequest) (*CreateRouteResponse, error) {
	// if requestBody.Account != nil && requestBody.Alias != nil {
	// 	return nil, fmt.Errorf("get_alias: quering for account and for alias at the same time its not allowed")
	// }
	path := "/routes"
	url := c.client.ResolvePath(path)

	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("create_route: failed encoding request: %w", err)
	}

	res, err := c.client.Post(ctx, url, body)
	if err != nil {
		return nil, fmt.Errorf("create_route: failed performing request: %w", err)
	}

	if !res.Successful {
		return nil, fmt.Errorf("create_route: failed creating service status code %d, body %s", res.StatusCode, string(res.Body))
	}

	decodedResponse := &CreateRouteResponse{}
	if err := json.Unmarshal(res.Body, decodedResponse); err != nil {
		return nil, fmt.Errorf("create_route: failed decoding response: %w, body: %s", err, string(res.Body))
	}
	return decodedResponse, nil
}
