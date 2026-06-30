package netappvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/netappvolume/internal"
)

type NetappVolumeCacheParametersOutputReference interface {
	cdktf.ComplexObject
	CacheConfig() NetappVolumeCacheParametersCacheConfigOutputReference
	CacheConfigInput() *NetappVolumeCacheParametersCacheConfig
	CacheState() *string
	Command() *string
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
	EnableGlobalFileLock() any
	SetEnableGlobalFileLock(val any)
	EnableGlobalFileLockInput() any
	// Experimental.
	Fqn() *string
	InternalValue() *NetappVolumeCacheParameters
	SetInternalValue(val *NetappVolumeCacheParameters)
	Passphrase() *string
	PeerClusterName() *string
	SetPeerClusterName(val *string)
	PeerClusterNameInput() *string
	PeeringCommandExpiryTime() *string
	SetPeeringCommandExpiryTime(val *string)
	PeeringCommandExpiryTimeInput() *string
	PeerIpAddresses() *[]*string
	SetPeerIpAddresses(val *[]*string)
	PeerIpAddressesInput() *[]*string
	PeerSvmName() *string
	SetPeerSvmName(val *string)
	PeerSvmNameInput() *string
	PeerVolumeName() *string
	SetPeerVolumeName(val *string)
	PeerVolumeNameInput() *string
	StateDetails() *string
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
	PutCacheConfig(value *NetappVolumeCacheParametersCacheConfig)
	ResetCacheConfig()
	ResetEnableGlobalFileLock()
	ResetPeerClusterName()
	ResetPeeringCommandExpiryTime()
	ResetPeerIpAddresses()
	ResetPeerSvmName()
	ResetPeerVolumeName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetappVolumeCacheParametersOutputReference
type jsiiProxy_NetappVolumeCacheParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) CacheConfig() NetappVolumeCacheParametersCacheConfigOutputReference {
	var returns NetappVolumeCacheParametersCacheConfigOutputReference
	_jsii_.Get(
		j,
		"cacheConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) CacheConfigInput() *NetappVolumeCacheParametersCacheConfig {
	var returns *NetappVolumeCacheParametersCacheConfig
	_jsii_.Get(
		j,
		"cacheConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) CacheState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) Command() *string {
	var returns *string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) EnableGlobalFileLock() any {
	var returns any
	_jsii_.Get(
		j,
		"enableGlobalFileLock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) EnableGlobalFileLockInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableGlobalFileLockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) InternalValue() *NetappVolumeCacheParameters {
	var returns *NetappVolumeCacheParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) Passphrase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passphrase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) PeerClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerClusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) PeerClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerClusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) PeeringCommandExpiryTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peeringCommandExpiryTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) PeeringCommandExpiryTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peeringCommandExpiryTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) PeerIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"peerIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) PeerIpAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"peerIpAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) PeerSvmName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerSvmName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) PeerSvmNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerSvmNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) PeerVolumeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerVolumeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) PeerVolumeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerVolumeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) StateDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewNetappVolumeCacheParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetappVolumeCacheParametersOutputReference {
	_init_.Initialize()

	if err := validateNewNetappVolumeCacheParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetappVolumeCacheParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.netappVolume.NetappVolumeCacheParametersOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetappVolumeCacheParametersOutputReference_Override(n NetappVolumeCacheParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.netappVolume.NetappVolumeCacheParametersOutputReference",
		[]any{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) SetEnableGlobalFileLock(val any) {
	if err := j.validateSetEnableGlobalFileLockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableGlobalFileLock",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) SetInternalValue(val *NetappVolumeCacheParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) SetPeerClusterName(val *string) {
	if err := j.validateSetPeerClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerClusterName",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) SetPeeringCommandExpiryTime(val *string) {
	if err := j.validateSetPeeringCommandExpiryTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peeringCommandExpiryTime",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) SetPeerIpAddresses(val *[]*string) {
	if err := j.validateSetPeerIpAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerIpAddresses",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) SetPeerSvmName(val *string) {
	if err := j.validateSetPeerSvmNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerSvmName",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) SetPeerVolumeName(val *string) {
	if err := j.validateSetPeerVolumeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerVolumeName",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeCacheParametersOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetappVolumeCacheParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeCacheParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeCacheParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeCacheParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeCacheParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeCacheParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeCacheParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeCacheParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeCacheParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeCacheParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeCacheParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeCacheParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeCacheParametersOutputReference) PutCacheConfig(value *NetappVolumeCacheParametersCacheConfig) {
	if err := n.validatePutCacheConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putCacheConfig",
		[]any{value},
	)
}

func (n *jsiiProxy_NetappVolumeCacheParametersOutputReference) ResetCacheConfig() {
	_jsii_.InvokeVoid(
		n,
		"resetCacheConfig",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeCacheParametersOutputReference) ResetEnableGlobalFileLock() {
	_jsii_.InvokeVoid(
		n,
		"resetEnableGlobalFileLock",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeCacheParametersOutputReference) ResetPeerClusterName() {
	_jsii_.InvokeVoid(
		n,
		"resetPeerClusterName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeCacheParametersOutputReference) ResetPeeringCommandExpiryTime() {
	_jsii_.InvokeVoid(
		n,
		"resetPeeringCommandExpiryTime",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeCacheParametersOutputReference) ResetPeerIpAddresses() {
	_jsii_.InvokeVoid(
		n,
		"resetPeerIpAddresses",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeCacheParametersOutputReference) ResetPeerSvmName() {
	_jsii_.InvokeVoid(
		n,
		"resetPeerSvmName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeCacheParametersOutputReference) ResetPeerVolumeName() {
	_jsii_.InvokeVoid(
		n,
		"resetPeerVolumeName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeCacheParametersOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		n,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeCacheParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
