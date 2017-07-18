// Code generated by cdpgen. DO NOT EDIT.

package security

// HandleCertificateErrorArgs represents the arguments for HandleCertificateError in the Security domain.
type HandleCertificateErrorArgs struct {
	EventID int                    `json:"eventId"` // The ID of the event.
	Action  CertificateErrorAction `json:"action"`  // The action to take on the certificate error.
}

// NewHandleCertificateErrorArgs initializes HandleCertificateErrorArgs with the required arguments.
func NewHandleCertificateErrorArgs(eventID int, action CertificateErrorAction) *HandleCertificateErrorArgs {
	args := new(HandleCertificateErrorArgs)
	args.EventID = eventID
	args.Action = action
	return args
}

// SetOverrideCertificateErrorsArgs represents the arguments for SetOverrideCertificateErrors in the Security domain.
type SetOverrideCertificateErrorsArgs struct {
	Override bool `json:"override"` // If true, certificate errors will be overridden.
}

// NewSetOverrideCertificateErrorsArgs initializes SetOverrideCertificateErrorsArgs with the required arguments.
func NewSetOverrideCertificateErrorsArgs(override bool) *SetOverrideCertificateErrorsArgs {
	args := new(SetOverrideCertificateErrorsArgs)
	args.Override = override
	return args
}
