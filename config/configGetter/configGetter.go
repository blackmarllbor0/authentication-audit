package configGetter

type ConfigGetter interface {
	GetValueByKey(key string) string
}
