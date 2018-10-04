// Code generated by go-swagger; DO NOT EDIT.

package cluster_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// DeleteKeyPairsReader is a Reader for the DeleteKeyPairs structure.
type DeleteKeyPairsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteKeyPairsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteKeyPairsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteKeyPairsOK creates a DeleteKeyPairsOK with default headers values
func NewDeleteKeyPairsOK() *DeleteKeyPairsOK {
	return &DeleteKeyPairsOK{}
}

/*DeleteKeyPairsOK handles this case with default header values.

A successful response.
*/
type DeleteKeyPairsOK struct {
	Payload *models.OpenpitrixDeleteKeyPairsResponse
}

func (o *DeleteKeyPairsOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/clusters/key_pairs][%d] deleteKeyPairsOK  %+v", 200, o.Payload)
}

func (o *DeleteKeyPairsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixDeleteKeyPairsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
