package decohandler

import "net/http"

// DecoratorFunc -
type DecoratorFunc func (http.ResponseWriter, *http.Request, http.Handler)

// DecoHandler -
type DecoHandler struct {
	fn DecoratorFunc
	h http.Handler
}

func (d *DecoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d.fn(w, r, d.h)
}

