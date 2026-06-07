package rating

import (
	"context"
	"errors"

	model "movieexample.com/rating/pkg"
)

var ErrNotFound = errors.New("ratings not found for a record")

type RatingRepository interface {
	Get(ctx context.Context, recordID model.RecordID, recordType model.RecordType) ([]model.Rating, error)
	Put(ctx context.Context, recodID model.RecordID, recordType model.RecordType, rating *model.Rating) error
}

type Controller struct {
	repo RatingRepository
}

func (c *Controller) New(r RatingRepository) *Controller {
	return &Controller{r}
}

func (c *Controller) GetAggregatedRating(ctx context.Context, recordID model.RecordID, recordType model.RecordType) (float64, error) {
	ratings, err := c.repo.Get(ctx, recordID, recordType)
	if err != nil && err.Error() == "ratings not found for a record" {
		return 0, ErrNotFound
	} else if err != nil {
		return 0, err
	}
	sum := float64(0)
	for _, val := range ratings {
		sum += float64(val.Value)
	}
	return sum / float64(len(ratings)), nil

}

func (c *Controller) PutRating(ctx context.Context, recordID model.RecordID, recordType model.RecordType, rating *model.Rating) error {
	return c.PutRating(ctx, recordID, recordType, rating)
}
