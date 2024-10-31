package dto

type Usertype string

const (
	// user roles
	JobSeeker Usertype = "Employee"
	Employer  Usertype = "Recuriter"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	RoleId   int    `json:"role_id"`
}
