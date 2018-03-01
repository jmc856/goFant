package yahoo

type YahooApiError struct {
	Status 		string					`json:"status"`
	YahooStatus int						`json:"yahoo_status"`
	Error       string					`json:"error"`
	Result      map[string]interface{}	`json:"result"`
}