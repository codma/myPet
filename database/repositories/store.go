package repositories

import (
	"context"
	"main/constant"
	"main/database"
	"main/ent/store"
	"main/model"
	"time"

	"github.com/spf13/viper"
)

func CreateStore(storeRequest model.AppInstalled) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration(constant.CONTEXT_TIMEOUT_DATABASE)*time.Second)
	tx, err := database.Client.WriteClient.Tx(ctx)

	defer cancel()

	if err != nil {
		return
	}

	_, err = tx.Store.Create().
		SetStoreID(storeRequest.StoreId).
		SetAppAccessKey(storeRequest.AppAccessKey).
		Save(ctx)

	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}

func UpdateStoreState(storeRequest model.AppInstalled, isUsed bool) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration(constant.CONTEXT_TIMEOUT_DATABASE)*time.Second)
	tx, err := database.Client.WriteClient.Tx(ctx)

	defer cancel()

	if err != nil {
		return
	}

	_, err = tx.Store.Update().
		Where(
			store.StoreID(storeRequest.StoreId),
		).
		SetIsUsed(isUsed).
		Save(ctx)

	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}

func DeleteStore(storeRequest model.AppInstalled) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration(constant.CONTEXT_TIMEOUT_DATABASE)*time.Second)
	tx, err := database.Client.WriteClient.Tx(ctx)

	defer cancel()

	if err != nil {
		return
	}

	//앱 엑세스키 제거 밑 isUsed false

	_, err = tx.Store.Update().
		Where(
			store.StoreID(storeRequest.StoreId),
		).
		SetIsUsed(false).
		SetAppAccessKey("").
		Save(ctx)

	//TODO 관련 정보 제거
	//스토어 - 유저 찾아서 유저정보 삭제(소프트딜리트), order 정보 삭제(소프트딜리트)

	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}

func ValidationAppKey(storeName string, appToken string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration(constant.CONTEXT_TIMEOUT_DATABASE)*time.Second)
	tx, err := database.Client.WriteClient.Tx(ctx)

	defer cancel()

	if err != nil {
		return
	}

	isAvailable, err := tx.Store.
		Query().
		Where(store.And(
			store.StoreIDEQ(storeName),
			store.AppAccessKeyEQ(appToken),
			store.IsUsedEQ(true),
		),
		).Exist(ctx)

	if !isAvailable || err != nil {
		return err
	}
	return
}
