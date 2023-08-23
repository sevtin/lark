package cr_wallet

import (
	"lark/domain/cache"
	"lark/domain/pdo"
	"lark/domain/repo"
	"lark/pkg/constant"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_wallet"
)

func GetAccountInfoByWalletId(walletCache cache.WalletCache, walletRepo repo.WalletRepository, walletId int64) (info *pb_wallet.AccountInfo, err error) {
	info, err = walletCache.GetAccountInfoByWalletId(walletId)
	if info.WalletId > 0 {
		return
	}
	var (
		q           = entity.NewMysqlQuery()
		accountInfo *pdo.AccountInfo
	)
	q.SetFilter("wallet_id=?", walletId)
	accountInfo, err = walletRepo.GetAccountInfo(q)
	if err != nil {
		return
	}
	info.WalletId = accountInfo.WalletId
	info.WalletType = pb_enum.WALLET_TYPE(accountInfo.WalletType)
	info.Uid = accountInfo.Uid
	info.Balance = accountInfo.Balance
	info.FrozenAmount = accountInfo.FrozenAmount
	info.Status = pb_enum.WALLET_STATUS(accountInfo.Status)
	walletCache.SetAccountInfo(info, constant.WALLET_KEY_TYPE_WALLET_ID)
	return
}

func GetAccountInfo(walletCache cache.WalletCache, walletRepo repo.WalletRepository, uid int64, walletType pb_enum.WALLET_TYPE) (info *pb_wallet.AccountInfo, err error) {
	info, err = walletCache.GetAccountInfo(uid, walletType)
	if info.WalletId > 0 {
		return
	}
	var (
		q           = entity.NewMysqlQuery()
		accountInfo *pdo.AccountInfo
	)
	q.SetFilter("uid=?", uid)
	q.SetFilter("wallet_type=?", walletType)
	accountInfo, err = walletRepo.GetAccountInfo(q)
	if err != nil {
		return
	}
	info.WalletId = accountInfo.WalletId
	info.WalletType = pb_enum.WALLET_TYPE(accountInfo.WalletType)
	info.Uid = accountInfo.Uid
	info.Balance = accountInfo.Balance
	info.FrozenAmount = accountInfo.FrozenAmount
	info.Status = pb_enum.WALLET_STATUS(accountInfo.Status)
	return
	//walletCache.SetAccountInfo(info, constant.WALLET_KEY_TYPE_UID_WALLET_TYPE)
}
