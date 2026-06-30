package osconfigospolicyassignment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/osconfigospolicyassignment/internal"
)

type OsConfigOsPolicyAssignmentInstanceFilterOutputReference interface {
	cdktf.ComplexObject
	All() any
	SetAll(val any)
	AllInput() any
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
	ExclusionLabels() OsConfigOsPolicyAssignmentInstanceFilterExclusionLabelsList
	ExclusionLabelsInput() any
	// Experimental.
	Fqn() *string
	InclusionLabels() OsConfigOsPolicyAssignmentInstanceFilterInclusionLabelsList
	InclusionLabelsInput() any
	InternalValue() *OsConfigOsPolicyAssignmentInstanceFilter
	SetInternalValue(val *OsConfigOsPolicyAssignmentInstanceFilter)
	Inventories() OsConfigOsPolicyAssignmentInstanceFilterInventoriesList
	InventoriesInput() any
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
	PutExclusionLabels(value any)
	PutInclusionLabels(value any)
	PutInventories(value any)
	ResetAll()
	ResetExclusionLabels()
	ResetInclusionLabels()
	ResetInventories()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OsConfigOsPolicyAssignmentInstanceFilterOutputReference
type jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) All() any {
	var returns any
	_jsii_.Get(
		j,
		"all",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) AllInput() any {
	var returns any
	_jsii_.Get(
		j,
		"allInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) ExclusionLabels() OsConfigOsPolicyAssignmentInstanceFilterExclusionLabelsList {
	var returns OsConfigOsPolicyAssignmentInstanceFilterExclusionLabelsList
	_jsii_.Get(
		j,
		"exclusionLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) ExclusionLabelsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"exclusionLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) InclusionLabels() OsConfigOsPolicyAssignmentInstanceFilterInclusionLabelsList {
	var returns OsConfigOsPolicyAssignmentInstanceFilterInclusionLabelsList
	_jsii_.Get(
		j,
		"inclusionLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) InclusionLabelsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"inclusionLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) InternalValue() *OsConfigOsPolicyAssignmentInstanceFilter {
	var returns *OsConfigOsPolicyAssignmentInstanceFilter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) Inventories() OsConfigOsPolicyAssignmentInstanceFilterInventoriesList {
	var returns OsConfigOsPolicyAssignmentInstanceFilterInventoriesList
	_jsii_.Get(
		j,
		"inventories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) InventoriesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"inventoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewOsConfigOsPolicyAssignmentInstanceFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OsConfigOsPolicyAssignmentInstanceFilterOutputReference {
	_init_.Initialize()

	if err := validateNewOsConfigOsPolicyAssignmentInstanceFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.osConfigOsPolicyAssignment.OsConfigOsPolicyAssignmentInstanceFilterOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOsConfigOsPolicyAssignmentInstanceFilterOutputReference_Override(o OsConfigOsPolicyAssignmentInstanceFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.osConfigOsPolicyAssignment.OsConfigOsPolicyAssignmentInstanceFilterOutputReference",
		[]any{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) SetAll(val any) {
	if err := j.validateSetAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"all",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) SetInternalValue(val *OsConfigOsPolicyAssignmentInstanceFilter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) PutExclusionLabels(value any) {
	if err := o.validatePutExclusionLabelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putExclusionLabels",
		[]any{value},
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) PutInclusionLabels(value any) {
	if err := o.validatePutInclusionLabelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putInclusionLabels",
		[]any{value},
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) PutInventories(value any) {
	if err := o.validatePutInventoriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putInventories",
		[]any{value},
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) ResetAll() {
	_jsii_.InvokeVoid(
		o,
		"resetAll",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) ResetExclusionLabels() {
	_jsii_.InvokeVoid(
		o,
		"resetExclusionLabels",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) ResetInclusionLabels() {
	_jsii_.InvokeVoid(
		o,
		"resetInclusionLabels",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) ResetInventories() {
	_jsii_.InvokeVoid(
		o,
		"resetInventories",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentInstanceFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
