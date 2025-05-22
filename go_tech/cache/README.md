## Thread-Safe LRU Cache
**Time**: 50 minutes | Difficulty: Hard
**Problem**: Implement a thread-safe LRU cache that supports concurrent reads/writes with high performance.

**Requirements**:
- Get(key) (value, bool) - O(1) average case
- Put(key, value) - O(1) average case
- Delete(key) bool - O(1) average case
- Thread-safe for concurrent access
- Efficient memory usage
- Handle cache capacity limits

**Key Challenges**:
- Choose between sync.RWMutex vs sync.Map vs channels
- Implement doubly-linked list + hashmap
- Handle race conditions in eviction
- Optimize for read-heavy vs write-heavy workloads

```
type LRUCache struct {
    // Your implementation here
}

func NewLRUCache(capacity int) *LRUCache

func (c *LRUCache) Get(key string) (interface{}, bool)
func (c *LRUCache) Put(key string, value interface{})
func (c *LRUCache) Delete(key string) bool
```