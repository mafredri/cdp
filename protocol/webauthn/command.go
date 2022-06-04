// Code generated by cdpgen. DO NOT EDIT.

package webauthn

// EnableArgs represents the arguments for Enable in the WebAuthn domain.
type EnableArgs struct {
	EnableUI *bool `json:"enableUI,omitempty"` // Whether to enable the WebAuthn user interface. Enabling the UI is recommended for debugging and demo purposes, as it is closer to the real experience. Disabling the UI is recommended for automated testing. Supported at the embedder's discretion if UI is available. Defaults to false.
}

// NewEnableArgs initializes EnableArgs with the required arguments.
func NewEnableArgs() *EnableArgs {
	args := new(EnableArgs)

	return args
}

// SetEnableUI sets the EnableUI optional argument. Whether to enable
// the WebAuthn user interface. Enabling the UI is recommended for
// debugging and demo purposes, as it is closer to the real experience.
// Disabling the UI is recommended for automated testing. Supported at
// the embedder's discretion if UI is available. Defaults to false.
func (a *EnableArgs) SetEnableUI(enableUI bool) *EnableArgs {
	a.EnableUI = &enableUI
	return a
}

// AddVirtualAuthenticatorArgs represents the arguments for AddVirtualAuthenticator in the WebAuthn domain.
type AddVirtualAuthenticatorArgs struct {
	Options VirtualAuthenticatorOptions `json:"options"` // No description.
}

// NewAddVirtualAuthenticatorArgs initializes AddVirtualAuthenticatorArgs with the required arguments.
func NewAddVirtualAuthenticatorArgs(options VirtualAuthenticatorOptions) *AddVirtualAuthenticatorArgs {
	args := new(AddVirtualAuthenticatorArgs)
	args.Options = options
	return args
}

// AddVirtualAuthenticatorReply represents the return values for AddVirtualAuthenticator in the WebAuthn domain.
type AddVirtualAuthenticatorReply struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"` // No description.
}

// RemoveVirtualAuthenticatorArgs represents the arguments for RemoveVirtualAuthenticator in the WebAuthn domain.
type RemoveVirtualAuthenticatorArgs struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"` // No description.
}

// NewRemoveVirtualAuthenticatorArgs initializes RemoveVirtualAuthenticatorArgs with the required arguments.
func NewRemoveVirtualAuthenticatorArgs(authenticatorID AuthenticatorID) *RemoveVirtualAuthenticatorArgs {
	args := new(RemoveVirtualAuthenticatorArgs)
	args.AuthenticatorID = authenticatorID
	return args
}

// AddCredentialArgs represents the arguments for AddCredential in the WebAuthn domain.
type AddCredentialArgs struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"` // No description.
	Credential      Credential      `json:"credential"`      // No description.
}

// NewAddCredentialArgs initializes AddCredentialArgs with the required arguments.
func NewAddCredentialArgs(authenticatorID AuthenticatorID, credential Credential) *AddCredentialArgs {
	args := new(AddCredentialArgs)
	args.AuthenticatorID = authenticatorID
	args.Credential = credential
	return args
}

// GetCredentialArgs represents the arguments for GetCredential in the WebAuthn domain.
type GetCredentialArgs struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"` // No description.
	CredentialID    string          `json:"credentialId"`    // No description.
}

// NewGetCredentialArgs initializes GetCredentialArgs with the required arguments.
func NewGetCredentialArgs(authenticatorID AuthenticatorID, credentialID string) *GetCredentialArgs {
	args := new(GetCredentialArgs)
	args.AuthenticatorID = authenticatorID
	args.CredentialID = credentialID
	return args
}

// GetCredentialReply represents the return values for GetCredential in the WebAuthn domain.
type GetCredentialReply struct {
	Credential Credential `json:"credential"` // No description.
}

// GetCredentialsArgs represents the arguments for GetCredentials in the WebAuthn domain.
type GetCredentialsArgs struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"` // No description.
}

// NewGetCredentialsArgs initializes GetCredentialsArgs with the required arguments.
func NewGetCredentialsArgs(authenticatorID AuthenticatorID) *GetCredentialsArgs {
	args := new(GetCredentialsArgs)
	args.AuthenticatorID = authenticatorID
	return args
}

// GetCredentialsReply represents the return values for GetCredentials in the WebAuthn domain.
type GetCredentialsReply struct {
	Credentials []Credential `json:"credentials"` // No description.
}

// RemoveCredentialArgs represents the arguments for RemoveCredential in the WebAuthn domain.
type RemoveCredentialArgs struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"` // No description.
	CredentialID    string          `json:"credentialId"`    // No description.
}

// NewRemoveCredentialArgs initializes RemoveCredentialArgs with the required arguments.
func NewRemoveCredentialArgs(authenticatorID AuthenticatorID, credentialID string) *RemoveCredentialArgs {
	args := new(RemoveCredentialArgs)
	args.AuthenticatorID = authenticatorID
	args.CredentialID = credentialID
	return args
}

// ClearCredentialsArgs represents the arguments for ClearCredentials in the WebAuthn domain.
type ClearCredentialsArgs struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"` // No description.
}

// NewClearCredentialsArgs initializes ClearCredentialsArgs with the required arguments.
func NewClearCredentialsArgs(authenticatorID AuthenticatorID) *ClearCredentialsArgs {
	args := new(ClearCredentialsArgs)
	args.AuthenticatorID = authenticatorID
	return args
}

// SetUserVerifiedArgs represents the arguments for SetUserVerified in the WebAuthn domain.
type SetUserVerifiedArgs struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"` // No description.
	IsUserVerified  bool            `json:"isUserVerified"`  // No description.
}

// NewSetUserVerifiedArgs initializes SetUserVerifiedArgs with the required arguments.
func NewSetUserVerifiedArgs(authenticatorID AuthenticatorID, isUserVerified bool) *SetUserVerifiedArgs {
	args := new(SetUserVerifiedArgs)
	args.AuthenticatorID = authenticatorID
	args.IsUserVerified = isUserVerified
	return args
}

// SetAutomaticPresenceSimulationArgs represents the arguments for SetAutomaticPresenceSimulation in the WebAuthn domain.
type SetAutomaticPresenceSimulationArgs struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"` // No description.
	Enabled         bool            `json:"enabled"`         // No description.
}

// NewSetAutomaticPresenceSimulationArgs initializes SetAutomaticPresenceSimulationArgs with the required arguments.
func NewSetAutomaticPresenceSimulationArgs(authenticatorID AuthenticatorID, enabled bool) *SetAutomaticPresenceSimulationArgs {
	args := new(SetAutomaticPresenceSimulationArgs)
	args.AuthenticatorID = authenticatorID
	args.Enabled = enabled
	return args
}
