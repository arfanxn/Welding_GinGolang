package entity

type OrderTrendPeriod struct {
	Period string  `json:"period"`
	Count  int     `json:"count"`
	Sum    float64 `json:"sum"`
}

type OrderTrends = []*OrderTrendPeriod
