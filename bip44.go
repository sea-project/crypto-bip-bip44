package bip44

import (
	bip32 "github.com/sea-project/crypto-bip-bip32"
	bip39 "github.com/sea-project/crypto-bip-bip39"
)

// https://github.com/satoshilabs/slips/blob/master/slip-0044.md
func NewKeyFromMnemonic(mnemonic string, coin, account, chain, address uint32) (*bip32.Key, error) {
	seed, err := bip39.NewSeedWithMnemonic(mnemonic, "")
	if err != nil {
		return nil, err
	}

	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, err
	}

	return NewKeyFromMasterKey(masterKey, coin, account, chain, address)
}

// m / purpose' / coin_type' / account' / change / address_index
// purpose 根据BIP43建议将常量设置为44'（或0x8000002C）。它指示根据此规范使用了此节点的子树
// coin_type 特指币种并且允许多元货币 HD 钱包中的货币在第二个层级下有自己的亚树状结构
// account  将密钥空间划分为独立的用户身份
// change 0用于外部接收地址 1用于找零地址
// address_index  地址索引
func NewKeyFromMasterKey(masterKey *bip32.Key, coin_type, account, change, address_index uint32) (*bip32.Key, error) {
	child, err := masterKey.NewChildKey(Purpose)
	if err != nil {
		return nil, err
	}

	child, err = child.NewChildKey(coin_type)
	if err != nil {
		return nil, err
	}

	child, err = child.NewChildKey(account)
	if err != nil {
		return nil, err
	}

	child, err = child.NewChildKey(change)
	if err != nil {
		return nil, err
	}

	child, err = child.NewChildKey(address_index)
	if err != nil {
		return nil, err
	}

	return child, nil
}
