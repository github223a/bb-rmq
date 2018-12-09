package structures

import uuid "github.com/satori/go.uuid"

type ErrorResponse struct {
	ID            uuid.UUID              `json:"id"`
	Namespace     string                 `json:"namespace"`
	Method        string                 `json:"method"`
	Domain        string                 `json:"domain"`
	Locale        string                 `json:"locale"`
	Error         map[string]interface{} `json:"error"`
	Source        string                 `json:"source"`
	ResponseQueue *string                `json:"responseQueue"`
}
