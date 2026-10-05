package dictionary

type Dictionary map[string]string

var (
	ErrorDictKeyNotFound  = DictionaryErr("could not found the word you were looking for")
	ErrorDictKeyExists    = DictionaryErr("word already exists")
	ErrorDictKeyNotExists = DictionaryErr("cannot perform operation on word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (dic *Dictionary) Search(key string) (string, error) {
	value, ok := (*dic)[key]

	if !ok {
		return "", ErrorDictKeyNotFound
	}

	return value, nil
}

func (dic *Dictionary) Add(key, value string) error {
	_, err := dic.Search(key)

	switch err {
	case ErrorDictKeyNotFound:
		(*dic)[key] = value
	case nil:
		return ErrorDictKeyExists
	default:
		return err
	}

	return nil
}

func (dic *Dictionary) Update(key, value string) error {
	_, exist := dic.Search(key)

	switch exist {
	case ErrorDictKeyNotFound:
		return ErrorDictKeyNotExists
	case nil:
		(*dic)[key] = value
	default:
		return exist
	}

	return nil
}

func (dic *Dictionary) Delete(key string) error {
	_, err := (*dic).Search(key)

	switch err {
	case ErrorDictKeyNotFound:
		return ErrorDictKeyNotExists
	case nil:
		delete((*dic), key)
	default:
		return err
	}

	return nil
}
