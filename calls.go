package ringover

// Get All Calls
func (client *Client) ListAllCalls() ([]byte, error) {
  req, _ := client.NewRequest("GET", "/calls", nil)

  data, _, err := client.Do(req)

  if err != nil {
    return nil, err
  }

  return data, nil
}
