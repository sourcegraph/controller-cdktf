package accessgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/accessgroup/internal"
)

type AccessGroupExcludeOutputReference interface {
	cdktf.ComplexObject
	AnyValidServiceToken() interface{}
	SetAnyValidServiceToken(val interface{})
	AnyValidServiceTokenInput() interface{}
	AuthMethod() *string
	SetAuthMethod(val *string)
	AuthMethodInput() *string
	Azure() AccessGroupExcludeAzureList
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
	ExternalEvaluation() AccessGroupExcludeExternalEvaluationOutputReference
	ExternalEvaluationInput() *AccessGroupExcludeExternalEvaluation
	// Experimental.
	Fqn() *string
	Geo() *[]*string
	SetGeo(val *[]*string)
	GeoInput() *[]*string
	Github() AccessGroupExcludeGithubList
	GithubInput() interface{}
	Group() *[]*string
	SetGroup(val *[]*string)
	GroupInput() *[]*string
	Gsuite() AccessGroupExcludeGsuiteList
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
	Okta() AccessGroupExcludeOktaList
	OktaInput() interface{}
	Saml() AccessGroupExcludeSamlList
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
	PutAzure(value interface{})
	PutExternalEvaluation(value *AccessGroupExcludeExternalEvaluation)
	PutGithub(value interface{})
	PutGsuite(value interface{})
	PutOkta(value interface{})
	PutSaml(value interface{})
	ResetAnyValidServiceToken()
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

// The jsii proxy struct for AccessGroupExcludeOutputReference
type jsiiProxy_AccessGroupExcludeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) AnyValidServiceToken() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anyValidServiceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) AnyValidServiceTokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anyValidServiceTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) AuthMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) AuthMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) Azure() AccessGroupExcludeAzureList {
	var returns AccessGroupExcludeAzureList
	_jsii_.Get(
		j,
		"azure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) AzureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) Certificate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) CertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) CommonName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) CommonNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commonNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) DevicePosture() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"devicePosture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) DevicePostureInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"devicePostureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) Email() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) EmailDomain() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) EmailDomainInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) EmailInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) Everyone() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"everyone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) EveryoneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"everyoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) ExternalEvaluation() AccessGroupExcludeExternalEvaluationOutputReference {
	var returns AccessGroupExcludeExternalEvaluationOutputReference
	_jsii_.Get(
		j,
		"externalEvaluation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) ExternalEvaluationInput() *AccessGroupExcludeExternalEvaluation {
	var returns *AccessGroupExcludeExternalEvaluation
	_jsii_.Get(
		j,
		"externalEvaluationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) Geo() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"geo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) GeoInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"geoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) Github() AccessGroupExcludeGithubList {
	var returns AccessGroupExcludeGithubList
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) GithubInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"githubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) Group() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) GroupInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) Gsuite() AccessGroupExcludeGsuiteList {
	var returns AccessGroupExcludeGsuiteList
	_jsii_.Get(
		j,
		"gsuite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) GsuiteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gsuiteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) Ip() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) IpInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) IpList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) IpListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) LoginMethod() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loginMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) LoginMethodInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loginMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) Okta() AccessGroupExcludeOktaList {
	var returns AccessGroupExcludeOktaList
	_jsii_.Get(
		j,
		"okta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) OktaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oktaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) Saml() AccessGroupExcludeSamlList {
	var returns AccessGroupExcludeSamlList
	_jsii_.Get(
		j,
		"saml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) SamlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"samlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) ServiceToken() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) ServiceTokenInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAccessGroupExcludeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AccessGroupExcludeOutputReference {
	_init_.Initialize()

	if err := validateNewAccessGroupExcludeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessGroupExcludeOutputReference{}

	_jsii_.Create(
		"cloudflare.accessGroup.AccessGroupExcludeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAccessGroupExcludeOutputReference_Override(a AccessGroupExcludeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudflare.accessGroup.AccessGroupExcludeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference)SetAnyValidServiceToken(val interface{}) {
	if err := j.validateSetAnyValidServiceTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anyValidServiceToken",
		val,
	)
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference)SetAuthMethod(val *string) {
	if err := j.validateSetAuthMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authMethod",
		val,
	)
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference)SetCertificate(val interface{}) {
	if err := j.validateSetCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference)SetCommonName(val *string) {
	if err := j.validateSetCommonNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commonName",
		val,
	)
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference)SetDevicePosture(val *[]*string) {
	if err := j.validateSetDevicePostureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"devicePosture",
		val,
	)
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference)SetEmail(val *[]*string) {
	if err := j.validateSetEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference)SetEmailDomain(val *[]*string) {
	if err := j.validateSetEmailDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailDomain",
		val,
	)
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference)SetEveryone(val interface{}) {
	if err := j.validateSetEveryoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"everyone",
		val,
	)
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference)SetGeo(val *[]*string) {
	if err := j.validateSetGeoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geo",
		val,
	)
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference)SetGroup(val *[]*string) {
	if err := j.validateSetGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"group",
		val,
	)
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference)SetIp(val *[]*string) {
	if err := j.validateSetIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ip",
		val,
	)
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference)SetIpList(val *[]*string) {
	if err := j.validateSetIpListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipList",
		val,
	)
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference)SetLoginMethod(val *[]*string) {
	if err := j.validateSetLoginMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginMethod",
		val,
	)
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference)SetServiceToken(val *[]*string) {
	if err := j.validateSetServiceTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceToken",
		val,
	)
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccessGroupExcludeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccessGroupExcludeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessGroupExcludeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccessGroupExcludeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccessGroupExcludeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccessGroupExcludeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccessGroupExcludeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccessGroupExcludeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccessGroupExcludeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccessGroupExcludeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccessGroupExcludeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccessGroupExcludeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessGroupExcludeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccessGroupExcludeOutputReference) PutAzure(value interface{}) {
	if err := a.validatePutAzureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAzure",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessGroupExcludeOutputReference) PutExternalEvaluation(value *AccessGroupExcludeExternalEvaluation) {
	if err := a.validatePutExternalEvaluationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExternalEvaluation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessGroupExcludeOutputReference) PutGithub(value interface{}) {
	if err := a.validatePutGithubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGithub",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessGroupExcludeOutputReference) PutGsuite(value interface{}) {
	if err := a.validatePutGsuiteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGsuite",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessGroupExcludeOutputReference) PutOkta(value interface{}) {
	if err := a.validatePutOktaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOkta",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessGroupExcludeOutputReference) PutSaml(value interface{}) {
	if err := a.validatePutSamlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSaml",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessGroupExcludeOutputReference) ResetAnyValidServiceToken() {
	_jsii_.InvokeVoid(
		a,
		"resetAnyValidServiceToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessGroupExcludeOutputReference) ResetAuthMethod() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthMethod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessGroupExcludeOutputReference) ResetAzure() {
	_jsii_.InvokeVoid(
		a,
		"resetAzure",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessGroupExcludeOutputReference) ResetCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessGroupExcludeOutputReference) ResetCommonName() {
	_jsii_.InvokeVoid(
		a,
		"resetCommonName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessGroupExcludeOutputReference) ResetDevicePosture() {
	_jsii_.InvokeVoid(
		a,
		"resetDevicePosture",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessGroupExcludeOutputReference) ResetEmail() {
	_jsii_.InvokeVoid(
		a,
		"resetEmail",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessGroupExcludeOutputReference) ResetEmailDomain() {
	_jsii_.InvokeVoid(
		a,
		"resetEmailDomain",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessGroupExcludeOutputReference) ResetEveryone() {
	_jsii_.InvokeVoid(
		a,
		"resetEveryone",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessGroupExcludeOutputReference) ResetExternalEvaluation() {
	_jsii_.InvokeVoid(
		a,
		"resetExternalEvaluation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessGroupExcludeOutputReference) ResetGeo() {
	_jsii_.InvokeVoid(
		a,
		"resetGeo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessGroupExcludeOutputReference) ResetGithub() {
	_jsii_.InvokeVoid(
		a,
		"resetGithub",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessGroupExcludeOutputReference) ResetGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessGroupExcludeOutputReference) ResetGsuite() {
	_jsii_.InvokeVoid(
		a,
		"resetGsuite",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessGroupExcludeOutputReference) ResetIp() {
	_jsii_.InvokeVoid(
		a,
		"resetIp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessGroupExcludeOutputReference) ResetIpList() {
	_jsii_.InvokeVoid(
		a,
		"resetIpList",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessGroupExcludeOutputReference) ResetLoginMethod() {
	_jsii_.InvokeVoid(
		a,
		"resetLoginMethod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessGroupExcludeOutputReference) ResetOkta() {
	_jsii_.InvokeVoid(
		a,
		"resetOkta",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessGroupExcludeOutputReference) ResetSaml() {
	_jsii_.InvokeVoid(
		a,
		"resetSaml",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessGroupExcludeOutputReference) ResetServiceToken() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessGroupExcludeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AccessGroupExcludeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

