package device

import (
	"math/rand"
	"time"
)

type Device struct {
	Id string `json:"id"`
	Language string `json:"language"`
	Latitude float64 `json: "latitude"` 
	Longitude float64 `json: "longitude"`
	Os_name string `json: "os_name"`
	Os_version string `json: "os_version"`
}

var devices = []Device {
	{	Id: "000000-00000-00000-000001", Language : "ru", Latitude : 80.023, Longitude: 101.23,	Os_name : "Android", Os_version : "1.23.3"},
	{	Id: "000000-00000-00000-000002", Language : "en", Latitude : 56.023, Longitude: 121.23,	Os_name : "Android", Os_version : "1.24.3"},
	{	Id: "000000-00000-00000-000003", Language : "si", Latitude : 80.85, Longitude: 301.23,	Os_name : "Android", Os_version : "1.23.4"},
	{	Id: "000000-00000-00000-000004", Language : "ru", Latitude : 80.44, Longitude: 101.45,	Os_name : "Android", Os_version : "1.27.6"},
	{	Id: "000000-00000-00000-000005", Language : "fr", Latitude : 80.21, Longitude: 121.23,	Os_name : "Android", Os_version : "1.21.3"},
	{	Id: "000000-00000-00000-000006", Language : "ru", Latitude : 81.6, Longitude: 141.23,	Os_name : "Android", Os_version : "1.26.9"},
	{	Id: "000000-00000-00000-000007", Language : "la", Latitude : 80.023, Longitude: 101.25,	Os_name : "Android", Os_version : "1.23.3"},
	{	Id: "000000-00000-00000-000008", Language : "ru", Latitude : 36.023, Longitude: 111.23,	Os_name : "Android", Os_version : "2.01.3"},
	{	Id: "000000-00000-00000-000009", Language : "en", Latitude : 13.023, Longitude: 1651.23,	Os_name : "Android", Os_version : "3.02.3"},
	{	Id: "000000-00000-00000-000010", Language : "ru", Latitude : 45.023, Longitude: 16.23,	Os_name : "Android", Os_version : "1.34.3"},
}

var source = rand.NewSource(time.Now().UnixMicro())														
var sRand = rand.New(source)


//отправляемые данные
func DeviceGet () Device {
	i := sRand.Intn(10)
	d := devices[i]
	return d
}