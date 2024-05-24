package googlespannerinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlespannerinstance/internal"
)

type GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference interface {
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
	HighPriorityCpuUtilizationPercent() *float64
	SetHighPriorityCpuUtilizationPercent(val *float64)
	HighPriorityCpuUtilizationPercentInput() *float64
	InternalValue() *GoogleSpannerInstanceAutoscalingConfigAutoscalingTargets
	SetInternalValue(val *GoogleSpannerInstanceAutoscalingConfigAutoscalingTargets)
	StorageUtilizationPercent() *float64
	SetStorageUtilizationPercent(val *float64)
	StorageUtilizationPercentInput() *float64
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
	ResetHighPriorityCpuUtilizationPercent()
	ResetStorageUtilizationPercent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference
type jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference) HighPriorityCpuUtilizationPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"highPriorityCpuUtilizationPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference) HighPriorityCpuUtilizationPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"highPriorityCpuUtilizationPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference) InternalValue() *GoogleSpannerInstanceAutoscalingConfigAutoscalingTargets {
	var returns *GoogleSpannerInstanceAutoscalingConfigAutoscalingTargets
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference) StorageUtilizationPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageUtilizationPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference) StorageUtilizationPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageUtilizationPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleSpannerInstance.GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference_Override(g GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleSpannerInstance.GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference)SetHighPriorityCpuUtilizationPercent(val *float64) {
	if err := j.validateSetHighPriorityCpuUtilizationPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"highPriorityCpuUtilizationPercent",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference)SetInternalValue(val *GoogleSpannerInstanceAutoscalingConfigAutoscalingTargets) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference)SetStorageUtilizationPercent(val *float64) {
	if err := j.validateSetStorageUtilizationPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageUtilizationPercent",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference) ResetHighPriorityCpuUtilizationPercent() {
	_jsii_.InvokeVoid(
		g,
		"resetHighPriorityCpuUtilizationPercent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference) ResetStorageUtilizationPercent() {
	_jsii_.InvokeVoid(
		g,
		"resetStorageUtilizationPercent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

