package v3

import (
	"fmt"
	"strings"

	"github.com/NexClipper/sudory/pkg/manager/database/vanilla"
)

type RabbitMqConfig struct {
	//amqp Dial
	Url string `column:"url,default(0)" json:"url"`

	//amqp.Channel.Publish
	ChannelPublish struct {
		Exchange   vanilla.NullString `column:"exchange"    json:"exchange,omitempty"    swaggertype:"string"`
		RoutingKey vanilla.NullString `column:"routing_key" json:"routing_key,omitempty" swaggertype:"string"`
		Mandatory  vanilla.NullBool   `column:"mandatory"   json:"mandatory,omitempty"   swaggertype:"boolean"`
		Immediate  vanilla.NullBool   `column:"immediate"   json:"immediate,omitempty"   swaggertype:"boolean"`
	} `json:"channel_publish"`

	//amqp.Publishing
	Publishing struct {
		MessageHeaders         vanilla.NullObject `column:"message_headers"          json:"message_headers,omitempty"          swaggertype:"object"`  // Application or header exchange table
		MessageContentType     vanilla.NullString `column:"message_content_type"     json:"message_content_type,omitempty"     swaggertype:"string"`  // MIME content type
		MessageContentEncoding vanilla.NullString `column:"message_content_encoding" json:"message_content_encoding,omitempty" swaggertype:"string"`  // MIME content encoding
		MessageDeliveryMode    vanilla.NullUint8  `column:"message_delivery_mode"    json:"message_delivery_mode,omitempty"    swaggertype:"integer"` // queue implementation use - Transient (1) or Persistent (2)
		MessagePriority        vanilla.NullUint8  `column:"message_priority"         json:"message_priority,omitempty"         swaggertype:"integer"` // queue implementation use - 0 to 9
		MessageCorrelationId   vanilla.NullString `column:"message_correlation_id"   json:"message_correlation_id,omitempty"   swaggertype:"string"`  // application use - correlation identifier
		MessageReplyTo         vanilla.NullString `column:"message_reply_to"         json:"message_reply_to,omitempty"         swaggertype:"string"`  // application use - address to to reply to (ex: RPC)
		MessageExpiration      vanilla.NullString `column:"message_expiration"       json:"message_expiration,omitempty"       swaggertype:"string"`  // implementation use - message expiration spec
		MessageMessageId       vanilla.NullString `column:"message_message_id"       json:"message_message_id,omitempty"       swaggertype:"string"`  // application use - message identifier
		MessageTimestamp       vanilla.NullBool   `column:"message_timestamp"        json:"message_timestamp,omitempty"        swaggertype:"boolean"` // application use - message timestamp
		MessageType            vanilla.NullString `column:"message_type"             json:"message_type,omitempty"             swaggertype:"string"`  // application use - message type name
		MessageUserId          vanilla.NullString `column:"message_user_id"          json:"message_user_id,omitempty"          swaggertype:"string"`  // application use - creating user id
		MessageAppId           vanilla.NullString `column:"message_app_id"           json:"message_app_id,omitempty"           swaggertype:"string"`  // application use - creating application
	} `json:"publishing"`
}

func (RabbitMqConfig) Type() NotifierType {
	return NotifierTypeRabbitmq
}

func (cfg RabbitMqConfig) Valid() error {
	const protocol = "amqp://"
	if strings.Index(cfg.Url, protocol) != 0 {
		return fmt.Errorf("url is not an expression of the RabbitMQ protocol")
	}
	if len(cfg.ChannelPublish.Exchange.String) == 0 &&
		len(cfg.ChannelPublish.RoutingKey.String) == 0 {
		return fmt.Errorf("missing exchange or routing-key")
	}
	return nil
}

type NotifierRabbitMq_update = RabbitMqConfig

type NotifierRabbitMq_property = RabbitMqConfig

type NotifierRabbitMq struct {
	Uuid string `column:"uuid" json:"uuid,omitempty"` // pk

	NotifierRabbitMq_property `json:",inline"`

	// Created vanilla.NullTime `column:"created" json:"created,omitempty" swaggertype:"string"`
	// Updated vanilla.NullTime `column:"updated" json:"updated,omitempty" swaggertype:"string"`
}

func (NotifierRabbitMq) TableName() string {
	return "managed_channel_notifier_rabbitmq"
}
