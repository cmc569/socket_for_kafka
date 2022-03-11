package migrate

import (
	"api/database/db_gorm"
	"api/model"
)

func Seed() {
	createCoin()
	createTransFormMarket()
	createReceiveMethod()
	createWithdrawMarket()
}

func createReceiveMethod() {
	payMethod1 := map[string]interface{}{}
	db_gorm.GetDB().Model(&model.ReceiveMethods{}).Where("id = ?", 1).Find(&payMethod1)
	if len(payMethod1) == 0 {
		payMethod := []model.ReceiveMethods{
			{
				Id:   1,
				Name: "銀行轉帳",
			},
		}
		db_gorm.GetDB().Create(&payMethod)
	}
}

func createTransFormMarket() {
	coin1 := map[string]interface{}{}
	db_gorm.GetDB().Model(&model.TransformMarkets{}).Where("from_coins_id = ?", 1).Find(&coin1)

	if coin1["id"] == nil {
		transFormMarket := []model.TransformMarkets{
			{
				FromCoinsId: 1, ToCoinsId: 1, ExchangeRate: 1, FeeType: 1,
			},
			{
				FromCoinsId: 1, ToCoinsId: 2, ExchangeRate: 1, FeeType: 1,
			},
			{
				FromCoinsId: 1, ToCoinsId: 3, ExchangeRate: 1, FeeType: 1,
			},
		}
		db_gorm.GetDB().Create(&transFormMarket)
	}

	coin2 := map[string]interface{}{}
	db_gorm.GetDB().Model(&model.TransformMarkets{}).Where("from_coins_id = ?", 2).Find(&coin2)
	if coin2["id"] == nil {
		transFormMarket := []model.TransformMarkets{
			{
				FromCoinsId: 2, ToCoinsId: 1, ExchangeRate: 1, FeeType: 1,
			},
			{
				FromCoinsId: 2, ToCoinsId: 2, ExchangeRate: 1, FeeType: 1,
			},
			{
				FromCoinsId: 2, ToCoinsId: 3, ExchangeRate: 1, FeeType: 1,
			},
		}
		db_gorm.GetDB().Create(&transFormMarket)
	}

	coin3 := map[string]interface{}{}
	db_gorm.GetDB().Model(&model.TransformMarkets{}).Where("from_coins_id = ?", 3).Find(&coin3)
	if coin3["id"] == nil {
		transFormMarket := []model.TransformMarkets{
			{
				FromCoinsId: 3, ToCoinsId: 1, ExchangeRate: 1, FeeType: 1,
			},
			{
				FromCoinsId: 3, ToCoinsId: 2, ExchangeRate: 1, FeeType: 1,
			},
			{
				FromCoinsId: 3, ToCoinsId: 3, ExchangeRate: 1, FeeType: 1,
			},
		}
		db_gorm.GetDB().Create(&transFormMarket)
	}

}

func createCoin() {
	coin1 := map[string]interface{}{}
	db_gorm.GetDB().Model(&model.Coins{}).Where("id = ?", 1).Find(&coin1)
	if len(coin1) == 0 {
		var coins model.Coins
		coins.Name = "USDT"
		coins.CoinType = "ETH"
		db_gorm.GetDB().Create(&coins)
	}

	coin2 := map[string]interface{}{}
	db_gorm.GetDB().Model(&model.Coins{}).Where("id = ?", 2).Find(&coin2)
	if len(coin2) == 0 {
		var coins model.Coins
		coins.Name = "ETH"
		coins.CoinType = "ETH"
		db_gorm.GetDB().Create(&coins)
	}

	coin3 := map[string]interface{}{}
	db_gorm.GetDB().Model(&model.Coins{}).Where("id = ?", 3).Find(&coin3)
	if len(coin3) == 0 {
		var coins model.Coins
		coins.Name = "BTC"
		coins.CoinType = "BTC"
		db_gorm.GetDB().Create(&coins)
	}
}

func createWithdrawMarket() {
	coin := map[string]interface{}{}
	db_gorm.GetDB().Model(&model.WithdrawMarkets{}).Where("coins_id = ?", 1).Find(&coin)
	if coin["id"] == nil {
		transFormMarket := []model.WithdrawMarkets{
			{
				CoinsId: 1, FeePercent: 0, FeeFix: 0, MinAmount: 1, MaxAmount: 100, FeeType: 1,
			},
			{
				CoinsId: 2, FeePercent: 0, FeeFix: 0, MinAmount: 1, MaxAmount: 100, FeeType: 1,
			},
			{
				CoinsId: 3, FeePercent: 0, FeeFix: 0, MinAmount: 1, MaxAmount: 100, FeeType: 1,
			},
			{
				CoinsId: 4, FeePercent: 0, FeeFix: 0, MinAmount: 1, MaxAmount: 100, FeeType: 1,
			},
		}
		db_gorm.GetDB().Create(&transFormMarket)
	}

}
