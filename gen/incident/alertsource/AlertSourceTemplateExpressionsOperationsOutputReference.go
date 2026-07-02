package alertsource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/incident/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/incident/alertsource/internal"
)

type AlertSourceTemplateExpressionsOperationsOutputReference interface {
	cdktf.ComplexObject
	Branches() AlertSourceTemplateExpressionsOperationsBranchesOutputReference
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
	Filter() AlertSourceTemplateExpressionsOperationsFilterOutputReference
	FilterInput() any
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
	Navigate() AlertSourceTemplateExpressionsOperationsNavigateOutputReference
	NavigateInput() any
	OperationType() *string
	SetOperationType(val *string)
	OperationTypeInput() *string
	Parse() AlertSourceTemplateExpressionsOperationsParseOutputReference
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
	PutBranches(value *AlertSourceTemplateExpressionsOperationsBranches)
	PutFilter(value *AlertSourceTemplateExpressionsOperationsFilter)
	PutNavigate(value *AlertSourceTemplateExpressionsOperationsNavigate)
	PutParse(value *AlertSourceTemplateExpressionsOperationsParse)
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

// The jsii proxy struct for AlertSourceTemplateExpressionsOperationsOutputReference
type jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) Branches() AlertSourceTemplateExpressionsOperationsBranchesOutputReference {
	var returns AlertSourceTemplateExpressionsOperationsBranchesOutputReference
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) BranchesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"branchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) Filter() AlertSourceTemplateExpressionsOperationsFilterOutputReference {
	var returns AlertSourceTemplateExpressionsOperationsFilterOutputReference
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) FilterInput() any {
	var returns any
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) Navigate() AlertSourceTemplateExpressionsOperationsNavigateOutputReference {
	var returns AlertSourceTemplateExpressionsOperationsNavigateOutputReference
	_jsii_.Get(
		j,
		"navigate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) NavigateInput() any {
	var returns any
	_jsii_.Get(
		j,
		"navigateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) OperationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) OperationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) Parse() AlertSourceTemplateExpressionsOperationsParseOutputReference {
	var returns AlertSourceTemplateExpressionsOperationsParseOutputReference
	_jsii_.Get(
		j,
		"parse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) ParseInput() any {
	var returns any
	_jsii_.Get(
		j,
		"parseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAlertSourceTemplateExpressionsOperationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AlertSourceTemplateExpressionsOperationsOutputReference {
	_init_.Initialize()

	if err := validateNewAlertSourceTemplateExpressionsOperationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-incident.alertSource.AlertSourceTemplateExpressionsOperationsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAlertSourceTemplateExpressionsOperationsOutputReference_Override(a AlertSourceTemplateExpressionsOperationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-incident.alertSource.AlertSourceTemplateExpressionsOperationsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) SetOperationType(val *string) {
	if err := j.validateSetOperationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationType",
		val,
	)
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) PutBranches(value *AlertSourceTemplateExpressionsOperationsBranches) {
	if err := a.validatePutBranchesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBranches",
		[]any{value},
	)
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) PutFilter(value *AlertSourceTemplateExpressionsOperationsFilter) {
	if err := a.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFilter",
		[]any{value},
	)
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) PutNavigate(value *AlertSourceTemplateExpressionsOperationsNavigate) {
	if err := a.validatePutNavigateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNavigate",
		[]any{value},
	)
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) PutParse(value *AlertSourceTemplateExpressionsOperationsParse) {
	if err := a.validatePutParseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putParse",
		[]any{value},
	)
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) ResetBranches() {
	_jsii_.InvokeVoid(
		a,
		"resetBranches",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) ResetFilter() {
	_jsii_.InvokeVoid(
		a,
		"resetFilter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) ResetNavigate() {
	_jsii_.InvokeVoid(
		a,
		"resetNavigate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) ResetParse() {
	_jsii_.InvokeVoid(
		a,
		"resetParse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		a,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
