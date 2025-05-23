package reportsystemhealthreview

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/nobl9/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/nobl9/reportsystemhealthreview/internal"
)

type ReportSystemHealthReviewFiltersOutputReference interface {
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
	InternalValue() *ReportSystemHealthReviewFilters
	SetInternalValue(val *ReportSystemHealthReviewFilters)
	Label() ReportSystemHealthReviewFiltersLabelList
	LabelInput() interface{}
	Projects() *[]*string
	SetProjects(val *[]*string)
	ProjectsInput() *[]*string
	Service() ReportSystemHealthReviewFiltersServiceList
	ServiceInput() interface{}
	Slo() ReportSystemHealthReviewFiltersSloList
	SloInput() interface{}
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
	PutLabel(value interface{})
	PutService(value interface{})
	PutSlo(value interface{})
	ResetLabel()
	ResetProjects()
	ResetService()
	ResetSlo()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ReportSystemHealthReviewFiltersOutputReference
type jsiiProxy_ReportSystemHealthReviewFiltersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) InternalValue() *ReportSystemHealthReviewFilters {
	var returns *ReportSystemHealthReviewFilters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) Label() ReportSystemHealthReviewFiltersLabelList {
	var returns ReportSystemHealthReviewFiltersLabelList
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) LabelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) Projects() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"projects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) ProjectsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"projectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) Service() ReportSystemHealthReviewFiltersServiceList {
	var returns ReportSystemHealthReviewFiltersServiceList
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) ServiceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) Slo() ReportSystemHealthReviewFiltersSloList {
	var returns ReportSystemHealthReviewFiltersSloList
	_jsii_.Get(
		j,
		"slo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) SloInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sloInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewReportSystemHealthReviewFiltersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ReportSystemHealthReviewFiltersOutputReference {
	_init_.Initialize()

	if err := validateNewReportSystemHealthReviewFiltersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ReportSystemHealthReviewFiltersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-nobl9.reportSystemHealthReview.ReportSystemHealthReviewFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewReportSystemHealthReviewFiltersOutputReference_Override(r ReportSystemHealthReviewFiltersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-nobl9.reportSystemHealthReview.ReportSystemHealthReviewFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference)SetInternalValue(val *ReportSystemHealthReviewFilters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference)SetProjects(val *[]*string) {
	if err := j.validateSetProjectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projects",
		val,
	)
}

func (j *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) PutLabel(value interface{}) {
	if err := r.validatePutLabelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putLabel",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) PutService(value interface{}) {
	if err := r.validatePutServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putService",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) PutSlo(value interface{}) {
	if err := r.validatePutSloParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putSlo",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) ResetLabel() {
	_jsii_.InvokeVoid(
		r,
		"resetLabel",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) ResetProjects() {
	_jsii_.InvokeVoid(
		r,
		"resetProjects",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) ResetService() {
	_jsii_.InvokeVoid(
		r,
		"resetService",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) ResetSlo() {
	_jsii_.InvokeVoid(
		r,
		"resetSlo",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReportSystemHealthReviewFiltersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

