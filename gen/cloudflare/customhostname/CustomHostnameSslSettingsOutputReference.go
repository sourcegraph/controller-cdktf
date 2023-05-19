package customhostname

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/customhostname/internal"
)

type CustomHostnameSslSettingsOutputReference interface {
	cdktf.ComplexObject
	Ciphers() *[]*string
	SetCiphers(val *[]*string)
	CiphersInput() *[]*string
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
	EarlyHints() *string
	SetEarlyHints(val *string)
	EarlyHintsInput() *string
	// Experimental.
	Fqn() *string
	Http2() *string
	SetHttp2(val *string)
	Http2Input() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MinTlsVersion() *string
	SetMinTlsVersion(val *string)
	MinTlsVersionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Tls13() *string
	SetTls13(val *string)
	Tls13Input() *string
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
	ResetCiphers()
	ResetEarlyHints()
	ResetHttp2()
	ResetMinTlsVersion()
	ResetTls13()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CustomHostnameSslSettingsOutputReference
type jsiiProxy_CustomHostnameSslSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CustomHostnameSslSettingsOutputReference) Ciphers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ciphers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslSettingsOutputReference) CiphersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ciphersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslSettingsOutputReference) EarlyHints() *string {
	var returns *string
	_jsii_.Get(
		j,
		"earlyHints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslSettingsOutputReference) EarlyHintsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"earlyHintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslSettingsOutputReference) Http2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"http2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslSettingsOutputReference) Http2Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"http2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslSettingsOutputReference) MinTlsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTlsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslSettingsOutputReference) MinTlsVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTlsVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslSettingsOutputReference) Tls13() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tls13",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslSettingsOutputReference) Tls13Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tls13Input",
		&returns,
	)
	return returns
}


func NewCustomHostnameSslSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CustomHostnameSslSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewCustomHostnameSslSettingsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomHostnameSslSettingsOutputReference{}

	_jsii_.Create(
		"cloudflare.customHostname.CustomHostnameSslSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCustomHostnameSslSettingsOutputReference_Override(c CustomHostnameSslSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudflare.customHostname.CustomHostnameSslSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CustomHostnameSslSettingsOutputReference)SetCiphers(val *[]*string) {
	if err := j.validateSetCiphersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ciphers",
		val,
	)
}

func (j *jsiiProxy_CustomHostnameSslSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CustomHostnameSslSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CustomHostnameSslSettingsOutputReference)SetEarlyHints(val *string) {
	if err := j.validateSetEarlyHintsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"earlyHints",
		val,
	)
}

func (j *jsiiProxy_CustomHostnameSslSettingsOutputReference)SetHttp2(val *string) {
	if err := j.validateSetHttp2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"http2",
		val,
	)
}

func (j *jsiiProxy_CustomHostnameSslSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CustomHostnameSslSettingsOutputReference)SetMinTlsVersion(val *string) {
	if err := j.validateSetMinTlsVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTlsVersion",
		val,
	)
}

func (j *jsiiProxy_CustomHostnameSslSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CustomHostnameSslSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CustomHostnameSslSettingsOutputReference)SetTls13(val *string) {
	if err := j.validateSetTls13Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tls13",
		val,
	)
}

func (c *jsiiProxy_CustomHostnameSslSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostnameSslSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostnameSslSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostnameSslSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostnameSslSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostnameSslSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostnameSslSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostnameSslSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostnameSslSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostnameSslSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostnameSslSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostnameSslSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostnameSslSettingsOutputReference) ResetCiphers() {
	_jsii_.InvokeVoid(
		c,
		"resetCiphers",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomHostnameSslSettingsOutputReference) ResetEarlyHints() {
	_jsii_.InvokeVoid(
		c,
		"resetEarlyHints",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomHostnameSslSettingsOutputReference) ResetHttp2() {
	_jsii_.InvokeVoid(
		c,
		"resetHttp2",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomHostnameSslSettingsOutputReference) ResetMinTlsVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetMinTlsVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomHostnameSslSettingsOutputReference) ResetTls13() {
	_jsii_.InvokeVoid(
		c,
		"resetTls13",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomHostnameSslSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostnameSslSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

