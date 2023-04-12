package googleprivatecacertificate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googleprivatecacertificate/internal"
)

type GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference interface {
	cdktf.ComplexObject
	ClientAuth() interface{}
	SetClientAuth(val interface{})
	ClientAuthInput() interface{}
	CodeSigning() interface{}
	SetCodeSigning(val interface{})
	CodeSigningInput() interface{}
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
	EmailProtection() interface{}
	SetEmailProtection(val interface{})
	EmailProtectionInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsage
	SetInternalValue(val *GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsage)
	OcspSigning() interface{}
	SetOcspSigning(val interface{})
	OcspSigningInput() interface{}
	ServerAuth() interface{}
	SetServerAuth(val interface{})
	ServerAuthInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeStamping() interface{}
	SetTimeStamping(val interface{})
	TimeStampingInput() interface{}
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
	ResetClientAuth()
	ResetCodeSigning()
	ResetEmailProtection()
	ResetOcspSigning()
	ResetServerAuth()
	ResetTimeStamping()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference
type jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) ClientAuth() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) ClientAuthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) CodeSigning() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeSigning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) CodeSigningInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeSigningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) EmailProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) EmailProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) InternalValue() *GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsage {
	var returns *GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsage
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) OcspSigning() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ocspSigning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) OcspSigningInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ocspSigningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) ServerAuth() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) ServerAuthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) TimeStamping() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeStamping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) TimeStampingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeStampingInput",
		&returns,
	)
	return returns
}


func NewGooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference {
	_init_.Initialize()

	if err := validateNewGooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googlePrivatecaCertificate.GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference_Override(g GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googlePrivatecaCertificate.GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference)SetClientAuth(val interface{}) {
	if err := j.validateSetClientAuthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientAuth",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference)SetCodeSigning(val interface{}) {
	if err := j.validateSetCodeSigningParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codeSigning",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference)SetEmailProtection(val interface{}) {
	if err := j.validateSetEmailProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailProtection",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference)SetInternalValue(val *GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsage) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference)SetOcspSigning(val interface{}) {
	if err := j.validateSetOcspSigningParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ocspSigning",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference)SetServerAuth(val interface{}) {
	if err := j.validateSetServerAuthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverAuth",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference)SetTimeStamping(val interface{}) {
	if err := j.validateSetTimeStampingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeStamping",
		val,
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) ResetClientAuth() {
	_jsii_.InvokeVoid(
		g,
		"resetClientAuth",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) ResetCodeSigning() {
	_jsii_.InvokeVoid(
		g,
		"resetCodeSigning",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) ResetEmailProtection() {
	_jsii_.InvokeVoid(
		g,
		"resetEmailProtection",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) ResetOcspSigning() {
	_jsii_.InvokeVoid(
		g,
		"resetOcspSigning",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) ResetServerAuth() {
	_jsii_.InvokeVoid(
		g,
		"resetServerAuth",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) ResetTimeStamping() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeStamping",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

