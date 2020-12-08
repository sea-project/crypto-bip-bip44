package bip44

import (
	"encoding/json"
	"fmt"
	"github.com/sea-project/crypto-bip-bip32"
	"github.com/sea-project/crypto-bip-bip39"
	"testing"
)

func Test_NewKeyFromMnemonic(t *testing.T) {
	seed, err := bip39.NewSeedWithMnemonic("gorilla easy one advance lesson name math clog awake private aerobic canvas kidney attend food amazing upper interest chicken shadow hip giraffe food curious", "")
	if err != nil {
		t.Log(err)
	}

	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		t.Log(err)
	}

	key, err := json.Marshal(masterKey)
	if err != nil {
		t.Log(err)
	}
	t.Log("key：" + string(key))

	xkey, _ := NewKeyFromMasterKey(masterKey, 0x8000003c, 60, 0, 0, 0, 0)
	fmt.Println("私钥：" + xkey.B58Serialize())
	fmt.Println("地址：" + bip32.PubKeyToAddr(xkey.PublicKey().Key))
}

