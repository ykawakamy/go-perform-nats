package main

type TopicNameSupplier struct {
	seq    uint
	topics []string
}

func CreateTopicNameSupplier() TopicNameSupplier {
	return TopicNameSupplier{
		seq:    0,
		topics: []string{"test"},
	}
}

func (s *TopicNameSupplier) Get() string {
	ret := s.topics[int(s.seq)%len(s.topics)]
	s.seq += 1

	return ret
}
