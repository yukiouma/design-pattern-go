package biz

type Publisher interface {
	Subsrcibe(channel string, client Subscriber)
	UnSubsrcibe(channel string, client Subscriber)
	Publish(channel string, newsContent string)
}

func NewNewsPublihser(channels ...string) Publisher {
	subscriber := make(map[string][]Subscriber)
	for _, v := range channels {
		subscriber[v] = make([]Subscriber, 0)
	}
	return &newsPublisher{
		subscriber: subscriber,
	}
}

type newsPublisher struct {
	subscriber map[string][]Subscriber
}

func (p *newsPublisher) Subsrcibe(channel string, client Subscriber) {
	list, ok := p.subscriber[channel]
	if !ok {
		return
	}
	list = append(list, client)
	p.subscriber[channel] = list
}

func (p *newsPublisher) UnSubsrcibe(channel string, client Subscriber) {
	list, ok := p.subscriber[channel]
	if !ok {
		return
	}
	for k, v := range list {
		if v == client {
			tmpList := make([]Subscriber, len(list)-1)
			copy(tmpList[:k], list[:k])
			copy(tmpList[k:], list[k+1:])
			p.subscriber[channel] = tmpList
			return
		}
	}
}

func (p *newsPublisher) Publish(channel string, newsContent string) {
	list, ok := p.subscriber[channel]
	if !ok {
		return
	}
	for _, subscriber := range list {
		subscriber.Receive(newsContent)
	}
}
