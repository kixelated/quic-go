package congestion

import (
	"time"

	"github.com/kixelated/quic-go/internal/protocol"
)

// A SendAlgorithm performs congestion control
type SendAlgorithm interface {
	OnPacketSent(sentTime time.Time, bytesInFlight protocol.ByteCount, packetNumber protocol.PacketNumber, bytes protocol.ByteCount, isRetransmittable bool)
	CanSend(bytesInFlight protocol.ByteCount) bool
	MaybeExitSlowStart()
	OnPacketAcked(number protocol.PacketNumber, ackedBytes protocol.ByteCount, priorInFlight protocol.ByteCount, eventTime time.Time)
	OnPacketLost(number protocol.PacketNumber, lostBytes protocol.ByteCount, priorInFlight protocol.ByteCount)
	OnRetransmissionTimeout(packetsRetransmitted bool)
	SetMaxDatagramSize(protocol.ByteCount)
}

// A SendAlgorithmWithDebugInfos is a SendAlgorithm that exposes some debug infos
type SendAlgorithmWithDebugInfos interface {
	SendAlgorithm
	InSlowStart() bool
	InRecovery() bool
	GetCongestionWindow() protocol.ByteCount
	BandwidthEstimate() Bandwidth
}

type Pacer interface {
	SentPacket(when time.Time, size protocol.ByteCount)
	HasBudget() bool
	TimeUntilSend() time.Time
	SetMaxDatagramSize(size protocol.ByteCount)
}
