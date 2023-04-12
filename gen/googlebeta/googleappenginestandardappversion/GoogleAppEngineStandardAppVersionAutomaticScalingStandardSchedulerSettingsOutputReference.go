package googleappenginestandardappversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googleappenginestandardappversion/internal"
)

type GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettings
	SetInternalValue(val *GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettings)
	MaxInstances() *float64
	SetMaxInstances(val *float64)
	MaxInstancesInput() *float64
	MinInstances() *float64
	SetMinInstances(val *float64)
	MinInstancesInput() *float64
	TargetCpuUtilization() *float64
	SetTargetCpuUtilization(val *float64)
	TargetCpuUtilizationInput() *float64
	TargetThroughputUtilization() *float64
	SetTargetThroughputUtilization(val *float64)
	TargetThroughputUtilizationInput() *float64
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
	ResetMaxInstances()
	ResetMinInstances()
	ResetTargetCpuUtilization()
	ResetTargetThroughputUtilization()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference
type jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) InternalValue() *GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettings {
	var returns *GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) MaxInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) MaxInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) MinInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) MinInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) TargetCpuUtilization() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetCpuUtilization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) TargetCpuUtilizationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetCpuUtilizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) TargetThroughputUtilization() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetThroughputUtilization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) TargetThroughputUtilizationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetThroughputUtilizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleAppEngineStandardAppVersion.GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference_Override(g GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleAppEngineStandardAppVersion.GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference)SetInternalValue(val *GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference)SetMaxInstances(val *float64) {
	if err := j.validateSetMaxInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxInstances",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference)SetMinInstances(val *float64) {
	if err := j.validateSetMinInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minInstances",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference)SetTargetCpuUtilization(val *float64) {
	if err := j.validateSetTargetCpuUtilizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetCpuUtilization",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference)SetTargetThroughputUtilization(val *float64) {
	if err := j.validateSetTargetThroughputUtilizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetThroughputUtilization",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) ResetMaxInstances() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxInstances",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) ResetMinInstances() {
	_jsii_.InvokeVoid(
		g,
		"resetMinInstances",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) ResetTargetCpuUtilization() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetCpuUtilization",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) ResetTargetThroughputUtilization() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetThroughputUtilization",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

