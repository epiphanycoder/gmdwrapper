package model

type userStatus int8

type User struct {
	ID           string     `json:"id" gorm:"index"`
	FirstName    string     `json:"firstName" gorm:"-"`
	LastName     string     `json:"lastName" gorm:"-"`
	Email        string     `json:"email" gorm:"primaryKey;index;size:128"`
	Title        string     `json:"title" gorm:"-"`
	Gender       string     `json:"gender,omitempty" gorm:"-"`
	DateOfBirth  string     `json:"dateOfBirth" gorm:"-"`
	Password     string     `json:"-"`
	RegisterDate string     `json:"registerDate" gorm:"-"`
	Phone        string     `json:"phone,omitempty" gorm:"-"`
	Location     Location   `json:"location,omitempty" gorm:"-"`
	Status       userStatus `json:"-"`
}

type Location struct {
	Street   string `json:"street"`
	City     string `json:"city"`
	State    string `json:"state"`
	Country  string `json:"country"`
	Timezone string `json:"timezone""`
}

func (u *User) String() string {
	return "User"
}

const (
	userInitialized userStatus = 0
	userCreated
)
