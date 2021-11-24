package entities


type FundLatestInfo struct {
	// 基金代码
	FundCode string `json:"fundcode"`
	// 单位净值
	LastValue string `json:"dwjz"`
	// 净值日期
	LastValueDate string `json:"jzrq"`
	// 估值
	EstimatedValue string `json:"gsz"`
	// 估计增长率
	EstimatedPerGain string `json:"gszzl"`
	// 估值日期
	EstimatedDate string `json:"gztime"`
}
