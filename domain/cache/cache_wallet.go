package cache

import (
	"lark/pkg/common/xredis"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_wallet"
	"lark/pkg/utils"
	"strconv"
)

type WalletCache interface {
	GetAccountInfoByWalletId(walletId int64) (info *pb_wallet.AccountInfo, err error)
	GetAccountInfo(uid int64, walletType pb_enum.WALLET_TYPE) (info *pb_wallet.AccountInfo, err error)
	SetAccountInfo(info *pb_wallet.AccountInfo, keyType int) (err error)
	DeleteWallet(walletId int64, uid int64, walletType pb_enum.WALLET_TYPE) (err error)
}

type walletCache struct {
}

func NewWalletCache() WalletCache {
	return &walletCache{}
}

func (c *walletCache) GetAccountInfoByWalletId(walletId int64) (info *pb_wallet.AccountInfo, err error) {
	info = new(pb_wallet.AccountInfo)
	var (
		key = constant.RK_SYNC_WALLET_ACCOUNT_INFO + utils.GetHashTagKey(walletId)
	)
	err = Get(key, info)
	return
}

func (c *walletCache) GetAccountInfo(uid int64, walletType pb_enum.WALLET_TYPE) (info *pb_wallet.AccountInfo, err error) {
	info = new(pb_wallet.AccountInfo)
	var (
		key = constant.RK_SYNC_WALLET_ACCOUNT_INFO + utils.GetHashTagKey(uid) + "-" + strconv.Itoa(int(walletType))
	)
	err = Get(key, info)
	return
}

func (c *walletCache) SetAccountInfo(info *pb_wallet.AccountInfo, keyType int) (err error) {
	var (
		key string
	)
	switch keyType {
	case constant.WALLET_KEY_TYPE_WALLET_ID:
		key = constant.RK_SYNC_WALLET_ACCOUNT_INFO + utils.GetHashTagKey(info.WalletId)
	case constant.WALLET_KEY_TYPE_UID_WALLET_TYPE:
		key = constant.RK_SYNC_WALLET_ACCOUNT_INFO + utils.GetHashTagKey(info.Uid) + "-" + strconv.Itoa(int(info.WalletType))
	}
	err = Set(key, info, constant.CONST_DURATION_WALLET_ACCOUNT_INFO_SECOND)
	return
}

func (c *walletCache) DeleteWallet(walletId int64, uid int64, walletType pb_enum.WALLET_TYPE) (err error) {
	var (
		key1 = constant.RK_SYNC_WALLET_ACCOUNT_INFO + utils.GetHashTagKey(walletId)
		key2 = constant.RK_SYNC_WALLET_ACCOUNT_INFO + utils.GetHashTagKey(uid) + "-" + strconv.Itoa(int(walletType))
	)
	err = xredis.CUnlink([]string{key1, key2})
	return
}
