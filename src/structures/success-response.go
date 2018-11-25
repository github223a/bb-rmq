package structures

import "github.com/satori/go.uuid"

type SuccessResponse struct {
	Id uuid.UUID `json:"id"`
	Namespace string `json:"namespace"`
	Method string `json:"method"`
	Domain string `json:"domain"`
	Locale string `json:"locale"`
	Result map[string] interface{} `json:"result"`
	Source string `json:"source"`
	ResponseQueue *string `json:"responseQueue"`
	CacheKey *string `json:"cacheKey"`
	Token *string `json:"token"`
}
