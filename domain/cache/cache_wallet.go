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
	GetUserWallets(uid int64) (wallets []*pb_wallet.WalletInfo, err error)
	SetUserWallets(uid int64, wallets []*pb_wallet.WalletInfo) (err error)
	GetUpdatePasswordCode(uid int64, walletType pb_enum.WALLET_TYPE) (code string, err error)
	DelUpdatePasswordCode(uid int64, walletType pb_enum.WALLET_TYPE) (err error)
	SetUpdatePasswordCode(uid int64, walletType pb_enum.WALLET_TYPE, code string) (err error)
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

func (c *walletCache) GetUserWallets(uid int64) (wallets []*pb_wallet.WalletInfo, err error) {
	var (
		key = constant.RK_SYNC_USER_WALLETS + utils.GetHashTagKey(uid)
	)
	err = Get(key, &wallets)
	return
}

func (c *walletCache) SetUserWallets(uid int64, wallets []*pb_wallet.WalletInfo) (err error) {
	var (
		key = constant.RK_SYNC_USER_WALLETS + utils.GetHashTagKey(uid)
	)
	err = Set(key, wallets, constant.CONST_DURATION_USER_WALLETS_SECOND)
	return
}

func (c *walletCache) GetUpdatePasswordCode(uid int64, walletType pb_enum.WALLET_TYPE) (code string, err error) {
	var (
		key = constant.RK_SYNC_WALLET_UPDATE_PWD_CODE + utils.GetHashTagKey(uid) + strconv.Itoa(int(walletType))
	)
	code, err = xredis.Get(key)
	if err != nil {
		return
	}
	return
}

func (c *walletCache) DelUpdatePasswordCode(uid int64, walletType pb_enum.WALLET_TYPE) (err error) {
	var (
		key = constant.RK_SYNC_WALLET_UPDATE_PWD_CODE + utils.GetHashTagKey(uid) + strconv.Itoa(int(walletType))
	)
	err = Delete(key)
	return
}

func (c *walletCache) SetUpdatePasswordCode(uid int64, walletType pb_enum.WALLET_TYPE, code string) (err error) {
	var (
		key = constant.RK_SYNC_WALLET_UPDATE_PWD_CODE + utils.GetHashTagKey(uid) + strconv.Itoa(int(walletType))
	)
	Set(key, code, constant.CONST_DURATION_WALLET_RESET_PWD_CODE_SECOND)
	return
}
