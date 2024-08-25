package did

import (
	"strings"

	"github.com/reiver/go-did/core/v1/syntax"
)


// DID represents a decentralized-identifier (DID).
//
// DIDs are defined here:
// https://www.w3.org/TR/did-core/#did-syntax
type DID string

func constructDID(methodName string, identifier string) ([]byte, error) {
	if !syntax.IsMethodName(methodName) {
		return nil, errBadDIDMethodName
	}
	if !syntax.IsMethodSpecificID(identifier) {
		return nil, errBadDIDIdentifier
	}

	var p []byte
	{
		var buffer [256]byte
		p = buffer[0:0]

		p = append(p, "did:"...)
		p = append(p, methodName...)
		p = append(p, ':')
		p = append(p, identifier...)
	}

	return p, nil
}

// ConstructDID creates a did.DID based on a method-name and identifier.
//
// ConstructDID also validates both method-name and identifier, and returns an error if either (or both) of them fail validation.
func ConstructDID(methodName string, identifier string) (DID, error) {
	var empty DID

	p, err := constructDID(methodName, identifier)
	if nil != err {
		return empty, err
	}

	return DID(string(p)), nil
}

// MustConstructDID is similar to [ConstructDID] except it panic()s if there is an error.
func MustConstructDID(methodName string, identifier string) DID {
	decentralizedID, err := ConstructDID(methodName, identifier)
	if nil != err {
		panic(err)
	}

	return decentralizedID
}

func (receiver DID) DecompileDID() (methodName string, identifier string, err error) {
	var str string = string(receiver)

	{
		if !strings.HasPrefix(str, didprefix) {
			return "", "", errMissingDIDPrefix
		}

		str = str[len(didprefix):]
	}

	var index int
	{
		index = strings.Index(str, ":")

		if index < 0 {
			return "", "", errMissingDIDIdentifier
		}

	}

	{
		methodName = str[:index]

		if !syntax.IsMethodName(methodName) {
			return "", "", errBadDIDMethodName
		}

		str = str[1+index:]
	}

	{
		identifier = str

		if !syntax.IsMethodSpecificID(str) {
			return "", "", errBadDIDIdentifier
		}
	}

	return methodName, identifier, nil

}

func (receiver DID) MarshalText() ([]byte, error) {
	err := receiver.Validate()
	if nil != err {
		return nil, err
	}

	var str string  = string(receiver)

	return []byte(str), nil
}

func (receiver DID) MethodName() string {
	methodName, _, err := receiver.DecompileDID()
	if nil != err {
		return ""
	}

	return methodName
}

func (receiver DID) Identifier() string {
	_, identifier, err := receiver.DecompileDID()
	if nil != err {
		return ""
	}

	return identifier
}

func (receiver DID) String() string {
	_, _, err := receiver.DecompileDID()
	if nil != err {
		return ""
	}

	return string(receiver)
}

// Validator returns an error did the did.DID is not valid.
func (receiver DID) Validate() error {
	_, _, err := receiver.DecompileDID()
	return err
}

func (receiver *DID) UnmarshalText(text []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	if nil == text {
		return errNilBytes
	}

	*receiver = DID(string(text))
	return nil
}
