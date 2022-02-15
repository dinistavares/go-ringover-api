package ringover

import (
	"encoding/json"
	"errors"
	"time"
)

type SearchedContacts struct {
	UserID            int       `json:"user_id"`
	TeamID            int       `json:"team_id"`
	LimitOffsetSetted int       `json:"limit_offset_setted"`
	LimitCountSetted  int       `json:"limit_count_setted"`
	TotalContactCount int       `json:"total_contact_count"`
	ContactListCount  int       `json:"contact_list_count"`
	ContactList       []Contact `json:"contact_list"`
}

type Contact struct {
	ContactID        int         `json:"contact_id"`
	IsShared         bool        `json:"is_shared"`
	ImOwner          bool        `json:"im_owner"`
	SocialService    interface{} `json:"social_service"`
	SocialServiceID  interface{} `json:"social_service_id"`
	SocialProfileURL interface{} `json:"social_profile_url"`
	SocialData       interface{} `json:"social_data"`
	Firstname        string      `json:"firstname"`
	Lastname         string      `json:"lastname"`
	Company          string      `json:"company"`
	ConcatName       string      `json:"concat_name"`
	Color            string      `json:"color"`
	Initial          string      `json:"initial"`
	ProfilePicture   string      `json:"profile_picture"`
	CreationDate     time.Time   `json:"creation_date"`
	Numbers          []Number    `json:"numbers"`
}

type Number struct {
	Number int64        `json:"number"`
	Type   string       `json:"type"`
	Format NumberFormat `json:"format"`
}

type NumberFormat struct {
	Raw              int64  `json:"raw"`
	CountryCode      string `json:"country_code"`
	Country          string `json:"country"`
	E164             string `json:"e164"`
	International    string `json:"international"`
	InternationalAlt string `json:"international_alt"`
	National         string `json:"national"`
	NationalAlt      string `json:"national_alt"`
	Rfc3966          string `json:"rfc3966"`
}

type NewContacts struct {
	Contacts []Contact `json:"contacts"`
}

type ContactFilter struct {
	Pagination        string `json:"pagination"`
	AlphabeticalOrder string `json:"alphabetical_order"`
	Search            string `json:"search"`
	LimitCount        int64  `json:"limit_count"`
	LimitOffset       int64  `json:"limit_offset"`
}

// Search for contacts
func (client *Client) GetorCreateContacts(contactfilter ContactFilter) ([]Contact, error) {

	req, _ := client.NewRequest("POST", "/contacts", contactfilter)

	searchedContacts := SearchedContacts{}
	contacts := []Contact{}

	data, err := client.Do(req)

	if err != nil {
		return contacts, err
	}

	json.Unmarshal(data, &searchedContacts)

	if searchedContacts.ContactList == nil {
		return contacts, errors.New("No contacts found")
	}

	return contacts, nil
}
