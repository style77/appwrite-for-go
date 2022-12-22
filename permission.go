package appwrite

type Permission struct{}

func (Permission) Read(role string) string {
	return "read(\"" + role + "\")"
}

func (Permission) Write(role string) string {
	return "write(\"" + role + "\")"
}

func (Permission) Create(role string) string {
	return "create(\"" + role + "\")"
}

func (Permission) Update(role string) string {
	return "update(\"" + role + "\")"
}

func (Permission) Delete(role string) string {
	return "delete(\"" + role + "\")"
}