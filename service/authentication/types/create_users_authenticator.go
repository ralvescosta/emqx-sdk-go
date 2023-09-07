package types

type (
	RequestCreateUserForAuthenticator struct {
		Password string `json:"password"`
		UserID   string `json:"user_id"`
	}

	ResponseCreateUserForAuthenticator struct {
		UserID      string `json:"user_id"`
		IsSuperUser bool   `json:"is_superuser"`
	}
)
