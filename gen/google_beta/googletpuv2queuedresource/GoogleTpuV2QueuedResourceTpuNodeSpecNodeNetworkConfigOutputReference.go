package googletpuv2queuedresource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googletpuv2queuedresource/internal"
)

type GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference interface {
	cdktf.ComplexObject
	CanIpForward() interface{}
	SetCanIpForward(val interface{})
	CanIpForwardInput() interface{}
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
	EnableExternalIps() interface{}
	SetEnableExternalIps(val interface{})
	EnableExternalIpsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfig
	SetInternalValue(val *GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfig)
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	QueueCount() *float64
	SetQueueCount(val *float64)
	QueueCountInput() *float64
	Subnetwork() *string
	SetSubnetwork(val *string)
	SubnetworkInput() *string
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
	ResetCanIpForward()
	ResetEnableExternalIps()
	ResetNetwork()
	ResetQueueCount()
	ResetSubnetwork()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference
type jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) CanIpForward() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canIpForward",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) CanIpForwardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canIpForwardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) EnableExternalIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExternalIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) EnableExternalIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExternalIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) InternalValue() *GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfig {
	var returns *GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) QueueCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queueCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) QueueCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queueCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) Subnetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) SubnetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleTpuV2QueuedResource.GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference_Override(g GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleTpuV2QueuedResource.GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference)SetCanIpForward(val interface{}) {
	if err := j.validateSetCanIpForwardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"canIpForward",
		val,
	)
}

func (j *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference)SetEnableExternalIps(val interface{}) {
	if err := j.validateSetEnableExternalIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableExternalIps",
		val,
	)
}

func (j *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference)SetInternalValue(val *GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference)SetQueueCount(val *float64) {
	if err := j.validateSetQueueCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queueCount",
		val,
	)
}

func (j *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference)SetSubnetwork(val *string) {
	if err := j.validateSetSubnetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetwork",
		val,
	)
}

func (j *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) ResetCanIpForward() {
	_jsii_.InvokeVoid(
		g,
		"resetCanIpForward",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) ResetEnableExternalIps() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableExternalIps",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) ResetNetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetNetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) ResetQueueCount() {
	_jsii_.InvokeVoid(
		g,
		"resetQueueCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) ResetSubnetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetSubnetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

