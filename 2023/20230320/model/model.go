package model

import "sync"

var meterMapIns = &meterMapManager{
	meters:make(map[string]string),
	mutex: new(sync.RWMutex),
}

type meterMapManager struct {
	meters map[string]string
	mutex *sync.RWMutex
}

func (m *meterMapManager) setMetersMap(meter map[string]string) {
	m.mutex.Lock()
	m.meters = meter
	m.mutex.Unlock()
}

func (m *meterMapManager) getMetersMap() map[string]string {
	m.mutex.RLock()
	meters := m.meters
	m.mutex.RUnlock()
	return meters
}

func (m *meterMapManager) getMeterByKey(key string) string {
	m.mutex.RLock()
	value := m.meters[key]
	m.mutex.RUnlock()
	return value
}

func (m *meterMapManager) clearMeterByKey(key string) {
	m.mutex.Lock()
	delete(m.meters, key)
	m.mutex.Unlock()
}

func SetMetersMap(meter map[string]string) {
	meterMapIns.setMetersMap(meter)
}

func GetMetersMap() map[string]string {
	return meterMapIns.getMetersMap()
}

func GetMeterByKey(key string) string {
	return meterMapIns.getMeterByKey(key)
}

func CheckMeterExists(key string) bool {
	metersMap := meterMapIns.getMetersMap()
	if _, exists := metersMap[key]; exists {
		return exists
	}
	return false
}

func ClearMetersMapByKey(key string) {
	meterMapIns.clearMeterByKey(key)
}