package migrate

import (
	"api/database/db_gorm"
	"api/model"
	"api/util/log"
)

func Migrate() {
	log.Error(db_gorm.GetDB().AutoMigrate(&model.Admins{}))
	log.Error(db_gorm.GetDB().AutoMigrate(&model.Users{}))
	log.Error(db_gorm.GetDB().AutoMigrate(&model.UsersWallets{}))
	log.Error(db_gorm.GetDB().AutoMigrate(&model.VerifyCodes{}))
	log.Error(db_gorm.GetDB().AutoMigrate(&model.LogWallets{}))
	log.Error(db_gorm.GetDB().AutoMigrate(&model.InvitationLists{}))
	log.Error(db_gorm.GetDB().AutoMigrate(&model.ServiceFiles{}))
	log.Error(db_gorm.GetDB().AutoMigrate(&model.UsersRealNames{}))
	log.Error(db_gorm.GetDB().AutoMigrate(&model.Coins{}))
	log.Error(db_gorm.GetDB().AutoMigrate(&model.TransformMarkets{}))
	log.Error(db_gorm.GetDB().AutoMigrate(&model.Orders{}))
	log.Error(db_gorm.GetDB().AutoMigrate(&model.OrdersTransactions{}))
	log.Error(db_gorm.GetDB().AutoMigrate(&model.ReceiveMethods{}))
	log.Error(db_gorm.GetDB().AutoMigrate(&model.UsersReceiveMethod{}))
	log.Error(db_gorm.GetDB().AutoMigrate(&model.FreezeWallets{}))
	log.Error(db_gorm.GetDB().AutoMigrate(&model.WithdrawMarkets{}))
	log.Error(db_gorm.GetDB().AutoMigrate(&model.WithdrawCoins{}))
	log.Error(db_gorm.GetDB().AutoMigrate(&model.OrdersTransactionsComplaints{}))
	log.Error(db_gorm.GetDB().AutoMigrate(&model.LogComplaintCompromises{}))
	Seed()
}
