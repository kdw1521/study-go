package pubsub

import "context"

type Publisher struct {
	ctx context.Context
	// 채널 인데 타입이 채널에 string type을 넣기만 되는 타입 (write only)
	// 만일 read only가 필요하면 <-chan string 으로 하면 된다.
	// 기본적으로 chan string 하면 양방향, 즉 넣고,빼기가 다 된다.
	// 넣거나 빼는걸 강제하려고 할때 사용하면 된다.
	subscribeCh chan chan<- string
	publishCh   chan string
	subscribers []chan<- string
}

func NewPublisher(ctx context.Context) *Publisher {
	return &Publisher{
		ctx:         ctx,
		subscribeCh: make(chan chan<- string),
		publishCh:   make(chan string),
		subscribers: make([]chan<- string, 0),
	}
}

func (p *Publisher) Subscribe(sub chan<- string) {
	p.subscribeCh <- sub
}

func (p *Publisher) Publish(msg string) {
	p.publishCh <- msg
}

func (p *Publisher) Update() {
	for {
		select {
		case sub := <-p.subscribeCh:
			p.subscribers = append(p.subscribers, sub)
		case msg := <-p.publishCh:
			for _, subscriber := range p.subscribers {
				subscriber <- msg
			}
		case <-p.ctx.Done():
			wg.Done()
			return
		}
	}
}
