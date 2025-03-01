package versions

import (
	"ocpi-cpo-mock-server/src/core/modules/controls"
	"ocpi-cpo-mock-server/src/core/modules/response"
	"time"
)

type EndpointRole string

const (
	EndpointRoleSender   EndpointRole = "SENDER"
	EndpointRoleReceiver EndpointRole = "RECEIVER"
)

type ModuleID string

const (
	ModuleIDCdrs             ModuleID = "cdrs"
	ModuleIDChargingProfiles ModuleID = "chargingprofiles"
	ModuleIDCommands         ModuleID = "commands"
	ModuleIDCredentials      ModuleID = "credentials"
	ModuleIDHubclientInfo    ModuleID = "hubclientinfo"
	ModuleIDLocations        ModuleID = "locations"
	ModuleIDSessions         ModuleID = "sessions"
	ModuleIDTariffs          ModuleID = "tariffs"
	ModuleIDTokens           ModuleID = "tokens"
	ModuleIDVersions         ModuleID = "versions"
)

type Endpoint struct {
	Identifier ModuleID     `json:"identifier"`
	Role       EndpointRole `json:"role"`
	URL        string       `json:"url"`
}

type VersionDetails struct {
	VersionNumber VersionNumber `json:"version_number"`
	Endpoints     []Endpoint    `json:"endpoints"`
}

type VersionDetailsResponse = response.BaseResponse[VersionDetails]

type GetVersionDetailsUsecase struct {
}

func NewGetVersionDetailsUsecase() *GetVersionDetailsUsecase {
	return &GetVersionDetailsUsecase{}
}

func (uc *GetVersionDetailsUsecase) Execute(version VersionNumber, control *controls.Control) VersionDetailsResponse {
	endpoints := []Endpoint{
		{
			Identifier: ModuleIDCdrs,
			Role:       EndpointRoleSender,
			URL:        "https://example.com/ocpi/2.0/cdrs",
		},
		{
			Identifier: ModuleIDChargingProfiles,
			Role:       EndpointRoleReceiver,
			URL:        "https://example.com/ocpi/2.0/chargingprofiles",
		},
		{
			Identifier: ModuleIDCommands,
			Role:       EndpointRoleReceiver,
			URL:        "https://example.com/ocpi/2.0/commands",
		},
		{
			Identifier: ModuleIDCredentials,
			Role:       EndpointRoleReceiver,
			URL:        "https://example.com/ocpi/2.0/credentials",
		},
		{
			Identifier: ModuleIDCredentials,
			Role:       EndpointRoleSender,
			URL:        "https://example.com/ocpi/2.0/credentials",
		},
		{
			Identifier: ModuleIDHubclientInfo,
			Role:       EndpointRoleReceiver,
			URL:        "https://example.com/ocpi/2.0/hubclientinfo",
		},
		{
			Identifier: ModuleIDLocations,
			Role:       EndpointRoleSender,
			URL:        "https://example.com/ocpi/2.0/locations",
		},
		{
			Identifier: ModuleIDSessions,
			Role:       EndpointRoleSender,
			URL:        "https://example.com/ocpi/2.0/sessions",
		},
		{
			Identifier: ModuleIDTariffs,
			Role:       EndpointRoleSender,
			URL:        "https://example.com/ocpi/2.0/tariffs",
		},
		{
			Identifier: ModuleIDTokens,
			Role:       EndpointRoleReceiver,
			URL:        "https://example.com/ocpi/2.0/tokens",
		},
		{
			Identifier: ModuleIDVersions,
			Role:       EndpointRoleReceiver,
			URL:        "https://example.com/ocpi/2.0/versions",
		},
		{
			Identifier: ModuleIDVersions,
			Role:       EndpointRoleSender,
			URL:        "https://example.com/ocpi/2.0/versions",
		},
	}

	return VersionDetailsResponse{
		Data: VersionDetails{
			VersionNumber: version,
			Endpoints:     endpoints,
		},
		StatusCode:    response.StatusCodeGenericSuccess,
		StatusMessage: "Success",
		TimeStamp:     time.Now().Format(time.RFC3339),
	}
}
