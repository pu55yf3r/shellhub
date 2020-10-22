package nsadm

import (
	"context"

	"github.com/shellhub-io/shellhub/api/store"
	"github.com/shellhub-io/shellhub/pkg/api/paginator"
	"github.com/shellhub-io/shellhub/pkg/models"
)

type Service interface {
	ListNamespaces(ctx context.Context, pagination paginator.Query) ([]models.Namespaces, int, error)
	CreateNamespace(ctx context.Context, namespace *models.Namespace) (*models.Namespace, error)
}

type service struct {
}

func NewService(store store.Store) Service {
	return &service{store}
}

func (s *service) ListNamespaces(ctx context.Context, pagination paginator.Query) ([]models.Namespaces, int, error) {
	return nil, 0, nil
}

func (s *service) CreateNamespace(ctx context.Context, namespace *models.Namespace) (*models.Namespace, error) {
	return nil, nil
}
