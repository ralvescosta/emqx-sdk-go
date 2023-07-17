package types

type (
	AlarmsDataDetailsSummary struct {
		HighWatermark int `json:"high_watermark"`
	}

	AlarmsDataSummary struct {
		Node         string                    `json:"node"`
		Name         string                    `json:"name"`
		Message      string                    `json:"message"`
		Duration     int                       `json:"duration"`
		ActivateAt   string                    `json:"activate_at"`
		DeactivateAt string                    `json:"deactivate_at"`
		Details      *AlarmsDataDetailsSummary `json:"details"`
	}

	AlarmsMetaSummary struct {
		Page    int `json:"page"`
		Limit   int `json:"limit"`
		Count   int `json:"count"`
		HasNext int `json:"hasnext"`
	}

	AlarmsResponse struct {
		Data []*AlarmsDataSummary `json:"data"`
		Meta *AlarmsMetaSummary   `json:"meta"`
	}
)
