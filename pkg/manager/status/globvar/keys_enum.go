// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package globvar

import (
	"fmt"
	"strings"
)

const (
	// KeyBearerTokenSignatureSecret is a Key of type Bearer-Token-Signature-Secret.
	KeyBearerTokenSignatureSecret Key = iota
	// KeyBearerTokenExpirationTime is a Key of type Bearer-Token-Expiration-Time.
	KeyBearerTokenExpirationTime
	// KeyClientSessionSignatureSecret is a Key of type Client-Session-Signature-Secret.
	KeyClientSessionSignatureSecret
	// KeyClientSessionExpirationTime is a Key of type Client-Session-Expiration-Time.
	KeyClientSessionExpirationTime
	// KeyClientConfigPollInterval is a Key of type Client-Config-Poll-Interval.
	KeyClientConfigPollInterval
	// KeyClientConfigLoglevel is a Key of type Client-Config-Loglevel.
	KeyClientConfigLoglevel
	// KeyClientConfigServiceValidTimeLimit is a Key of type Client-Config-Service-Valid-Time-Limit.
	KeyClientConfigServiceValidTimeLimit
	// KeyEventNotifierStatusRotateLimit is a Key of type Event-Notifier-Status-Rotate-Limit.
	KeyEventNotifierStatusRotateLimit
	// KeyEventNotifierRabbitMqTimeout is a Key of type Event-Notifier-RabbitMq-Timeout.
	KeyEventNotifierRabbitMqTimeout
	// KeyServiceSessionSignatureSecret is a Key of type Service-Session-Signature-Secret.
	KeyServiceSessionSignatureSecret
	// KeyServiceSessionExpirationTime is a Key of type Service-Session-Expiration-Time.
	KeyServiceSessionExpirationTime
)

var ErrInvalidKey = fmt.Errorf("not a valid Key, try [%s]", strings.Join(_KeyNames, ", "))

const _KeyName = "bearer-token-signature-secretbearer-token-expiration-timeclient-session-signature-secretclient-session-expiration-timeclient-config-poll-intervalclient-config-loglevelclient-config-service-valid-time-limitevent-notifier-status-rotate-limitevent-notifier-rabbitMq-timeoutservice-session-signature-secretservice-session-expiration-time"

var _KeyNames = []string{
	_KeyName[0:29],
	_KeyName[29:57],
	_KeyName[57:88],
	_KeyName[88:118],
	_KeyName[118:145],
	_KeyName[145:167],
	_KeyName[167:205],
	_KeyName[205:239],
	_KeyName[239:270],
	_KeyName[270:302],
	_KeyName[302:333],
}

// KeyNames returns a list of possible string values of Key.
func KeyNames() []string {
	tmp := make([]string, len(_KeyNames))
	copy(tmp, _KeyNames)
	return tmp
}

var _KeyMap = map[Key]string{
	KeyBearerTokenSignatureSecret:        _KeyName[0:29],
	KeyBearerTokenExpirationTime:         _KeyName[29:57],
	KeyClientSessionSignatureSecret:      _KeyName[57:88],
	KeyClientSessionExpirationTime:       _KeyName[88:118],
	KeyClientConfigPollInterval:          _KeyName[118:145],
	KeyClientConfigLoglevel:              _KeyName[145:167],
	KeyClientConfigServiceValidTimeLimit: _KeyName[167:205],
	KeyEventNotifierStatusRotateLimit:    _KeyName[205:239],
	KeyEventNotifierRabbitMqTimeout:      _KeyName[239:270],
	KeyServiceSessionSignatureSecret:     _KeyName[270:302],
	KeyServiceSessionExpirationTime:      _KeyName[302:333],
}

// String implements the Stringer interface.
func (x Key) String() string {
	if str, ok := _KeyMap[x]; ok {
		return str
	}
	return fmt.Sprintf("Key(%d)", x)
}

var _KeyValue = map[string]Key{
	_KeyName[0:29]:                     KeyBearerTokenSignatureSecret,
	strings.ToLower(_KeyName[0:29]):    KeyBearerTokenSignatureSecret,
	_KeyName[29:57]:                    KeyBearerTokenExpirationTime,
	strings.ToLower(_KeyName[29:57]):   KeyBearerTokenExpirationTime,
	_KeyName[57:88]:                    KeyClientSessionSignatureSecret,
	strings.ToLower(_KeyName[57:88]):   KeyClientSessionSignatureSecret,
	_KeyName[88:118]:                   KeyClientSessionExpirationTime,
	strings.ToLower(_KeyName[88:118]):  KeyClientSessionExpirationTime,
	_KeyName[118:145]:                  KeyClientConfigPollInterval,
	strings.ToLower(_KeyName[118:145]): KeyClientConfigPollInterval,
	_KeyName[145:167]:                  KeyClientConfigLoglevel,
	strings.ToLower(_KeyName[145:167]): KeyClientConfigLoglevel,
	_KeyName[167:205]:                  KeyClientConfigServiceValidTimeLimit,
	strings.ToLower(_KeyName[167:205]): KeyClientConfigServiceValidTimeLimit,
	_KeyName[205:239]:                  KeyEventNotifierStatusRotateLimit,
	strings.ToLower(_KeyName[205:239]): KeyEventNotifierStatusRotateLimit,
	_KeyName[239:270]:                  KeyEventNotifierRabbitMqTimeout,
	strings.ToLower(_KeyName[239:270]): KeyEventNotifierRabbitMqTimeout,
	_KeyName[270:302]:                  KeyServiceSessionSignatureSecret,
	strings.ToLower(_KeyName[270:302]): KeyServiceSessionSignatureSecret,
	_KeyName[302:333]:                  KeyServiceSessionExpirationTime,
	strings.ToLower(_KeyName[302:333]): KeyServiceSessionExpirationTime,
}

// ParseKey attempts to convert a string to a Key.
func ParseKey(name string) (Key, error) {
	if x, ok := _KeyValue[name]; ok {
		return x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := _KeyValue[strings.ToLower(name)]; ok {
		return x, nil
	}
	return Key(0), fmt.Errorf("%s is %w", name, ErrInvalidKey)
}