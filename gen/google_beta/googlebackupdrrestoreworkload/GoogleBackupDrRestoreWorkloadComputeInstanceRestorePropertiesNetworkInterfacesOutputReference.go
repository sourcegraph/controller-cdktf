package googlebackupdrrestoreworkload

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebackupdrrestoreworkload/internal"
)

type GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference interface {
	cdktn.ComplexObject
	AccessConfigs() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAccessConfigsList
	AccessConfigsInput() interface{}
	AliasIpRanges() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesList
	AliasIpRangesInput() interface{}
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
	InternalIpv6PrefixLength() *float64
	SetInternalIpv6PrefixLength(val *float64)
	InternalIpv6PrefixLengthInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpAddress() *string
	SetIpAddress(val *string)
	IpAddressInput() *string
	Ipv6AccessConfigs() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsList
	Ipv6AccessConfigsInput() interface{}
	Ipv6AccessType() *string
	SetIpv6AccessType(val *string)
	Ipv6AccessTypeInput() *string
	Ipv6Address() *string
	SetIpv6Address(val *string)
	Ipv6AddressInput() *string
	Network() *string
	SetNetwork(val *string)
	NetworkAttachment() *string
	SetNetworkAttachment(val *string)
	NetworkAttachmentInput() *string
	NetworkInput() *string
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
	PutAccessConfigs(value interface{})
	PutAliasIpRanges(value interface{})
	PutIpv6AccessConfigs(value interface{})
	ResetAccessConfigs()
	ResetAliasIpRanges()
	ResetInternalIpv6PrefixLength()
	ResetIpAddress()
	ResetIpv6AccessConfigs()
	ResetIpv6AccessType()
	ResetIpv6Address()
	ResetNetwork()
	ResetNetworkAttachment()
	ResetNicType()
	ResetQueueCount()
	ResetStackType()
	ResetSubnetwork()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference
type jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) AccessConfigs() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAccessConfigsList {
	var returns GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAccessConfigsList
	_jsii_.Get(
		j,
		"accessConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) AccessConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) AliasIpRanges() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesList {
	var returns GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesList
	_jsii_.Get(
		j,
		"aliasIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) AliasIpRangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aliasIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) InternalIpv6PrefixLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"internalIpv6PrefixLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) InternalIpv6PrefixLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"internalIpv6PrefixLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) IpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) Ipv6AccessConfigs() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsList {
	var returns GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsList
	_jsii_.Get(
		j,
		"ipv6AccessConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) Ipv6AccessConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6AccessConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) Ipv6AccessType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6AccessType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) Ipv6AccessTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6AccessTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) Ipv6Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) Ipv6AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6AddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) NetworkAttachment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkAttachment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) NetworkAttachmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkAttachmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) NicType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nicType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) NicTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nicTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) QueueCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queueCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) QueueCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queueCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) StackType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) StackTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) Subnetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) SubnetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleBackupDrRestoreWorkload.GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference_Override(g GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleBackupDrRestoreWorkload.GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference)SetInternalIpv6PrefixLength(val *float64) {
	if err := j.validateSetInternalIpv6PrefixLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalIpv6PrefixLength",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference)SetIpAddress(val *string) {
	if err := j.validateSetIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddress",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference)SetIpv6AccessType(val *string) {
	if err := j.validateSetIpv6AccessTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6AccessType",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference)SetIpv6Address(val *string) {
	if err := j.validateSetIpv6AddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Address",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference)SetNetworkAttachment(val *string) {
	if err := j.validateSetNetworkAttachmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkAttachment",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference)SetNicType(val *string) {
	if err := j.validateSetNicTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nicType",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference)SetQueueCount(val *float64) {
	if err := j.validateSetQueueCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queueCount",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference)SetStackType(val *string) {
	if err := j.validateSetStackTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stackType",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference)SetSubnetwork(val *string) {
	if err := j.validateSetSubnetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetwork",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) PutAccessConfigs(value interface{}) {
	if err := g.validatePutAccessConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAccessConfigs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) PutAliasIpRanges(value interface{}) {
	if err := g.validatePutAliasIpRangesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAliasIpRanges",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) PutIpv6AccessConfigs(value interface{}) {
	if err := g.validatePutIpv6AccessConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIpv6AccessConfigs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) ResetAccessConfigs() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessConfigs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) ResetAliasIpRanges() {
	_jsii_.InvokeVoid(
		g,
		"resetAliasIpRanges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) ResetInternalIpv6PrefixLength() {
	_jsii_.InvokeVoid(
		g,
		"resetInternalIpv6PrefixLength",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) ResetIpAddress() {
	_jsii_.InvokeVoid(
		g,
		"resetIpAddress",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) ResetIpv6AccessConfigs() {
	_jsii_.InvokeVoid(
		g,
		"resetIpv6AccessConfigs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) ResetIpv6AccessType() {
	_jsii_.InvokeVoid(
		g,
		"resetIpv6AccessType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) ResetIpv6Address() {
	_jsii_.InvokeVoid(
		g,
		"resetIpv6Address",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) ResetNetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetNetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) ResetNetworkAttachment() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkAttachment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) ResetNicType() {
	_jsii_.InvokeVoid(
		g,
		"resetNicType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) ResetQueueCount() {
	_jsii_.InvokeVoid(
		g,
		"resetQueueCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) ResetStackType() {
	_jsii_.InvokeVoid(
		g,
		"resetStackType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) ResetSubnetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetSubnetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

