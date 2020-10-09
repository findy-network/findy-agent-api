package resolver

import (
	"context"
	"fmt"

	"github.com/findy-network/findy-agent-api/graph/model"
)

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
