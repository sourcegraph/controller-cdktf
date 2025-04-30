package googlevmwareenginecluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlevmwareenginecluster/internal"
)

type GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference interface {
	cdktf.ComplexObject
	AutoscalePolicyId() *string
	SetAutoscalePolicyId(val *string)
	AutoscalePolicyIdInput() *string
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
	ConsumedMemoryThresholds() GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference
	ConsumedMemoryThresholdsInput() *GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholds
	CpuThresholds() GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesCpuThresholdsOutputReference
	CpuThresholdsInput() *GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesCpuThresholds
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NodeTypeId() *string
	SetNodeTypeId(val *string)
	NodeTypeIdInput() *string
	ScaleOutSize() *float64
	SetScaleOutSize(val *float64)
	ScaleOutSizeInput() *float64
	StorageThresholds() GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesStorageThresholdsOutputReference
	StorageThresholdsInput() *GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesStorageThresholds
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
	PutConsumedMemoryThresholds(value *GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholds)
	PutCpuThresholds(value *GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesCpuThresholds)
	PutStorageThresholds(value *GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesStorageThresholds)
	ResetConsumedMemoryThresholds()
	ResetCpuThresholds()
	ResetStorageThresholds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference
type jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) AutoscalePolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoscalePolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) AutoscalePolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoscalePolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ConsumedMemoryThresholds() GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference {
	var returns GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference
	_jsii_.Get(
		j,
		"consumedMemoryThresholds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ConsumedMemoryThresholdsInput() *GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholds {
	var returns *GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholds
	_jsii_.Get(
		j,
		"consumedMemoryThresholdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) CpuThresholds() GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesCpuThresholdsOutputReference {
	var returns GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesCpuThresholdsOutputReference
	_jsii_.Get(
		j,
		"cpuThresholds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) CpuThresholdsInput() *GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesCpuThresholds {
	var returns *GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesCpuThresholds
	_jsii_.Get(
		j,
		"cpuThresholdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) NodeTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) NodeTypeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ScaleOutSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleOutSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ScaleOutSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleOutSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) StorageThresholds() GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesStorageThresholdsOutputReference {
	var returns GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesStorageThresholdsOutputReference
	_jsii_.Get(
		j,
		"storageThresholds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) StorageThresholdsInput() *GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesStorageThresholds {
	var returns *GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesStorageThresholds
	_jsii_.Get(
		j,
		"storageThresholdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVmwareengineCluster.GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference_Override(g GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVmwareengineCluster.GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference)SetAutoscalePolicyId(val *string) {
	if err := j.validateSetAutoscalePolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscalePolicyId",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference)SetNodeTypeId(val *string) {
	if err := j.validateSetNodeTypeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeTypeId",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference)SetScaleOutSize(val *float64) {
	if err := j.validateSetScaleOutSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleOutSize",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) PutConsumedMemoryThresholds(value *GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholds) {
	if err := g.validatePutConsumedMemoryThresholdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConsumedMemoryThresholds",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) PutCpuThresholds(value *GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesCpuThresholds) {
	if err := g.validatePutCpuThresholdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCpuThresholds",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) PutStorageThresholds(value *GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesStorageThresholds) {
	if err := g.validatePutStorageThresholdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStorageThresholds",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ResetConsumedMemoryThresholds() {
	_jsii_.InvokeVoid(
		g,
		"resetConsumedMemoryThresholds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ResetCpuThresholds() {
	_jsii_.InvokeVoid(
		g,
		"resetCpuThresholds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ResetStorageThresholds() {
	_jsii_.InvokeVoid(
		g,
		"resetStorageThresholds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

