package handler_redis

import (
	"context"
	"fmt"
	"pc-shop-final-project/domain/entity"
	sqlConn "pc-shop-final-project/internal/config/database/mysql"
	redisConn "pc-shop-final-project/internal/config/database/redis"
	"pc-shop-final-project/internal/repository/mysql"
	"pc-shop-final-project/internal/repository/redis"
)

var (
	mysqlConn        = sqlConn.InitMysqlDB()
	repoUser         = mysql.NewUserMysql(mysqlConn)
	repoCust         = mysql.NewCustomerMysql(mysqlConn)
	repoSetP         = mysql.NewSettlePurchaseMysql(mysqlConn)
	redisConnect     = redisConn.InitRedisClient()
	HandlerShopRedis = redis.NewRepoShopRedis(redisConnect, repoUser, repoCust, repoSetP)
)

func GetAllRedis(ctx context.Context) (*entity.User, *entity.Customer, []*entity.SettlePurchase) {
	u, c, i, err := HandlerShopRedis.GetAllRedis(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, nil, nil
	}
	return u, c, i
}

func GetUserRedis(ctx context.Context) *entity.User {
	u, err := HandlerShopRedis.GetUserRedis(ctx)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return u
}

func GetUserCustomerRedis(ctx context.Context) (*entity.User, *entity.Customer) {
	u, c, err := HandlerShopRedis.GetUserCustomerRedis(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	return u, c
}

func SetUserRedis(ctx context.Context, u *entity.User) {
	HandlerShopRedis.StoreUser(ctx, u)
}

func SetCustomerRedis(ctx context.Context, c *entity.Customer) {
	HandlerShopRedis.StoreCustomer(ctx, c)
}

func SetSetItRedis(ctx context.Context, i []*entity.SettlePurchase) {
	HandlerShopRedis.StoreSetIt(ctx, i)
}
