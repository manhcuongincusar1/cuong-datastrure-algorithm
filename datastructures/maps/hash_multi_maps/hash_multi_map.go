package hashmultimaps

type IHashMultiMap[K, V comparable] interface {
	Put(key K, value V)
	PutAll(key K, value ...V)

	GetKeys() []K
	GetValues(key K) []V

	Contains(key K) bool
	ContainsAll(keys ...K) bool
	ContainsAny(keys ...K) bool
	IsEmpty() bool
	Size() int

	Clear()
	Remove(key K, value V)
	RemoveKey(keys ...K)

	Merge(maps ...*HashMultiMap[K, V])
}

type HashMultiMap[K, V comparable] struct {
	data map[K][]V
}

func New[K, V comparable]() IHashMultiMap[K, V] {
	_map := HashMultiMap[K, V]{
		data: make(map[K][]V),
	}
	return &_map
}

func (p *HashMultiMap[K, V]) Put(key K, value V)
func (p *HashMultiMap[K, V]) PutAll(key K, value ...V)
func (p *HashMultiMap[K, V]) GetKeys() []K
func (p *HashMultiMap[K, V]) GetValues(key K) []V
func (p *HashMultiMap[K, V]) Contains(key K) bool
func (p *HashMultiMap[K, V]) ContainsAll(keys ...K) bool
func (p *HashMultiMap[K, V]) ContainsAny(keys ...K) bool
func (p *HashMultiMap[K, V]) IsEmpty() bool
func (p *HashMultiMap[K, V]) Size() int
func (p *HashMultiMap[K, V]) Clear()
func (p *HashMultiMap[K, V]) Remove(key K, value V)
func (p *HashMultiMap[K, V]) RemoveKey(keys ...K)
func (p *HashMultiMap[K, V]) Merge(maps ...*HashMultiMap[K, V])
