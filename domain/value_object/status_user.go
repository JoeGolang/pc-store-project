package value_object

import "fmt"

type StatusUser struct {
	value int
}

func NewStatusUser(user int) (*StatusUser, error) {
	su := StatusUser{value: user}
	err := su.validate()
	if err != nil {
		return &StatusUser{}, err
	}
	return &su, nil
}

func StatusUserInt(user string) int {
	switch user {
	case "Owner":
		return 1
	case "Employee":
		return 2
	default:
		return 0
	}
}

func (su *StatusUser) validate() error {
	switch su.value {
	case 1:
		return nil
	case 2:
		return nil
	default:
		return fmt.Errorf("status user tidak ditemukan")
	}
}

func (su *StatusUser) StatusUserToString() string {
	switch su.value {
	case 1:
		return "Owner"
	case 2:
		return "Employee"
	default:
		return ""
	}
}
