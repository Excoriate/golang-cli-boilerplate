package cloudaws

import (
	"context"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/o11y"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

type ECSConnector struct {
	Client *ecs.Client
	Ctx    context.Context
	Logger o11y.LoggerInterface
}

type ECSReader interface {
	ListECSServices() (*ecs.ListServicesOutput, error)
	ListECSClusters() (*ecs.ListClustersOutput, error)
	ListTaskDefinitions() (*ecs.ListTaskDefinitionsOutput, error)
}

func NewECSConnector(ctx context.Context, client *ecs.Client, logger o11y.LoggerInterface) *ECSConnector {
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
