package customer

import (
	"errors"
	"strings"
	"time"
)

type Customer struct {
	uniqId   string
	name     string
	joinDate time.Time
}

type DTOCustomer struct {
	UniqId   string
	Name     string
	JoinDate string
}

func NewCustomer(cust *DTOCustomer, store string) (*Customer, error) {
	var timeJoinDate time.Time

	if cust.Name == "" {
		return nil, errors.New("name cannot be empty")
	}
	if len([]rune(cust.Name)) < 3 {
		return nil, errors.New("name cannot be less than 3 character")
	}
	if cust.JoinDate != "" {
		timeJoinDate, _ = time.Parse("2006-01-02 15:04:05.", cust.JoinDate)
	}

	if cust.UniqId == "" {
		timeJoinDate = time.Now()

		charsName := []rune(cust.Name)
		charName := string(charsName[0]) + string(charsName[1]) + string(charsName[2])
		upCharName := strings.ToUpper(charName)

		cust.UniqId = upCharName + "-" + store + "-" + timeJoinDate.Format("020106150405")
	}

	return &Customer{
		uniqId:   cust.UniqId,
		name:     cust.Name,
		joinDate: timeJoinDate,
	}, nil
}

func (cust *Customer) GetValueUniqIdCust() string {
	return cust.uniqId
}

func (cust *Customer) GetValueNameCust() string {
	return cust.name
}

func (cust *Customer) GetValueJoinDateCust() string {
	return cust.joinDate.Format("2006-01-02 15:04:05.")
}
