package googleclouddeploydeploypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleclouddeploydeploypolicy/internal"
)

type GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference interface {
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
	EndDate() GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndDateOutputReference
	EndDateInput() *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndDate
	EndTime() GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndTimeOutputReference
	EndTimeInput() *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndTime
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	StartDate() GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartDateOutputReference
	StartDateInput() *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartDate
	StartTime() GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartTimeOutputReference
	StartTimeInput() *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartTime
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
	PutEndDate(value *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndDate)
	PutEndTime(value *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndTime)
	PutStartDate(value *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartDate)
	PutStartTime(value *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartTime)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference
type jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) EndDate() GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndDateOutputReference {
	var returns GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndDateOutputReference
	_jsii_.Get(
		j,
		"endDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) EndDateInput() *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndDate {
	var returns *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndDate
	_jsii_.Get(
		j,
		"endDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) EndTime() GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndTimeOutputReference {
	var returns GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndTimeOutputReference
	_jsii_.Get(
		j,
		"endTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) EndTimeInput() *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndTime {
	var returns *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndTime
	_jsii_.Get(
		j,
		"endTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) StartDate() GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartDateOutputReference {
	var returns GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartDateOutputReference
	_jsii_.Get(
		j,
		"startDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) StartDateInput() *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartDate {
	var returns *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartDate
	_jsii_.Get(
		j,
		"startDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) StartTime() GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartTimeOutputReference {
	var returns GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartTimeOutputReference
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) StartTimeInput() *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartTime {
	var returns *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartTime
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleClouddeployDeployPolicy.GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference_Override(g GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleClouddeployDeployPolicy.GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) PutEndDate(value *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndDate) {
	if err := g.validatePutEndDateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEndDate",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) PutEndTime(value *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndTime) {
	if err := g.validatePutEndTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEndTime",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) PutStartDate(value *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartDate) {
	if err := g.validatePutStartDateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStartDate",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) PutStartTime(value *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartTime) {
	if err := g.validatePutStartTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStartTime",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

