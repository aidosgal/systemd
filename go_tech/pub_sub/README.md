## Pub/Sub Message Broker
#### **Time**: 50 minutes | Difficulty: Very Hard
#### **Problem**: Build an in-memory message broker supporting topics, subscriptions, and delivery guarantees.

**Requirements**:
- Publish(topic, message) - publish to topic
- Subscribe(topic, handler) - subscribe to topic
- Support wildcard subscriptions (user.*, order.#)
- At-least-once delivery guarantee
- Message persistence with configurable retention
- Subscriber backpressure handling

**Key Challenges**:
- Implement topic matching with wildcards
- Handle slow subscribers without blocking publishers
- Implement message acknowledgment and retry logic
- Coordinate subscriber lifecycle management

```
type Message struct {
    ID        string
    Topic     string
    Payload   []byte
    Timestamp time.Time
}

type MessageBroker struct {
    // Your implementation here
}

func NewMessageBroker() *MessageBroker

func (mb *MessageBroker) Publish(topic string, payload []byte) error
func (mb *MessageBroker) Subscribe(pattern string, handler func(Message) error) (Subscription, error)
func (mb *MessageBroker) Close() error
```