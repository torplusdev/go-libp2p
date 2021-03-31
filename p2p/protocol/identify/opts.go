package identify

type config struct {
	userAgent               string
	disableSignedPeerRecord bool
	stellarPublicKey string
}

// Option is an option function for identify.
type Option func(*config)

// UserAgent sets the user agent this node will identify itself with to peers.
func UserAgent(ua string) Option {
	return func(cfg *config) {
		cfg.userAgent = ua
	}
}

// StellarPublicKey sets the stellar public key of the node for receiving payment transactions.
func StellarPublicKey(ua string) Option {
	return func(cfg *config) {
		cfg.stellarPublicKey = ua
	}
}

// DisableSignedPeerRecord disables populating signed peer records on the outgoing Identify response
// and ONLY sends the unsigned addresses.
func DisableSignedPeerRecord() Option {
	return func(cfg *config) {
		cfg.disableSignedPeerRecord = true
	}
}
