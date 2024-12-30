package util

func CopyMap[K comparable, V any](orig map[K]V) map[K]V {
	mapCopy := make(map[K]V)
	for k, v := range orig {
		mapCopy[k] = v
	}
	return mapCopy
}
