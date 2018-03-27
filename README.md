# field-mask-test

## Running

```bash
$ go run main.go
```

In a terminal:
```bash
$ curl -H "Content-Type: application/json" -X PATCH -d '{"display_name": "user3"}' localhost:11000/api/v1/user2?update_mask=display_name
```
