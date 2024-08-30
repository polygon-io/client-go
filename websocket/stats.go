package polygonws

import "encoding/json"

const (
	ReconnectCount    = "reconnectCount"
	MessagesRecv      = "messagesRecv"
	SubscriptionCount = "subscriptionCount"
)

type ClientStatsCollector interface {
	Get(name string) uint64
	Increment(name string)
	Update(name string, value uint64)
	ToJSON() ([]byte, error)
}

type ClientStats struct {
	stats map[string]uint64
}

func NewClientStats() *ClientStats {
	return &ClientStats{stats: map[string]uint64{}}
}

func (s *ClientStats) Get(name string) uint64 {
	return s.stats[name]
}

func (s *ClientStats) Increment(name string) {
	if _, ok := s.stats[name]; ok {
		s.stats[name]++
	} else {
		s.stats[name] = 1
	}
}

func (s *ClientStats) ToJSON() ([]byte, error) {
	return json.Marshal(s.stats)
}
