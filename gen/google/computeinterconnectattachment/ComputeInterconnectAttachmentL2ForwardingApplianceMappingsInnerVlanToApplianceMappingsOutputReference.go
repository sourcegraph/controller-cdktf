package computeinterconnectattachment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/computeinterconnectattachment/internal"
)

type ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference interface {
	cdktf.ComplexObject
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
	// Experimental.
	Fqn() *string
	InnerApplianceIpAddress() *string
	SetInnerApplianceIpAddress(val *string)
	InnerApplianceIpAddressInput() *string
	InnerVlanTags() *[]*string
	SetInnerVlanTags(val *[]*string)
	InnerVlanTagsInput() *[]*string
	InternalValue() any
	SetInternalValue(val any)
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
	ResetInnerApplianceIpAddress()
	ResetInnerVlanTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference
type jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) InnerApplianceIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"innerApplianceIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) InnerApplianceIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"innerApplianceIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) InnerVlanTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"innerVlanTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) InnerVlanTagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"innerVlanTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference {
	_init_.Initialize()

	if err := validateNewComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeInterconnectAttachment.ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference_Override(c ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeInterconnectAttachment.ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) SetInnerApplianceIpAddress(val *string) {
	if err := j.validateSetInnerApplianceIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"innerApplianceIpAddress",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) SetInnerVlanTags(val *[]*string) {
	if err := j.validateSetInnerVlanTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"innerVlanTags",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) ResetInnerApplianceIpAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetInnerApplianceIpAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) ResetInnerVlanTags() {
	_jsii_.InvokeVoid(
		c,
		"resetInnerVlanTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
