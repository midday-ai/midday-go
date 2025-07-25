// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/midday-ai/midday-go/models/components"
)

type UpdateTagRequestBody struct {
	// The new name of the tag.
	Name string `json:"name"`
}

func (o *UpdateTagRequestBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type UpdateTagRequest struct {
	ID          string                `pathParam:"style=simple,explode=false,name=id"`
	RequestBody *UpdateTagRequestBody `request:"mediaType=application/json"`
}

func (o *UpdateTagRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateTagRequest) GetRequestBody() *UpdateTagRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type UpdateTagResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Tag updated
	TagResponse *components.TagResponse
}

func (o *UpdateTagResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateTagResponse) GetTagResponse() *components.TagResponse {
	if o == nil {
		return nil
	}
	return o.TagResponse
}
