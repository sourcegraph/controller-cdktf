package googlecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googlecontainercluster/internal"
)

type GoogleContainerClusterPrivateClusterConfigOutputReference interface {
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
	EnablePrivateEndpoint() interface{}
	SetEnablePrivateEndpoint(val interface{})
	EnablePrivateEndpointInput() interface{}
	EnablePrivateNodes() interface{}
	SetEnablePrivateNodes(val interface{})
	EnablePrivateNodesInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleContainerClusterPrivateClusterConfig
	SetInternalValue(val *GoogleContainerClusterPrivateClusterConfig)
	MasterGlobalAccessConfig() GoogleContainerClusterPrivateClusterConfigMasterGlobalAccessConfigOutputReference
	MasterGlobalAccessConfigInput() *GoogleContainerClusterPrivateClusterConfigMasterGlobalAccessConfig
	MasterIpv4CidrBlock() *string
	SetMasterIpv4CidrBlock(val *string)
	MasterIpv4CidrBlockInput() *string
	PeeringName() *string
	PrivateEndpoint() *string
	PublicEndpoint() *string
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
	PutMasterGlobalAccessConfig(value *GoogleContainerClusterPrivateClusterConfigMasterGlobalAccessConfig)
	ResetEnablePrivateNodes()
	ResetMasterGlobalAccessConfig()
	ResetMasterIpv4CidrBlock()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleContainerClusterPrivateClusterConfigOutputReference
type jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) EnablePrivateEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivateEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) EnablePrivateEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivateEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) EnablePrivateNodes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivateNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) EnablePrivateNodesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivateNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) InternalValue() *GoogleContainerClusterPrivateClusterConfig {
	var returns *GoogleContainerClusterPrivateClusterConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) MasterGlobalAccessConfig() GoogleContainerClusterPrivateClusterConfigMasterGlobalAccessConfigOutputReference {
	var returns GoogleContainerClusterPrivateClusterConfigMasterGlobalAccessConfigOutputReference
	_jsii_.Get(
		j,
		"masterGlobalAccessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) MasterGlobalAccessConfigInput() *GoogleContainerClusterPrivateClusterConfigMasterGlobalAccessConfig {
	var returns *GoogleContainerClusterPrivateClusterConfigMasterGlobalAccessConfig
	_jsii_.Get(
		j,
		"masterGlobalAccessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) MasterIpv4CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterIpv4CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) MasterIpv4CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterIpv4CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) PeeringName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peeringName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) PrivateEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) PublicEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleContainerClusterPrivateClusterConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleContainerClusterPrivateClusterConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerClusterPrivateClusterConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleContainerCluster.GoogleContainerClusterPrivateClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContainerClusterPrivateClusterConfigOutputReference_Override(g GoogleContainerClusterPrivateClusterConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleContainerCluster.GoogleContainerClusterPrivateClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference)SetEnablePrivateEndpoint(val interface{}) {
	if err := j.validateSetEnablePrivateEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePrivateEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference)SetEnablePrivateNodes(val interface{}) {
	if err := j.validateSetEnablePrivateNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePrivateNodes",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference)SetInternalValue(val *GoogleContainerClusterPrivateClusterConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference)SetMasterIpv4CidrBlock(val *string) {
	if err := j.validateSetMasterIpv4CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterIpv4CidrBlock",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) PutMasterGlobalAccessConfig(value *GoogleContainerClusterPrivateClusterConfigMasterGlobalAccessConfig) {
	if err := g.validatePutMasterGlobalAccessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMasterGlobalAccessConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) ResetEnablePrivateNodes() {
	_jsii_.InvokeVoid(
		g,
		"resetEnablePrivateNodes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) ResetMasterGlobalAccessConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetMasterGlobalAccessConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) ResetMasterIpv4CidrBlock() {
	_jsii_.InvokeVoid(
		g,
		"resetMasterIpv4CidrBlock",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterPrivateClusterConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

