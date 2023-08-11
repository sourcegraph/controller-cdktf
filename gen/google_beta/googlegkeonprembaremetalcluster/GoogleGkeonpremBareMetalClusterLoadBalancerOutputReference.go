package googlegkeonprembaremetalcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlegkeonprembaremetalcluster/internal"
)

type GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference interface {
	cdktf.ComplexObject
	BgpLbConfig() GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference
	BgpLbConfigInput() *GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfig
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
	InternalValue() *GoogleGkeonpremBareMetalClusterLoadBalancer
	SetInternalValue(val *GoogleGkeonpremBareMetalClusterLoadBalancer)
	ManualLbConfig() GoogleGkeonpremBareMetalClusterLoadBalancerManualLbConfigOutputReference
	ManualLbConfigInput() *GoogleGkeonpremBareMetalClusterLoadBalancerManualLbConfig
	MetalLbConfig() GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigOutputReference
	MetalLbConfigInput() *GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfig
	PortConfig() GoogleGkeonpremBareMetalClusterLoadBalancerPortConfigOutputReference
	PortConfigInput() *GoogleGkeonpremBareMetalClusterLoadBalancerPortConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VipConfig() GoogleGkeonpremBareMetalClusterLoadBalancerVipConfigOutputReference
	VipConfigInput() *GoogleGkeonpremBareMetalClusterLoadBalancerVipConfig
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
	PutBgpLbConfig(value *GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfig)
	PutManualLbConfig(value *GoogleGkeonpremBareMetalClusterLoadBalancerManualLbConfig)
	PutMetalLbConfig(value *GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfig)
	PutPortConfig(value *GoogleGkeonpremBareMetalClusterLoadBalancerPortConfig)
	PutVipConfig(value *GoogleGkeonpremBareMetalClusterLoadBalancerVipConfig)
	ResetBgpLbConfig()
	ResetManualLbConfig()
	ResetMetalLbConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference
type jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) BgpLbConfig() GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference {
	var returns GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference
	_jsii_.Get(
		j,
		"bgpLbConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) BgpLbConfigInput() *GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfig {
	var returns *GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfig
	_jsii_.Get(
		j,
		"bgpLbConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) InternalValue() *GoogleGkeonpremBareMetalClusterLoadBalancer {
	var returns *GoogleGkeonpremBareMetalClusterLoadBalancer
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) ManualLbConfig() GoogleGkeonpremBareMetalClusterLoadBalancerManualLbConfigOutputReference {
	var returns GoogleGkeonpremBareMetalClusterLoadBalancerManualLbConfigOutputReference
	_jsii_.Get(
		j,
		"manualLbConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) ManualLbConfigInput() *GoogleGkeonpremBareMetalClusterLoadBalancerManualLbConfig {
	var returns *GoogleGkeonpremBareMetalClusterLoadBalancerManualLbConfig
	_jsii_.Get(
		j,
		"manualLbConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) MetalLbConfig() GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigOutputReference {
	var returns GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigOutputReference
	_jsii_.Get(
		j,
		"metalLbConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) MetalLbConfigInput() *GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfig {
	var returns *GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfig
	_jsii_.Get(
		j,
		"metalLbConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) PortConfig() GoogleGkeonpremBareMetalClusterLoadBalancerPortConfigOutputReference {
	var returns GoogleGkeonpremBareMetalClusterLoadBalancerPortConfigOutputReference
	_jsii_.Get(
		j,
		"portConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) PortConfigInput() *GoogleGkeonpremBareMetalClusterLoadBalancerPortConfig {
	var returns *GoogleGkeonpremBareMetalClusterLoadBalancerPortConfig
	_jsii_.Get(
		j,
		"portConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) VipConfig() GoogleGkeonpremBareMetalClusterLoadBalancerVipConfigOutputReference {
	var returns GoogleGkeonpremBareMetalClusterLoadBalancerVipConfigOutputReference
	_jsii_.Get(
		j,
		"vipConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) VipConfigInput() *GoogleGkeonpremBareMetalClusterLoadBalancerVipConfig {
	var returns *GoogleGkeonpremBareMetalClusterLoadBalancerVipConfig
	_jsii_.Get(
		j,
		"vipConfigInput",
		&returns,
	)
	return returns
}


func NewGoogleGkeonpremBareMetalClusterLoadBalancerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleGkeonpremBareMetalClusterLoadBalancerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleGkeonpremBareMetalCluster.GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleGkeonpremBareMetalClusterLoadBalancerOutputReference_Override(g GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleGkeonpremBareMetalCluster.GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference)SetInternalValue(val *GoogleGkeonpremBareMetalClusterLoadBalancer) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) PutBgpLbConfig(value *GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfig) {
	if err := g.validatePutBgpLbConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBgpLbConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) PutManualLbConfig(value *GoogleGkeonpremBareMetalClusterLoadBalancerManualLbConfig) {
	if err := g.validatePutManualLbConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putManualLbConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) PutMetalLbConfig(value *GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfig) {
	if err := g.validatePutMetalLbConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMetalLbConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) PutPortConfig(value *GoogleGkeonpremBareMetalClusterLoadBalancerPortConfig) {
	if err := g.validatePutPortConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPortConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) PutVipConfig(value *GoogleGkeonpremBareMetalClusterLoadBalancerVipConfig) {
	if err := g.validatePutVipConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVipConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) ResetBgpLbConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetBgpLbConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) ResetManualLbConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetManualLbConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) ResetMetalLbConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetMetalLbConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

