package did

import (
	"github.com/reiver/go-erorr"
)

const (
	errBadDIDIdentifier     = erorr.Error("did: bad did identifier")
	errBadDIDMethodName     = erorr.Error("did: bad did method-name")
	errMissingDIDIdentifier = erorr.Error("did: missing did identifier")
	errMissingDIDPrefix     = erorr.Error("did: missing did prefix")
	errNilBytes             = erorr.Error("did: nil bytes")
	errNilReceiver          = erorr.Error("did: nil receiver")
)
