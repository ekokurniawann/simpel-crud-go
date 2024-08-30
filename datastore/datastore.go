package datastore

import (
	"crud-simpel/models"
	"sync"
)

const shardCount = 3

type Shard struct {
	mu   sync.RWMutex
	cond *sync.Cond
	data map[string]interface{}
}

var (
	UserShards        [shardCount]*Shard
	ProductShards     [shardCount]*Shard
	TransactionShards [shardCount]*Shard
)

func init() {
	for i := 0; i < shardCount; i++ {
		shard := &Shard{
			data: make(map[string]interface{}),
		}
		shard.cond = sync.NewCond(&shard.mu)
		UserShards[i] = shard
		ProductShards[i] = shard
		TransactionShards[i] = shard
	}
}

func GetShard(id string) int {
	hash := 0
	for _, c := range id {
		hash = int(c) + hash*31
	}
	if hash < 0 {
		hash = -hash
	}
	return hash % shardCount
}

func toInterface(model interface{}) interface{} {
	return model
}

func CreateUser(user models.User) {
	shard := UserShards[GetShard(user.ID)]
	shard.mu.Lock()
	defer shard.mu.Unlock()

	shard.data[user.ID] = toInterface(user)
	shard.cond.Broadcast() 
}

func GetUser(id string) (models.User, bool) {
	shard := UserShards[GetShard(id)]
	shard.mu.RLock()
	defer shard.mu.RUnlock()

	user, exists := shard.data[id].(models.User)
	return user, exists
}

func UpdateUser(user models.User) bool {
	shard := UserShards[GetShard(user.ID)]
	shard.mu.Lock()
	defer shard.mu.Unlock()

	if _, exists := shard.data[user.ID].(models.User); exists {
		shard.data[user.ID] = toInterface(user)
		shard.cond.Broadcast()
		return true
	}
	return false
}

func DeleteUser(id string) bool {
	shard := UserShards[GetShard(id)]
	shard.mu.Lock()
	defer shard.mu.Unlock()

	if _, exists := shard.data[id]; exists {
		delete(shard.data, id)
		shard.cond.Broadcast()
		return true
	}
	return false
}

func CreateProduct(product models.Product) {
	shard := ProductShards[GetShard(product.ID)]
	shard.mu.Lock()
	defer shard.mu.Unlock()

	shard.data[product.ID] = toInterface(product)
	shard.cond.Broadcast() 
}

func GetProduct(id string) (models.Product, bool) {
	shard := ProductShards[GetShard(id)]
	shard.mu.RLock()
	defer shard.mu.RUnlock()

	product, exists := shard.data[id].(models.Product)
	return product, exists
}

func UpdateProduct(product models.Product) bool {
	shard := ProductShards[GetShard(product.ID)]
	shard.mu.Lock()
	defer shard.mu.Unlock()

	if _, exists := shard.data[product.ID].(models.Product); exists {
		shard.data[product.ID] = toInterface(product)
		shard.cond.Broadcast()
		return true
	}
	return false
}

func DeleteProduct(id string) bool {
	shard := ProductShards[GetShard(id)]
	shard.mu.Lock()
	defer shard.mu.Unlock()

	if _, exists := shard.data[id]; exists {
		delete(shard.data, id)
		shard.cond.Broadcast()
		return true
	}
	return false
}

func CreateTransaction(transaction models.Transaction) {
	shard := TransactionShards[GetShard(transaction.ID)]
	shard.mu.Lock()
	defer shard.mu.Unlock()

	shard.data[transaction.ID] = toInterface(transaction)
	shard.cond.Broadcast() 
}

func GetTransaction(id string) (models.Transaction, bool) {
	shard := TransactionShards[GetShard(id)]
	shard.mu.RLock()
	defer shard.mu.RUnlock()

	transaction, exists := shard.data[id].(models.Transaction)
	return transaction, exists
}

func UpdateTransaction(transaction models.Transaction) bool {
	shard := TransactionShards[GetShard(transaction.ID)]
	shard.mu.Lock()
	defer shard.mu.Unlock()

	if _, exists := shard.data[transaction.ID].(models.Transaction); exists {
		shard.data[transaction.ID] = toInterface(transaction)
		shard.cond.Broadcast()
		return true
	}
	return false
}

func DeleteTransaction(id string) bool {
	shard := TransactionShards[GetShard(id)]
	shard.mu.Lock()
	defer shard.mu.Unlock()

	if _, exists := shard.data[id]; exists {
		delete(shard.data, id)
		shard.cond.Broadcast()
		return true
	}
	return false
}
