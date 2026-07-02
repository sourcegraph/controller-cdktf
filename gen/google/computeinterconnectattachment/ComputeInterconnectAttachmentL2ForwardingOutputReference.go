package computeinterconnectattachment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/computeinterconnectattachment/internal"
)

type ComputeInterconnectAttachmentL2ForwardingOutputReference interface {
	cdktf.ComplexObject
	ApplianceMappings() ComputeInterconnectAttachmentL2ForwardingApplianceMappingsList
	ApplianceMappingsInput() any
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
	DefaultApplianceIpAddress() *string
	SetDefaultApplianceIpAddress(val *string)
	DefaultApplianceIpAddressInput() *string
	// Experimental.
	Fqn() *string
	GeneveHeader() ComputeInterconnectAttachmentL2ForwardingGeneveHeaderOutputReference
	GeneveHeaderInput() *ComputeInterconnectAttachmentL2ForwardingGeneveHeader
	InternalValue() *ComputeInterconnectAttachmentL2Forwarding
	SetInternalValue(val *ComputeInterconnectAttachmentL2Forwarding)
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TunnelEndpointIpAddress() *string
	SetTunnelEndpointIpAddress(val *string)
	TunnelEndpointIpAddressInput() *string
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
	PutApplianceMappings(value any)
	PutGeneveHeader(value *ComputeInterconnectAttachmentL2ForwardingGeneveHeader)
	ResetApplianceMappings()
	ResetDefaultApplianceIpAddress()
	ResetGeneveHeader()
	ResetNetwork()
	ResetTunnelEndpointIpAddress()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeInterconnectAttachmentL2ForwardingOutputReference
type jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) ApplianceMappings() ComputeInterconnectAttachmentL2ForwardingApplianceMappingsList {
	var returns ComputeInterconnectAttachmentL2ForwardingApplianceMappingsList
	_jsii_.Get(
		j,
		"applianceMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) ApplianceMappingsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"applianceMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) DefaultApplianceIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultApplianceIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) DefaultApplianceIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultApplianceIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) GeneveHeader() ComputeInterconnectAttachmentL2ForwardingGeneveHeaderOutputReference {
	var returns ComputeInterconnectAttachmentL2ForwardingGeneveHeaderOutputReference
	_jsii_.Get(
		j,
		"geneveHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) GeneveHeaderInput() *ComputeInterconnectAttachmentL2ForwardingGeneveHeader {
	var returns *ComputeInterconnectAttachmentL2ForwardingGeneveHeader
	_jsii_.Get(
		j,
		"geneveHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) InternalValue() *ComputeInterconnectAttachmentL2Forwarding {
	var returns *ComputeInterconnectAttachmentL2Forwarding
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) TunnelEndpointIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelEndpointIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) TunnelEndpointIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelEndpointIpAddressInput",
		&returns,
	)
	return returns
}

func NewComputeInterconnectAttachmentL2ForwardingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeInterconnectAttachmentL2ForwardingOutputReference {
	_init_.Initialize()

	if err := validateNewComputeInterconnectAttachmentL2ForwardingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeInterconnectAttachment.ComputeInterconnectAttachmentL2ForwardingOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeInterconnectAttachmentL2ForwardingOutputReference_Override(c ComputeInterconnectAttachmentL2ForwardingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeInterconnectAttachment.ComputeInterconnectAttachmentL2ForwardingOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) SetDefaultApplianceIpAddress(val *string) {
	if err := j.validateSetDefaultApplianceIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultApplianceIpAddress",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) SetInternalValue(val *ComputeInterconnectAttachmentL2Forwarding) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) SetTunnelEndpointIpAddress(val *string) {
	if err := j.validateSetTunnelEndpointIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnelEndpointIpAddress",
		val,
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) PutApplianceMappings(value any) {
	if err := c.validatePutApplianceMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putApplianceMappings",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) PutGeneveHeader(value *ComputeInterconnectAttachmentL2ForwardingGeneveHeader) {
	if err := c.validatePutGeneveHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGeneveHeader",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) ResetApplianceMappings() {
	_jsii_.InvokeVoid(
		c,
		"resetApplianceMappings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) ResetDefaultApplianceIpAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultApplianceIpAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) ResetGeneveHeader() {
	_jsii_.InvokeVoid(
		c,
		"resetGeneveHeader",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) ResetNetwork() {
	_jsii_.InvokeVoid(
		c,
		"resetNetwork",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) ResetTunnelEndpointIpAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetTunnelEndpointIpAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_ComputeInterconnectAttachmentL2ForwardingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
