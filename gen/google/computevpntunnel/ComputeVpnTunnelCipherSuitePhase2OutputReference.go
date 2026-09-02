package computevpntunnel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/computevpntunnel/internal"
)

type ComputeVpnTunnelCipherSuitePhase2OutputReference interface {
	cdktn.ComplexObject
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
	Encryption() *[]*string
	SetEncryption(val *[]*string)
	EncryptionInput() *[]*string
	// Experimental.
	Fqn() *string
	Integrity() *[]*string
	SetIntegrity(val *[]*string)
	IntegrityInput() *[]*string
	InternalValue() *ComputeVpnTunnelCipherSuitePhase2
	SetInternalValue(val *ComputeVpnTunnelCipherSuitePhase2)
	Pfs() *[]*string
	SetPfs(val *[]*string)
	PfsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	ResetEncryption()
	ResetIntegrity()
	ResetPfs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeVpnTunnelCipherSuitePhase2OutputReference
type jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference) Encryption() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"encryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference) EncryptionInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"encryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference) Integrity() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"integrity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference) IntegrityInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"integrityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference) InternalValue() *ComputeVpnTunnelCipherSuitePhase2 {
	var returns *ComputeVpnTunnelCipherSuitePhase2
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference) Pfs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference) PfsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeVpnTunnelCipherSuitePhase2OutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ComputeVpnTunnelCipherSuitePhase2OutputReference {
	_init_.Initialize()

	if err := validateNewComputeVpnTunnelCipherSuitePhase2OutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.computeVpnTunnel.ComputeVpnTunnelCipherSuitePhase2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeVpnTunnelCipherSuitePhase2OutputReference_Override(c ComputeVpnTunnelCipherSuitePhase2OutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.computeVpnTunnel.ComputeVpnTunnelCipherSuitePhase2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference)SetEncryption(val *[]*string) {
	if err := j.validateSetEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryption",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference)SetIntegrity(val *[]*string) {
	if err := j.validateSetIntegrityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrity",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference)SetInternalValue(val *ComputeVpnTunnelCipherSuitePhase2) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference)SetPfs(val *[]*string) {
	if err := j.validateSetPfsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pfs",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference) ResetEncryption() {
	_jsii_.InvokeVoid(
		c,
		"resetEncryption",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference) ResetIntegrity() {
	_jsii_.InvokeVoid(
		c,
		"resetIntegrity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference) ResetPfs() {
	_jsii_.InvokeVoid(
		c,
		"resetPfs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeVpnTunnelCipherSuitePhase2OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

