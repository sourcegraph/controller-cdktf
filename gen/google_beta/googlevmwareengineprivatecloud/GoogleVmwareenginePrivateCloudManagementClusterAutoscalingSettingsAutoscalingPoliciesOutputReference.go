package googlevmwareengineprivatecloud

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlevmwareengineprivatecloud/internal"
)

type GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference interface {
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
	ConsumedMemoryThresholds() GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference
	ConsumedMemoryThresholdsInput() *GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholds
	CpuThresholds() GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesCpuThresholdsOutputReference
	CpuThresholdsInput() *GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesCpuThresholds
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
	StorageThresholds() GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesStorageThresholdsOutputReference
	StorageThresholdsInput() *GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesStorageThresholds
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
	PutConsumedMemoryThresholds(value *GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholds)
	PutCpuThresholds(value *GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesCpuThresholds)
	PutStorageThresholds(value *GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesStorageThresholds)
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

// The jsii proxy struct for GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference
type jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) AutoscalePolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoscalePolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) AutoscalePolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoscalePolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ConsumedMemoryThresholds() GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference {
	var returns GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference
	_jsii_.Get(
		j,
		"consumedMemoryThresholds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ConsumedMemoryThresholdsInput() *GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholds {
	var returns *GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholds
	_jsii_.Get(
		j,
		"consumedMemoryThresholdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) CpuThresholds() GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesCpuThresholdsOutputReference {
	var returns GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesCpuThresholdsOutputReference
	_jsii_.Get(
		j,
		"cpuThresholds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) CpuThresholdsInput() *GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesCpuThresholds {
	var returns *GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesCpuThresholds
	_jsii_.Get(
		j,
		"cpuThresholdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) NodeTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) NodeTypeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ScaleOutSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleOutSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ScaleOutSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleOutSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) StorageThresholds() GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesStorageThresholdsOutputReference {
	var returns GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesStorageThresholdsOutputReference
	_jsii_.Get(
		j,
		"storageThresholds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) StorageThresholdsInput() *GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesStorageThresholds {
	var returns *GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesStorageThresholds
	_jsii_.Get(
		j,
		"storageThresholdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVmwareenginePrivateCloud.GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference_Override(g GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVmwareenginePrivateCloud.GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference)SetAutoscalePolicyId(val *string) {
	if err := j.validateSetAutoscalePolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscalePolicyId",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference)SetNodeTypeId(val *string) {
	if err := j.validateSetNodeTypeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeTypeId",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference)SetScaleOutSize(val *float64) {
	if err := j.validateSetScaleOutSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleOutSize",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) PutConsumedMemoryThresholds(value *GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholds) {
	if err := g.validatePutConsumedMemoryThresholdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConsumedMemoryThresholds",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) PutCpuThresholds(value *GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesCpuThresholds) {
	if err := g.validatePutCpuThresholdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCpuThresholds",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) PutStorageThresholds(value *GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesStorageThresholds) {
	if err := g.validatePutStorageThresholdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStorageThresholds",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ResetConsumedMemoryThresholds() {
	_jsii_.InvokeVoid(
		g,
		"resetConsumedMemoryThresholds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ResetCpuThresholds() {
	_jsii_.InvokeVoid(
		g,
		"resetCpuThresholds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ResetStorageThresholds() {
	_jsii_.InvokeVoid(
		g,
		"resetStorageThresholds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

