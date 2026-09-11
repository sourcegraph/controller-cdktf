package googlecomputefuturereservation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputefuturereservation/internal"
)

type GoogleComputeFutureReservationCommitmentInfoOutputReference interface {
	cdktn.ComplexObject
	CommitmentName() *string
	SetCommitmentName(val *string)
	CommitmentNameInput() *string
	CommitmentPlan() *string
	SetCommitmentPlan(val *string)
	CommitmentPlanInput() *string
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
	InternalValue() *GoogleComputeFutureReservationCommitmentInfo
	SetInternalValue(val *GoogleComputeFutureReservationCommitmentInfo)
	PreviousCommitmentTerms() *string
	SetPreviousCommitmentTerms(val *string)
	PreviousCommitmentTermsInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	ResetCommitmentName()
	ResetCommitmentPlan()
	ResetPreviousCommitmentTerms()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeFutureReservationCommitmentInfoOutputReference
type jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference) CommitmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference) CommitmentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitmentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference) CommitmentPlan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitmentPlan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference) CommitmentPlanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitmentPlanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference) InternalValue() *GoogleComputeFutureReservationCommitmentInfo {
	var returns *GoogleComputeFutureReservationCommitmentInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference) PreviousCommitmentTerms() *string {
	var returns *string
	_jsii_.Get(
		j,
		"previousCommitmentTerms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference) PreviousCommitmentTermsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"previousCommitmentTermsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeFutureReservationCommitmentInfoOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleComputeFutureReservationCommitmentInfoOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeFutureReservationCommitmentInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeFutureReservation.GoogleComputeFutureReservationCommitmentInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeFutureReservationCommitmentInfoOutputReference_Override(g GoogleComputeFutureReservationCommitmentInfoOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeFutureReservation.GoogleComputeFutureReservationCommitmentInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference)SetCommitmentName(val *string) {
	if err := j.validateSetCommitmentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commitmentName",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference)SetCommitmentPlan(val *string) {
	if err := j.validateSetCommitmentPlanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commitmentPlan",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference)SetInternalValue(val *GoogleComputeFutureReservationCommitmentInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference)SetPreviousCommitmentTerms(val *string) {
	if err := j.validateSetPreviousCommitmentTermsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"previousCommitmentTerms",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference) ResetCommitmentName() {
	_jsii_.InvokeVoid(
		g,
		"resetCommitmentName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference) ResetCommitmentPlan() {
	_jsii_.InvokeVoid(
		g,
		"resetCommitmentPlan",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference) ResetPreviousCommitmentTerms() {
	_jsii_.InvokeVoid(
		g,
		"resetPreviousCommitmentTerms",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeFutureReservationCommitmentInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

