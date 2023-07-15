package types

import "encoding/json"

type (
	StatsResponse struct {
		Node            string `json:"node"`
		Channels        *StatsChannelsSummary
		Connections     *StatsConnectionsSummary
		Delayed         *StatsDelayedSummary
		LiveConnections *StatsLiveConnectionsSummary
		Retained        *StatsRetainedSummarySummary
		Sessions        *StatsSessionsSummary
		SubOptions      *StatsSubOptionsSummary
		Subscribers     *StatsSubscribersSummary
		Subscriptions   *StatsSubscriptionsSummary
		Topics          *StatsTopicsSummary
	}

	StatsChannelsSummary struct {
		Count int `json:"channels.count"`
		Max   int `json:"channels.max"`
	}

	StatsConnectionsSummary struct {
		Count int `json:"connections.count"`
		Max   int `json:"connections.max"`
	}

	StatsDelayedSummary struct {
		Count int `json:"delayed.count"`
		Max   int `json:"delayed.max"`
	}

	StatsLiveConnectionsSummary struct {
		Count int `json:"live_connections.count"`
		Max   int `json:"live_connections.max"`
	}

	StatsRetainedSummarySummary struct {
		Count int `json:"retained.count"`
		Max   int `json:"retained.max"`
	}

	StatsSessionsSummary struct {
		Count int `json:"sessions.count"`
		Max   int `json:"sessions.max"`
	}

	StatsSubOptionsSummary struct {
		Count int `json:"suboptions.count"`
		Max   int `json:"suboptions.max"`
	}

	StatsSubscribersSummary struct {
		Count int `json:"subscribers.count"`
		Max   int `json:"subscribers.max"`
	}

	StatsSubscriptionsSummary struct {
		Count         int `json:"subscriptions.count"`
		Max           int `json:"subscriptions.max"`
		SharedSummary *StatsSubscriptionsSharedSummary
	}

	StatsSubscriptionsSharedSummary struct {
		Count int `json:"subscriptions.shared.count"`
		Max   int `json:"subscriptions.shared.max"`
	}

	StatsTopicsSummary struct {
		Count int `json:"topics.count"`
		Max   int `json:"topics.max"`
	}
)

func (r *StatsResponse) UnmarshalJSON(b []byte) error {
	var v map[string]interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}

	r.Node = v["node"].(string)

	r.Channels = &StatsChannelsSummary{}
	r.Channels.Map(v)

	r.Connections = &StatsConnectionsSummary{}
	r.Connections.Map(v)

	r.Delayed = &StatsDelayedSummary{}
	r.Delayed.Map(v)

	r.LiveConnections = &StatsLiveConnectionsSummary{}
	r.LiveConnections.Map(v)

	r.Retained = &StatsRetainedSummarySummary{}
	r.Retained.Map(v)

	r.Sessions = &StatsSessionsSummary{}
	r.Sessions.Map(v)

	r.SubOptions = &StatsSubOptionsSummary{}
	r.SubOptions.Map(v)

	r.Subscribers = &StatsSubscribersSummary{}
	r.Subscribers.Map(v)

	r.Subscriptions = &StatsSubscriptionsSummary{}
	r.Subscriptions.Map(v)

	r.Topics = &StatsTopicsSummary{}
	r.Topics.Map(v)

	return nil
}

func (s *StatsChannelsSummary) Map(v map[string]interface{}) {
	s.Count = int(v["channels.count"].(float64))
	s.Max = int(v["channels.max"].(float64))
}

func (s *StatsConnectionsSummary) Map(v map[string]interface{}) {
	s.Count = int(v["channels.count"].(float64))
	s.Max = int(v["channels.max"].(float64))
}

func (s *StatsDelayedSummary) Map(v map[string]interface{}) {
	s.Count = int(v["channels.count"].(float64))
	s.Max = int(v["channels.max"].(float64))
}

func (s *StatsLiveConnectionsSummary) Map(v map[string]interface{}) {
	s.Count = int(v["channels.count"].(float64))
	s.Max = int(v["channels.max"].(float64))
}

func (s *StatsRetainedSummarySummary) Map(v map[string]interface{}) {
	s.Count = int(v["channels.count"].(float64))
	s.Max = int(v["channels.max"].(float64))
}

func (s *StatsSessionsSummary) Map(v map[string]interface{}) {
	s.Count = int(v["channels.count"].(float64))
	s.Max = int(v["channels.max"].(float64))
}

func (s *StatsSubOptionsSummary) Map(v map[string]interface{}) {
	s.Count = int(v["channels.count"].(float64))
	s.Max = int(v["channels.max"].(float64))
}

func (s *StatsSubscribersSummary) Map(v map[string]interface{}) {
	s.Count = int(v["channels.count"].(float64))
	s.Max = int(v["channels.max"].(float64))
}

func (s *StatsSubscriptionsSummary) Map(v map[string]interface{}) {
	s.Count = int(v["channels.count"].(float64))
	s.Max = int(v["channels.max"].(float64))
}

func (s *StatsTopicsSummary) Map(v map[string]interface{}) {
	s.Count = int(v["channels.count"].(float64))
	s.Max = int(v["channels.max"].(float64))
}
