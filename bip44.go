package bip44

import (
	"github.com/sea-project/crypto-bip-bip32"
)

const (
	Purpose   uint32 = 0x8000002C
	Purpose45 uint32 = 0x8000002d
)

// bip44 m / purpose' / coin_type' / account' / change / address_index
// bip45 m / purpose' / algorithmType'/ orgOrCoinType' / account' / change / address_index
// CKD: m: 使用 CKDpriv, M 则表示使用 CKDPub
// purpose 根据BIP43建议将常量设置为44'（或0x8000002C）。它指示根据此规范使用了此节点的子树
// purpose参见：https://github.com/satoshilabs/slips/blob/master/slip-0044.md
// coinType 币种
// org 组织
// account  将密钥空间划分为独立的用户身份
// change 0用于外部接收地址 1用于找零地址
// addressIndex  地址索引
func NewKeyFromMasterKey(masterKey *bip32.Key, purpose, coinType, org, account, change, addressIndex uint32) (*bip32.Key, error) {
	child, err := masterKey.NewChildKey(purpose)
	if err != nil {
		return nil, err
	}

	child, err = child.NewChildKey(coinType)
	if err != nil {
		return nil, err
	}

	if purpose != Purpose {
		child, err = child.NewChildKey(org)
		if err != nil {
			return nil, err
		}
	}

	child, err = child.NewChildKey(account)
	if err != nil {
		return nil, err
	}

	child, err = child.NewChildKey(change)
	if err != nil {
		return nil, err
	}

	child, err = child.NewChildKey(addressIndex)
	if err != nil {
		return nil, err
	}

	return child, nil
}
