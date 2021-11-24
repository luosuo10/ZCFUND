package entities


type FundBaseInfo struct {
	// 基金代码
	FundCode string `json:"fundcode"`
	// 基金名
	Name string `json:"name"`
	Type string
}
