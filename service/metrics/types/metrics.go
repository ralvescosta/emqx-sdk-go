package types

import (
	"encoding/json"
)

type (
	DeliveryDroppedSummary struct {
		QueueFull int `json:"delivery.dropped.queue_full"`
		Qos0Msg   int `json:"delivery.dropped.qos0_msg"`
		Expired   int `json:"delivery.dropped.expired"`
		NoLocal   int `json:"delivery.dropped.no_local"`
		TooLarge  int `json:"delivery.dropped.too_large"`
	}

	DeliverySummary struct {
		DroppedSummary *DeliveryDroppedSummary
		Dropped        int `json:"delivery.dropped"`
	}

	MatchedSummary struct {
		Allow int `json:"authorization.matched.allow"`
		Deny  int `json:"authorization.matched.deny"`
	}

	AuthorizationSummary struct {
		Matched   *MatchedSummary
		CacheHit  int `json:"authorization.cache_hit"`
		NoMatch   int `json:"authorization.nomatch"`
		Deny      int `json:"authorization.deny"`
		Allow     int `json:"authorization.allow"`
		SuperUser int `json:"authorization.superuser"`
	}

	MessageDroppedSummary struct {
		AwaitPubrelTimeout int `json:"messages.dropped.await_pubrel_timeout"`
		NoSubscribers      int `json:"messages.dropped.no_subscribers"`
	}

	MessageQoS0Summary struct {
		Received int `json:"messages.qos0.received"`
		Sent     int `json:"messages.qos0.sent"`
	}

	MessageQoS1Summary struct {
		Received int `json:"messages.qos1.received"`
		Sent     int `json:"messages.qos1.sent"`
	}

	MessageQoS2Summary struct {
		Received int `json:"messages.qos2.received"`
		Sent     int `json:"messages.qos2.sent"`
	}

	MessageSummary struct {
		Received       int `json:"messages.received"`
		Sent           int `json:"messages.sent"`
		Publish        int `json:"messages.publish"`
		Acked          int `json:"messages.acked"`
		Forward        int `json:"messages.forward"`
		Delayed        int `json:"messages.delayed"`
		Delivered      int `json:"messages.delivered"`
		Dropped        int `json:"messages.dropped"`
		DroppedSummary *MessageDroppedSummary
		QoS0SUmmary    *MessageQoS0Summary
		QoS1SUmmary    *MessageQoS1Summary
		QoS2SUmmary    *MessageQoS2Summary
	}

	ClientAuthSummary struct {
		Anonymous int // "client.auth.anonymous"
	}

	ClientSummary struct {
		Unsubscribe  int `json:"client.unsubscribe"`
		Connected    int `json:"client.connected"`
		Connack      int `json:"client.connack"`
		Disconnected int `json:"client.disconnected"`
		Authorize    int `json:"client.authorize"`
		Subscribe    int `json:"client.subscribe"`
		Authenticate int `json:"client.authenticate"`
		Connect      int `json:"client.connect"`
		Auth         *ClientAuthSummary
	}

	PacketsSubscribeSummary struct {
		Error     int `json:"packets.subscribe.error"`
		Received  int `json:"packets.subscribe.received"`
		AuthError int `json:"packets.subscribe.auth_error"`
	}

	PacketsPubrecSummary struct {
		Missed   int `json:"packets.pubrec.missed"`
		Received int `json:"packets.pubrec.received"`
		Inuse    int `json:"packets.pubrec.inuse"`
		Sent     int `json:"packets.pubrec.sent"`
	}

	PacketsPublishSummary struct {
		Sent      int `json:"packets.publish.sent"`
		Inuse     int `json:"packets.publish.inuse"`
		Error     int `json:"packets.publish.error"`
		Dropped   int `json:"packets.publish.dropped"`
		AuthError int `json:"packets.publish.auth_error"`
		Received  int `json:"packets.publish.received"`
	}

	PacketsPubrelSummary struct {
		Received int `json:"packets.pubrel.received"`
		Missed   int `json:"packets.pubrel.missed"`
		Sent     int `json:"packets.pubrel.sent"`
	}

	PacketsPingreqSummary struct {
		Received int `json:"packets.pingreq.received"`
	}

	PacketsConnackSummary struct {
		Error     int `json:"packets.connack.error"`
		AuthError int `json:"packets.connack.auth_error"`
		Sent      int `json:"packets.connack.sent"`
	}

	PacketsDisconnectSummary struct {
		Sent     int `json:"packets.disconnect.sent"`
		Received int `json:"packets.disconnect.received"`
	}

	PacketsPubackSummary struct {
		Inuse    int `json:"packets.puback.inuse"`
		Sent     int `json:"packets.puback.sent"`
		Missed   int `json:"packets.puback.missed"`
		Received int `json:"packets.puback.received"`
	}

	PacketsPubcompSummary struct {
		Sent     int `json:"packets.pubcomp.sent"`
		Inuse    int `json:"packets.pubcomp.inuse"`
		Received int `json:"packets.pubcomp.received"`
		Missed   int `json:"packets.pubcomp.missed"`
	}

	PacketsSubackSummary struct {
		Sent int `json:"packets.suback.sent"`
	}

	PacketsAuthSummary struct {
		Sent     int `json:"packets.auth.sent"`
		Received int `json:"packets.auth.received"`
	}

	PacketsConnectSummary struct {
		Received int `json:"packets.connect.received"`
	}

	PacketsUnsubackSummary struct {
		Sent int `json:"packets.unsuback.sent"`
	}

	PacketsUnsubscribeSummary struct {
		Received int `json:"packets.unsubscribe.received"`
		Error    int `json:"packets.unsubscribe.error"`
	}

	PacketsPingrespSummary struct {
		Sent int `json:"packets.pingresp.sent"`
	}

	PacketsSummary struct {
		Sent        int `json:"packets.sent"`
		Received    int `json:"packets.received"`
		Subscribe   *PacketsSubscribeSummary
		Pubrec      *PacketsPubrecSummary
		Publish     *PacketsPublishSummary
		Pubrel      *PacketsPubrelSummary
		PingreqS    *PacketsPingreqSummary
		Connack     *PacketsConnackSummary
		Disconnect  *PacketsDisconnectSummary
		Puback      *PacketsPubackSummary
		Pubcomp     *PacketsPubcompSummary
		Suback      *PacketsSubackSummary
		Auth        *PacketsAuthSummary
		Connect     *PacketsConnectSummary
		Unsuback    *PacketsUnsubackSummary
		Unsubscribe *PacketsUnsubscribeSummary
		Pingresp    *PacketsPingrespSummary
	}

	AuthenticationSuccessSummary struct {
		Anonymous int `json:"authentication.success.anonymous"`
	}

	AuthenticationSummary struct {
		SuccessSummary *AuthenticationSuccessSummary
		Success        int `json:"authentication.success"`
		Failure        int `json:"authentication.failure"`
	}

	OLPDelaySummary struct {
		Timeout int `json:"olp.delay.timeout"`
		Ok      int `json:"olp.delay.ok"`
	}

	OLPSummary struct {
		DelaySummary *OLPDelaySummary
		Hbn          int `json:"olp.hbn"`
		Gc           int `json:"olp.gc"`
		NewConn      int `json:"olp.new_conn"`
	}

	SessionSummary struct {
		Created    int `json:"session.created"`
		Discarded  int `json:"session.discarded"`
		Resumed    int `json:"session.resumed"`
		Takenover  int `json:"session.takenover"`
		Terminated int `json:"session.terminated"`
	}

	BytesSummary struct {
		Received int `json:"bytes.received"`
		Sent     int `json:"bytes.sent"`
	}

	MetricsResponse struct {
		Delivery       *DeliverySummary
		Authorization  *AuthorizationSummary
		Message        *MessageSummary
		Client         *ClientSummary
		Packets        *PacketsSummary
		Authentication *AuthenticationSummary
		OLP            *OLPSummary
		Session        *SessionSummary
		Bytes          *BytesSummary
		Node           string `json:"node"`
	}
)

func (r *MetricsResponse) UnmarshalJSON(b []byte) error {
	var v map[string]interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}

	r.Node = v["node"].(string)

	r.Delivery = &DeliverySummary{}
	r.Delivery.Map(v)

	r.Authorization = &AuthorizationSummary{}
	r.Authorization.Map(v)

	r.Message = &MessageSummary{}
	r.Message.Map(v)

	r.Client = &ClientSummary{}
	r.Client.Map(v)

	r.Packets = &PacketsSummary{}
	r.Packets.Map(v)

	//AuthenticationSummary
	r.Authentication = &AuthenticationSummary{}
	r.Authentication.Map(v)

	//OLPSummary
	r.OLP = &OLPSummary{}
	r.OLP.Map(v)

	//SessionSummary
	r.Session = &SessionSummary{}
	r.Session.Map(v)

	//BytesSummary
	r.Bytes = &BytesSummary{}
	r.Bytes.Map(v)

	return nil
}

func (s *DeliverySummary) Map(v map[string]interface{}) {
	s.Dropped = int(v["delivery.dropped"].(float64))
	s.DroppedSummary = &DeliveryDroppedSummary{}
	s.DroppedSummary.Map(v)
}

func (s *DeliveryDroppedSummary) Map(v map[string]interface{}) {
	s.QueueFull = int(v["delivery.dropped.queue_full"].(float64))
	s.Qos0Msg = int(v["delivery.dropped.qos0_msg"].(float64))
	s.Expired = int(v["delivery.dropped.expired"].(float64))
	s.NoLocal = int(v["delivery.dropped.no_local"].(float64))
	s.TooLarge = int(v["delivery.dropped.too_large"].(float64))
}

func (s *AuthorizationSummary) Map(v map[string]interface{}) {
	s.CacheHit = int(v["authorization.cache_hit"].(float64))
	s.NoMatch = int(v["authorization.nomatch"].(float64))
	s.Deny = int(v["authorization.deny"].(float64))
	s.Allow = int(v["authorization.allow"].(float64))
	s.SuperUser = int(v["authorization.superuser"].(float64))
	s.Matched.Map(v)
}

func (s *MatchedSummary) Map(v map[string]interface{}) {
	s.Allow = int(v["authorization.matched.allow"].(float64))
	s.Deny = int(v["authorization.matched.deny"].(float64))
}

func (s *MessageSummary) Map(v map[string]interface{}) {
	s.Received = int(v["messages.received"].(float64))
	s.Sent = int(v["messages.sent"].(float64))
	s.Publish = int(v["messages.acked"].(float64))
	s.Acked = int(v["messages.forward"].(float64))
	s.Forward = int(v["messages.delayed"].(float64))
	s.Delayed = int(v["messages.delayed"].(float64))
	s.Delivered = int(v["messages.delivered"].(float64))
	s.Dropped = int(v["messages.dropped"].(float64))
	s.QoS0SUmmary.Map(v)
	s.QoS1SUmmary.Map(v)
	s.QoS2SUmmary.Map(v)
}

func (s *MessageDroppedSummary) Map(v map[string]interface{}) {
	s.AwaitPubrelTimeout = int(v["messages.dropped.await_pubrel_timeout"].(float64))
	s.NoSubscribers = int(v["messages.dropped.no_subscribers"].(float64))
}

func (s *MessageQoS0Summary) Map(v map[string]interface{}) {
	s.Received = int(v["messages.qos0.received"].(float64))
	s.Sent = int(v["messages.qos0.sent"].(float64))
}

func (s *MessageQoS1Summary) Map(v map[string]interface{}) {
	s.Received = int(v["messages.qos1.received"].(float64))
	s.Sent = int(v["messages.qos1.sent"].(float64))
}

func (s *MessageQoS2Summary) Map(v map[string]interface{}) {
	s.Received = int(v["messages.qos2.received"].(float64))
	s.Sent = int(v["messages.qos2.sent"].(float64))
}

func (s *ClientSummary) Map(v map[string]interface{}) {
	s.Unsubscribe = int(v["client.unsubscribe"].(float64))
	s.Connected = int(v["client.connected"].(float64))
	s.Connack = int(v["client.connack"].(float64))
	s.Disconnected = int(v["client.disconnected"].(float64))
	s.Authorize = int(v["client.authorize"].(float64))
	s.Subscribe = int(v["client.subscribe"].(float64))
	s.Authenticate = int(v["client.authenticate"].(float64))
	s.Connect = int(v["client.connect"].(float64))
	s.Auth.Map(v)
}

func (s *ClientAuthSummary) Map(v map[string]interface{}) {
	s.Anonymous = int(v["client.auth.anonymous"].(float64))
}

func (s *PacketsSummary) Map(v map[string]interface{}) {
	s.Sent = int(v["packets.sent"].(float64))
	s.Received = int(v["packets.received"].(float64))
	s.Subscribe.Map(v)
	s.Pubrec.Map(v)
	s.Publish.Map(v)
	s.Pubrel.Map(v)
	s.PingreqS.Map(v)
	s.Connack.Map(v)
	s.Disconnect.Map(v)
	s.Puback.Map(v)
	s.Pubcomp.Map(v)
	s.Suback.Map(v)
	s.Auth.Map(v)
	s.Connect.Map(v)
	s.Unsuback.Map(v)
	s.Unsubscribe.Map(v)
	s.Pingresp.Map(v)
}

func (s *PacketsSubscribeSummary) Map(v map[string]interface{}) {
	s.Error = int(v["packets.subscribe.error"].(float64))
	s.Received = int(v["packets.subscribe.received"].(float64))
	s.AuthError = int(v["packets.subscribe.auth_error"].(float64))
}

func (s *PacketsPubrecSummary) Map(v map[string]interface{}) {
	s.Missed = int(v["packets.pubrec.missed"].(float64))
	s.Received = int(v["packets.pubrec.received"].(float64))
	s.Inuse = int(v["packets.pubrec.inuse"].(float64))
	s.Sent = int(v["packets.pubrec.sent"].(float64))
}

func (s *PacketsPublishSummary) Map(v map[string]interface{}) {
	s.Sent = int(v["packets.publish.sent"].(float64))
	s.Inuse = int(v["packets.publish.inuse"].(float64))
	s.Error = int(v["packets.publish.error"].(float64))
	s.Dropped = int(v["packets.publish.dropped"].(float64))
	s.AuthError = int(v["packets.publish.auth_error"].(float64))
	s.Received = int(v["packets.publish.received"].(float64))
}

func (s *PacketsPubrelSummary) Map(v map[string]interface{}) {
	s.Received = int(v["packets.pubrel.received"].(float64))
	s.Missed = int(v["packets.pubrel.missed"].(float64))
	s.Sent = int(v["packets.pubrel.sent"].(float64))
}

func (s *PacketsPingreqSummary) Map(v map[string]interface{}) {
	s.Received = int(v["packets.pingreq.received"].(float64))
}

func (s *PacketsConnackSummary) Map(v map[string]interface{}) {
	s.Error = int(v["packets.connack.error"].(float64))
	s.AuthError = int(v["packets.connack.auth_error"].(float64))
	s.Sent = int(v["packets.connack.sent"].(float64))
}

func (s *PacketsDisconnectSummary) Map(v map[string]interface{}) {
	s.Sent = int(v["packets.disconnect.sent"].(float64))
	s.Received = int(v["packets.disconnect.received"].(float64))
}

func (s *PacketsPubackSummary) Map(v map[string]interface{}) {
	s.Inuse = int(v["packets.puback.inuse"].(float64))
	s.Sent = int(v["packets.puback.sent"].(float64))
	s.Missed = int(v["packets.puback.missed"].(float64))
	s.Received = int(v["packets.puback.received"].(float64))
}

func (s *PacketsPubcompSummary) Map(v map[string]interface{}) {
	s.Sent = int(v["packets.pubcomp.sent"].(float64))
	s.Inuse = int(v["packets.pubcomp.inuse"].(float64))
	s.Received = int(v["packets.pubcomp.received"].(float64))
	s.Missed = int(v["packets.pubcomp.missed"].(float64))
}

func (s *PacketsSubackSummary) Map(v map[string]interface{}) {
	s.Sent = int(v["packets.suback.sent"].(float64))
}

func (s *PacketsAuthSummary) Map(v map[string]interface{}) {
	s.Sent = int(v["packets.auth.sent"].(float64))
	s.Received = int(v["ackets.auth.received"].(float64))
}

func (s *PacketsConnectSummary) Map(v map[string]interface{}) {
	s.Received = int(v["packets.connect.received"].(float64))
}

func (s *PacketsUnsubackSummary) Map(v map[string]interface{}) {
	s.Sent = int(v["packets.unsuback.sent"].(float64))
}

func (s *PacketsUnsubscribeSummary) Map(v map[string]interface{}) {
	s.Received = int(v["packets.unsubscribe.received"].(float64))
	s.Error = int(v["packets.unsubscribe.error"].(float64))
}

func (s *PacketsPingrespSummary) Map(v map[string]interface{}) {
	s.Sent = int(v["packets.pingresp.sent"].(float64))
}

func (s *AuthenticationSummary) Map(v map[string]interface{}) {
	s.Success = int(v["authentication.success"].(float64))
	s.Failure = int(v["authentication.failure"].(float64))
	s.SuccessSummary.Map(v)
}

func (s *AuthenticationSuccessSummary) Map(v map[string]interface{}) {
	s.Anonymous = int(v["authentication.success.anonymous"].(float64))
}

func (s *OLPSummary) Map(v map[string]interface{}) {
	s.Hbn = int(v["olp.hbn"].(float64))
	s.Gc = int(v["olp.gc"].(float64))
	s.NewConn = int(v["olp.new_conn"].(float64))
	s.DelaySummary.Map(v)
}

func (s *OLPDelaySummary) Map(v map[string]interface{}) {
	s.Timeout = int(v["olp.delay.timeout"].(float64))
	s.Ok = int(v["olp.delay.ok"].(float64))
}

func (s *SessionSummary) Map(v map[string]interface{}) {
	s.Created = int(v["session.created"].(float64))
	s.Discarded = int(v["session.discarded"].(float64))
	s.Resumed = int(v["session.resumed"].(float64))
	s.Takenover = int(v["session.takenover"].(float64))
	s.Terminated = int(v["session.terminated"].(float64))
}

func (s *BytesSummary) Map(v map[string]interface{}) {
	s.Received = int(v["bytes.received"].(float64))
	s.Sent = int(v["bytes.sent"].(float64))
}
