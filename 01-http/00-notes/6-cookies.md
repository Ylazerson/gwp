# B"H


### Cookies

- A cookie is a small piece of information that’s stored at the **client**, originally sent from the **server** through an **HTTP response** message. 
- **Every time** the client sends an HTTP request to the server, the cookie is sent along with it. 
- Cookies are designed to overcome the stateless-ness of HTTP. Although it’s not the only mechanism that can be used, it’s one of the most common and popular methods.

---

- There are a number of types of cookies, such as: 
    - **super cookies** 
    - **third-party cookies** 
    - **zombie cookies** 
    - etc.

- But generally there are only **two** classes of cookies: 
    1. **session cookies** 
    2. **persistent cookies** (most other types of cookies are variants of the persistent cookies).

---

The `Cookie` struct is the representation of a cookie in Go:

```go
type Cookie struct {
  Name       string
  Value      string
  Path       string
  Domain     string
  Expires    time.Time
  RawExpires string
  MaxAge     int
  Secure     bool
  HttpOnly   bool
  Raw        string
  Unparsed   []string
}
```

- If the `Expires` field isn’t set, then the cookie is a temp **session** cookie
- Session cookies are removed from the browser when the browser is closed. 

- There are two ways of specifying the expiry time: `Expires` and `MaxAge`: 
    - `Expires` tells us exactly when the cookie will expire. 
    - `MaxAge` tells us how long the cookie should last (in seconds). 
    - `Expires` was deprecated but almost all browsers still support it. `MaxAge` isn’t supported by Microsoft Internet Explorer < 8. The pragmatic solution is to use only `Expires` or to use both in order to support all browsers.