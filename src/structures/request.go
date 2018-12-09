package structures

import uuid "github.com/satori/go.uuid"

type Request struct {
	ID            uuid.UUID              `json:"id"`
	Namespace     string                 `json:"namespace"`
	Method        string                 `json:"method"`
	Domain        *string                `json:"domain"`
	Locale        *string                `json:"locale"`
	Params        map[string]interface{} `json:"params"`
	Source        string                 `json:"source"`
	ResponseQueue *string                `json:"responseQueue"`
}
