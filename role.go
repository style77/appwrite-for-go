package appwrite

import ("fmt")

type Role struct{}

func (Role) Any() string {
	return "any"
}

func (Role) User(id string, status string) string {
	if status == "" {
		return fmt.Sprintf("user:%s", id)
	}
	return fmt.Sprintf("user:%s/%s", id, status)
}

func (Role) Users(status string) string {
	if status == "" {
		return "users"
	}
	return fmt.Sprintf("users/%s", status)
}

func (Role) Guests() string {
	return "guests"
}

func (Role) Team(id string, role string) string {
	if role == "" {
		return fmt.Sprintf("team:%s", id)
	}
	return fmt.Sprintf("team:%s/%s", id, role)
}

func (Role) Member(id string) string {
	return fmt.Sprintf("member:%s", id)
}