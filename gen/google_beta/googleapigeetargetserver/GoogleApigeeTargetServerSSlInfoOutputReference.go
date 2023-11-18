package googleapigeetargetserver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleapigeetargetserver/internal"
)

type GoogleApigeeTargetServerSSlInfoOutputReference interface {
	cdktf.ComplexObject
	Ciphers() *[]*string
	SetCiphers(val *[]*string)
	CiphersInput() *[]*string
	ClientAuthEnabled() interface{}
	SetClientAuthEnabled(val interface{})
	ClientAuthEnabledInput() interface{}
	CommonName() GoogleApigeeTargetServerSSlInfoCommonNameOutputReference
	CommonNameInput() *GoogleApigeeTargetServerSSlInfoCommonName
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	IgnoreValidationErrors() interface{}
	SetIgnoreValidationErrors(val interface{})
	IgnoreValidationErrorsInput() interface{}
	InternalValue() *GoogleApigeeTargetServerSSlInfo
	SetInternalValue(val *GoogleApigeeTargetServerSSlInfo)
	KeyAlias() *string
	SetKeyAlias(val *string)
	KeyAliasInput() *string
	KeyStore() *string
	SetKeyStore(val *string)
	KeyStoreInput() *string
	Protocols() *[]*string
	SetProtocols(val *[]*string)
	ProtocolsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrustStore() *string
	SetTrustStore(val *string)
	TrustStoreInput() *string
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
	PutCommonName(value *GoogleApigeeTargetServerSSlInfoCommonName)
	ResetCiphers()
	ResetClientAuthEnabled()
	ResetCommonName()
	ResetIgnoreValidationErrors()
	ResetKeyAlias()
	ResetKeyStore()
	ResetProtocols()
	ResetTrustStore()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleApigeeTargetServerSSlInfoOutputReference
type jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) Ciphers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ciphers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) CiphersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ciphersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) ClientAuthEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientAuthEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) ClientAuthEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientAuthEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) CommonName() GoogleApigeeTargetServerSSlInfoCommonNameOutputReference {
	var returns GoogleApigeeTargetServerSSlInfoCommonNameOutputReference
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) CommonNameInput() *GoogleApigeeTargetServerSSlInfoCommonName {
	var returns *GoogleApigeeTargetServerSSlInfoCommonName
	_jsii_.Get(
		j,
		"commonNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) IgnoreValidationErrors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreValidationErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) IgnoreValidationErrorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreValidationErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) InternalValue() *GoogleApigeeTargetServerSSlInfo {
	var returns *GoogleApigeeTargetServerSSlInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) KeyAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) KeyAliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyAliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) KeyStore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) KeyStoreInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) Protocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) ProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) TrustStore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) TrustStoreInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStoreInput",
		&returns,
	)
	return returns
}


func NewGoogleApigeeTargetServerSSlInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleApigeeTargetServerSSlInfoOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleApigeeTargetServerSSlInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleApigeeTargetServer.GoogleApigeeTargetServerSSlInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleApigeeTargetServerSSlInfoOutputReference_Override(g GoogleApigeeTargetServerSSlInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleApigeeTargetServer.GoogleApigeeTargetServerSSlInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference)SetCiphers(val *[]*string) {
	if err := j.validateSetCiphersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ciphers",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference)SetClientAuthEnabled(val interface{}) {
	if err := j.validateSetClientAuthEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientAuthEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference)SetIgnoreValidationErrors(val interface{}) {
	if err := j.validateSetIgnoreValidationErrorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreValidationErrors",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference)SetInternalValue(val *GoogleApigeeTargetServerSSlInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference)SetKeyAlias(val *string) {
	if err := j.validateSetKeyAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyAlias",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference)SetKeyStore(val *string) {
	if err := j.validateSetKeyStoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyStore",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference)SetProtocols(val *[]*string) {
	if err := j.validateSetProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocols",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference)SetTrustStore(val *string) {
	if err := j.validateSetTrustStoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustStore",
		val,
	)
}

func (g *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) PutCommonName(value *GoogleApigeeTargetServerSSlInfoCommonName) {
	if err := g.validatePutCommonNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCommonName",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) ResetCiphers() {
	_jsii_.InvokeVoid(
		g,
		"resetCiphers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) ResetClientAuthEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetClientAuthEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) ResetCommonName() {
	_jsii_.InvokeVoid(
		g,
		"resetCommonName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) ResetIgnoreValidationErrors() {
	_jsii_.InvokeVoid(
		g,
		"resetIgnoreValidationErrors",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) ResetKeyAlias() {
	_jsii_.InvokeVoid(
		g,
		"resetKeyAlias",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) ResetKeyStore() {
	_jsii_.InvokeVoid(
		g,
		"resetKeyStore",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) ResetProtocols() {
	_jsii_.InvokeVoid(
		g,
		"resetProtocols",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) ResetTrustStore() {
	_jsii_.InvokeVoid(
		g,
		"resetTrustStore",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeTargetServerSSlInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

