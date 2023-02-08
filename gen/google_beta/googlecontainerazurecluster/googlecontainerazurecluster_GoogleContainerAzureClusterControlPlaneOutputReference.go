package googlecontainerazurecluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecontainerazurecluster/internal"
)

type GoogleContainerAzureClusterControlPlaneOutputReference interface {
	cdktf.ComplexObject
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
	DatabaseEncryption() GoogleContainerAzureClusterControlPlaneDatabaseEncryptionOutputReference
	DatabaseEncryptionInput() *GoogleContainerAzureClusterControlPlaneDatabaseEncryption
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleContainerAzureClusterControlPlane
	SetInternalValue(val *GoogleContainerAzureClusterControlPlane)
	MainVolume() GoogleContainerAzureClusterControlPlaneMainVolumeOutputReference
	MainVolumeInput() *GoogleContainerAzureClusterControlPlaneMainVolume
	ProxyConfig() GoogleContainerAzureClusterControlPlaneProxyConfigOutputReference
	ProxyConfigInput() *GoogleContainerAzureClusterControlPlaneProxyConfig
	ReplicaPlacements() GoogleContainerAzureClusterControlPlaneReplicaPlacementsList
	ReplicaPlacementsInput() interface{}
	RootVolume() GoogleContainerAzureClusterControlPlaneRootVolumeOutputReference
	RootVolumeInput() *GoogleContainerAzureClusterControlPlaneRootVolume
	SshConfig() GoogleContainerAzureClusterControlPlaneSshConfigOutputReference
	SshConfigInput() *GoogleContainerAzureClusterControlPlaneSshConfig
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
	VmSize() *string
	SetVmSize(val *string)
	VmSizeInput() *string
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
	PutDatabaseEncryption(value *GoogleContainerAzureClusterControlPlaneDatabaseEncryption)
	PutMainVolume(value *GoogleContainerAzureClusterControlPlaneMainVolume)
	PutProxyConfig(value *GoogleContainerAzureClusterControlPlaneProxyConfig)
	PutReplicaPlacements(value interface{})
	PutRootVolume(value *GoogleContainerAzureClusterControlPlaneRootVolume)
	PutSshConfig(value *GoogleContainerAzureClusterControlPlaneSshConfig)
	ResetDatabaseEncryption()
	ResetMainVolume()
	ResetProxyConfig()
	ResetReplicaPlacements()
	ResetRootVolume()
	ResetTags()
	ResetVmSize()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleContainerAzureClusterControlPlaneOutputReference
type jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) DatabaseEncryption() GoogleContainerAzureClusterControlPlaneDatabaseEncryptionOutputReference {
	var returns GoogleContainerAzureClusterControlPlaneDatabaseEncryptionOutputReference
	_jsii_.Get(
		j,
		"databaseEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) DatabaseEncryptionInput() *GoogleContainerAzureClusterControlPlaneDatabaseEncryption {
	var returns *GoogleContainerAzureClusterControlPlaneDatabaseEncryption
	_jsii_.Get(
		j,
		"databaseEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) InternalValue() *GoogleContainerAzureClusterControlPlane {
	var returns *GoogleContainerAzureClusterControlPlane
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) MainVolume() GoogleContainerAzureClusterControlPlaneMainVolumeOutputReference {
	var returns GoogleContainerAzureClusterControlPlaneMainVolumeOutputReference
	_jsii_.Get(
		j,
		"mainVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) MainVolumeInput() *GoogleContainerAzureClusterControlPlaneMainVolume {
	var returns *GoogleContainerAzureClusterControlPlaneMainVolume
	_jsii_.Get(
		j,
		"mainVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) ProxyConfig() GoogleContainerAzureClusterControlPlaneProxyConfigOutputReference {
	var returns GoogleContainerAzureClusterControlPlaneProxyConfigOutputReference
	_jsii_.Get(
		j,
		"proxyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) ProxyConfigInput() *GoogleContainerAzureClusterControlPlaneProxyConfig {
	var returns *GoogleContainerAzureClusterControlPlaneProxyConfig
	_jsii_.Get(
		j,
		"proxyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) ReplicaPlacements() GoogleContainerAzureClusterControlPlaneReplicaPlacementsList {
	var returns GoogleContainerAzureClusterControlPlaneReplicaPlacementsList
	_jsii_.Get(
		j,
		"replicaPlacements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) ReplicaPlacementsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replicaPlacementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) RootVolume() GoogleContainerAzureClusterControlPlaneRootVolumeOutputReference {
	var returns GoogleContainerAzureClusterControlPlaneRootVolumeOutputReference
	_jsii_.Get(
		j,
		"rootVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) RootVolumeInput() *GoogleContainerAzureClusterControlPlaneRootVolume {
	var returns *GoogleContainerAzureClusterControlPlaneRootVolume
	_jsii_.Get(
		j,
		"rootVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) SshConfig() GoogleContainerAzureClusterControlPlaneSshConfigOutputReference {
	var returns GoogleContainerAzureClusterControlPlaneSshConfigOutputReference
	_jsii_.Get(
		j,
		"sshConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) SshConfigInput() *GoogleContainerAzureClusterControlPlaneSshConfig {
	var returns *GoogleContainerAzureClusterControlPlaneSshConfig
	_jsii_.Get(
		j,
		"sshConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) VmSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) VmSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSizeInput",
		&returns,
	)
	return returns
}


func NewGoogleContainerAzureClusterControlPlaneOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleContainerAzureClusterControlPlaneOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerAzureClusterControlPlaneOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference{}

	_jsii_.Create(
		"google-beta.googleContainerAzureCluster.GoogleContainerAzureClusterControlPlaneOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContainerAzureClusterControlPlaneOutputReference_Override(g GoogleContainerAzureClusterControlPlaneOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleContainerAzureCluster.GoogleContainerAzureClusterControlPlaneOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference)SetInternalValue(val *GoogleContainerAzureClusterControlPlane) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference)SetVmSize(val *string) {
	if err := j.validateSetVmSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmSize",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) PutDatabaseEncryption(value *GoogleContainerAzureClusterControlPlaneDatabaseEncryption) {
	if err := g.validatePutDatabaseEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDatabaseEncryption",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) PutMainVolume(value *GoogleContainerAzureClusterControlPlaneMainVolume) {
	if err := g.validatePutMainVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMainVolume",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) PutProxyConfig(value *GoogleContainerAzureClusterControlPlaneProxyConfig) {
	if err := g.validatePutProxyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putProxyConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) PutReplicaPlacements(value interface{}) {
	if err := g.validatePutReplicaPlacementsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putReplicaPlacements",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) PutRootVolume(value *GoogleContainerAzureClusterControlPlaneRootVolume) {
	if err := g.validatePutRootVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRootVolume",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) PutSshConfig(value *GoogleContainerAzureClusterControlPlaneSshConfig) {
	if err := g.validatePutSshConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSshConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) ResetDatabaseEncryption() {
	_jsii_.InvokeVoid(
		g,
		"resetDatabaseEncryption",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) ResetMainVolume() {
	_jsii_.InvokeVoid(
		g,
		"resetMainVolume",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) ResetProxyConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetProxyConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) ResetReplicaPlacements() {
	_jsii_.InvokeVoid(
		g,
		"resetReplicaPlacements",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) ResetRootVolume() {
	_jsii_.InvokeVoid(
		g,
		"resetRootVolume",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) ResetVmSize() {
	_jsii_.InvokeVoid(
		g,
		"resetVmSize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContainerAzureClusterControlPlaneOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

