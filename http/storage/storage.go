package storage

import (
	"fmt"
	"strings"
	"sync"

	"github.com/Grifonhard/les/http/device"
)

var devices []device.Device

var mu sync.Mutex

//добавление очередного девайса в хранилище
func Push(d device.Device) bool {
	if seach(d.Id) != -1 {
		return false
	}
	mu.Lock()
	devices = append(devices, d)
	mu.Unlock()
	fmt.Println(devices)
	return true
}

//поиск девайса в хранилище
func seach(id string) int {
	for i, di := range devices {
		b := strings.EqualFold(di.Id, id)
		if b {
			return i
		}
	}
	return -1
}

//выдача данных по запросу
func TakeFrom (id string, d* device.Device) bool {
	i := seach(id)
	if i == -1 {
		return false
	}
	*d = devices[i]
	return true
} 