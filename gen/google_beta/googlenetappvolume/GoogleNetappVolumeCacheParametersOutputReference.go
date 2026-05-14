package googlenetappvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlenetappvolume/internal"
)

type GoogleNetappVolumeCacheParametersOutputReference interface {
	cdktf.ComplexObject
	CacheConfig() GoogleNetappVolumeCacheParametersCacheConfigOutputReference
	CacheConfigInput() *GoogleNetappVolumeCacheParametersCacheConfig
	CacheState() *string
	Command() *string
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
	EnableGlobalFileLock() interface{}
	SetEnableGlobalFileLock(val interface{})
	EnableGlobalFileLockInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleNetappVolumeCacheParameters
	SetInternalValue(val *GoogleNetappVolumeCacheParameters)
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
	PutCacheConfig(value *GoogleNetappVolumeCacheParametersCacheConfig)
	ResetCacheConfig()
	ResetEnableGlobalFileLock()
	ResetPeerClusterName()
	ResetPeeringCommandExpiryTime()
	ResetPeerIpAddresses()
	ResetPeerSvmName()
	ResetPeerVolumeName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNetappVolumeCacheParametersOutputReference
type jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) CacheConfig() GoogleNetappVolumeCacheParametersCacheConfigOutputReference {
	var returns GoogleNetappVolumeCacheParametersCacheConfigOutputReference
	_jsii_.Get(
		j,
		"cacheConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) CacheConfigInput() *GoogleNetappVolumeCacheParametersCacheConfig {
	var returns *GoogleNetappVolumeCacheParametersCacheConfig
	_jsii_.Get(
		j,
		"cacheConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) CacheState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) Command() *string {
	var returns *string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) EnableGlobalFileLock() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableGlobalFileLock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) EnableGlobalFileLockInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableGlobalFileLockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) InternalValue() *GoogleNetappVolumeCacheParameters {
	var returns *GoogleNetappVolumeCacheParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) Passphrase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passphrase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) PeerClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerClusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) PeerClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerClusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) PeeringCommandExpiryTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peeringCommandExpiryTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) PeeringCommandExpiryTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peeringCommandExpiryTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) PeerIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"peerIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) PeerIpAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"peerIpAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) PeerSvmName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerSvmName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) PeerSvmNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerSvmNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) PeerVolumeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerVolumeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) PeerVolumeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerVolumeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) StateDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleNetappVolumeCacheParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleNetappVolumeCacheParametersOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetappVolumeCacheParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetappVolume.GoogleNetappVolumeCacheParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleNetappVolumeCacheParametersOutputReference_Override(g GoogleNetappVolumeCacheParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetappVolume.GoogleNetappVolumeCacheParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference)SetEnableGlobalFileLock(val interface{}) {
	if err := j.validateSetEnableGlobalFileLockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableGlobalFileLock",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference)SetInternalValue(val *GoogleNetappVolumeCacheParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference)SetPeerClusterName(val *string) {
	if err := j.validateSetPeerClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerClusterName",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference)SetPeeringCommandExpiryTime(val *string) {
	if err := j.validateSetPeeringCommandExpiryTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peeringCommandExpiryTime",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference)SetPeerIpAddresses(val *[]*string) {
	if err := j.validateSetPeerIpAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerIpAddresses",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference)SetPeerSvmName(val *string) {
	if err := j.validateSetPeerSvmNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerSvmName",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference)SetPeerVolumeName(val *string) {
	if err := j.validateSetPeerVolumeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerVolumeName",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) PutCacheConfig(value *GoogleNetappVolumeCacheParametersCacheConfig) {
	if err := g.validatePutCacheConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCacheConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) ResetCacheConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetCacheConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) ResetEnableGlobalFileLock() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableGlobalFileLock",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) ResetPeerClusterName() {
	_jsii_.InvokeVoid(
		g,
		"resetPeerClusterName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) ResetPeeringCommandExpiryTime() {
	_jsii_.InvokeVoid(
		g,
		"resetPeeringCommandExpiryTime",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) ResetPeerIpAddresses() {
	_jsii_.InvokeVoid(
		g,
		"resetPeerIpAddresses",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) ResetPeerSvmName() {
	_jsii_.InvokeVoid(
		g,
		"resetPeerSvmName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) ResetPeerVolumeName() {
	_jsii_.InvokeVoid(
		g,
		"resetPeerVolumeName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleNetappVolumeCacheParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

