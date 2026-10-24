Short-lived bearer keys, kept in memory, for a job acting as someone on an API:
e.g. an AI agent (becka) calling the API back as the user it answers.

- `Store[T]`: issues a secret (`<prefix>` + 256 random bits, base64url) standing
  for a value `T` (a user, a principal…) until its TTL passes or it is revoked.
  Only the secret's SHA-256 is kept; a restart ends every key.
- `Callers`: whether a request comes straight from the job's machine (loopback,
  no `X-Forwarded-For`), so a key that leaks is of no use elsewhere.

```go
type actor struct{ UserID, KeyID string }

keys := runkey.New[actor]("kar_")
callers, err := runkey.ParseCallers(os.Getenv("AGENT_CALLER_NETWORKS")) // empty: loopback

// when the job starts
secret, revoke, err := keys.Issue(actor{UserID: user.ID, KeyID: keyID}, 5*time.Minute)
defer revoke()
// hand secret to the agent, which sends "Authorization: Bearer kar_…"

// in the API's auth middleware
if keys.IsKey(secret) {
	a, err := keys.Authenticate(secret) // runkey.ErrInvalidKey: unknown, revoked or expired
	if err != nil || !callers.Allow(r) {
		// 401
	}
}

// a user unlinks an integration: end its keys
keys.RevokeWhere(func(a actor) bool { return a.KeyID == keyID })
```

`Callers.Allow` refuses any request with `X-Forwarded-For`: the reverse proxy in
front of the API must set it (Go's `httputil.ReverseProxy` always does), or
outside traffic it forwards from the same machine looks local.
