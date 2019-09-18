package maps

type Content map[string]string

const (
	ErrNotFound = ContentErr("word not found")
	ErrWordExists = ContentErr("cannot add word because it already exists")
	ErrWordNotExists = ContentErr("cannot update word because it does not exist")
)

type ContentErr string

func (e ContentErr) Error() string {
	return string(e)
}

func (c Content) Search(word string) (string, error) {
	found := c[word]

	if len(found) > 1 {
		return found, nil
	}
	return "", ErrNotFound
}

func (c Content) Add(word string, definition string) error{
	_, err := c.Search(word)

	switch err {
	case ErrNotFound:
		c[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (c Content) Update(word, definition string) error{
	_, err := c.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordNotExists
	case nil:
		c[word] = definition
	default:
		return err
	}
	return nil
}
