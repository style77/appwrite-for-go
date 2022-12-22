package appwrite

type ID struct{}

func (ID) Unique() string {
	return "unique()"
}

func (ID) Custom(id string) string {
	return id
}