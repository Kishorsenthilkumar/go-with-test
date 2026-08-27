package Maps

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot perform operation on word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (dict Dictionary) Search(key string) (string, error) {

	value, noKey := dict[key]

	if !noKey {
		return "", ErrNotFound
	}
	return value, nil
}

func (dict Dictionary) Add(key, value string) error {
	_, err := dict.Search(key)

	switch err {
	case ErrNotFound:
		dict[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (dict Dictionary) Update(key, value string) error {
	_, err := dict.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		dict[key] = value
	default:
		return err
	}
	return nil
}

func (dict Dictionary) Delete(key string) error {
	_, err := dict.Search(key)

	switch err {
	case nil:
		delete(dict, key)
	case ErrNotFound:
		return ErrWordExists
	default:
		return err

	}
	return nil
}
