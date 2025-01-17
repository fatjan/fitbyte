package duck

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/fatjan/fitbyte/internal/dto"
	"github.com/fatjan/fitbyte/internal/repositories/duck"
)

type useCase struct {
	duckRepo duck.Repository
}

func NewDuckUsecase(duckRepo duck.Repository) Usecase {
	return &useCase{
		duckRepo: duckRepo,
	}
}

func (u *useCase) GetDucks(ctx context.Context) ([]*dto.Duck, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	ducks, err := u.duckRepo.GetDucks(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*dto.Duck, len(ducks))
	for i, duck := range ducks {
		result[i] = &dto.Duck{
			ID:   fmt.Sprint(duck.ID),
			Name: duck.Name,
		}
	}

	return result, nil
}

func (u *useCase) GetDuckByID(ctx context.Context, id string) (*dto.Duck, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %v", err)
	}

	duck, err := u.duckRepo.GetDuckByID(ctx, idInt)
	if err != nil {
		return nil, err
	}

	return &dto.Duck{
		ID:   strconv.Itoa(duck.ID),
		Name: duck.Name,
	}, nil
}
