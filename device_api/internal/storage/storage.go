package storage

import (
	"sync"
	"log/slog"

	"github.com/Grifonhard/les/device_api/internal/devices"
)

var devices = make(map[string] devices.Device)

var mu sync.Mutex

//adding another device to the storage
func Push(d device.Device, logg *slog.Logger) bool {
	_, b := devices[d.Id]
	if !b {
		return b
	}
	mu.Lock()
	devices[d.Id] = d
	mu.Unlock()
	logg.Info("device push: ", slog.String("ID", d.ID), slog.String("Language", d.Language), slog.Float64("Latitude", d.Latitude), slog.Longitude("Longitude", d.Longitude), slog.String("Os_name", d.Os_name), slog.String("Os_version", d.Os_version))
	return b
}

//issuing data upon request
func TakeFrom (id string, d* device.Device) bool {
	*d, b = devices[id]
	return b
} 