# `github.com/cdzombak/gotfy`

`gotfy` is an API client for sending notifications using [ntfy](https://ntfy.sh) servers.

## Installation

```shell
go get github.com/cdzombak/gotfy
```

## Example Usage

```go
serverURL, _ := url.Parse("https://ntfy.example.com")

publisher := gotfy.NewPublisher(PublisherOpts{
    Server:  serverURL,
    Auth:    gotfy.AccessToken("tk_0123456789"),
	Headers: http.Header{
		"User-Agent": []string{"my-app / 1.0"},
    },
})

ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
defer cancel()

publisher.SendMessage(ctx, &gotfy.Message{
    Topic:    "topic",
    Email:    "me@example.com",
    Message:  "message",
    Title:    "title",
    Tags:     []string{"emoji1","emoji2","some text"},
    Priority: gotfy.PriorityHigh,
    Actions:  []gotfy.ViewAction {
	    Label: "View Btn",
	    Link:  "https://view.example.com",
	    Clear: true,
    },
    ClickURL: "https://click.example.com",
    IconURL:  "https://icon.example.com",
    Delay:    time.Minute * 5,
})
```

## License & Authors

gotfy is licensed under the Apache 2.0 license; see [LICENSE](LICENSE) in this repository.

- Original library copyright [AnthonyHewins](https://github.com/AnthonyHewins)
- Fork changes copyright 2023-2024 [Chris Dzombak](https://www.dzombak.com)
