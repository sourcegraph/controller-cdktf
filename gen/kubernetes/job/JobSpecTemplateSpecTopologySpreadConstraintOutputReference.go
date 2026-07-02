package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/kubernetes/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/kubernetes/job/internal"
)

type JobSpecTemplateSpecTopologySpreadConstraintOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
	LabelSelector() JobSpecTemplateSpecTopologySpreadConstraintLabelSelectorList
	LabelSelectorInput() any
	MaxSkew() *float64
	SetMaxSkew(val *float64)
	MaxSkewInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TopologyKey() *string
	SetTopologyKey(val *string)
	TopologyKeyInput() *string
	WhenUnsatisfiable() *string
	SetWhenUnsatisfiable(val *string)
	WhenUnsatisfiableInput() *string
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
	PutLabelSelector(value any)
	ResetLabelSelector()
	ResetMaxSkew()
	ResetTopologyKey()
	ResetWhenUnsatisfiable()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobSpecTemplateSpecTopologySpreadConstraintOutputReference
type jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) LabelSelector() JobSpecTemplateSpecTopologySpreadConstraintLabelSelectorList {
	var returns JobSpecTemplateSpecTopologySpreadConstraintLabelSelectorList
	_jsii_.Get(
		j,
		"labelSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) LabelSelectorInput() any {
	var returns any
	_jsii_.Get(
		j,
		"labelSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) MaxSkew() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSkew",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) MaxSkewInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSkewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) TopologyKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topologyKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) TopologyKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topologyKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) WhenUnsatisfiable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whenUnsatisfiable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) WhenUnsatisfiableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whenUnsatisfiableInput",
		&returns,
	)
	return returns
}

func NewJobSpecTemplateSpecTopologySpreadConstraintOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) JobSpecTemplateSpecTopologySpreadConstraintOutputReference {
	_init_.Initialize()

	if err := validateNewJobSpecTemplateSpecTopologySpreadConstraintOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.job.JobSpecTemplateSpecTopologySpreadConstraintOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewJobSpecTemplateSpecTopologySpreadConstraintOutputReference_Override(j JobSpecTemplateSpecTopologySpreadConstraintOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.job.JobSpecTemplateSpecTopologySpreadConstraintOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		j,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) SetMaxSkew(val *float64) {
	if err := j.validateSetMaxSkewParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSkew",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) SetTopologyKey(val *string) {
	if err := j.validateSetTopologyKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topologyKey",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) SetWhenUnsatisfiable(val *string) {
	if err := j.validateSetWhenUnsatisfiableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whenUnsatisfiable",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := j.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		j,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := j.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := j.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		j,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := j.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		j,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := j.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		j,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := j.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		j,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := j.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		j,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := j.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := j.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		j,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := j.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) PutLabelSelector(value any) {
	if err := j.validatePutLabelSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putLabelSelector",
		[]any{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) ResetLabelSelector() {
	_jsii_.InvokeVoid(
		j,
		"resetLabelSelector",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) ResetMaxSkew() {
	_jsii_.InvokeVoid(
		j,
		"resetMaxSkew",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) ResetTopologyKey() {
	_jsii_.InvokeVoid(
		j,
		"resetTopologyKey",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) ResetWhenUnsatisfiable() {
	_jsii_.InvokeVoid(
		j,
		"resetWhenUnsatisfiable",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := j.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		j,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
