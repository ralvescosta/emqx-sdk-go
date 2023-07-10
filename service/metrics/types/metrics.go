package types

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
		PUblish        int `json:"messages.publish"`
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
		Node           string `json:"none"`
	}
)
