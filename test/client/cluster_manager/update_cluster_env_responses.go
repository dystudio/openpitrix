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

// UpdateClusterEnvReader is a Reader for the UpdateClusterEnv structure.
type UpdateClusterEnvReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateClusterEnvReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateClusterEnvOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateClusterEnvOK creates a UpdateClusterEnvOK with default headers values
func NewUpdateClusterEnvOK() *UpdateClusterEnvOK {
	return &UpdateClusterEnvOK{}
}

/*UpdateClusterEnvOK handles this case with default header values.

A successful response.
*/
type UpdateClusterEnvOK struct {
	Payload *models.OpenpitrixUpdateClusterEnvResponse
}

func (o *UpdateClusterEnvOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/clusters/update_env][%d] updateClusterEnvOK  %+v", 200, o.Payload)
}

func (o *UpdateClusterEnvOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixUpdateClusterEnvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
