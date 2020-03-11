# B"H



`curl` commnads to run:

```sh
# Note how even though you didn’t set the content type, it was detected and set correctly:
curl -i 127.0.0.1:8080/write

# Return the 501 Not Implemented status code
curl -i 127.0.0.1:8080/writeheader

# Redirect example (try in browser as well).
curl -i 127.0.0.1:8080/redirect

curl -i 127.0.0.1:8080/json

```