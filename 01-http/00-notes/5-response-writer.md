# B"H


### `ResponseWriter`

Let's revisit the `http.ResponseWriter` used in handler functions. 

Example:

```go
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}
```

---

**`ResponseWriter`** 
- It is an **interface** that a **handler** uses to create an **HTTP response**.
- The actual **struct** backing up `ResponseWriter` is the nonexported struct `http.response` - you can’t use it directly; you can only use it through the `ResponseWriter` interface. (Note, it is actually a pointer to a stuct behind the scenes).

---

The `ResponseWriter` interface has three methods:

1. `Write`
    - This takes in an array of bytes, and this gets written into the body of the HTTP response. 
    - If the header doesn’t have a content type by the time Write is called, the first 512 bytes of the data are used to detect the content type. 

2. `WriteHeader`
    - **Note** name is **misleading**. It doesn’t allow you to write any headers (you use Header for that), but it takes an integer that represents the **status code** of the **HTTP response**.
    - After calling this method, you can still write to the ResponseWriter, **though you can no longer write to the header**. 
    - If you don’t call this method, by default when you call the Write method, 200 OK will be sent as the response code.

3. `Header`
    - This method returns a map of **HTTP Response Headers** that you can modify (**see** `3-http-response.md`).
    - Note, this can be used to create a HTTP redirect.

---

**SEE:** `11-response-writer/server.go` for details on usage.   