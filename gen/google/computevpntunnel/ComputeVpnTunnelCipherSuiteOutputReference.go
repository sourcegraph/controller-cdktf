package computevpntunnel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/computevpntunnel/internal"
)

type ComputeVpnTunnelCipherSuiteOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeVpnTunnelCipherSuite
	SetInternalValue(val *ComputeVpnTunnelCipherSuite)
	Phase1() ComputeVpnTunnelCipherSuitePhase1OutputReference
	Phase1Input() *ComputeVpnTunnelCipherSuitePhase1
	Phase2() ComputeVpnTunnelCipherSuitePhase2OutputReference
	Phase2Input() *ComputeVpnTunnelCipherSuitePhase2
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
	PutPhase1(value *ComputeVpnTunnelCipherSuitePhase1)
	PutPhase2(value *ComputeVpnTunnelCipherSuitePhase2)
	ResetPhase1()
	ResetPhase2()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeVpnTunnelCipherSuiteOutputReference
type jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference) InternalValue() *ComputeVpnTunnelCipherSuite {
	var returns *ComputeVpnTunnelCipherSuite
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference) Phase1() ComputeVpnTunnelCipherSuitePhase1OutputReference {
	var returns ComputeVpnTunnelCipherSuitePhase1OutputReference
	_jsii_.Get(
		j,
		"phase1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference) Phase1Input() *ComputeVpnTunnelCipherSuitePhase1 {
	var returns *ComputeVpnTunnelCipherSuitePhase1
	_jsii_.Get(
		j,
		"phase1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference) Phase2() ComputeVpnTunnelCipherSuitePhase2OutputReference {
	var returns ComputeVpnTunnelCipherSuitePhase2OutputReference
	_jsii_.Get(
		j,
		"phase2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference) Phase2Input() *ComputeVpnTunnelCipherSuitePhase2 {
	var returns *ComputeVpnTunnelCipherSuitePhase2
	_jsii_.Get(
		j,
		"phase2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeVpnTunnelCipherSuiteOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ComputeVpnTunnelCipherSuiteOutputReference {
	_init_.Initialize()

	if err := validateNewComputeVpnTunnelCipherSuiteOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.computeVpnTunnel.ComputeVpnTunnelCipherSuiteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeVpnTunnelCipherSuiteOutputReference_Override(c ComputeVpnTunnelCipherSuiteOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.computeVpnTunnel.ComputeVpnTunnelCipherSuiteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference)SetInternalValue(val *ComputeVpnTunnelCipherSuite) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference) PutPhase1(value *ComputeVpnTunnelCipherSuitePhase1) {
	if err := c.validatePutPhase1Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPhase1",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference) PutPhase2(value *ComputeVpnTunnelCipherSuitePhase2) {
	if err := c.validatePutPhase2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPhase2",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference) ResetPhase1() {
	_jsii_.InvokeVoid(
		c,
		"resetPhase1",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference) ResetPhase2() {
	_jsii_.InvokeVoid(
		c,
		"resetPhase2",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeVpnTunnelCipherSuiteOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

