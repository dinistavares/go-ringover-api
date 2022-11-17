package ringover

import (
  "encoding/json"
  "errors"
  "strings"
  "time"
)

type NewContacts struct {
  Contacts []NewContact `json:"contacts"`
}

type NewContact struct {
  Firstname string       `json:"firstname"`
  Lastname  string       `json:"lastname"`
  Company   string       `json:"company"`
  IsShared  bool         `json:"is_shared"`
  Numbers   []NewNumber `json:"numbers"`
}

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
  SocialService    string      `json:"social_service"`
  SocialServiceID  string      `json:"social_service_id"`
  SocialProfileURL string      `json:"social_profile_url"`
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

type ContactFilter struct {
  Pagination        string `json:"pagination"`
  AlphabeticalOrder string `json:"alphabetical_order"`
  Search            string `json:"search"`
  LimitCount        int64  `json:"limit_count"`
  LimitOffset       int64  `json:"limit_offset"`
}

type User struct {
  UserID     int      `json:"user_id"`
  TeamID     int      `json:"team_id"`
  Initial    string   `json:"initial"`
  Color      string   `json:"color"`
  Firstname  string   `json:"firstname"`
  Lastname   string   `json:"lastname"`
  Company    string   `json:"company"`
  Email      string   `json:"email"`
  Picture    string   `json:"picture"`
  ConcatName string   `json:"concat_name"`
  Numbers    []Number `json:"numbers,omitempty"`
}

type UserResponse struct {
  UUID      string `json:"uuid"`
  Firstname string `json:"firstname"`
  Lastname  string `json:"lastname"`
  Company   string `json:"company"`
  Picture   string `json:"picture"`
  URL       string `json:"url"`
  Data      []byte   `json:"data"`
  IsShared  int    `json:"is_shared"`
}

// List Contacts by Filter
func (client *Client) ListContactsByFilter(contactfilter ContactFilter) (*[]Contact, error) {
  req, _ := client.NewRequest("POST", "/contacts", contactfilter)

  searchedContacts := SearchedContacts{}

  data, _, err := client.Do(req)

  if err != nil {
    return nil, err
  }

  json.Unmarshal(data, &searchedContacts)

  if searchedContacts.ContactList == nil || searchedContacts.ContactList[0].Numbers == nil {
    return nil, errors.New("no contacts found")
  }

  return &searchedContacts.ContactList, nil
}


// Create A New Contact
func (client *Client) CreateNewContact(newConacts NewContacts) error {
  req, _ := client.NewRequest("POST", "/contacts", newConacts)

  _, _, err := client.Do(req)

  if err != nil {
    return err
  }

  return nil
}

// Update A Contact
func (client *Client) UpdateContactByID(contactID string, newContact NewContact) error {
  url := "/contacts/" + contactID

  req, _ := client.NewRequest("PUT", url, newContact)

  _, _, err := client.Do(req)

  if err != nil {
    return err
  }

  return nil
}

// Add number to specific contact
func (client *Client) AddNewNumberToExistingContact(contactID string, newNumber NewNumber) error {
  numbers := []NewNumber{}
  numbers = append(numbers, newNumber)

  url := strings.Join([]string{"/contacts", contactID, "numbers"}, "/")

  req, _ := client.NewRequest("POST", url, numbers)

  _, _, err := client.Do(req)

  if err != nil {
    return err
  }

  return nil
}


// Delete A number
func (client *Client) DeleteNumberFromContact(contactID string, number string) error {
  url := strings.Join([]string{"/contacts", contactID, "numbers", number}, "/")

  req, _ := client.NewRequest("DELETE", url, nil)

  _, _, err := client.Do(req)

  if err != nil {
    return err
  }

  return nil
}

// Delete A Contact
func (client *Client) DeleteContactByID(contactID string) error  {
  url := "/contacts/" + contactID

  req, _ := client.NewRequest("DELETE", url, nil)

  _, _, err := client.Do(req)

  if err != nil {
    return err
  }

  return nil

}