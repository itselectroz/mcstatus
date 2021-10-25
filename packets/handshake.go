package packets

type HandshakePacket = PacketData

// Make HandshakePacket implement OutgoingPacket interface

func (packet *HandshakePacket) writePayload() {
	// Handshake packet has no payload
}

type HandshakeResponsePacket struct {
	*PacketData
	payload string
}

func (packet *HandshakeResponsePacket) readPayload() {

}
