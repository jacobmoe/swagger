package swagger

import (
	"fmt"
	"reflect"
)

type TRANSFER_PROTOCOL int64

const (
	TRANSFER_PROTOCOL_HTTP  TRANSFER_PROTOCOL = 0
	TRANSFER_PROTOCOL_HTTPS TRANSFER_PROTOCOL = 1
	TRANSFER_PROTOCOL_WS    TRANSFER_PROTOCOL = 2
	TRANSFER_PROTOCOL_WSS   TRANSFER_PROTOCOL = 3
)

func (p TRANSFER_PROTOCOL) String() string {
	switch p {
	case TRANSFER_PROTOCOL_HTTP:
		return "http"
	case TRANSFER_PROTOCOL_HTTPS:
		return "https"
	case TRANSFER_PROTOCOL_WS:
		return "ws"
	case TRANSFER_PROTOCOL_WSS:
		return "wss"
	}
	return "<UNSET>"
}

func (p TRANSFER_PROTOCOL) Type() reflect.Type {
	return reflect.TypeOf(p)
}

func (p TRANSFER_PROTOCOL) Slice() reflect.Type {
	return reflect.SliceOf(p.Type())
}

func TRANSFER_PROTOCOLFromString(s string) (TRANSFER_PROTOCOL, error) {
	switch s {
	case "http":
		return TRANSFER_PROTOCOL_HTTP, nil
	case "https":
		return TRANSFER_PROTOCOL_HTTPS, nil
	case "ws":
		return TRANSFER_PROTOCOL_WS, nil
	case "wss":
		return TRANSFER_PROTOCOL_WSS, nil
	}
	return TRANSFER_PROTOCOL(0), fmt.Errorf("not a valid TRANSFER_PROTOCOL string: " + s)
}
