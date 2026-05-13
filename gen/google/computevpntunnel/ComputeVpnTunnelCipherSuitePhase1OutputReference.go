package computevpntunnel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/computevpntunnel/internal"
)

type ComputeVpnTunnelCipherSuitePhase1OutputReference interface {
	cdktf.ComplexObject
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
	Dh() *[]*string
	SetDh(val *[]*string)
	DhInput() *[]*string
	Encryption() *[]*string
	SetEncryption(val *[]*string)
	EncryptionInput() *[]*string
	// Experimental.
	Fqn() *string
	Integrity() *[]*string
	SetIntegrity(val *[]*string)
	IntegrityInput() *[]*string
	InternalValue() *ComputeVpnTunnelCipherSuitePhase1
	SetInternalValue(val *ComputeVpnTunnelCipherSuitePhase1)
	Prf() *[]*string
	SetPrf(val *[]*string)
	PrfInput() *[]*string
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
	ResetDh()
	ResetEncryption()
	ResetIntegrity()
	ResetPrf()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeVpnTunnelCipherSuitePhase1OutputReference
type jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) Dh() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) DhInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dhInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) Encryption() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"encryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) EncryptionInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"encryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) Integrity() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"integrity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) IntegrityInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"integrityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) InternalValue() *ComputeVpnTunnelCipherSuitePhase1 {
	var returns *ComputeVpnTunnelCipherSuitePhase1
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) Prf() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"prf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) PrfInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"prfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeVpnTunnelCipherSuitePhase1OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeVpnTunnelCipherSuitePhase1OutputReference {
	_init_.Initialize()

	if err := validateNewComputeVpnTunnelCipherSuitePhase1OutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeVpnTunnel.ComputeVpnTunnelCipherSuitePhase1OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeVpnTunnelCipherSuitePhase1OutputReference_Override(c ComputeVpnTunnelCipherSuitePhase1OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeVpnTunnel.ComputeVpnTunnelCipherSuitePhase1OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference)SetDh(val *[]*string) {
	if err := j.validateSetDhParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dh",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference)SetEncryption(val *[]*string) {
	if err := j.validateSetEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryption",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference)SetIntegrity(val *[]*string) {
	if err := j.validateSetIntegrityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrity",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference)SetInternalValue(val *ComputeVpnTunnelCipherSuitePhase1) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference)SetPrf(val *[]*string) {
	if err := j.validateSetPrfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prf",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) ResetDh() {
	_jsii_.InvokeVoid(
		c,
		"resetDh",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) ResetEncryption() {
	_jsii_.InvokeVoid(
		c,
		"resetEncryption",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) ResetIntegrity() {
	_jsii_.InvokeVoid(
		c,
		"resetIntegrity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) ResetPrf() {
	_jsii_.InvokeVoid(
		c,
		"resetPrf",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase1OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

