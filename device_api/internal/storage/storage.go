package storage

import (
	"sync"
	"log/slog"

	"github.com/Grifonhard/les/device_api/internal/devices"
)

var deviceStorage = make(map[string] devices.Device)

var mu sync.Mutex

//adding another device to the storage
func Push(d* devices.Device, logg *slog.Logger) bool {
	_, b := deviceStorage[d.Id]
	if b {
		return !b
	}
	mu.Lock()
	deviceStorage[d.Id] = *d
	mu.Unlock()
	logg.Info("device push: ", slog.String("ID", d.Id), slog.String("Language", d.Language), slog.Float64("Latitude", d.Latitude), slog.Float64("Longitude", d.Longitude), slog.String("Os_name", d.Os_name), slog.String("Os_version", d.Os_version))
	return !b
}

//issuing data upon request
func TakeFrom (id string, d* devices.Device) bool {
	var b bool
	*d, b = deviceStorage[id]
	return b
} 