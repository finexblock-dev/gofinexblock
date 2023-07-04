package ethereum

const issue179FixEnvar = "GO_ETHEREUM_HDWALLET_FIX_ISSUE_179"

type GetReceiptOutput struct {
	TxHash           string `json:"tx_hash,omitempty"`
	Status           uint64 `json:"status,omitempty"`
	BlockHash        string `json:"block_hash,omitempty"`
	BlockNumber      string `json:"block_number,omitempty"`
	GasUsed          uint64 `json:"gas_used,omitempty"`
	TransactionIndex uint64 `json:"transaction_index,omitempty"`
}