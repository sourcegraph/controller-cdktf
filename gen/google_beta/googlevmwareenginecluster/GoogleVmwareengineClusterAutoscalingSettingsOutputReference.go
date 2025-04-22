package googlevmwareenginecluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlevmwareenginecluster/internal"
)

type GoogleVmwareengineClusterAutoscalingSettingsOutputReference interface {
	cdktf.ComplexObject
	AutoscalingPolicies() GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesList
	AutoscalingPoliciesInput() interface{}
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
	CoolDownPeriod() *string
	SetCoolDownPeriod(val *string)
	CoolDownPeriodInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleVmwareengineClusterAutoscalingSettings
	SetInternalValue(val *GoogleVmwareengineClusterAutoscalingSettings)
	MaxClusterNodeCount() *float64
	SetMaxClusterNodeCount(val *float64)
	MaxClusterNodeCountInput() *float64
	MinClusterNodeCount() *float64
	SetMinClusterNodeCount(val *float64)
	MinClusterNodeCountInput() *float64
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
	PutAutoscalingPolicies(value interface{})
	ResetCoolDownPeriod()
	ResetMaxClusterNodeCount()
	ResetMinClusterNodeCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVmwareengineClusterAutoscalingSettingsOutputReference
type jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) AutoscalingPolicies() GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesList {
	var returns GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesList
	_jsii_.Get(
		j,
		"autoscalingPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) AutoscalingPoliciesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscalingPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) CoolDownPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coolDownPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) CoolDownPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coolDownPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) InternalValue() *GoogleVmwareengineClusterAutoscalingSettings {
	var returns *GoogleVmwareengineClusterAutoscalingSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) MaxClusterNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxClusterNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) MaxClusterNodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxClusterNodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) MinClusterNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minClusterNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) MinClusterNodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minClusterNodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleVmwareengineClusterAutoscalingSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleVmwareengineClusterAutoscalingSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVmwareengineClusterAutoscalingSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVmwareengineCluster.GoogleVmwareengineClusterAutoscalingSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVmwareengineClusterAutoscalingSettingsOutputReference_Override(g GoogleVmwareengineClusterAutoscalingSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVmwareengineCluster.GoogleVmwareengineClusterAutoscalingSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference)SetCoolDownPeriod(val *string) {
	if err := j.validateSetCoolDownPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coolDownPeriod",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference)SetInternalValue(val *GoogleVmwareengineClusterAutoscalingSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference)SetMaxClusterNodeCount(val *float64) {
	if err := j.validateSetMaxClusterNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxClusterNodeCount",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference)SetMinClusterNodeCount(val *float64) {
	if err := j.validateSetMinClusterNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minClusterNodeCount",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) PutAutoscalingPolicies(value interface{}) {
	if err := g.validatePutAutoscalingPoliciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAutoscalingPolicies",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) ResetCoolDownPeriod() {
	_jsii_.InvokeVoid(
		g,
		"resetCoolDownPeriod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) ResetMaxClusterNodeCount() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxClusterNodeCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) ResetMinClusterNodeCount() {
	_jsii_.InvokeVoid(
		g,
		"resetMinClusterNodeCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

