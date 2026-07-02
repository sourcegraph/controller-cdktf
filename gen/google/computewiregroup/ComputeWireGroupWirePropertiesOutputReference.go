package computewiregroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/computewiregroup/internal"
)

type ComputeWireGroupWirePropertiesOutputReference interface {
	cdktf.ComplexObject
	BandwidthAllocation() *string
	SetBandwidthAllocation(val *string)
	BandwidthAllocationInput() *string
	BandwidthUnmetered() *float64
	SetBandwidthUnmetered(val *float64)
	BandwidthUnmeteredInput() *float64
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
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
	FaultResponse() *string
	SetFaultResponse(val *string)
	FaultResponseInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeWireGroupWireProperties
	SetInternalValue(val *ComputeWireGroupWireProperties)
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
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	ResetBandwidthUnmetered()
	ResetFaultResponse()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeWireGroupWirePropertiesOutputReference
type jsiiProxy_ComputeWireGroupWirePropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) BandwidthAllocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bandwidthAllocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) BandwidthAllocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bandwidthAllocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) BandwidthUnmetered() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthUnmetered",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) BandwidthUnmeteredInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthUnmeteredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) FaultResponse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"faultResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) FaultResponseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"faultResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) InternalValue() *ComputeWireGroupWireProperties {
	var returns *ComputeWireGroupWireProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewComputeWireGroupWirePropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeWireGroupWirePropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewComputeWireGroupWirePropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeWireGroupWirePropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeWireGroup.ComputeWireGroupWirePropertiesOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeWireGroupWirePropertiesOutputReference_Override(c ComputeWireGroupWirePropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeWireGroup.ComputeWireGroupWirePropertiesOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) SetBandwidthAllocation(val *string) {
	if err := j.validateSetBandwidthAllocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidthAllocation",
		val,
	)
}

func (j *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) SetBandwidthUnmetered(val *float64) {
	if err := j.validateSetBandwidthUnmeteredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidthUnmetered",
		val,
	)
}

func (j *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) SetFaultResponse(val *string) {
	if err := j.validateSetFaultResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"faultResponse",
		val,
	)
}

func (j *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) SetInternalValue(val *ComputeWireGroupWireProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) ResetBandwidthUnmetered() {
	_jsii_.InvokeVoid(
		c,
		"resetBandwidthUnmetered",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) ResetFaultResponse() {
	_jsii_.InvokeVoid(
		c,
		"resetFaultResponse",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		c,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeWireGroupWirePropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
