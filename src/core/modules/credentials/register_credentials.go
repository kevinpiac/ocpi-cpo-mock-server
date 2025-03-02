package credentials

import (
	"log"
	"ocpi-cpo-mock-server/src/core/modules/controls"
	"ocpi-cpo-mock-server/src/core/modules/env"
	"ocpi-cpo-mock-server/src/core/modules/response"
	"ocpi-cpo-mock-server/src/core/modules/versions"
	"strings"
	"time"
)

type Logo struct {
	Category  string `json:"category"`
	Type      string `json:"type"`
	URL       string `json:"url"`
	Thumbnail string `json:"thumbnail"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
}

type BusinessDetails struct {
	Name    string `json:"name"`
	Website string `json:"website"`
	Logo    Logo   `json:"logo"`
}

type Role string

// Defines values for CredentialsDataRolesRole.
const (
	RoleCPO   Role = "CPO"
	RoleEMSP  Role = "EMSP"
	RoleHUB   Role = "HUB"
	RoleNAP   Role = "NAP"
	RoleNSP   Role = "NSP"
	RoleOTHER Role = "OTHER"
	RoleSCSP  Role = "SCSP"
)

type CredentialsRole struct {
	BusinessDetails BusinessDetails `json:"business_details"`
	CountryCode     string          `json:"country_code"`
	PartyID         string          `json:"party_id"`
	Role            Role            `json:"role"`
}

type Credentials struct {
	Token string            `json:"token,omitempty"`
	URL   string            `json:"url,omitempty"`
	Roles []CredentialsRole `json:"roles,omitempty"`
}

type RegisterCredentialsResponse = response.BaseResponse[Credentials]

type RegisterCredentialsUsecase struct {
}

func NewRegisterCredentialsUsecase() *RegisterCredentialsUsecase {
	return &RegisterCredentialsUsecase{}
}

func (uc *RegisterCredentialsUsecase) Execute(version versions.VersionNumber, credentialTokenA string, control *controls.Control) RegisterCredentialsResponse {
	validToken := env.ValidateEnv().CredentialTokenA

	log.Println("Valid token: ", validToken, len(validToken))
	log.Println("Credential token A: ", credentialTokenA, len(credentialTokenA))

	isTokenValid := strings.Compare(credentialTokenA, validToken) == 0

	log.Println("Is token valid: ", isTokenValid)

	if strings.Compare(credentialTokenA, validToken) != 0 {
		//exhaustruct:ignore
		credentials := Credentials{}

		return RegisterCredentialsResponse{
			StatusCode:    response.StatusCodeGenericClientError,
			StatusMessage: "Unauthorized",
			TimeStamp:     time.Now().Format(time.RFC3339),
			//exhaustruct:ignore
			Data: credentials,
		}
	}

	return RegisterCredentialsResponse{
		StatusCode:    response.StatusCodeGenericSuccess,
		StatusMessage: "Success",
		Data: Credentials{
			Token: "1234567890",
			URL:   "https://example.com",
			Roles: []CredentialsRole{
				{
					BusinessDetails: BusinessDetails{
						Name:    "Test",
						Website: "https://example.com",
						Logo: Logo{
							Category:  "test",
							Type:      "test",
							URL:       "https://example.com",
							Thumbnail: "https://example.com",
							Width:     100,
							Height:    100,
						},
					},
					CountryCode: "FR",
					PartyID:     "1234567890",
					Role:        RoleCPO,
				},
			},
		},
	}
}
