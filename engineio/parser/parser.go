package parser

import (
	"bytes"

	"go-socket.io/engineio/protocol"
)

type decodeOpts struct {
	payloadType protocol.PayloadType
}

type DecodeOptions func(*decodeOpts)

type Parser interface {
	GetProtocolVersion() int

	EncodePacket(packet protocol.EnginePacket, supportsBinary bool) []byte

	EncodePayload(packets []protocol.EnginePacket, supportsBinary bool) bytes.Buffer

	DecodePacket(data any, opts ...DecodeOptions) protocol.EnginePacket

	DecodePayload(data any, opts ...DecodeOptions) []protocol.EnginePacket
}

func WithPlaintextDecode() DecodeOptions {
	return func(opts *decodeOpts) {
		opts.payloadType = protocol.PayloadPlaintext
	}
}
