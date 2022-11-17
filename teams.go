package ringover

import (
  "encoding/json"
  "errors"
)

type Team struct {
  TeamID            int         `json:"team_id"`
  Name              string      `json:"name"`
  TotalNumbersCount int         `json:"total_numbers_count"`
  Numbers           []Number    `json:"numbers"`
  TotalUsersCount   int         `json:"total_users_count"`
  Users             []User      `json:"users"`
  TotalIvrsCount    int         `json:"total_ivrs_count"`
  Ivrs              interface{} `json:"ivrs"`
  TotalTagsCount    int         `json:"total_tags_count"`
  Tags              interface{} `json:"tags"`
  TotalGroupsCount  int         `json:"total_groups_count"`
  Groups            []Groups    `json:"groups"`
}

type Groups struct {
  GroupID         int         `json:"group_id"`
  Name            string      `json:"name"`
  TotalUsersCount int         `json:"total_users_count"`
  Color           interface{} `json:"color"`
  IsJumper        bool        `json:"is_jumper"`
}

// Get Users in Teams
func (client *Client) GetUsersInTeams() (*[]User, error) {
  req, _ := client.NewRequest("GET", "/teams", nil)

  team := Team{}

  data, _, err := client.Do(req)

  if err != nil {
    return nil, err
  }

  json.Unmarshal(data, &team)

  if team.Users == nil {
    return nil, errors.New("no users found")
  }

  return &team.Users, nil
}
