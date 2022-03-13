# go-ringover-api

An Unofficial Ringover API Golang wrapper.

## Usage & Authentication

```go
import ringoverAPI "github.com/dinistavares/go-ringover-api"
```

To authenticate against the [Ringover REST API](https://developer.ringover.com), obtain your API Key from the Developer page in your [Ringover Dashboard](https://dashboard.ringover.com/developer). You'll need to create, copy and replace `API_KEY` in the code below.

**Keep your token API key private, and store them safely for long-term use.**

Construct a new Ringover client, and Authenticate. You can then call the availabl methods, for example: 

```go
  client := ringoverAPI.New()

  client.Authenticate(API_KEY)

	users, err := client.GetUsersInTeams()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Users: %+v\n", users ) 
	}
```
