package types

type (
	MonitorCurrentNodeResponse struct {
		DroppedMsgRate  int `json:"dropped_msg_rate"`
		ReceivedMsgRate int `json:"received_msg_rate"`
		SentMsgRate     int `json:"sent_msg_rate"`
		Subscriptions   int `json:"subscriptions"`
		Topics          int `json:"topics"`
		Connections     int `json:"connections"`
		LiveConnections int `json:"live_connections"`
	}
)
