package accesspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/accesspolicy/internal"
)

type AccessPolicyIncludeOutputReference interface {
	cdktf.ComplexObject
	AnyValidServiceToken() interface{}
	SetAnyValidServiceToken(val interface{})
	AnyValidServiceTokenInput() interface{}
	AuthContext() AccessPolicyIncludeAuthContextList
	AuthContextInput() interface{}
	AuthMethod() *string
	SetAuthMethod(val *string)
	AuthMethodInput() *string
	Azure() AccessPolicyIncludeAzureList
	AzureInput() interface{}
	Certificate() interface{}
	SetCertificate(val interface{})
	CertificateInput() interface{}
	CommonName() *string
	SetCommonName(val *string)
	CommonNameInput() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DevicePosture() *[]*string
	SetDevicePosture(val *[]*string)
	DevicePostureInput() *[]*string
	Email() *[]*string
	SetEmail(val *[]*string)
	EmailDomain() *[]*string
	SetEmailDomain(val *[]*string)
	EmailDomainInput() *[]*string
	EmailInput() *[]*string
	Everyone() interface{}
	SetEveryone(val interface{})
	EveryoneInput() interface{}
	ExternalEvaluation() AccessPolicyIncludeExternalEvaluationOutputReference
	ExternalEvaluationInput() *AccessPolicyIncludeExternalEvaluation
	// Experimental.
	Fqn() *string
	Geo() *[]*string
	SetGeo(val *[]*string)
	GeoInput() *[]*string
	Github() AccessPolicyIncludeGithubList
	GithubInput() interface{}
	Group() *[]*string
	SetGroup(val *[]*string)
	GroupInput() *[]*string
	Gsuite() AccessPolicyIncludeGsuiteList
	GsuiteInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Ip() *[]*string
	SetIp(val *[]*string)
	IpInput() *[]*string
	IpList() *[]*string
	SetIpList(val *[]*string)
	IpListInput() *[]*string
	LoginMethod() *[]*string
	SetLoginMethod(val *[]*string)
	LoginMethodInput() *[]*string
	Okta() AccessPolicyIncludeOktaList
	OktaInput() interface{}
	Saml() AccessPolicyIncludeSamlList
	SamlInput() interface{}
	ServiceToken() *[]*string
	SetServiceToken(val *[]*string)
	ServiceTokenInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutAuthContext(value interface{})
	PutAzure(value interface{})
	PutExternalEvaluation(value *AccessPolicyIncludeExternalEvaluation)
	PutGithub(value interface{})
	PutGsuite(value interface{})
	PutOkta(value interface{})
	PutSaml(value interface{})
	ResetAnyValidServiceToken()
	ResetAuthContext()
	ResetAuthMethod()
	ResetAzure()
	ResetCertificate()
	ResetCommonName()
	ResetDevicePosture()
	ResetEmail()
	ResetEmailDomain()
	ResetEveryone()
	ResetExternalEvaluation()
	ResetGeo()
	ResetGithub()
	ResetGroup()
	ResetGsuite()
	ResetIp()
	ResetIpList()
	ResetLoginMethod()
	ResetOkta()
	ResetSaml()
	ResetServiceToken()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccessPolicyIncludeOutputReference
type jsiiProxy_AccessPolicyIncludeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) AnyValidServiceToken() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anyValidServiceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) AnyValidServiceTokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anyValidServiceTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) AuthContext() AccessPolicyIncludeAuthContextList {
	var returns AccessPolicyIncludeAuthContextList
	_jsii_.Get(
		j,
		"authContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) AuthContextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) AuthMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) AuthMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) Azure() AccessPolicyIncludeAzureList {
	var returns AccessPolicyIncludeAzureList
	_jsii_.Get(
		j,
		"azure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) AzureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) Certificate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) CertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) CommonName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) CommonNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commonNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) DevicePosture() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"devicePosture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) DevicePostureInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"devicePostureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) Email() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) EmailDomain() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) EmailDomainInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) EmailInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) Everyone() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"everyone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) EveryoneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"everyoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) ExternalEvaluation() AccessPolicyIncludeExternalEvaluationOutputReference {
	var returns AccessPolicyIncludeExternalEvaluationOutputReference
	_jsii_.Get(
		j,
		"externalEvaluation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) ExternalEvaluationInput() *AccessPolicyIncludeExternalEvaluation {
	var returns *AccessPolicyIncludeExternalEvaluation
	_jsii_.Get(
		j,
		"externalEvaluationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) Geo() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"geo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) GeoInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"geoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) Github() AccessPolicyIncludeGithubList {
	var returns AccessPolicyIncludeGithubList
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) GithubInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"githubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) Group() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) GroupInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) Gsuite() AccessPolicyIncludeGsuiteList {
	var returns AccessPolicyIncludeGsuiteList
	_jsii_.Get(
		j,
		"gsuite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) GsuiteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gsuiteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) Ip() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) IpInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) IpList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) IpListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) LoginMethod() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loginMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) LoginMethodInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loginMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) Okta() AccessPolicyIncludeOktaList {
	var returns AccessPolicyIncludeOktaList
	_jsii_.Get(
		j,
		"okta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) OktaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oktaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) Saml() AccessPolicyIncludeSamlList {
	var returns AccessPolicyIncludeSamlList
	_jsii_.Get(
		j,
		"saml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) SamlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"samlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) ServiceToken() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) ServiceTokenInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAccessPolicyIncludeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AccessPolicyIncludeOutputReference {
	_init_.Initialize()

	if err := validateNewAccessPolicyIncludeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessPolicyIncludeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.accessPolicy.AccessPolicyIncludeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAccessPolicyIncludeOutputReference_Override(a AccessPolicyIncludeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.accessPolicy.AccessPolicyIncludeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference)SetAnyValidServiceToken(val interface{}) {
	if err := j.validateSetAnyValidServiceTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anyValidServiceToken",
		val,
	)
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference)SetAuthMethod(val *string) {
	if err := j.validateSetAuthMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authMethod",
		val,
	)
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference)SetCertificate(val interface{}) {
	if err := j.validateSetCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference)SetCommonName(val *string) {
	if err := j.validateSetCommonNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commonName",
		val,
	)
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference)SetDevicePosture(val *[]*string) {
	if err := j.validateSetDevicePostureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"devicePosture",
		val,
	)
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference)SetEmail(val *[]*string) {
	if err := j.validateSetEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference)SetEmailDomain(val *[]*string) {
	if err := j.validateSetEmailDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailDomain",
		val,
	)
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference)SetEveryone(val interface{}) {
	if err := j.validateSetEveryoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"everyone",
		val,
	)
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference)SetGeo(val *[]*string) {
	if err := j.validateSetGeoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geo",
		val,
	)
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference)SetGroup(val *[]*string) {
	if err := j.validateSetGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"group",
		val,
	)
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference)SetIp(val *[]*string) {
	if err := j.validateSetIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ip",
		val,
	)
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference)SetIpList(val *[]*string) {
	if err := j.validateSetIpListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipList",
		val,
	)
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference)SetLoginMethod(val *[]*string) {
	if err := j.validateSetLoginMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginMethod",
		val,
	)
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference)SetServiceToken(val *[]*string) {
	if err := j.validateSetServiceTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceToken",
		val,
	)
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccessPolicyIncludeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) PutAuthContext(value interface{}) {
	if err := a.validatePutAuthContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAuthContext",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) PutAzure(value interface{}) {
	if err := a.validatePutAzureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAzure",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) PutExternalEvaluation(value *AccessPolicyIncludeExternalEvaluation) {
	if err := a.validatePutExternalEvaluationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExternalEvaluation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) PutGithub(value interface{}) {
	if err := a.validatePutGithubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGithub",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) PutGsuite(value interface{}) {
	if err := a.validatePutGsuiteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGsuite",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) PutOkta(value interface{}) {
	if err := a.validatePutOktaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOkta",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) PutSaml(value interface{}) {
	if err := a.validatePutSamlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSaml",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) ResetAnyValidServiceToken() {
	_jsii_.InvokeVoid(
		a,
		"resetAnyValidServiceToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) ResetAuthContext() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthContext",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) ResetAuthMethod() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthMethod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) ResetAzure() {
	_jsii_.InvokeVoid(
		a,
		"resetAzure",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) ResetCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) ResetCommonName() {
	_jsii_.InvokeVoid(
		a,
		"resetCommonName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) ResetDevicePosture() {
	_jsii_.InvokeVoid(
		a,
		"resetDevicePosture",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) ResetEmail() {
	_jsii_.InvokeVoid(
		a,
		"resetEmail",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) ResetEmailDomain() {
	_jsii_.InvokeVoid(
		a,
		"resetEmailDomain",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) ResetEveryone() {
	_jsii_.InvokeVoid(
		a,
		"resetEveryone",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) ResetExternalEvaluation() {
	_jsii_.InvokeVoid(
		a,
		"resetExternalEvaluation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) ResetGeo() {
	_jsii_.InvokeVoid(
		a,
		"resetGeo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) ResetGithub() {
	_jsii_.InvokeVoid(
		a,
		"resetGithub",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) ResetGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) ResetGsuite() {
	_jsii_.InvokeVoid(
		a,
		"resetGsuite",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) ResetIp() {
	_jsii_.InvokeVoid(
		a,
		"resetIp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) ResetIpList() {
	_jsii_.InvokeVoid(
		a,
		"resetIpList",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) ResetLoginMethod() {
	_jsii_.InvokeVoid(
		a,
		"resetLoginMethod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) ResetOkta() {
	_jsii_.InvokeVoid(
		a,
		"resetOkta",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) ResetSaml() {
	_jsii_.InvokeVoid(
		a,
		"resetSaml",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) ResetServiceToken() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AccessPolicyIncludeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

