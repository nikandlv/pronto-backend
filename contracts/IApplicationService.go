package contracts

import "nikan.dev/pronto/payloads"

type IApplicationService interface {
	Info() (payloads.ApplicationInfoPayload, error)
	Ping() (payloads.ApplicationPingPayload, error)
}
