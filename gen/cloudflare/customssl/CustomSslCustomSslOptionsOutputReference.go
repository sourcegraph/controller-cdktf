package customssl

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/customssl/internal"
)

type CustomSslCustomSslOptionsOutputReference interface {
	cdktf.ComplexObject
	BundleMethod() *string
	SetBundleMethod(val *string)
	BundleMethodInput() *string
	Certificate() *string
	SetCertificate(val *string)
	CertificateInput() *string
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
	GeoRestrictions() *string
	SetGeoRestrictions(val *string)
	GeoRestrictionsInput() *string
	InternalValue() *CustomSslCustomSslOptions
	SetInternalValue(val *CustomSslCustomSslOptions)
	PrivateKey() *string
	SetPrivateKey(val *string)
	PrivateKeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetBundleMethod()
	ResetCertificate()
	ResetGeoRestrictions()
	ResetPrivateKey()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CustomSslCustomSslOptionsOutputReference
type jsiiProxy_CustomSslCustomSslOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CustomSslCustomSslOptionsOutputReference) BundleMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSslCustomSslOptionsOutputReference) BundleMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSslCustomSslOptionsOutputReference) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSslCustomSslOptionsOutputReference) CertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSslCustomSslOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSslCustomSslOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSslCustomSslOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSslCustomSslOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSslCustomSslOptionsOutputReference) GeoRestrictions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geoRestrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSslCustomSslOptionsOutputReference) GeoRestrictionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geoRestrictionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSslCustomSslOptionsOutputReference) InternalValue() *CustomSslCustomSslOptions {
	var returns *CustomSslCustomSslOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSslCustomSslOptionsOutputReference) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSslCustomSslOptionsOutputReference) PrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSslCustomSslOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSslCustomSslOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSslCustomSslOptionsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSslCustomSslOptionsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewCustomSslCustomSslOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CustomSslCustomSslOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewCustomSslCustomSslOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomSslCustomSslOptionsOutputReference{}

	_jsii_.Create(
		"cloudflare.customSsl.CustomSslCustomSslOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCustomSslCustomSslOptionsOutputReference_Override(c CustomSslCustomSslOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudflare.customSsl.CustomSslCustomSslOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CustomSslCustomSslOptionsOutputReference)SetBundleMethod(val *string) {
	if err := j.validateSetBundleMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bundleMethod",
		val,
	)
}

func (j *jsiiProxy_CustomSslCustomSslOptionsOutputReference)SetCertificate(val *string) {
	if err := j.validateSetCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_CustomSslCustomSslOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CustomSslCustomSslOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CustomSslCustomSslOptionsOutputReference)SetGeoRestrictions(val *string) {
	if err := j.validateSetGeoRestrictionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geoRestrictions",
		val,
	)
}

func (j *jsiiProxy_CustomSslCustomSslOptionsOutputReference)SetInternalValue(val *CustomSslCustomSslOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CustomSslCustomSslOptionsOutputReference)SetPrivateKey(val *string) {
	if err := j.validateSetPrivateKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateKey",
		val,
	)
}

func (j *jsiiProxy_CustomSslCustomSslOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CustomSslCustomSslOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CustomSslCustomSslOptionsOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (c *jsiiProxy_CustomSslCustomSslOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomSslCustomSslOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CustomSslCustomSslOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CustomSslCustomSslOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CustomSslCustomSslOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CustomSslCustomSslOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CustomSslCustomSslOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CustomSslCustomSslOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CustomSslCustomSslOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CustomSslCustomSslOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CustomSslCustomSslOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomSslCustomSslOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CustomSslCustomSslOptionsOutputReference) ResetBundleMethod() {
	_jsii_.InvokeVoid(
		c,
		"resetBundleMethod",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomSslCustomSslOptionsOutputReference) ResetCertificate() {
	_jsii_.InvokeVoid(
		c,
		"resetCertificate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomSslCustomSslOptionsOutputReference) ResetGeoRestrictions() {
	_jsii_.InvokeVoid(
		c,
		"resetGeoRestrictions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomSslCustomSslOptionsOutputReference) ResetPrivateKey() {
	_jsii_.InvokeVoid(
		c,
		"resetPrivateKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomSslCustomSslOptionsOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		c,
		"resetType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomSslCustomSslOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CustomSslCustomSslOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

