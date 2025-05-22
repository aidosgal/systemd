## Token Bucket Rate Limiter
#### **Time**: 50 minutes | Difficulty: Hard
#### **Problem**: Implement a distributed-ready token bucket rate limiter with burst capability.

**Requirements**:
- Allow(tokens int) bool - check if request is allowed
- Wait(ctx context.Context, tokens int) error - block until tokens available
- Support burst traffic (bucket capacity > rate)
- Configurable refill rate and bucket size
- Handle time precision and clock drift
- Graceful shutdown

Key Challenges:
- Implement precise token refilling with time.Ticker
- Handle concurrent access to token bucket
- Implement context cancellation for Wait()
- Avoid busy waiting and optimize CPU usage

```
gotype TokenBucket struct {
    // Your implementation here
}

func NewTokenBucket(rate float64, burst int) *TokenBucket

func (tb *TokenBucket) Allow(tokens int) bool
func (tb *TokenBucket) Wait(ctx context.Context, tokens int) error
func (tb *TokenBucket) Close()
```