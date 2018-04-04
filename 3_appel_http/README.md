# APPEL HTTP

Ecrivez un programme qui appelle votre service ecrit lors de l'étape précédente.

Dans le package `net/http` vous trouverez les fonctions associées aux methodes HTTP (GET, POST, DELETE).

Notament:

```
func Get(url string) (resp *Response, err error)
```


```
func Post(url string, contentType string, body io.Reader) (resp *Response, err error)
```

Ces fonctions renvoient une erreur ainsi qu'un pointeur sur une `http.Response` qui contient les champs suivant:

```
type Response struct {
        Status              string // e.g. "200 OK"
        StatusCode          int    // e.g. 200
        Proto               string // e.g. "HTTP/1.0"
        ProtoMajor          int    // e.g. 1
        ProtoMinor          int    // e.g. 0

        Header              Header
        Body                io.ReadCloser
        ContentLength       int64
        TransferEncoding    []string
        Close               bool
        Uncompressed        bool
        Trailer             Header
        Request             *Request
        TLS                 *tls.ConnectionState
}
```