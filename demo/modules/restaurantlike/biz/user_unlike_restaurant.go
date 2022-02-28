package rstlikebiz

import (
	"context"
	"demo/common"
	"demo/component/asyncjob"
	restaurantlikemodel "demo/modules/restaurantlike/model"
)

type UserUnlikeRestaurantStore interface {
	Delete(ctx context.Context, userId, restaurantId int) error
}

type DecreaseLikeCountStore interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

type userUnlikeRestaurantBiz struct {
	store    UserUnlikeRestaurantStore
	decStore DecreaseLikeCountStore
}

func NewUserUnlikeRestaurantBiz(store UserUnlikeRestaurantStore, decStore DecreaseLikeCountStore) *userUnlikeRestaurantBiz {
	return &userUnlikeRestaurantBiz{store: store, decStore: decStore}
}

func (biz *userUnlikeRestaurantBiz) UnlikeRestaurant(
	ctx context.Context,
	userId,
	restaurantId int,
) error {
	err := biz.store.Delete(ctx, userId, restaurantId)

	if err != nil {
		return restaurantlikemodel.ErrCannotUnlikeRestaurant(err)
	}

	// side effect
	go func() {
		defer common.AppRecover()
		job := asyncjob.NewJob(func(ctx context.Context) error {
			return biz.decStore.DecreaseLikeCount(ctx, restaurantId)
		})

		//job.SetRetryDurations([]time.Duration{time.Second * 3})

		_ = asyncjob.NewGroup(true, job).Run(ctx)
	}()

	return nil
}
