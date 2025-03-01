package versions

import (
	"ocpi-cpo-mock-server/src/core/modules/controls"
	"ocpi-cpo-mock-server/src/core/modules/response"
	"time"
)

type Version struct {
	URL           string        `json:"url"`
	VersionNumber VersionNumber `json:"version_number"`
}

type VersionNumber string

const (
	VersionN20  VersionNumber = "2.0"
	VersionN21  VersionNumber = "2.1"
	VersionN211 VersionNumber = "2.1.1"
	VersionN22  VersionNumber = "2.2"
	VersionN221 VersionNumber = "2.2.1"
)

type VersionsResponse = response.BaseResponse[[]Version]

func NewVersionsResponse(data []Version, statusCode response.StatusCode, statusMessage string) VersionsResponse {
	return VersionsResponse{
		Data:          data,
		StatusCode:    statusCode,
		StatusMessage: statusMessage,
		TimeStamp:     time.Now().Format(time.RFC3339),
	}
}

type ListVersionsUsecase struct {
}

func NewListVersionsUsecase() *ListVersionsUsecase {
	return &ListVersionsUsecase{}
}

func (uc *ListVersionsUsecase) Execute(control *controls.Control) VersionsResponse {

	versions := []Version{
		{URL: "https://example.com/ocpi/2.0", VersionNumber: VersionN20},
		{URL: "https://example.com/ocpi/2.1", VersionNumber: VersionN21},
		{URL: "https://example.com/ocpi/2.2", VersionNumber: VersionN22},
	}

	if control.ResponseType == controls.ResponseTypeControlValueEmpty {
		versions = []Version{}
	}

	if control.ResponseType == controls.ResponseTypeControlValueError {
		return NewVersionsResponse([]Version{}, response.StatusCodeGenericServerError, "An error occurred")
	}

	return NewVersionsResponse(versions, response.StatusCodeGenericSuccess, "Success")
}
