package httpex

import (
	"net/http"
)

type Response struct {
	http.Response
	BinaryBody []byte // 包体二进制
}

func newResponse(response *http.Response, binaryBody []byte) *Response {
	r := &Response{}
	r.setResponse(response, binaryBody)

	return r
}

func (r *Response) setResponse(response *http.Response, binaryBody []byte) {
	r.BinaryBody = binaryBody

	r.Status = response.Status
	r.StatusCode = response.StatusCode
	r.Proto = response.Proto
	r.ProtoMajor = response.ProtoMajor
	r.ProtoMinor = response.ProtoMinor
	r.Header = response.Header
	r.ContentLength = response.ContentLength
	r.TransferEncoding = response.TransferEncoding
	r.Close = response.Close
	r.Uncompressed = response.Uncompressed
	r.Trailer = response.Trailer
	r.Request = response.Request
	r.TLS = response.TLS
}
