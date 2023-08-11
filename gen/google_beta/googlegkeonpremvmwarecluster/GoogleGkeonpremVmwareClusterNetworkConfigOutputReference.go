package googlegkeonpremvmwarecluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlegkeonpremvmwarecluster/internal"
)

type GoogleGkeonpremVmwareClusterNetworkConfigOutputReference interface {
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
	ControlPlaneV2Config() GoogleGkeonpremVmwareClusterNetworkConfigControlPlaneV2ConfigOutputReference
	ControlPlaneV2ConfigInput() *GoogleGkeonpremVmwareClusterNetworkConfigControlPlaneV2Config
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DhcpIpConfig() GoogleGkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference
	DhcpIpConfigInput() *GoogleGkeonpremVmwareClusterNetworkConfigDhcpIpConfig
	// Experimental.
	Fqn() *string
	HostConfig() GoogleGkeonpremVmwareClusterNetworkConfigHostConfigOutputReference
	HostConfigInput() *GoogleGkeonpremVmwareClusterNetworkConfigHostConfig
	InternalValue() *GoogleGkeonpremVmwareClusterNetworkConfig
	SetInternalValue(val *GoogleGkeonpremVmwareClusterNetworkConfig)
	PodAddressCidrBlocks() *[]*string
	SetPodAddressCidrBlocks(val *[]*string)
	PodAddressCidrBlocksInput() *[]*string
	ServiceAddressCidrBlocks() *[]*string
	SetServiceAddressCidrBlocks(val *[]*string)
	ServiceAddressCidrBlocksInput() *[]*string
	StaticIpConfig() GoogleGkeonpremVmwareClusterNetworkConfigStaticIpConfigOutputReference
	StaticIpConfigInput() *GoogleGkeonpremVmwareClusterNetworkConfigStaticIpConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VcenterNetwork() *string
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
	PutControlPlaneV2Config(value *GoogleGkeonpremVmwareClusterNetworkConfigControlPlaneV2Config)
	PutDhcpIpConfig(value *GoogleGkeonpremVmwareClusterNetworkConfigDhcpIpConfig)
	PutHostConfig(value *GoogleGkeonpremVmwareClusterNetworkConfigHostConfig)
	PutStaticIpConfig(value *GoogleGkeonpremVmwareClusterNetworkConfigStaticIpConfig)
	ResetControlPlaneV2Config()
	ResetDhcpIpConfig()
	ResetHostConfig()
	ResetStaticIpConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleGkeonpremVmwareClusterNetworkConfigOutputReference
type jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) ControlPlaneV2Config() GoogleGkeonpremVmwareClusterNetworkConfigControlPlaneV2ConfigOutputReference {
	var returns GoogleGkeonpremVmwareClusterNetworkConfigControlPlaneV2ConfigOutputReference
	_jsii_.Get(
		j,
		"controlPlaneV2Config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) ControlPlaneV2ConfigInput() *GoogleGkeonpremVmwareClusterNetworkConfigControlPlaneV2Config {
	var returns *GoogleGkeonpremVmwareClusterNetworkConfigControlPlaneV2Config
	_jsii_.Get(
		j,
		"controlPlaneV2ConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) DhcpIpConfig() GoogleGkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference {
	var returns GoogleGkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference
	_jsii_.Get(
		j,
		"dhcpIpConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) DhcpIpConfigInput() *GoogleGkeonpremVmwareClusterNetworkConfigDhcpIpConfig {
	var returns *GoogleGkeonpremVmwareClusterNetworkConfigDhcpIpConfig
	_jsii_.Get(
		j,
		"dhcpIpConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) HostConfig() GoogleGkeonpremVmwareClusterNetworkConfigHostConfigOutputReference {
	var returns GoogleGkeonpremVmwareClusterNetworkConfigHostConfigOutputReference
	_jsii_.Get(
		j,
		"hostConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) HostConfigInput() *GoogleGkeonpremVmwareClusterNetworkConfigHostConfig {
	var returns *GoogleGkeonpremVmwareClusterNetworkConfigHostConfig
	_jsii_.Get(
		j,
		"hostConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) InternalValue() *GoogleGkeonpremVmwareClusterNetworkConfig {
	var returns *GoogleGkeonpremVmwareClusterNetworkConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) PodAddressCidrBlocks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"podAddressCidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) PodAddressCidrBlocksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"podAddressCidrBlocksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) ServiceAddressCidrBlocks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceAddressCidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) ServiceAddressCidrBlocksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceAddressCidrBlocksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) StaticIpConfig() GoogleGkeonpremVmwareClusterNetworkConfigStaticIpConfigOutputReference {
	var returns GoogleGkeonpremVmwareClusterNetworkConfigStaticIpConfigOutputReference
	_jsii_.Get(
		j,
		"staticIpConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) StaticIpConfigInput() *GoogleGkeonpremVmwareClusterNetworkConfigStaticIpConfig {
	var returns *GoogleGkeonpremVmwareClusterNetworkConfigStaticIpConfig
	_jsii_.Get(
		j,
		"staticIpConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) VcenterNetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vcenterNetwork",
		&returns,
	)
	return returns
}


func NewGoogleGkeonpremVmwareClusterNetworkConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleGkeonpremVmwareClusterNetworkConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleGkeonpremVmwareClusterNetworkConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleGkeonpremVmwareCluster.GoogleGkeonpremVmwareClusterNetworkConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleGkeonpremVmwareClusterNetworkConfigOutputReference_Override(g GoogleGkeonpremVmwareClusterNetworkConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleGkeonpremVmwareCluster.GoogleGkeonpremVmwareClusterNetworkConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference)SetInternalValue(val *GoogleGkeonpremVmwareClusterNetworkConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference)SetPodAddressCidrBlocks(val *[]*string) {
	if err := j.validateSetPodAddressCidrBlocksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"podAddressCidrBlocks",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference)SetServiceAddressCidrBlocks(val *[]*string) {
	if err := j.validateSetServiceAddressCidrBlocksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAddressCidrBlocks",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) PutControlPlaneV2Config(value *GoogleGkeonpremVmwareClusterNetworkConfigControlPlaneV2Config) {
	if err := g.validatePutControlPlaneV2ConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putControlPlaneV2Config",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) PutDhcpIpConfig(value *GoogleGkeonpremVmwareClusterNetworkConfigDhcpIpConfig) {
	if err := g.validatePutDhcpIpConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDhcpIpConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) PutHostConfig(value *GoogleGkeonpremVmwareClusterNetworkConfigHostConfig) {
	if err := g.validatePutHostConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHostConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) PutStaticIpConfig(value *GoogleGkeonpremVmwareClusterNetworkConfigStaticIpConfig) {
	if err := g.validatePutStaticIpConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStaticIpConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) ResetControlPlaneV2Config() {
	_jsii_.InvokeVoid(
		g,
		"resetControlPlaneV2Config",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) ResetDhcpIpConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetDhcpIpConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) ResetHostConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetHostConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) ResetStaticIpConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetStaticIpConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterNetworkConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

