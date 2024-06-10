package kong

import (
	"context"
)

type IKongClient interface {
	CreateService(ctx context.Context, requestBody CreateServiceRequest) (*CreateServiceRespose, error)
	CreateRoute(ctx context.Context, requestBody CreateRouteRequest) (*CreateRouteResponse, error)
}
