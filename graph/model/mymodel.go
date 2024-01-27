package model

import (
	"fmt"
	"io"
	"net/url"

	"github.com/99designs/gqlgen/graphql"
)

// MarshalURI marshals a url.URL into a graphql.Marshaler
// third party code can use this to implement custom scalars
// format: Marshal{Xxx}
func MarshalURI(u url.URL) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, fmt.Sprintf("%s", u.String()))
	})
}

// UnmarshalURI unmarshals a graphql.Marshaler into a url.URL
// format: Unmarshal{Xxx}
func UnmarshalURI(v interface{}) (url.URL, error) {
	switch v := v.(type) {
	case string:
		u, err := url.Parse(v)
		if err != nil {
			return url.URL{}, err
		}

		return *u, nil
	case []byte:
		u := url.URL{}
		if err := u.UnmarshalBinary(v); err != nil {
			return url.URL{}, err
		}

		return u, nil
	default:
		return url.URL{}, fmt.Errorf("%T is not a valid url.URL", v)
	}
}
