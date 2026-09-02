package developerconnectaccountconnector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/developerconnectaccountconnector/internal"
)

type DeveloperConnectAccountConnectorCustomOauthConfigOutputReference interface {
	cdktn.ComplexObject
	AuthUri() *string
	SetAuthUri(val *string)
	AuthUriInput() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	HostUri() *string
	SetHostUri(val *string)
	HostUriInput() *string
	InternalValue() *DeveloperConnectAccountConnectorCustomOauthConfig
	SetInternalValue(val *DeveloperConnectAccountConnectorCustomOauthConfig)
	PkceDisabled() interface{}
	SetPkceDisabled(val interface{})
	PkceDisabledInput() interface{}
	ScmProvider() *string
	SetScmProvider(val *string)
	ScmProviderInput() *string
	Scopes() *[]*string
	SetScopes(val *[]*string)
	ScopesInput() *[]*string
	ServerVersion() *string
	ServiceDirectoryConfig() DeveloperConnectAccountConnectorCustomOauthConfigServiceDirectoryConfigOutputReference
	ServiceDirectoryConfigInput() *DeveloperConnectAccountConnectorCustomOauthConfigServiceDirectoryConfig
	SslCaCertificate() *string
	SetSslCaCertificate(val *string)
	SslCaCertificateInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TokenUri() *string
	SetTokenUri(val *string)
	TokenUriInput() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutServiceDirectoryConfig(value *DeveloperConnectAccountConnectorCustomOauthConfigServiceDirectoryConfig)
	ResetPkceDisabled()
	ResetServiceDirectoryConfig()
	ResetSslCaCertificate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DeveloperConnectAccountConnectorCustomOauthConfigOutputReference
type jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) AuthUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) AuthUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) HostUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) HostUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) InternalValue() *DeveloperConnectAccountConnectorCustomOauthConfig {
	var returns *DeveloperConnectAccountConnectorCustomOauthConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) PkceDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pkceDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) PkceDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pkceDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ScmProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ScmProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ServiceDirectoryConfig() DeveloperConnectAccountConnectorCustomOauthConfigServiceDirectoryConfigOutputReference {
	var returns DeveloperConnectAccountConnectorCustomOauthConfigServiceDirectoryConfigOutputReference
	_jsii_.Get(
		j,
		"serviceDirectoryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ServiceDirectoryConfigInput() *DeveloperConnectAccountConnectorCustomOauthConfigServiceDirectoryConfig {
	var returns *DeveloperConnectAccountConnectorCustomOauthConfigServiceDirectoryConfig
	_jsii_.Get(
		j,
		"serviceDirectoryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) SslCaCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) SslCaCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) TokenUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) TokenUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUriInput",
		&returns,
	)
	return returns
}


func NewDeveloperConnectAccountConnectorCustomOauthConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DeveloperConnectAccountConnectorCustomOauthConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDeveloperConnectAccountConnectorCustomOauthConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.developerConnectAccountConnector.DeveloperConnectAccountConnectorCustomOauthConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDeveloperConnectAccountConnectorCustomOauthConfigOutputReference_Override(d DeveloperConnectAccountConnectorCustomOauthConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.developerConnectAccountConnector.DeveloperConnectAccountConnectorCustomOauthConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference)SetAuthUri(val *string) {
	if err := j.validateSetAuthUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authUri",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference)SetHostUri(val *string) {
	if err := j.validateSetHostUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostUri",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference)SetInternalValue(val *DeveloperConnectAccountConnectorCustomOauthConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference)SetPkceDisabled(val interface{}) {
	if err := j.validateSetPkceDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pkceDisabled",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference)SetScmProvider(val *string) {
	if err := j.validateSetScmProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scmProvider",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference)SetScopes(val *[]*string) {
	if err := j.validateSetScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference)SetSslCaCertificate(val *string) {
	if err := j.validateSetSslCaCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCaCertificate",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference)SetTokenUri(val *string) {
	if err := j.validateSetTokenUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenUri",
		val,
	)
}

func (d *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) PutServiceDirectoryConfig(value *DeveloperConnectAccountConnectorCustomOauthConfigServiceDirectoryConfig) {
	if err := d.validatePutServiceDirectoryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putServiceDirectoryConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ResetPkceDisabled() {
	_jsii_.InvokeVoid(
		d,
		"resetPkceDisabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ResetServiceDirectoryConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceDirectoryConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ResetSslCaCertificate() {
	_jsii_.InvokeVoid(
		d,
		"resetSslCaCertificate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

