package workflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/incident/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/incident/workflow/internal"
)

type WorkflowExpressionsOperationsOutputReference interface {
	cdktf.ComplexObject
	Branches() WorkflowExpressionsOperationsBranchesOutputReference
	BranchesInput() any
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
	Filter() WorkflowExpressionsOperationsFilterOutputReference
	FilterInput() any
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
	Navigate() WorkflowExpressionsOperationsNavigateOutputReference
	NavigateInput() any
	OperationType() *string
	SetOperationType(val *string)
	OperationTypeInput() *string
	Parse() WorkflowExpressionsOperationsParseOutputReference
	ParseInput() any
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
	PutBranches(value *WorkflowExpressionsOperationsBranches)
	PutFilter(value *WorkflowExpressionsOperationsFilter)
	PutNavigate(value *WorkflowExpressionsOperationsNavigate)
	PutParse(value *WorkflowExpressionsOperationsParse)
	ResetBranches()
	ResetFilter()
	ResetNavigate()
	ResetParse()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkflowExpressionsOperationsOutputReference
type jsiiProxy_WorkflowExpressionsOperationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkflowExpressionsOperationsOutputReference) Branches() WorkflowExpressionsOperationsBranchesOutputReference {
	var returns WorkflowExpressionsOperationsBranchesOutputReference
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkflowExpressionsOperationsOutputReference) BranchesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"branchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkflowExpressionsOperationsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkflowExpressionsOperationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkflowExpressionsOperationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkflowExpressionsOperationsOutputReference) Filter() WorkflowExpressionsOperationsFilterOutputReference {
	var returns WorkflowExpressionsOperationsFilterOutputReference
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkflowExpressionsOperationsOutputReference) FilterInput() any {
	var returns any
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkflowExpressionsOperationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkflowExpressionsOperationsOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkflowExpressionsOperationsOutputReference) Navigate() WorkflowExpressionsOperationsNavigateOutputReference {
	var returns WorkflowExpressionsOperationsNavigateOutputReference
	_jsii_.Get(
		j,
		"navigate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkflowExpressionsOperationsOutputReference) NavigateInput() any {
	var returns any
	_jsii_.Get(
		j,
		"navigateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkflowExpressionsOperationsOutputReference) OperationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkflowExpressionsOperationsOutputReference) OperationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkflowExpressionsOperationsOutputReference) Parse() WorkflowExpressionsOperationsParseOutputReference {
	var returns WorkflowExpressionsOperationsParseOutputReference
	_jsii_.Get(
		j,
		"parse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkflowExpressionsOperationsOutputReference) ParseInput() any {
	var returns any
	_jsii_.Get(
		j,
		"parseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkflowExpressionsOperationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkflowExpressionsOperationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewWorkflowExpressionsOperationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WorkflowExpressionsOperationsOutputReference {
	_init_.Initialize()

	if err := validateNewWorkflowExpressionsOperationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkflowExpressionsOperationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-incident.workflow.WorkflowExpressionsOperationsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWorkflowExpressionsOperationsOutputReference_Override(w WorkflowExpressionsOperationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-incident.workflow.WorkflowExpressionsOperationsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WorkflowExpressionsOperationsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkflowExpressionsOperationsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkflowExpressionsOperationsOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkflowExpressionsOperationsOutputReference) SetOperationType(val *string) {
	if err := j.validateSetOperationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationType",
		val,
	)
}

func (j *jsiiProxy_WorkflowExpressionsOperationsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkflowExpressionsOperationsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WorkflowExpressionsOperationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkflowExpressionsOperationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkflowExpressionsOperationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkflowExpressionsOperationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkflowExpressionsOperationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkflowExpressionsOperationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkflowExpressionsOperationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkflowExpressionsOperationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkflowExpressionsOperationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkflowExpressionsOperationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkflowExpressionsOperationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkflowExpressionsOperationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkflowExpressionsOperationsOutputReference) PutBranches(value *WorkflowExpressionsOperationsBranches) {
	if err := w.validatePutBranchesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putBranches",
		[]any{value},
	)
}

func (w *jsiiProxy_WorkflowExpressionsOperationsOutputReference) PutFilter(value *WorkflowExpressionsOperationsFilter) {
	if err := w.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putFilter",
		[]any{value},
	)
}

func (w *jsiiProxy_WorkflowExpressionsOperationsOutputReference) PutNavigate(value *WorkflowExpressionsOperationsNavigate) {
	if err := w.validatePutNavigateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putNavigate",
		[]any{value},
	)
}

func (w *jsiiProxy_WorkflowExpressionsOperationsOutputReference) PutParse(value *WorkflowExpressionsOperationsParse) {
	if err := w.validatePutParseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putParse",
		[]any{value},
	)
}

func (w *jsiiProxy_WorkflowExpressionsOperationsOutputReference) ResetBranches() {
	_jsii_.InvokeVoid(
		w,
		"resetBranches",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkflowExpressionsOperationsOutputReference) ResetFilter() {
	_jsii_.InvokeVoid(
		w,
		"resetFilter",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkflowExpressionsOperationsOutputReference) ResetNavigate() {
	_jsii_.InvokeVoid(
		w,
		"resetNavigate",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkflowExpressionsOperationsOutputReference) ResetParse() {
	_jsii_.InvokeVoid(
		w,
		"resetParse",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkflowExpressionsOperationsOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		w,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkflowExpressionsOperationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
