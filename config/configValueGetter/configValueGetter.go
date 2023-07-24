package configValueGetter

type ConfigValueGetter interface {
	GetValueByKeys(key string) string
}
