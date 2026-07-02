package fisexperimenttemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/fisexperimenttemplate/internal"
)

type FisExperimentTemplateTargetOutputReference interface {
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
	Filter() FisExperimentTemplateTargetFilterList
	FilterInput() any
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
	Name() *string
	SetName(val *string)
	NameInput() *string
	ResourceArns() *[]*string
	SetResourceArns(val *[]*string)
	ResourceArnsInput() *[]*string
	ResourceTag() FisExperimentTemplateTargetResourceTagList
	ResourceTagInput() any
	ResourceType() *string
	SetResourceType(val *string)
	ResourceTypeInput() *string
	SelectionMode() *string
	SetSelectionMode(val *string)
	SelectionModeInput() *string
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
	PutFilter(value any)
	PutResourceTag(value any)
	ResetFilter()
	ResetResourceArns()
	ResetResourceTag()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FisExperimentTemplateTargetOutputReference
type jsiiProxy_FisExperimentTemplateTargetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FisExperimentTemplateTargetOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateTargetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateTargetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateTargetOutputReference) Filter() FisExperimentTemplateTargetFilterList {
	var returns FisExperimentTemplateTargetFilterList
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateTargetOutputReference) FilterInput() any {
	var returns any
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateTargetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateTargetOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateTargetOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateTargetOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateTargetOutputReference) ResourceArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateTargetOutputReference) ResourceArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateTargetOutputReference) ResourceTag() FisExperimentTemplateTargetResourceTagList {
	var returns FisExperimentTemplateTargetResourceTagList
	_jsii_.Get(
		j,
		"resourceTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateTargetOutputReference) ResourceTagInput() any {
	var returns any
	_jsii_.Get(
		j,
		"resourceTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateTargetOutputReference) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateTargetOutputReference) ResourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateTargetOutputReference) SelectionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selectionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateTargetOutputReference) SelectionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selectionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateTargetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateTargetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewFisExperimentTemplateTargetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FisExperimentTemplateTargetOutputReference {
	_init_.Initialize()

	if err := validateNewFisExperimentTemplateTargetOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_FisExperimentTemplateTargetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.fisExperimentTemplate.FisExperimentTemplateTargetOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFisExperimentTemplateTargetOutputReference_Override(f FisExperimentTemplateTargetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.fisExperimentTemplate.FisExperimentTemplateTargetOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FisExperimentTemplateTargetOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FisExperimentTemplateTargetOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FisExperimentTemplateTargetOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FisExperimentTemplateTargetOutputReference) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FisExperimentTemplateTargetOutputReference) SetResourceArns(val *[]*string) {
	if err := j.validateSetResourceArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceArns",
		val,
	)
}

func (j *jsiiProxy_FisExperimentTemplateTargetOutputReference) SetResourceType(val *string) {
	if err := j.validateSetResourceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceType",
		val,
	)
}

func (j *jsiiProxy_FisExperimentTemplateTargetOutputReference) SetSelectionMode(val *string) {
	if err := j.validateSetSelectionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selectionMode",
		val,
	)
}

func (j *jsiiProxy_FisExperimentTemplateTargetOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FisExperimentTemplateTargetOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FisExperimentTemplateTargetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FisExperimentTemplateTargetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FisExperimentTemplateTargetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FisExperimentTemplateTargetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FisExperimentTemplateTargetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FisExperimentTemplateTargetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FisExperimentTemplateTargetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FisExperimentTemplateTargetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FisExperimentTemplateTargetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FisExperimentTemplateTargetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FisExperimentTemplateTargetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FisExperimentTemplateTargetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FisExperimentTemplateTargetOutputReference) PutFilter(value any) {
	if err := f.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putFilter",
		[]any{value},
	)
}

func (f *jsiiProxy_FisExperimentTemplateTargetOutputReference) PutResourceTag(value any) {
	if err := f.validatePutResourceTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putResourceTag",
		[]any{value},
	)
}

func (f *jsiiProxy_FisExperimentTemplateTargetOutputReference) ResetFilter() {
	_jsii_.InvokeVoid(
		f,
		"resetFilter",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FisExperimentTemplateTargetOutputReference) ResetResourceArns() {
	_jsii_.InvokeVoid(
		f,
		"resetResourceArns",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FisExperimentTemplateTargetOutputReference) ResetResourceTag() {
	_jsii_.InvokeVoid(
		f,
		"resetResourceTag",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FisExperimentTemplateTargetOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := f.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		f,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FisExperimentTemplateTargetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
