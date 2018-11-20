package templates

import "github.com/satori/go.uuid"

type SuccessResponse struct {
	Id uuid.UUID `json:"id"`
	Namespace string `json:"namespace"`
	Method string `json:"method"`
	Domain string `json:"domain"`
	Locale string `json:"locale"`
	Result map[string] interface{} `json:"result"`
	Source string `json:"source"`
}
