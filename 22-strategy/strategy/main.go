package strategy

import "fmt"

const (
	strategyDefault = strategyETCD
	strategyETCD = iota + 1
	strategyRedis
)

type LockFactory struct {}
func (*LockFactory) newLock(strategy int) Lock {
	strategyLockMap := map[int]func()Lock{
		strategyETCD: func()Lock{return &ETCDLock{}},
		strategyRedis: func()Lock{return &RedisLock{}},
	}
	if strategyFunc, ok := strategyLockMap[strategy]; ok {
		return strategyFunc()
	}
	return strategyLockMap[strategyDefault]()
}

type Lock interface {
	Lock(key string) error
	Unlock(key string)
}
type ETCDLock struct {}
func (e *ETCDLock) Lock(key string) error {
	fmt.Printf("ETCD Locking %s\n", key)
	return nil
}
func (e *ETCDLock) Unlock(key string) {
	fmt.Printf("ETCD Unlocking %s\n", key)
}

type RedisLock struct {}
func (r *RedisLock) Lock(key string) error {
	fmt.Printf("Redis Locking %s\n", key)
	return nil
}
func (r *RedisLock) Unlock(key string) {
	fmt.Printf("Redis Unlocking %s\n", key)
}

type App struct {
	region string
	lockFactory *LockFactory
}
func (a *App) orderQueueProcessor(orderID string)  {
	var lock Lock
	if a.region == "sg" {
		lock = a.lockFactory.newLock(strategyRedis)
	} else {
		lock = a.lockFactory.newLock(strategyETCD)
	}
	err := lock.Lock(orderID)
	if err != nil {
		fmt.Println("Unable to acquire lock after timeout")
		return
	}
	defer lock.Unlock(orderID)
	fmt.Println("Processing order...")
}
