package googlespannerinstancepartition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlespannerinstancepartition/internal"
)

type GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference interface {
	cdktn.ComplexObject
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
	InternalValue() *GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargets
	SetInternalValue(val *GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargets)
	StorageUtilizationPercent() *float64
	SetStorageUtilizationPercent(val *float64)
	StorageUtilizationPercentInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TotalCpuUtilizationPercent() *float64
	SetTotalCpuUtilizationPercent(val *float64)
	TotalCpuUtilizationPercentInput() *float64
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
	ResetHighPriorityCpuUtilizationPercent()
	ResetStorageUtilizationPercent()
	ResetTotalCpuUtilizationPercent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference
type jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) HighPriorityCpuUtilizationPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"highPriorityCpuUtilizationPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) HighPriorityCpuUtilizationPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"highPriorityCpuUtilizationPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) InternalValue() *GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargets {
	var returns *GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargets
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) StorageUtilizationPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageUtilizationPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) StorageUtilizationPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageUtilizationPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) TotalCpuUtilizationPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalCpuUtilizationPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) TotalCpuUtilizationPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalCpuUtilizationPercentInput",
		&returns,
	)
	return returns
}


func NewGoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleSpannerInstancePartition.GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference_Override(g GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleSpannerInstancePartition.GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference)SetHighPriorityCpuUtilizationPercent(val *float64) {
	if err := j.validateSetHighPriorityCpuUtilizationPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"highPriorityCpuUtilizationPercent",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference)SetInternalValue(val *GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargets) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference)SetStorageUtilizationPercent(val *float64) {
	if err := j.validateSetStorageUtilizationPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageUtilizationPercent",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference)SetTotalCpuUtilizationPercent(val *float64) {
	if err := j.validateSetTotalCpuUtilizationPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalCpuUtilizationPercent",
		val,
	)
}

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) ResetHighPriorityCpuUtilizationPercent() {
	_jsii_.InvokeVoid(
		g,
		"resetHighPriorityCpuUtilizationPercent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) ResetStorageUtilizationPercent() {
	_jsii_.InvokeVoid(
		g,
		"resetStorageUtilizationPercent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) ResetTotalCpuUtilizationPercent() {
	_jsii_.InvokeVoid(
		g,
		"resetTotalCpuUtilizationPercent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

