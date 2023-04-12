package googlecontainerawsnodepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googlecontainerawsnodepool/internal"
)

type GoogleContainerAwsNodePoolConfigAOutputReference interface {
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
	ConfigEncryption() GoogleContainerAwsNodePoolConfigConfigEncryptionOutputReference
	ConfigEncryptionInput() *GoogleContainerAwsNodePoolConfigConfigEncryption
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	IamInstanceProfile() *string
	SetIamInstanceProfile(val *string)
	IamInstanceProfileInput() *string
	ImageType() *string
	SetImageType(val *string)
	ImageTypeInput() *string
	InstancePlacement() GoogleContainerAwsNodePoolConfigInstancePlacementOutputReference
	InstancePlacementInput() *GoogleContainerAwsNodePoolConfigInstancePlacement
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	InternalValue() *GoogleContainerAwsNodePoolConfigA
	SetInternalValue(val *GoogleContainerAwsNodePoolConfigA)
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	ProxyConfig() GoogleContainerAwsNodePoolConfigProxyConfigOutputReference
	ProxyConfigInput() *GoogleContainerAwsNodePoolConfigProxyConfig
	RootVolume() GoogleContainerAwsNodePoolConfigRootVolumeOutputReference
	RootVolumeInput() *GoogleContainerAwsNodePoolConfigRootVolume
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	SshConfig() GoogleContainerAwsNodePoolConfigSshConfigOutputReference
	SshConfigInput() *GoogleContainerAwsNodePoolConfigSshConfig
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	Taints() GoogleContainerAwsNodePoolConfigTaintsList
	TaintsInput() interface{}
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
	PutConfigEncryption(value *GoogleContainerAwsNodePoolConfigConfigEncryption)
	PutInstancePlacement(value *GoogleContainerAwsNodePoolConfigInstancePlacement)
	PutProxyConfig(value *GoogleContainerAwsNodePoolConfigProxyConfig)
	PutRootVolume(value *GoogleContainerAwsNodePoolConfigRootVolume)
	PutSshConfig(value *GoogleContainerAwsNodePoolConfigSshConfig)
	PutTaints(value interface{})
	ResetImageType()
	ResetInstancePlacement()
	ResetInstanceType()
	ResetLabels()
	ResetProxyConfig()
	ResetRootVolume()
	ResetSecurityGroupIds()
	ResetSshConfig()
	ResetTags()
	ResetTaints()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleContainerAwsNodePoolConfigAOutputReference
type jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) ConfigEncryption() GoogleContainerAwsNodePoolConfigConfigEncryptionOutputReference {
	var returns GoogleContainerAwsNodePoolConfigConfigEncryptionOutputReference
	_jsii_.Get(
		j,
		"configEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) ConfigEncryptionInput() *GoogleContainerAwsNodePoolConfigConfigEncryption {
	var returns *GoogleContainerAwsNodePoolConfigConfigEncryption
	_jsii_.Get(
		j,
		"configEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) IamInstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) IamInstanceProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) ImageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) ImageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) InstancePlacement() GoogleContainerAwsNodePoolConfigInstancePlacementOutputReference {
	var returns GoogleContainerAwsNodePoolConfigInstancePlacementOutputReference
	_jsii_.Get(
		j,
		"instancePlacement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) InstancePlacementInput() *GoogleContainerAwsNodePoolConfigInstancePlacement {
	var returns *GoogleContainerAwsNodePoolConfigInstancePlacement
	_jsii_.Get(
		j,
		"instancePlacementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) InternalValue() *GoogleContainerAwsNodePoolConfigA {
	var returns *GoogleContainerAwsNodePoolConfigA
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) ProxyConfig() GoogleContainerAwsNodePoolConfigProxyConfigOutputReference {
	var returns GoogleContainerAwsNodePoolConfigProxyConfigOutputReference
	_jsii_.Get(
		j,
		"proxyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) ProxyConfigInput() *GoogleContainerAwsNodePoolConfigProxyConfig {
	var returns *GoogleContainerAwsNodePoolConfigProxyConfig
	_jsii_.Get(
		j,
		"proxyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) RootVolume() GoogleContainerAwsNodePoolConfigRootVolumeOutputReference {
	var returns GoogleContainerAwsNodePoolConfigRootVolumeOutputReference
	_jsii_.Get(
		j,
		"rootVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) RootVolumeInput() *GoogleContainerAwsNodePoolConfigRootVolume {
	var returns *GoogleContainerAwsNodePoolConfigRootVolume
	_jsii_.Get(
		j,
		"rootVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) SshConfig() GoogleContainerAwsNodePoolConfigSshConfigOutputReference {
	var returns GoogleContainerAwsNodePoolConfigSshConfigOutputReference
	_jsii_.Get(
		j,
		"sshConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) SshConfigInput() *GoogleContainerAwsNodePoolConfigSshConfig {
	var returns *GoogleContainerAwsNodePoolConfigSshConfig
	_jsii_.Get(
		j,
		"sshConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) Taints() GoogleContainerAwsNodePoolConfigTaintsList {
	var returns GoogleContainerAwsNodePoolConfigTaintsList
	_jsii_.Get(
		j,
		"taints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) TaintsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleContainerAwsNodePoolConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleContainerAwsNodePoolConfigAOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerAwsNodePoolConfigAOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleContainerAwsNodePool.GoogleContainerAwsNodePoolConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContainerAwsNodePoolConfigAOutputReference_Override(g GoogleContainerAwsNodePoolConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleContainerAwsNodePool.GoogleContainerAwsNodePoolConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference)SetIamInstanceProfile(val *string) {
	if err := j.validateSetIamInstanceProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamInstanceProfile",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference)SetImageType(val *string) {
	if err := j.validateSetImageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageType",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference)SetInternalValue(val *GoogleContainerAwsNodePoolConfigA) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) PutConfigEncryption(value *GoogleContainerAwsNodePoolConfigConfigEncryption) {
	if err := g.validatePutConfigEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConfigEncryption",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) PutInstancePlacement(value *GoogleContainerAwsNodePoolConfigInstancePlacement) {
	if err := g.validatePutInstancePlacementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInstancePlacement",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) PutProxyConfig(value *GoogleContainerAwsNodePoolConfigProxyConfig) {
	if err := g.validatePutProxyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putProxyConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) PutRootVolume(value *GoogleContainerAwsNodePoolConfigRootVolume) {
	if err := g.validatePutRootVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRootVolume",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) PutSshConfig(value *GoogleContainerAwsNodePoolConfigSshConfig) {
	if err := g.validatePutSshConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSshConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) PutTaints(value interface{}) {
	if err := g.validatePutTaintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTaints",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) ResetImageType() {
	_jsii_.InvokeVoid(
		g,
		"resetImageType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) ResetInstancePlacement() {
	_jsii_.InvokeVoid(
		g,
		"resetInstancePlacement",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		g,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) ResetProxyConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetProxyConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) ResetRootVolume() {
	_jsii_.InvokeVoid(
		g,
		"resetRootVolume",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) ResetSshConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSshConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) ResetTaints() {
	_jsii_.InvokeVoid(
		g,
		"resetTaints",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContainerAwsNodePoolConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

