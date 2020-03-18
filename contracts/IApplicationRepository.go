package contracts

import "nikan.dev/pronto/payloads"

type IApplicationRepository interface {
	Info() (payloads.ApplicationInfoPayload, error)
	Ping() (payloads.ApplicationPingPayload, error)
}
