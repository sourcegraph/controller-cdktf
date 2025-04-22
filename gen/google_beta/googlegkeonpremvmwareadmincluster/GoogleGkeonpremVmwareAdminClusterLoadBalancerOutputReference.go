package googlegkeonpremvmwareadmincluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlegkeonpremvmwareadmincluster/internal"
)

type GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference interface {
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
	F5Config() GoogleGkeonpremVmwareAdminClusterLoadBalancerF5ConfigOutputReference
	F5ConfigInput() *GoogleGkeonpremVmwareAdminClusterLoadBalancerF5Config
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleGkeonpremVmwareAdminClusterLoadBalancer
	SetInternalValue(val *GoogleGkeonpremVmwareAdminClusterLoadBalancer)
	ManualLbConfig() GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference
	ManualLbConfigInput() *GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfig
	MetalLbConfig() GoogleGkeonpremVmwareAdminClusterLoadBalancerMetalLbConfigOutputReference
	MetalLbConfigInput() *GoogleGkeonpremVmwareAdminClusterLoadBalancerMetalLbConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VipConfig() GoogleGkeonpremVmwareAdminClusterLoadBalancerVipConfigOutputReference
	VipConfigInput() *GoogleGkeonpremVmwareAdminClusterLoadBalancerVipConfig
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
	PutF5Config(value *GoogleGkeonpremVmwareAdminClusterLoadBalancerF5Config)
	PutManualLbConfig(value *GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfig)
	PutMetalLbConfig(value *GoogleGkeonpremVmwareAdminClusterLoadBalancerMetalLbConfig)
	PutVipConfig(value *GoogleGkeonpremVmwareAdminClusterLoadBalancerVipConfig)
	ResetF5Config()
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

// The jsii proxy struct for GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference
type jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) F5Config() GoogleGkeonpremVmwareAdminClusterLoadBalancerF5ConfigOutputReference {
	var returns GoogleGkeonpremVmwareAdminClusterLoadBalancerF5ConfigOutputReference
	_jsii_.Get(
		j,
		"f5Config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) F5ConfigInput() *GoogleGkeonpremVmwareAdminClusterLoadBalancerF5Config {
	var returns *GoogleGkeonpremVmwareAdminClusterLoadBalancerF5Config
	_jsii_.Get(
		j,
		"f5ConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) InternalValue() *GoogleGkeonpremVmwareAdminClusterLoadBalancer {
	var returns *GoogleGkeonpremVmwareAdminClusterLoadBalancer
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) ManualLbConfig() GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference {
	var returns GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference
	_jsii_.Get(
		j,
		"manualLbConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) ManualLbConfigInput() *GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfig {
	var returns *GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfig
	_jsii_.Get(
		j,
		"manualLbConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) MetalLbConfig() GoogleGkeonpremVmwareAdminClusterLoadBalancerMetalLbConfigOutputReference {
	var returns GoogleGkeonpremVmwareAdminClusterLoadBalancerMetalLbConfigOutputReference
	_jsii_.Get(
		j,
		"metalLbConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) MetalLbConfigInput() *GoogleGkeonpremVmwareAdminClusterLoadBalancerMetalLbConfig {
	var returns *GoogleGkeonpremVmwareAdminClusterLoadBalancerMetalLbConfig
	_jsii_.Get(
		j,
		"metalLbConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) VipConfig() GoogleGkeonpremVmwareAdminClusterLoadBalancerVipConfigOutputReference {
	var returns GoogleGkeonpremVmwareAdminClusterLoadBalancerVipConfigOutputReference
	_jsii_.Get(
		j,
		"vipConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) VipConfigInput() *GoogleGkeonpremVmwareAdminClusterLoadBalancerVipConfig {
	var returns *GoogleGkeonpremVmwareAdminClusterLoadBalancerVipConfig
	_jsii_.Get(
		j,
		"vipConfigInput",
		&returns,
	)
	return returns
}


func NewGoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleGkeonpremVmwareAdminCluster.GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference_Override(g GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleGkeonpremVmwareAdminCluster.GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference)SetInternalValue(val *GoogleGkeonpremVmwareAdminClusterLoadBalancer) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) PutF5Config(value *GoogleGkeonpremVmwareAdminClusterLoadBalancerF5Config) {
	if err := g.validatePutF5ConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putF5Config",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) PutManualLbConfig(value *GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfig) {
	if err := g.validatePutManualLbConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putManualLbConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) PutMetalLbConfig(value *GoogleGkeonpremVmwareAdminClusterLoadBalancerMetalLbConfig) {
	if err := g.validatePutMetalLbConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMetalLbConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) PutVipConfig(value *GoogleGkeonpremVmwareAdminClusterLoadBalancerVipConfig) {
	if err := g.validatePutVipConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVipConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) ResetF5Config() {
	_jsii_.InvokeVoid(
		g,
		"resetF5Config",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) ResetManualLbConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetManualLbConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) ResetMetalLbConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetMetalLbConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

