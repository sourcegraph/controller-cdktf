package googlecomputeinterconnectattachment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeinterconnectattachment/internal"
)

type GoogleComputeInterconnectAttachmentL2ForwardingOutputReference interface {
	cdktn.ComplexObject
	ApplianceMappings() GoogleComputeInterconnectAttachmentL2ForwardingApplianceMappingsList
	ApplianceMappingsInput() interface{}
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
	DefaultApplianceIpAddress() *string
	SetDefaultApplianceIpAddress(val *string)
	DefaultApplianceIpAddressInput() *string
	// Experimental.
	Fqn() *string
	GeneveHeader() GoogleComputeInterconnectAttachmentL2ForwardingGeneveHeaderOutputReference
	GeneveHeaderInput() *GoogleComputeInterconnectAttachmentL2ForwardingGeneveHeader
	InternalValue() *GoogleComputeInterconnectAttachmentL2Forwarding
	SetInternalValue(val *GoogleComputeInterconnectAttachmentL2Forwarding)
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TunnelEndpointIpAddress() *string
	SetTunnelEndpointIpAddress(val *string)
	TunnelEndpointIpAddressInput() *string
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
	PutApplianceMappings(value interface{})
	PutGeneveHeader(value *GoogleComputeInterconnectAttachmentL2ForwardingGeneveHeader)
	ResetApplianceMappings()
	ResetDefaultApplianceIpAddress()
	ResetGeneveHeader()
	ResetNetwork()
	ResetTunnelEndpointIpAddress()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeInterconnectAttachmentL2ForwardingOutputReference
type jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) ApplianceMappings() GoogleComputeInterconnectAttachmentL2ForwardingApplianceMappingsList {
	var returns GoogleComputeInterconnectAttachmentL2ForwardingApplianceMappingsList
	_jsii_.Get(
		j,
		"applianceMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) ApplianceMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applianceMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) DefaultApplianceIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultApplianceIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) DefaultApplianceIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultApplianceIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) GeneveHeader() GoogleComputeInterconnectAttachmentL2ForwardingGeneveHeaderOutputReference {
	var returns GoogleComputeInterconnectAttachmentL2ForwardingGeneveHeaderOutputReference
	_jsii_.Get(
		j,
		"geneveHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) GeneveHeaderInput() *GoogleComputeInterconnectAttachmentL2ForwardingGeneveHeader {
	var returns *GoogleComputeInterconnectAttachmentL2ForwardingGeneveHeader
	_jsii_.Get(
		j,
		"geneveHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) InternalValue() *GoogleComputeInterconnectAttachmentL2Forwarding {
	var returns *GoogleComputeInterconnectAttachmentL2Forwarding
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) TunnelEndpointIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelEndpointIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) TunnelEndpointIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelEndpointIpAddressInput",
		&returns,
	)
	return returns
}


func NewGoogleComputeInterconnectAttachmentL2ForwardingOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleComputeInterconnectAttachmentL2ForwardingOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeInterconnectAttachmentL2ForwardingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeInterconnectAttachment.GoogleComputeInterconnectAttachmentL2ForwardingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeInterconnectAttachmentL2ForwardingOutputReference_Override(g GoogleComputeInterconnectAttachmentL2ForwardingOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeInterconnectAttachment.GoogleComputeInterconnectAttachmentL2ForwardingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference)SetDefaultApplianceIpAddress(val *string) {
	if err := j.validateSetDefaultApplianceIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultApplianceIpAddress",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference)SetInternalValue(val *GoogleComputeInterconnectAttachmentL2Forwarding) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference)SetTunnelEndpointIpAddress(val *string) {
	if err := j.validateSetTunnelEndpointIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnelEndpointIpAddress",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) PutApplianceMappings(value interface{}) {
	if err := g.validatePutApplianceMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putApplianceMappings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) PutGeneveHeader(value *GoogleComputeInterconnectAttachmentL2ForwardingGeneveHeader) {
	if err := g.validatePutGeneveHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGeneveHeader",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) ResetApplianceMappings() {
	_jsii_.InvokeVoid(
		g,
		"resetApplianceMappings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) ResetDefaultApplianceIpAddress() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultApplianceIpAddress",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) ResetGeneveHeader() {
	_jsii_.InvokeVoid(
		g,
		"resetGeneveHeader",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) ResetNetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetNetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) ResetTunnelEndpointIpAddress() {
	_jsii_.InvokeVoid(
		g,
		"resetTunnelEndpointIpAddress",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachmentL2ForwardingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

