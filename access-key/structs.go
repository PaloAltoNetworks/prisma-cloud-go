package accesskey

type AccessKey struct {
	CreatedBy          string `json:"createdBy,omitempty"`
	CreatedTimestamp   int    `json:"createdTs,omitempty"`
	ExpiresOn          int    `json:"expiresOn,omitempty"`
	Id                 string `json:"id,omitempty,omitempty"`
	LastUsed           int    `json:"lastUsedTime,omitempty"`
	Name               string `json:"name,omitempty"`
	Role               Role   `json:"role,omitempty"`
	RoleType           string `json:"roleType,omitempty"`
	Status             bool   `json:"status,omitempty"`
	Username           string `json:"username,omitempty"`
	ServiceAccountName string `json:"serviceAccountName,omitempty"`
	SecretKey          string `json:"secretKey,omitempty"`
}

type Role struct {
	Name string `json:"name"`
}
