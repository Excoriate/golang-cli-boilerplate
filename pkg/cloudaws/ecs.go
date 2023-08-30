package cloudaws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"go.uber.org/zap"
)

type ECSConnector struct {
	Client *ecs.Client
	Ctx    context.Context
	Logger *zap.Logger
}

type ECSReader interface {
	ListECSServices() (*ecs.ListServicesOutput, error)
	ListECSClusters() (*ecs.ListClustersOutput, error)
	ListTaskDefinitions() (*ecs.ListTaskDefinitionsOutput, error)
}

func NewECSConnector(ctx context.Context, client *ecs.Client, logger *zap.Logger) *ECSConnector {
	return &ECSConnector{
		Client: client,
		Logger: logger,
		Ctx:    ctx,
	}
}

func (e *ECSConnector) ListECSServices() (*ecs.ListServicesOutput, error) {
	return e.Client.ListServices(e.Ctx, &ecs.ListServicesInput{})
}

func (e *ECSConnector) ListECSClusters() (*ecs.ListClustersOutput, error) {
	return e.Client.ListClusters(e.Ctx, &ecs.ListClustersInput{})
}

func (e *ECSConnector) ListTaskDefinitions() (*ecs.ListTaskDefinitionsOutput, error) {
	return e.Client.ListTaskDefinitions(e.Ctx, &ecs.ListTaskDefinitionsInput{})
}
