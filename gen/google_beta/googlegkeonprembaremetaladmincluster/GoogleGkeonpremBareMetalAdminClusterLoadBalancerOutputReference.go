package googlegkeonprembaremetaladmincluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlegkeonprembaremetaladmincluster/internal"
)

type GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference interface {
	cdktn.ComplexObject
	BgpLbConfig() GoogleGkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference
	BgpLbConfigInput() *GoogleGkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfig
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
	InternalValue() *GoogleGkeonpremBareMetalAdminClusterLoadBalancer
	SetInternalValue(val *GoogleGkeonpremBareMetalAdminClusterLoadBalancer)
	ManualLbConfig() GoogleGkeonpremBareMetalAdminClusterLoadBalancerManualLbConfigOutputReference
	ManualLbConfigInput() *GoogleGkeonpremBareMetalAdminClusterLoadBalancerManualLbConfig
	PortConfig() GoogleGkeonpremBareMetalAdminClusterLoadBalancerPortConfigOutputReference
	PortConfigInput() *GoogleGkeonpremBareMetalAdminClusterLoadBalancerPortConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VipConfig() GoogleGkeonpremBareMetalAdminClusterLoadBalancerVipConfigOutputReference
	VipConfigInput() *GoogleGkeonpremBareMetalAdminClusterLoadBalancerVipConfig
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutBgpLbConfig(value *GoogleGkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfig)
	PutManualLbConfig(value *GoogleGkeonpremBareMetalAdminClusterLoadBalancerManualLbConfig)
	PutPortConfig(value *GoogleGkeonpremBareMetalAdminClusterLoadBalancerPortConfig)
	PutVipConfig(value *GoogleGkeonpremBareMetalAdminClusterLoadBalancerVipConfig)
	ResetBgpLbConfig()
	ResetManualLbConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference
type jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) BgpLbConfig() GoogleGkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference {
	var returns GoogleGkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference
	_jsii_.Get(
		j,
		"bgpLbConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) BgpLbConfigInput() *GoogleGkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfig {
	var returns *GoogleGkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfig
	_jsii_.Get(
		j,
		"bgpLbConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) InternalValue() *GoogleGkeonpremBareMetalAdminClusterLoadBalancer {
	var returns *GoogleGkeonpremBareMetalAdminClusterLoadBalancer
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) ManualLbConfig() GoogleGkeonpremBareMetalAdminClusterLoadBalancerManualLbConfigOutputReference {
	var returns GoogleGkeonpremBareMetalAdminClusterLoadBalancerManualLbConfigOutputReference
	_jsii_.Get(
		j,
		"manualLbConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) ManualLbConfigInput() *GoogleGkeonpremBareMetalAdminClusterLoadBalancerManualLbConfig {
	var returns *GoogleGkeonpremBareMetalAdminClusterLoadBalancerManualLbConfig
	_jsii_.Get(
		j,
		"manualLbConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) PortConfig() GoogleGkeonpremBareMetalAdminClusterLoadBalancerPortConfigOutputReference {
	var returns GoogleGkeonpremBareMetalAdminClusterLoadBalancerPortConfigOutputReference
	_jsii_.Get(
		j,
		"portConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) PortConfigInput() *GoogleGkeonpremBareMetalAdminClusterLoadBalancerPortConfig {
	var returns *GoogleGkeonpremBareMetalAdminClusterLoadBalancerPortConfig
	_jsii_.Get(
		j,
		"portConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) VipConfig() GoogleGkeonpremBareMetalAdminClusterLoadBalancerVipConfigOutputReference {
	var returns GoogleGkeonpremBareMetalAdminClusterLoadBalancerVipConfigOutputReference
	_jsii_.Get(
		j,
		"vipConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) VipConfigInput() *GoogleGkeonpremBareMetalAdminClusterLoadBalancerVipConfig {
	var returns *GoogleGkeonpremBareMetalAdminClusterLoadBalancerVipConfig
	_jsii_.Get(
		j,
		"vipConfigInput",
		&returns,
	)
	return returns
}


func NewGoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleGkeonpremBareMetalAdminCluster.GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference_Override(g GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleGkeonpremBareMetalAdminCluster.GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference)SetInternalValue(val *GoogleGkeonpremBareMetalAdminClusterLoadBalancer) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) PutBgpLbConfig(value *GoogleGkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfig) {
	if err := g.validatePutBgpLbConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBgpLbConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) PutManualLbConfig(value *GoogleGkeonpremBareMetalAdminClusterLoadBalancerManualLbConfig) {
	if err := g.validatePutManualLbConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putManualLbConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) PutPortConfig(value *GoogleGkeonpremBareMetalAdminClusterLoadBalancerPortConfig) {
	if err := g.validatePutPortConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPortConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) PutVipConfig(value *GoogleGkeonpremBareMetalAdminClusterLoadBalancerVipConfig) {
	if err := g.validatePutVipConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVipConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) ResetBgpLbConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetBgpLbConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) ResetManualLbConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetManualLbConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalAdminClusterLoadBalancerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

