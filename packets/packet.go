package packets

type OutgoingPacket interface {
	writePayload()
}

type IncomingPacket interface {
	readPayload()
}

type PacketData struct {
	packetType uint8
	sessionId  uint32
}
