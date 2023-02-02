package repositories

import (
	"context"
	"errors"
	"main/constant"
	"main/database"
	"main/ent"
	"main/ent/user"
	"main/model"
	"time"

	"github.com/spf13/viper"
)

func CreateUser(userRequest model.User, storeId string) (memberId uint64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration(constant.CONTEXT_TIMEOUT_DATABASE)*time.Second)
	tx, err := database.Client.WriteClient.Tx(ctx)

	defer cancel()

	if err != nil {
		return
	}

	user, err := tx.User.
		Query().
		Where(user.And(
			user.UserIDEQ(userRequest.UserId),
			user.StoreID(storeId),
			user.IsUsedEQ(true),
		),
		).Exist(ctx)

	if err != nil {
		return
	}

	if user {
		err = errors.New("id is not available")
		return
	}

	newUser, err := tx.User.Create().
		SetUserID(userRequest.UserId).
		SetUserName(userRequest.UserName).
		SetBirthday(userRequest.Birthday).
		SetWeight(userRequest.Weight).
		SetType(userRequest.Type).
		SetStoreID(storeId).
		Save(ctx)

	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	memberId = newUser.ID
	return
}

func GetUser(userName string, storeId string) (result *ent.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration(constant.CONTEXT_TIMEOUT_DATABASE)*time.Second)
	tx, err := database.Client.WriteClient.Tx(ctx)

	defer cancel()

	if err != nil {
		return
	}

	result, err = tx.User.
		Query().
		Where(user.And(
			user.UserNameEQ(userName),
			user.StoreID(storeId),
			user.IsUsedEQ(true),
		),
		).First(ctx)

	return

}
