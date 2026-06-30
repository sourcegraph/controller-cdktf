package gkeonprembaremetaladmincluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/gkeonprembaremetaladmincluster/internal"
)

type GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference interface {
	cdktf.ComplexObject
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
	// Experimental.
	Fqn() *string
	InternalValue() *GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfig
	SetInternalValue(val *GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfig)
	KubeletConfig() GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigKubeletConfigOutputReference
	KubeletConfigInput() *GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigKubeletConfig
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	NodeConfigs() GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigNodeConfigsList
	NodeConfigsInput() any
	OperatingSystem() *string
	SetOperatingSystem(val *string)
	OperatingSystemInput() *string
	Taints() GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigTaintsList
	TaintsInput() any
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
	PutKubeletConfig(value *GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigKubeletConfig)
	PutNodeConfigs(value any)
	PutTaints(value any)
	ResetKubeletConfig()
	ResetLabels()
	ResetNodeConfigs()
	ResetOperatingSystem()
	ResetTaints()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference
type jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) InternalValue() *GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfig {
	var returns *GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) KubeletConfig() GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigKubeletConfigOutputReference {
	var returns GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigKubeletConfigOutputReference
	_jsii_.Get(
		j,
		"kubeletConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) KubeletConfigInput() *GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigKubeletConfig {
	var returns *GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigKubeletConfig
	_jsii_.Get(
		j,
		"kubeletConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) NodeConfigs() GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigNodeConfigsList {
	var returns GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigNodeConfigsList
	_jsii_.Get(
		j,
		"nodeConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) NodeConfigsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"nodeConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) OperatingSystem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatingSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) OperatingSystemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatingSystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) Taints() GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigTaintsList {
	var returns GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigTaintsList
	_jsii_.Get(
		j,
		"taints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) TaintsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"taintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremBareMetalAdminCluster.GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference_Override(g GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremBareMetalAdminCluster.GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) SetInternalValue(val *GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) SetOperatingSystem(val *string) {
	if err := j.validateSetOperatingSystemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operatingSystem",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) PutKubeletConfig(value *GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigKubeletConfig) {
	if err := g.validatePutKubeletConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putKubeletConfig",
		[]any{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) PutNodeConfigs(value any) {
	if err := g.validatePutNodeConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNodeConfigs",
		[]any{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) PutTaints(value any) {
	if err := g.validatePutTaintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTaints",
		[]any{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) ResetKubeletConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetKubeletConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) ResetNodeConfigs() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeConfigs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) ResetOperatingSystem() {
	_jsii_.InvokeVoid(
		g,
		"resetOperatingSystem",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) ResetTaints() {
	_jsii_.InvokeVoid(
		g,
		"resetTaints",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		g,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
