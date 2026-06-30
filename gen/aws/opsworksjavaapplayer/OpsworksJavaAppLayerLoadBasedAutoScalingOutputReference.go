package opsworksjavaapplayer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/opsworksjavaapplayer/internal"
)

type OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
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
	Downscaling() OpsworksJavaAppLayerLoadBasedAutoScalingDownscalingOutputReference
	DownscalingInput() *OpsworksJavaAppLayerLoadBasedAutoScalingDownscaling
	Enable() any
	SetEnable(val any)
	EnableInput() any
	// Experimental.
	Fqn() *string
	InternalValue() *OpsworksJavaAppLayerLoadBasedAutoScaling
	SetInternalValue(val *OpsworksJavaAppLayerLoadBasedAutoScaling)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Upscaling() OpsworksJavaAppLayerLoadBasedAutoScalingUpscalingOutputReference
	UpscalingInput() *OpsworksJavaAppLayerLoadBasedAutoScalingUpscaling
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	PutDownscaling(value *OpsworksJavaAppLayerLoadBasedAutoScalingDownscaling)
	PutUpscaling(value *OpsworksJavaAppLayerLoadBasedAutoScalingUpscaling)
	ResetDownscaling()
	ResetEnable()
	ResetUpscaling()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference
type jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) Downscaling() OpsworksJavaAppLayerLoadBasedAutoScalingDownscalingOutputReference {
	var returns OpsworksJavaAppLayerLoadBasedAutoScalingDownscalingOutputReference
	_jsii_.Get(
		j,
		"downscaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) DownscalingInput() *OpsworksJavaAppLayerLoadBasedAutoScalingDownscaling {
	var returns *OpsworksJavaAppLayerLoadBasedAutoScalingDownscaling
	_jsii_.Get(
		j,
		"downscalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) Enable() any {
	var returns any
	_jsii_.Get(
		j,
		"enable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) EnableInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) InternalValue() *OpsworksJavaAppLayerLoadBasedAutoScaling {
	var returns *OpsworksJavaAppLayerLoadBasedAutoScaling
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) Upscaling() OpsworksJavaAppLayerLoadBasedAutoScalingUpscalingOutputReference {
	var returns OpsworksJavaAppLayerLoadBasedAutoScalingUpscalingOutputReference
	_jsii_.Get(
		j,
		"upscaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) UpscalingInput() *OpsworksJavaAppLayerLoadBasedAutoScalingUpscaling {
	var returns *OpsworksJavaAppLayerLoadBasedAutoScalingUpscaling
	_jsii_.Get(
		j,
		"upscalingInput",
		&returns,
	)
	return returns
}

func NewOpsworksJavaAppLayerLoadBasedAutoScalingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference {
	_init_.Initialize()

	if err := validateNewOpsworksJavaAppLayerLoadBasedAutoScalingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.opsworksJavaAppLayer.OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOpsworksJavaAppLayerLoadBasedAutoScalingOutputReference_Override(o OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.opsworksJavaAppLayer.OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference",
		[]any{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) SetEnable(val any) {
	if err := j.validateSetEnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enable",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) SetInternalValue(val *OpsworksJavaAppLayerLoadBasedAutoScaling) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) PutDownscaling(value *OpsworksJavaAppLayerLoadBasedAutoScalingDownscaling) {
	if err := o.validatePutDownscalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDownscaling",
		[]any{value},
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) PutUpscaling(value *OpsworksJavaAppLayerLoadBasedAutoScalingUpscaling) {
	if err := o.validatePutUpscalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putUpscaling",
		[]any{value},
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) ResetDownscaling() {
	_jsii_.InvokeVoid(
		o,
		"resetDownscaling",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) ResetEnable() {
	_jsii_.InvokeVoid(
		o,
		"resetEnable",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) ResetUpscaling() {
	_jsii_.InvokeVoid(
		o,
		"resetUpscaling",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		o,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksJavaAppLayerLoadBasedAutoScalingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
