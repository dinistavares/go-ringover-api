package ringover

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

type NewNumber struct {
  Number int64  `json:"number"`
  Type   string `json:"type"`
}