package models

type Namespace struct {
	Name     string `json:"name"`
	Owner    string `json:"owner"`
	TenantID string `json:"tenant_id" bson:"tenant_id"`
}
