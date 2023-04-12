package googleappenginestandardappversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googleappenginestandardappversion/internal"
)

type GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference interface {
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
	InternalValue() *GoogleAppEngineStandardAppVersionAutomaticScaling
	SetInternalValue(val *GoogleAppEngineStandardAppVersionAutomaticScaling)
	MaxConcurrentRequests() *float64
	SetMaxConcurrentRequests(val *float64)
	MaxConcurrentRequestsInput() *float64
	MaxIdleInstances() *float64
	SetMaxIdleInstances(val *float64)
	MaxIdleInstancesInput() *float64
	MaxPendingLatency() *string
	SetMaxPendingLatency(val *string)
	MaxPendingLatencyInput() *string
	MinIdleInstances() *float64
	SetMinIdleInstances(val *float64)
	MinIdleInstancesInput() *float64
	MinPendingLatency() *string
	SetMinPendingLatency(val *string)
	MinPendingLatencyInput() *string
	StandardSchedulerSettings() GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference
	StandardSchedulerSettingsInput() *GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettings
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
	PutStandardSchedulerSettings(value *GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettings)
	ResetMaxConcurrentRequests()
	ResetMaxIdleInstances()
	ResetMaxPendingLatency()
	ResetMinIdleInstances()
	ResetMinPendingLatency()
	ResetStandardSchedulerSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference
type jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) InternalValue() *GoogleAppEngineStandardAppVersionAutomaticScaling {
	var returns *GoogleAppEngineStandardAppVersionAutomaticScaling
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) MaxConcurrentRequests() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentRequests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) MaxConcurrentRequestsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentRequestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) MaxIdleInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) MaxIdleInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) MaxPendingLatency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxPendingLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) MaxPendingLatencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxPendingLatencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) MinIdleInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIdleInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) MinIdleInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIdleInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) MinPendingLatency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minPendingLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) MinPendingLatencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minPendingLatencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) StandardSchedulerSettings() GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference {
	var returns GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference
	_jsii_.Get(
		j,
		"standardSchedulerSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) StandardSchedulerSettingsInput() *GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettings {
	var returns *GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettings
	_jsii_.Get(
		j,
		"standardSchedulerSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleAppEngineStandardAppVersionAutomaticScalingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleAppEngineStandardAppVersionAutomaticScalingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleAppEngineStandardAppVersion.GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleAppEngineStandardAppVersionAutomaticScalingOutputReference_Override(g GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleAppEngineStandardAppVersion.GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference)SetInternalValue(val *GoogleAppEngineStandardAppVersionAutomaticScaling) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference)SetMaxConcurrentRequests(val *float64) {
	if err := j.validateSetMaxConcurrentRequestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentRequests",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference)SetMaxIdleInstances(val *float64) {
	if err := j.validateSetMaxIdleInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxIdleInstances",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference)SetMaxPendingLatency(val *string) {
	if err := j.validateSetMaxPendingLatencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPendingLatency",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference)SetMinIdleInstances(val *float64) {
	if err := j.validateSetMinIdleInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minIdleInstances",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference)SetMinPendingLatency(val *string) {
	if err := j.validateSetMinPendingLatencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minPendingLatency",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) PutStandardSchedulerSettings(value *GoogleAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettings) {
	if err := g.validatePutStandardSchedulerSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStandardSchedulerSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) ResetMaxConcurrentRequests() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxConcurrentRequests",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) ResetMaxIdleInstances() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxIdleInstances",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) ResetMaxPendingLatency() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxPendingLatency",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) ResetMinIdleInstances() {
	_jsii_.InvokeVoid(
		g,
		"resetMinIdleInstances",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) ResetMinPendingLatency() {
	_jsii_.InvokeVoid(
		g,
		"resetMinPendingLatency",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) ResetStandardSchedulerSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetStandardSchedulerSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

