package dictionary

type Dictionary map[string]string

func (dic *Dictionary) Search(key string) string {
	return (*dic)[key]
}
