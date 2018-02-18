package smartcontract

const (
	FAIL = "FAIL"
	SUCCESS = "SUCCESS"
)

type SmartContractResponse struct {
	Result string			`json:"result"`
	Method string			`json:"method"` // query, invoke
	Data map[string]string	`json:"data"`
	Error error				`json:"error"`	// It's set when the result is fail
}
