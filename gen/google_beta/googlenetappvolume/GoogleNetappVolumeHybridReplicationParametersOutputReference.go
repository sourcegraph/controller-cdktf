package googlenetappvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlenetappvolume/internal"
)

type GoogleNetappVolumeHybridReplicationParametersOutputReference interface {
	cdktf.ComplexObject
	ClusterLocation() *string
	SetClusterLocation(val *string)
	ClusterLocationInput() *string
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	HybridReplicationType() *string
	SetHybridReplicationType(val *string)
	HybridReplicationTypeInput() *string
	InternalValue() *GoogleNetappVolumeHybridReplicationParameters
	SetInternalValue(val *GoogleNetappVolumeHybridReplicationParameters)
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	LargeVolumeConstituentCount() *float64
	SetLargeVolumeConstituentCount(val *float64)
	LargeVolumeConstituentCountInput() *float64
	PeerClusterName() *string
	SetPeerClusterName(val *string)
	PeerClusterNameInput() *string
	PeerIpAddresses() *[]*string
	SetPeerIpAddresses(val *[]*string)
	PeerIpAddressesInput() *[]*string
	PeerSvmName() *string
	SetPeerSvmName(val *string)
	PeerSvmNameInput() *string
	PeerVolumeName() *string
	SetPeerVolumeName(val *string)
	PeerVolumeNameInput() *string
	Replication() *string
	SetReplication(val *string)
	ReplicationInput() *string
	ReplicationSchedule() *string
	SetReplicationSchedule(val *string)
	ReplicationScheduleInput() *string
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
	ResetClusterLocation()
	ResetDescription()
	ResetHybridReplicationType()
	ResetLabels()
	ResetLargeVolumeConstituentCount()
	ResetPeerClusterName()
	ResetPeerIpAddresses()
	ResetPeerSvmName()
	ResetPeerVolumeName()
	ResetReplication()
	ResetReplicationSchedule()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNetappVolumeHybridReplicationParametersOutputReference
type jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) ClusterLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) ClusterLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) HybridReplicationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hybridReplicationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) HybridReplicationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hybridReplicationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) InternalValue() *GoogleNetappVolumeHybridReplicationParameters {
	var returns *GoogleNetappVolumeHybridReplicationParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) LargeVolumeConstituentCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"largeVolumeConstituentCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) LargeVolumeConstituentCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"largeVolumeConstituentCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) PeerClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerClusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) PeerClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerClusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) PeerIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"peerIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) PeerIpAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"peerIpAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) PeerSvmName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerSvmName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) PeerSvmNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerSvmNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) PeerVolumeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerVolumeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) PeerVolumeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerVolumeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) Replication() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) ReplicationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) ReplicationSchedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) ReplicationScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleNetappVolumeHybridReplicationParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleNetappVolumeHybridReplicationParametersOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetappVolumeHybridReplicationParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetappVolume.GoogleNetappVolumeHybridReplicationParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleNetappVolumeHybridReplicationParametersOutputReference_Override(g GoogleNetappVolumeHybridReplicationParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetappVolume.GoogleNetappVolumeHybridReplicationParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference)SetClusterLocation(val *string) {
	if err := j.validateSetClusterLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterLocation",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference)SetHybridReplicationType(val *string) {
	if err := j.validateSetHybridReplicationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hybridReplicationType",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference)SetInternalValue(val *GoogleNetappVolumeHybridReplicationParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference)SetLargeVolumeConstituentCount(val *float64) {
	if err := j.validateSetLargeVolumeConstituentCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"largeVolumeConstituentCount",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference)SetPeerClusterName(val *string) {
	if err := j.validateSetPeerClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerClusterName",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference)SetPeerIpAddresses(val *[]*string) {
	if err := j.validateSetPeerIpAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerIpAddresses",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference)SetPeerSvmName(val *string) {
	if err := j.validateSetPeerSvmNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerSvmName",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference)SetPeerVolumeName(val *string) {
	if err := j.validateSetPeerVolumeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerVolumeName",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference)SetReplication(val *string) {
	if err := j.validateSetReplicationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replication",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference)SetReplicationSchedule(val *string) {
	if err := j.validateSetReplicationScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationSchedule",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) ResetClusterLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetClusterLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) ResetHybridReplicationType() {
	_jsii_.InvokeVoid(
		g,
		"resetHybridReplicationType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) ResetLargeVolumeConstituentCount() {
	_jsii_.InvokeVoid(
		g,
		"resetLargeVolumeConstituentCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) ResetPeerClusterName() {
	_jsii_.InvokeVoid(
		g,
		"resetPeerClusterName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) ResetPeerIpAddresses() {
	_jsii_.InvokeVoid(
		g,
		"resetPeerIpAddresses",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) ResetPeerSvmName() {
	_jsii_.InvokeVoid(
		g,
		"resetPeerSvmName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) ResetPeerVolumeName() {
	_jsii_.InvokeVoid(
		g,
		"resetPeerVolumeName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) ResetReplication() {
	_jsii_.InvokeVoid(
		g,
		"resetReplication",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) ResetReplicationSchedule() {
	_jsii_.InvokeVoid(
		g,
		"resetReplicationSchedule",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleNetappVolumeHybridReplicationParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

