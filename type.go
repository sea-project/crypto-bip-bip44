package bip44

type CoinType struct {
	Index    int64
	Hexa     uint32
	Symbol   string
	CoinName string
}

const Purpose uint32 = 0x8000002C

var (
	// 数字资产所用币种
	TypeBTC  = CoinType{Index: 0, Hexa: 0x80000000, Symbol: "BTC", CoinName: "Bitcoin"}
	TypeETH  = CoinType{Index: 60, Hexa: 0x8000003c, Symbol: "ETH", CoinName: "Ethereum"}
	TypeSEA  = CoinType{Index: 888888, Hexa: 0x800d9038, Symbol: "SEA", CoinName: "SEA"}
	TypeBSC  = CoinType{Index: 888887, Hexa: 0x800d9037, Symbol: "BSC", CoinName: "BSC"}
	TypeSSEA = CoinType{Index: 888886, Hexa: 0x800d9036, Symbol: "SSEA", CoinName: "SSEA"}
	TypeSDG  = CoinType{Index: 888885, Hexa: 0x800d9035, Symbol: "SDG", CoinName: "SDG"}
	TypeIHC  = CoinType{Index: 888884, Hexa: 0x800d9034, Symbol: "IHC", CoinName: "IHC"}
	TypeUSDTETH  = CoinType{Index: 888883, Hexa: 0x800d9033, Symbol: "USDT-ERC20", CoinName: "USDT-ERC20"}
	TypeUSDTBTC  = CoinType{Index: 888882, Hexa: 0x800d9032, Symbol: "USDT-OMNI", CoinName: "USDT-OMNI"}
	TypeBTCTEST  = CoinType{Index: 1, Hexa: 0x80000001, Symbol: "TESTNET", CoinName: "Bitcoin Testnet"}
)

