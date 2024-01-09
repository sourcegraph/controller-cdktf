package accessidentityprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/accessidentityprovider/internal"
)

type AccessIdentityProviderConfigAOutputReference interface {
	cdktf.ComplexObject
	ApiToken() *string
	SetApiToken(val *string)
	ApiTokenInput() *string
	AppsDomain() *string
	SetAppsDomain(val *string)
	AppsDomainInput() *string
	Attributes() *[]*string
	SetAttributes(val *[]*string)
	AttributesInput() *[]*string
	AuthorizationServerId() *string
	SetAuthorizationServerId(val *string)
	AuthorizationServerIdInput() *string
	AuthUrl() *string
	SetAuthUrl(val *string)
	AuthUrlInput() *string
	CentrifyAccount() *string
	SetCentrifyAccount(val *string)
	CentrifyAccountInput() *string
	CentrifyAppId() *string
	SetCentrifyAppId(val *string)
	CentrifyAppIdInput() *string
	CertsUrl() *string
	SetCertsUrl(val *string)
	CertsUrlInput() *string
	Claims() *[]*string
	SetClaims(val *[]*string)
	ClaimsInput() *[]*string
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ConditionalAccessEnabled() interface{}
	SetConditionalAccessEnabled(val interface{})
	ConditionalAccessEnabledInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DirectoryId() *string
	SetDirectoryId(val *string)
	DirectoryIdInput() *string
	EmailAttributeName() *string
	SetEmailAttributeName(val *string)
	EmailAttributeNameInput() *string
	EmailClaimName() *string
	SetEmailClaimName(val *string)
	EmailClaimNameInput() *string
	// Experimental.
	Fqn() *string
	IdpPublicCert() *string
	SetIdpPublicCert(val *string)
	IdpPublicCertInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IssuerUrl() *string
	SetIssuerUrl(val *string)
	IssuerUrlInput() *string
	OktaAccount() *string
	SetOktaAccount(val *string)
	OktaAccountInput() *string
	OneloginAccount() *string
	SetOneloginAccount(val *string)
	OneloginAccountInput() *string
	PingEnvId() *string
	SetPingEnvId(val *string)
	PingEnvIdInput() *string
	PkceEnabled() interface{}
	SetPkceEnabled(val interface{})
	PkceEnabledInput() interface{}
	RedirectUrl() *string
	Scopes() *[]*string
	SetScopes(val *[]*string)
	ScopesInput() *[]*string
	SignRequest() interface{}
	SetSignRequest(val interface{})
	SignRequestInput() interface{}
	SsoTargetUrl() *string
	SetSsoTargetUrl(val *string)
	SsoTargetUrlInput() *string
	SupportGroups() interface{}
	SetSupportGroups(val interface{})
	SupportGroupsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TokenUrl() *string
	SetTokenUrl(val *string)
	TokenUrlInput() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetApiToken()
	ResetAppsDomain()
	ResetAttributes()
	ResetAuthorizationServerId()
	ResetAuthUrl()
	ResetCentrifyAccount()
	ResetCentrifyAppId()
	ResetCertsUrl()
	ResetClaims()
	ResetClientId()
	ResetClientSecret()
	ResetConditionalAccessEnabled()
	ResetDirectoryId()
	ResetEmailAttributeName()
	ResetEmailClaimName()
	ResetIdpPublicCert()
	ResetIssuerUrl()
	ResetOktaAccount()
	ResetOneloginAccount()
	ResetPingEnvId()
	ResetPkceEnabled()
	ResetScopes()
	ResetSignRequest()
	ResetSsoTargetUrl()
	ResetSupportGroups()
	ResetTokenUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccessIdentityProviderConfigAOutputReference
type jsiiProxy_AccessIdentityProviderConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ApiToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ApiTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) AppsDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appsDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) AppsDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appsDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) Attributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) AttributesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) AuthorizationServerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationServerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) AuthorizationServerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationServerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) AuthUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) AuthUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) CentrifyAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"centrifyAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) CentrifyAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"centrifyAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) CentrifyAppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"centrifyAppId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) CentrifyAppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"centrifyAppIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) CertsUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certsUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) CertsUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certsUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) Claims() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"claims",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ClaimsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"claimsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ConditionalAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionalAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ConditionalAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionalAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) DirectoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) DirectoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) EmailAttributeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAttributeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) EmailAttributeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAttributeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) EmailClaimName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailClaimName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) EmailClaimNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailClaimNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) IdpPublicCert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpPublicCert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) IdpPublicCertInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpPublicCertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) IssuerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) IssuerUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) OktaAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oktaAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) OktaAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oktaAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) OneloginAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oneloginAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) OneloginAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oneloginAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) PingEnvId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pingEnvId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) PingEnvIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pingEnvIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) PkceEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pkceEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) PkceEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pkceEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) RedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SignRequest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"signRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SignRequestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"signRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SsoTargetUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoTargetUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SsoTargetUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoTargetUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SupportGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SupportGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) TokenUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) TokenUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUrlInput",
		&returns,
	)
	return returns
}


func NewAccessIdentityProviderConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AccessIdentityProviderConfigAOutputReference {
	_init_.Initialize()

	if err := validateNewAccessIdentityProviderConfigAOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessIdentityProviderConfigAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.accessIdentityProvider.AccessIdentityProviderConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAccessIdentityProviderConfigAOutputReference_Override(a AccessIdentityProviderConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.accessIdentityProvider.AccessIdentityProviderConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference)SetApiToken(val *string) {
	if err := j.validateSetApiTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiToken",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference)SetAppsDomain(val *string) {
	if err := j.validateSetAppsDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appsDomain",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference)SetAttributes(val *[]*string) {
	if err := j.validateSetAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributes",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference)SetAuthorizationServerId(val *string) {
	if err := j.validateSetAuthorizationServerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationServerId",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference)SetAuthUrl(val *string) {
	if err := j.validateSetAuthUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authUrl",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference)SetCentrifyAccount(val *string) {
	if err := j.validateSetCentrifyAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"centrifyAccount",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference)SetCentrifyAppId(val *string) {
	if err := j.validateSetCentrifyAppIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"centrifyAppId",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference)SetCertsUrl(val *string) {
	if err := j.validateSetCertsUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certsUrl",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference)SetClaims(val *[]*string) {
	if err := j.validateSetClaimsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"claims",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference)SetConditionalAccessEnabled(val interface{}) {
	if err := j.validateSetConditionalAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"conditionalAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference)SetDirectoryId(val *string) {
	if err := j.validateSetDirectoryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directoryId",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference)SetEmailAttributeName(val *string) {
	if err := j.validateSetEmailAttributeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailAttributeName",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference)SetEmailClaimName(val *string) {
	if err := j.validateSetEmailClaimNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailClaimName",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference)SetIdpPublicCert(val *string) {
	if err := j.validateSetIdpPublicCertParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idpPublicCert",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference)SetIssuerUrl(val *string) {
	if err := j.validateSetIssuerUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuerUrl",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference)SetOktaAccount(val *string) {
	if err := j.validateSetOktaAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oktaAccount",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference)SetOneloginAccount(val *string) {
	if err := j.validateSetOneloginAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oneloginAccount",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference)SetPingEnvId(val *string) {
	if err := j.validateSetPingEnvIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pingEnvId",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference)SetPkceEnabled(val interface{}) {
	if err := j.validateSetPkceEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pkceEnabled",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference)SetScopes(val *[]*string) {
	if err := j.validateSetScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference)SetSignRequest(val interface{}) {
	if err := j.validateSetSignRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signRequest",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference)SetSsoTargetUrl(val *string) {
	if err := j.validateSetSsoTargetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ssoTargetUrl",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference)SetSupportGroups(val interface{}) {
	if err := j.validateSetSupportGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportGroups",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference)SetTokenUrl(val *string) {
	if err := j.validateSetTokenUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenUrl",
		val,
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetApiToken() {
	_jsii_.InvokeVoid(
		a,
		"resetApiToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetAppsDomain() {
	_jsii_.InvokeVoid(
		a,
		"resetAppsDomain",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetAttributes() {
	_jsii_.InvokeVoid(
		a,
		"resetAttributes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetAuthorizationServerId() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizationServerId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetAuthUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetCentrifyAccount() {
	_jsii_.InvokeVoid(
		a,
		"resetCentrifyAccount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetCentrifyAppId() {
	_jsii_.InvokeVoid(
		a,
		"resetCentrifyAppId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetCertsUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetCertsUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetClaims() {
	_jsii_.InvokeVoid(
		a,
		"resetClaims",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetClientId() {
	_jsii_.InvokeVoid(
		a,
		"resetClientId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetClientSecret() {
	_jsii_.InvokeVoid(
		a,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetConditionalAccessEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetConditionalAccessEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetDirectoryId() {
	_jsii_.InvokeVoid(
		a,
		"resetDirectoryId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetEmailAttributeName() {
	_jsii_.InvokeVoid(
		a,
		"resetEmailAttributeName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetEmailClaimName() {
	_jsii_.InvokeVoid(
		a,
		"resetEmailClaimName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetIdpPublicCert() {
	_jsii_.InvokeVoid(
		a,
		"resetIdpPublicCert",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetIssuerUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetIssuerUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetOktaAccount() {
	_jsii_.InvokeVoid(
		a,
		"resetOktaAccount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetOneloginAccount() {
	_jsii_.InvokeVoid(
		a,
		"resetOneloginAccount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetPingEnvId() {
	_jsii_.InvokeVoid(
		a,
		"resetPingEnvId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetPkceEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetPkceEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetScopes() {
	_jsii_.InvokeVoid(
		a,
		"resetScopes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetSignRequest() {
	_jsii_.InvokeVoid(
		a,
		"resetSignRequest",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetSsoTargetUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetSsoTargetUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetSupportGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetSupportGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetTokenUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

