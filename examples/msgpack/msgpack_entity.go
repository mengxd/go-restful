package restPack

import (
	restful "github.com/emicklei/go-restful"
	"gopkg.in/vmihailenco/msgpack.v2"
)

const MIME_MSGPACK = "application/x-msgpack" // Accept or Content-Type used in Consumes() and/or Produces()

// NewEntityAccessorMPack returns a new EntityReaderWriter for accessing MessagePack content.
// This package is not initialized with such an accessor using the MIME_MSGPACK contentType.
func NewEntityAccessorMsgPack(contentType string) restful.EntityReaderWriter {
	return entityMsgPackAccess{ContentType: contentType}
}

// entityOctetAccess is a EntityReaderWriter for Octet encoding
type entityMsgPackAccess struct {
	// This is used for setting the Content-Type header when writing
	ContentType string
}

// Read unmarshalls the value from byte slice and using msgpack to unmarshal
func (e entityMsgPackAccess) Read(req *restful.Request, v interface{}) error {
	return msgpack.NewDecoder(req.Request.Body).Decode(v)
}

// Write marshals the value to byte slice and set the Content-Type Header.
func (e entityMsgPackAccess) Write(resp *restful.Response, status int, v interface{}) error {
	return writeMsgPack(resp, status, e.ContentType, v)
}

// writeMsgPack marshals the value to byte slice and set the Content-Type Header.
func writeMsgPack(resp *restful.Response, status int, contentType string, v interface{}) error {
	if v == nil {
		resp.WriteHeader(status)
		// do not write a nil representation
		return nil
	}
	resp.Header().Set(restful.HEADER_ContentType, contentType)
	resp.WriteHeader(status)
	return msgpack.NewEncoder(resp).Encode(v)
}
