package coregeth

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params/types/ctypes"
)

var testConfig = &CoreGethChainConfig{
	NetworkID:  1,
	Ethash:     new(ctypes.EthashConfig),
	ChainID:    big.NewInt(61),
	EIP2FBlock: big.NewInt(1150000),
	EIP7FBlock: big.NewInt(1150000),
	//DAOForkBlock:        big.NewInt(1920000),
	EIP150Block:        big.NewInt(2500000),
	EIP155Block:        big.NewInt(3000000),
	EIP160FBlock:       big.NewInt(3000000),
	EIP161FBlock:       big.NewInt(8772000),
	EIP170FBlock:       big.NewInt(8772000),
	EIP100FBlock:       big.NewInt(8772000),
	EIP140FBlock:       big.NewInt(8772000),
	EIP198FBlock:       big.NewInt(8772000),
	EIP211FBlock:       big.NewInt(8772000),
	EIP212FBlock:       big.NewInt(8772000),
	EIP213FBlock:       big.NewInt(8772000),
	EIP214FBlock:       big.NewInt(8772000),
	EIP658FBlock:       big.NewInt(8772000),
	EIP145FBlock:       big.NewInt(9573000),
	EIP1014FBlock:      big.NewInt(9573000),
	EIP1052FBlock:      big.NewInt(9573000),
	EIP1283FBlock:      nil,
	PetersburgBlock:    nil, // Un1283
	EIP2200FBlock:      nil, // RePetersburg (== re-1283)
	DisposalBlock:      big.NewInt(5900000),
	ECIP1017FBlock:     big.NewInt(5000000),
	ECIP1017EraRounds:  big.NewInt(5000000),
	ECIP1010PauseBlock: big.NewInt(3000000),
	ECIP1010Length:     big.NewInt(2000000),
	RequireBlockHashes: map[uint64]common.Hash{
		2500000: common.HexToHash("0xca12c63534f565899681965528d536c52cb05b7c48e269c2a6cb77ad864d878a"),
	},
}

func TestCoreGethChainConfig_String(t *testing.T) {
	t.Skip("(noop) development use only")
	t.Log(testConfig.String())
}
