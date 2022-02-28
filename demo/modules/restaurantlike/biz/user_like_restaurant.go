package rstlikebiz

import (
	"context"
	"demo/component/asyncjob"
	restaurantlikemodel "demo/modules/restaurantlike/model"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

type IncreaseLikeCountStore interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type userLikeRestaurantBiz struct {
	store    UserLikeRestaurantStore
	incStore IncreaseLikeCountStore
}

func NewUserLikeRestaurantBiz(store UserLikeRestaurantStore, incStore IncreaseLikeCountStore) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store: store, incStore: incStore}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.Like,
) error {
	err := biz.store.Create(ctx, data)

	if err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	// side effect
	job := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId)
	})

	_ = asyncjob.NewGroup(true, job).Run(ctx)

	return nil
}
