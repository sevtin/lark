package repo

import (
	"gorm.io/gorm"
	"lark/domain/pdo"
	"lark/domain/po"
	"lark/pkg/common/xmysql"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_wallet"
)

type WalletRepository interface {
	TxCreateWallets(tx *gorm.DB, wallets []*po.Wallet) (err error)
	TxChangeWalletBalance(tx *gorm.DB, q *entity.MysqlUpdate) (rowsAffected int64)
	GetAccountInfo(q *entity.MysqlQuery) (info *pdo.AccountInfo, err error)
	GetAccountBalance(q *entity.MysqlQuery) (balance *pdo.AccountBalance, err error)
	UserWallets(q *entity.MysqlQuery) (wallets []*pb_wallet.WalletInfo, err error)
}

type walletRepository struct {
}

func NewWalletRepository() WalletRepository {
	return &walletRepository{}
}

func (r *walletRepository) TxCreateWallets(tx *gorm.DB, wallets []*po.Wallet) (err error) {
	return tx.Create(wallets).Error
}

func (r *walletRepository) TxChangeWalletBalance(tx *gorm.DB, q *entity.MysqlUpdate) (rowsAffected int64) {
	var w = new(po.Wallet)
	rowsAffected = tx.Model(w).Where(q.Query, q.Args...).Updates(q.Values).RowsAffected
	return
}

func (r *walletRepository) GetAccountInfo(q *entity.MysqlQuery) (info *pdo.AccountInfo, err error) {
	info = new(pdo.AccountInfo)
	db := xmysql.GetDB()
	err = db.Model(new(po.Wallet)).Select(info.GetFields()).Where(q.Query, q.Args...).Find(info).Error
	return
}

func (r *walletRepository) GetAccountBalance(q *entity.MysqlQuery) (balance *pdo.AccountBalance, err error) {
	balance = new(pdo.AccountBalance)
	db := xmysql.GetDB()
	err = db.Model(new(po.Wallet)).Select(balance.GetFields()).Where(q.Query, q.Args...).Find(balance).Error
	return
}

func (r *walletRepository) UserWallets(q *entity.MysqlQuery) (wallets []*pb_wallet.WalletInfo, err error) {
	db := xmysql.GetDB()
	err = db.Model(new(po.Wallet)).Where(q.Query, q.Args...).Find(&wallets).Error
	return
}
