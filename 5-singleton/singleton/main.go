// Singleton pattern on a package-level via global variables
package singleton

import (
	"fmt"
	"sync"
	"time"
)

type config struct {
	myKey int
}

var configInstance *config

func GetConfig() *config {
	if configInstance == nil {
		fmt.Println("Creating config...")
		configInstance = &config{}
	}

	return configInstance
}

func (c *config) GetMyKey() int {return c.myKey}
func (c *config) SetMyKey(value int) {c.myKey = value}


var ConfigLock = &sync.Mutex{}

func main() {
	for i := 0; i < 100; i++ {
		counter := i
		go func(j int) {
			ConfigLock.Lock()			// TODO: Try commenting out and see what happens
			defer ConfigLock.Unlock()   // TODO: Try commenting out and see what happens
			GetConfig().SetMyKey(j)
			key := GetConfig().GetMyKey()
			fmt.Println(fmt.Sprintf("j: %d, Key: %d", j, key))
		}(counter)
	}

	time.Sleep(5 * time.Second)
}
