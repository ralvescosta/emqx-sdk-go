package types

type (
	MonitorNodesResponse struct {
		Timestamp       int `json:"time_stamp"`
		Subscriptions   int `json:"subscriptions"`
		Topics          int `json:"topics"`
		Connections     int `json:"connections"`
		LiveConnections int `json:"live_connections"`
		Received        int `json:"received"`
		Sent            int `json:"sent"`
		Dropped         int `json:"dropped"`
	}
)
