package p2p

type HandShaker interface {
	Handshake() error
}

type HandshakerFunc func() error

type DefualtHandshaker struct{}
