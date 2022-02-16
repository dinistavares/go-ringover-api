# go-ringover-api

An Unofficial Ringover API Golang wrapper.

## Usage & Authentication

```go
import ringoverAPI "example.com/test/go-ringover-api"
```

To authenticate against the API, obtain your API Key from your [Ringover Dashboard](https://dashboard.ringover.com/developer). You'll need to create, copy and replace `API_KEY` in the code above.

**Keep your token API key private, and store them safely for long-term use.**

Construct a new Ringover client, and Authenticate for example: 

```go
client := ringoverAPI.New()

client.Authenticate(API_KEY)
```
