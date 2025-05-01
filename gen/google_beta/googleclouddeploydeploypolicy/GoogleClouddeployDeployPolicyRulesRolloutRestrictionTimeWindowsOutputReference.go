package googleclouddeploydeploypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleclouddeploydeploypolicy/internal"
)

type GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference interface {
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
	InternalValue() *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindows
	SetInternalValue(val *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindows)
	OneTimeWindows() GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsList
	OneTimeWindowsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeZone() *string
	SetTimeZone(val *string)
	TimeZoneInput() *string
	WeeklyWindows() GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsList
	WeeklyWindowsInput() interface{}
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
	PutOneTimeWindows(value interface{})
	PutWeeklyWindows(value interface{})
	ResetOneTimeWindows()
	ResetWeeklyWindows()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference
type jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) InternalValue() *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindows {
	var returns *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindows
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) OneTimeWindows() GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsList {
	var returns GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsList
	_jsii_.Get(
		j,
		"oneTimeWindows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) OneTimeWindowsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oneTimeWindowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) WeeklyWindows() GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsList {
	var returns GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsList
	_jsii_.Get(
		j,
		"weeklyWindows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) WeeklyWindowsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"weeklyWindowsInput",
		&returns,
	)
	return returns
}


func NewGoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleClouddeployDeployPolicy.GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference_Override(g GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleClouddeployDeployPolicy.GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference)SetInternalValue(val *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindows) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference)SetTimeZone(val *string) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) PutOneTimeWindows(value interface{}) {
	if err := g.validatePutOneTimeWindowsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOneTimeWindows",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) PutWeeklyWindows(value interface{}) {
	if err := g.validatePutWeeklyWindowsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWeeklyWindows",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) ResetOneTimeWindows() {
	_jsii_.InvokeVoid(
		g,
		"resetOneTimeWindows",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) ResetWeeklyWindows() {
	_jsii_.InvokeVoid(
		g,
		"resetWeeklyWindows",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

