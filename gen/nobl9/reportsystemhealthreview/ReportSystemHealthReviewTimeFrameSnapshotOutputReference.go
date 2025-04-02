package reportsystemhealthreview

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/nobl9/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/nobl9/reportsystemhealthreview/internal"
)

type ReportSystemHealthReviewTimeFrameSnapshotOutputReference interface {
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
	DateTime() *string
	SetDateTime(val *string)
	DateTimeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *ReportSystemHealthReviewTimeFrameSnapshot
	SetInternalValue(val *ReportSystemHealthReviewTimeFrameSnapshot)
	Point() *string
	SetPoint(val *string)
	PointInput() *string
	Rrule() *string
	SetRrule(val *string)
	RruleInput() *string
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
	ResetDateTime()
	ResetRrule()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ReportSystemHealthReviewTimeFrameSnapshotOutputReference
type jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference) DateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference) DateTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference) InternalValue() *ReportSystemHealthReviewTimeFrameSnapshot {
	var returns *ReportSystemHealthReviewTimeFrameSnapshot
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference) Point() *string {
	var returns *string
	_jsii_.Get(
		j,
		"point",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference) PointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference) Rrule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rrule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference) RruleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rruleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewReportSystemHealthReviewTimeFrameSnapshotOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ReportSystemHealthReviewTimeFrameSnapshotOutputReference {
	_init_.Initialize()

	if err := validateNewReportSystemHealthReviewTimeFrameSnapshotOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-nobl9.reportSystemHealthReview.ReportSystemHealthReviewTimeFrameSnapshotOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewReportSystemHealthReviewTimeFrameSnapshotOutputReference_Override(r ReportSystemHealthReviewTimeFrameSnapshotOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-nobl9.reportSystemHealthReview.ReportSystemHealthReviewTimeFrameSnapshotOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference)SetDateTime(val *string) {
	if err := j.validateSetDateTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dateTime",
		val,
	)
}

func (j *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference)SetInternalValue(val *ReportSystemHealthReviewTimeFrameSnapshot) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference)SetPoint(val *string) {
	if err := j.validateSetPointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"point",
		val,
	)
}

func (j *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference)SetRrule(val *string) {
	if err := j.validateSetRruleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rrule",
		val,
	)
}

func (j *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (r *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference) ResetDateTime() {
	_jsii_.InvokeVoid(
		r,
		"resetDateTime",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference) ResetRrule() {
	_jsii_.InvokeVoid(
		r,
		"resetRrule",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (r *jsiiProxy_ReportSystemHealthReviewTimeFrameSnapshotOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

