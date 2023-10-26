package do

import (
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_wallet"
)

type Wallets struct {
	maps map[pb_enum.WALLET_TYPE]*pb_wallet.WalletInfo
}

func (w *Wallets) Get(walletType pb_enum.WALLET_TYPE) *pb_wallet.WalletInfo {
	return w.maps[walletType]
}

func (w *Wallets) Set(maps map[pb_enum.WALLET_TYPE]*pb_wallet.WalletInfo) {
	w.maps = maps
}
