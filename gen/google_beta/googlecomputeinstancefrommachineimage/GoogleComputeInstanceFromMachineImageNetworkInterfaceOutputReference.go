package googlecomputeinstancefrommachineimage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeinstancefrommachineimage/internal"
)

type GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference interface {
	cdktf.ComplexObject
	AccessConfig() GoogleComputeInstanceFromMachineImageNetworkInterfaceAccessConfigList
	AccessConfigInput() interface{}
	AliasIpRange() GoogleComputeInstanceFromMachineImageNetworkInterfaceAliasIpRangeList
	AliasIpRangeInput() interface{}
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Ipv6AccessConfig() GoogleComputeInstanceFromMachineImageNetworkInterfaceIpv6AccessConfigList
	Ipv6AccessConfigInput() interface{}
	Ipv6AccessType() *string
	Name() *string
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	NetworkIp() *string
	SetNetworkIp(val *string)
	NetworkIpInput() *string
	NicType() *string
	SetNicType(val *string)
	NicTypeInput() *string
	QueueCount() *float64
	SetQueueCount(val *float64)
	QueueCountInput() *float64
	StackType() *string
	SetStackType(val *string)
	StackTypeInput() *string
	Subnetwork() *string
	SetSubnetwork(val *string)
	SubnetworkInput() *string
	SubnetworkProject() *string
	SetSubnetworkProject(val *string)
	SubnetworkProjectInput() *string
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
	PutAccessConfig(value interface{})
	PutAliasIpRange(value interface{})
	PutIpv6AccessConfig(value interface{})
	ResetAccessConfig()
	ResetAliasIpRange()
	ResetIpv6AccessConfig()
	ResetNetwork()
	ResetNetworkIp()
	ResetNicType()
	ResetQueueCount()
	ResetStackType()
	ResetSubnetwork()
	ResetSubnetworkProject()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference
type jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) AccessConfig() GoogleComputeInstanceFromMachineImageNetworkInterfaceAccessConfigList {
	var returns GoogleComputeInstanceFromMachineImageNetworkInterfaceAccessConfigList
	_jsii_.Get(
		j,
		"accessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) AccessConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) AliasIpRange() GoogleComputeInstanceFromMachineImageNetworkInterfaceAliasIpRangeList {
	var returns GoogleComputeInstanceFromMachineImageNetworkInterfaceAliasIpRangeList
	_jsii_.Get(
		j,
		"aliasIpRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) AliasIpRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aliasIpRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) Ipv6AccessConfig() GoogleComputeInstanceFromMachineImageNetworkInterfaceIpv6AccessConfigList {
	var returns GoogleComputeInstanceFromMachineImageNetworkInterfaceIpv6AccessConfigList
	_jsii_.Get(
		j,
		"ipv6AccessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) Ipv6AccessConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6AccessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) Ipv6AccessType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6AccessType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) NetworkIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) NetworkIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) NicType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nicType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) NicTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nicTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) QueueCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queueCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) QueueCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queueCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) StackType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) StackTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) Subnetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) SubnetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) SubnetworkProject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetworkProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) SubnetworkProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetworkProjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeInstanceFromMachineImage.GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference_Override(g GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeInstanceFromMachineImage.GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference)SetNetworkIp(val *string) {
	if err := j.validateSetNetworkIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkIp",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference)SetNicType(val *string) {
	if err := j.validateSetNicTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nicType",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference)SetQueueCount(val *float64) {
	if err := j.validateSetQueueCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queueCount",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference)SetStackType(val *string) {
	if err := j.validateSetStackTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stackType",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference)SetSubnetwork(val *string) {
	if err := j.validateSetSubnetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetwork",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference)SetSubnetworkProject(val *string) {
	if err := j.validateSetSubnetworkProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetworkProject",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) PutAccessConfig(value interface{}) {
	if err := g.validatePutAccessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAccessConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) PutAliasIpRange(value interface{}) {
	if err := g.validatePutAliasIpRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAliasIpRange",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) PutIpv6AccessConfig(value interface{}) {
	if err := g.validatePutIpv6AccessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIpv6AccessConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) ResetAccessConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) ResetAliasIpRange() {
	_jsii_.InvokeVoid(
		g,
		"resetAliasIpRange",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) ResetIpv6AccessConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetIpv6AccessConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) ResetNetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetNetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) ResetNetworkIp() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkIp",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) ResetNicType() {
	_jsii_.InvokeVoid(
		g,
		"resetNicType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) ResetQueueCount() {
	_jsii_.InvokeVoid(
		g,
		"resetQueueCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) ResetStackType() {
	_jsii_.InvokeVoid(
		g,
		"resetStackType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) ResetSubnetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetSubnetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) ResetSubnetworkProject() {
	_jsii_.InvokeVoid(
		g,
		"resetSubnetworkProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImageNetworkInterfaceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

