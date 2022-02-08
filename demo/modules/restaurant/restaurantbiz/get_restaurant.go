package restaurantbiz

import (
	"context"
	"demo/modules/restaurant/restaurantmodel"
	"errors"
)

type GetRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type getRestaurantBiz struct {
	store GetRestaurantStore
}

func NewGetRestaurantBiz(store GetRestaurantStore) *getRestaurantBiz {
	return &getRestaurantBiz{store: store}

}

func (biz *getRestaurantBiz) GetRestaurant(
	ctx context.Context,
	id int) (*restaurantmodel.Restaurant, error) {

	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}

	if data.Status == 0 {
		return nil, errors.New("data is deleted")
	}

	return data, err
}
