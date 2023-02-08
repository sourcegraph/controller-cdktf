package googlecontainerawscluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecontainerawscluster/internal"
)

type GoogleContainerAwsClusterControlPlaneOutputReference interface {
	cdktf.ComplexObject
	AwsServicesAuthentication() GoogleContainerAwsClusterControlPlaneAwsServicesAuthenticationOutputReference
	AwsServicesAuthenticationInput() *GoogleContainerAwsClusterControlPlaneAwsServicesAuthentication
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
	ConfigEncryption() GoogleContainerAwsClusterControlPlaneConfigEncryptionOutputReference
	ConfigEncryptionInput() *GoogleContainerAwsClusterControlPlaneConfigEncryption
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DatabaseEncryption() GoogleContainerAwsClusterControlPlaneDatabaseEncryptionOutputReference
	DatabaseEncryptionInput() *GoogleContainerAwsClusterControlPlaneDatabaseEncryption
	// Experimental.
	Fqn() *string
	IamInstanceProfile() *string
	SetIamInstanceProfile(val *string)
	IamInstanceProfileInput() *string
	InstancePlacement() GoogleContainerAwsClusterControlPlaneInstancePlacementOutputReference
	InstancePlacementInput() *GoogleContainerAwsClusterControlPlaneInstancePlacement
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	InternalValue() *GoogleContainerAwsClusterControlPlane
	SetInternalValue(val *GoogleContainerAwsClusterControlPlane)
	MainVolume() GoogleContainerAwsClusterControlPlaneMainVolumeOutputReference
	MainVolumeInput() *GoogleContainerAwsClusterControlPlaneMainVolume
	ProxyConfig() GoogleContainerAwsClusterControlPlaneProxyConfigOutputReference
	ProxyConfigInput() *GoogleContainerAwsClusterControlPlaneProxyConfig
	RootVolume() GoogleContainerAwsClusterControlPlaneRootVolumeOutputReference
	RootVolumeInput() *GoogleContainerAwsClusterControlPlaneRootVolume
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	SshConfig() GoogleContainerAwsClusterControlPlaneSshConfigOutputReference
	SshConfigInput() *GoogleContainerAwsClusterControlPlaneSshConfig
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	SubnetIdsInput() *[]*string
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
	PutAwsServicesAuthentication(value *GoogleContainerAwsClusterControlPlaneAwsServicesAuthentication)
	PutConfigEncryption(value *GoogleContainerAwsClusterControlPlaneConfigEncryption)
	PutDatabaseEncryption(value *GoogleContainerAwsClusterControlPlaneDatabaseEncryption)
	PutInstancePlacement(value *GoogleContainerAwsClusterControlPlaneInstancePlacement)
	PutMainVolume(value *GoogleContainerAwsClusterControlPlaneMainVolume)
	PutProxyConfig(value *GoogleContainerAwsClusterControlPlaneProxyConfig)
	PutRootVolume(value *GoogleContainerAwsClusterControlPlaneRootVolume)
	PutSshConfig(value *GoogleContainerAwsClusterControlPlaneSshConfig)
	ResetInstancePlacement()
	ResetInstanceType()
	ResetMainVolume()
	ResetProxyConfig()
	ResetRootVolume()
	ResetSecurityGroupIds()
	ResetSshConfig()
	ResetTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleContainerAwsClusterControlPlaneOutputReference
type jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) AwsServicesAuthentication() GoogleContainerAwsClusterControlPlaneAwsServicesAuthenticationOutputReference {
	var returns GoogleContainerAwsClusterControlPlaneAwsServicesAuthenticationOutputReference
	_jsii_.Get(
		j,
		"awsServicesAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) AwsServicesAuthenticationInput() *GoogleContainerAwsClusterControlPlaneAwsServicesAuthentication {
	var returns *GoogleContainerAwsClusterControlPlaneAwsServicesAuthentication
	_jsii_.Get(
		j,
		"awsServicesAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) ConfigEncryption() GoogleContainerAwsClusterControlPlaneConfigEncryptionOutputReference {
	var returns GoogleContainerAwsClusterControlPlaneConfigEncryptionOutputReference
	_jsii_.Get(
		j,
		"configEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) ConfigEncryptionInput() *GoogleContainerAwsClusterControlPlaneConfigEncryption {
	var returns *GoogleContainerAwsClusterControlPlaneConfigEncryption
	_jsii_.Get(
		j,
		"configEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) DatabaseEncryption() GoogleContainerAwsClusterControlPlaneDatabaseEncryptionOutputReference {
	var returns GoogleContainerAwsClusterControlPlaneDatabaseEncryptionOutputReference
	_jsii_.Get(
		j,
		"databaseEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) DatabaseEncryptionInput() *GoogleContainerAwsClusterControlPlaneDatabaseEncryption {
	var returns *GoogleContainerAwsClusterControlPlaneDatabaseEncryption
	_jsii_.Get(
		j,
		"databaseEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) IamInstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) IamInstanceProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) InstancePlacement() GoogleContainerAwsClusterControlPlaneInstancePlacementOutputReference {
	var returns GoogleContainerAwsClusterControlPlaneInstancePlacementOutputReference
	_jsii_.Get(
		j,
		"instancePlacement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) InstancePlacementInput() *GoogleContainerAwsClusterControlPlaneInstancePlacement {
	var returns *GoogleContainerAwsClusterControlPlaneInstancePlacement
	_jsii_.Get(
		j,
		"instancePlacementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) InternalValue() *GoogleContainerAwsClusterControlPlane {
	var returns *GoogleContainerAwsClusterControlPlane
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) MainVolume() GoogleContainerAwsClusterControlPlaneMainVolumeOutputReference {
	var returns GoogleContainerAwsClusterControlPlaneMainVolumeOutputReference
	_jsii_.Get(
		j,
		"mainVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) MainVolumeInput() *GoogleContainerAwsClusterControlPlaneMainVolume {
	var returns *GoogleContainerAwsClusterControlPlaneMainVolume
	_jsii_.Get(
		j,
		"mainVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) ProxyConfig() GoogleContainerAwsClusterControlPlaneProxyConfigOutputReference {
	var returns GoogleContainerAwsClusterControlPlaneProxyConfigOutputReference
	_jsii_.Get(
		j,
		"proxyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) ProxyConfigInput() *GoogleContainerAwsClusterControlPlaneProxyConfig {
	var returns *GoogleContainerAwsClusterControlPlaneProxyConfig
	_jsii_.Get(
		j,
		"proxyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) RootVolume() GoogleContainerAwsClusterControlPlaneRootVolumeOutputReference {
	var returns GoogleContainerAwsClusterControlPlaneRootVolumeOutputReference
	_jsii_.Get(
		j,
		"rootVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) RootVolumeInput() *GoogleContainerAwsClusterControlPlaneRootVolume {
	var returns *GoogleContainerAwsClusterControlPlaneRootVolume
	_jsii_.Get(
		j,
		"rootVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) SshConfig() GoogleContainerAwsClusterControlPlaneSshConfigOutputReference {
	var returns GoogleContainerAwsClusterControlPlaneSshConfigOutputReference
	_jsii_.Get(
		j,
		"sshConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) SshConfigInput() *GoogleContainerAwsClusterControlPlaneSshConfig {
	var returns *GoogleContainerAwsClusterControlPlaneSshConfig
	_jsii_.Get(
		j,
		"sshConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewGoogleContainerAwsClusterControlPlaneOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleContainerAwsClusterControlPlaneOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerAwsClusterControlPlaneOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference{}

	_jsii_.Create(
		"google-beta.googleContainerAwsCluster.GoogleContainerAwsClusterControlPlaneOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContainerAwsClusterControlPlaneOutputReference_Override(g GoogleContainerAwsClusterControlPlaneOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleContainerAwsCluster.GoogleContainerAwsClusterControlPlaneOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference)SetIamInstanceProfile(val *string) {
	if err := j.validateSetIamInstanceProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamInstanceProfile",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference)SetInternalValue(val *GoogleContainerAwsClusterControlPlane) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) PutAwsServicesAuthentication(value *GoogleContainerAwsClusterControlPlaneAwsServicesAuthentication) {
	if err := g.validatePutAwsServicesAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAwsServicesAuthentication",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) PutConfigEncryption(value *GoogleContainerAwsClusterControlPlaneConfigEncryption) {
	if err := g.validatePutConfigEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConfigEncryption",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) PutDatabaseEncryption(value *GoogleContainerAwsClusterControlPlaneDatabaseEncryption) {
	if err := g.validatePutDatabaseEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDatabaseEncryption",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) PutInstancePlacement(value *GoogleContainerAwsClusterControlPlaneInstancePlacement) {
	if err := g.validatePutInstancePlacementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInstancePlacement",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) PutMainVolume(value *GoogleContainerAwsClusterControlPlaneMainVolume) {
	if err := g.validatePutMainVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMainVolume",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) PutProxyConfig(value *GoogleContainerAwsClusterControlPlaneProxyConfig) {
	if err := g.validatePutProxyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putProxyConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) PutRootVolume(value *GoogleContainerAwsClusterControlPlaneRootVolume) {
	if err := g.validatePutRootVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRootVolume",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) PutSshConfig(value *GoogleContainerAwsClusterControlPlaneSshConfig) {
	if err := g.validatePutSshConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSshConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) ResetInstancePlacement() {
	_jsii_.InvokeVoid(
		g,
		"resetInstancePlacement",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		g,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) ResetMainVolume() {
	_jsii_.InvokeVoid(
		g,
		"resetMainVolume",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) ResetProxyConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetProxyConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) ResetRootVolume() {
	_jsii_.InvokeVoid(
		g,
		"resetRootVolume",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) ResetSshConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSshConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContainerAwsClusterControlPlaneOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

