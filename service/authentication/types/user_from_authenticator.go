package types

type (
	ResponseUserFromAuthenticator struct {
		UserID      string `json:"user_id"`
		IsSuperuser *bool  `json:"is_superuser,omitempty"`
	}
)
