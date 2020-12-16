package bip44

type CoinType struct {
	Index    int64
	Hexa     uint32
	Symbol   string
	CoinName string
}

const Purpose uint32 = 0x8000002C

var (
	TypeBTC  = CoinType{Index: 0, Hexa: 0x80000000, Symbol: "BTC", CoinName: "Bitcoin"}
	TypeETH  = CoinType{Index: 60, Hexa: 0x8000003c, Symbol: "ETH", CoinName: "Ethereum"}
	TypeSEA  = CoinType{Index: 888888, Hexa: 0x800d9038, Symbol: "SEA", CoinName: "SEA"}
	TypeBTCTEST  = CoinType{Index: 1, Hexa: 0x80000001, Symbol: "TESTNET", CoinName: "Bitcoin Testnet"}
)

