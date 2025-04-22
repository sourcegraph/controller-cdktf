package googleclouddeployautomation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleclouddeployautomation/internal"
)

type GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference interface {
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
	DestinationPhase() *string
	SetDestinationPhase(val *string)
	DestinationPhaseInput() *string
	DestinationTargetId() *string
	SetDestinationTargetId(val *string)
	DestinationTargetIdInput() *string
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() *GoogleClouddeployAutomationRulesTimedPromoteReleaseRule
	SetInternalValue(val *GoogleClouddeployAutomationRulesTimedPromoteReleaseRule)
	Schedule() *string
	SetSchedule(val *string)
	ScheduleInput() *string
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
	ResetDestinationPhase()
	ResetDestinationTargetId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference
type jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) DestinationPhase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPhase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) DestinationPhaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPhaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) DestinationTargetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationTargetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) DestinationTargetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationTargetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) InternalValue() *GoogleClouddeployAutomationRulesTimedPromoteReleaseRule {
	var returns *GoogleClouddeployAutomationRulesTimedPromoteReleaseRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) Schedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) ScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}


func NewGoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleClouddeployAutomation.GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference_Override(g GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleClouddeployAutomation.GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference)SetDestinationPhase(val *string) {
	if err := j.validateSetDestinationPhaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPhase",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference)SetDestinationTargetId(val *string) {
	if err := j.validateSetDestinationTargetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationTargetId",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference)SetInternalValue(val *GoogleClouddeployAutomationRulesTimedPromoteReleaseRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference)SetSchedule(val *string) {
	if err := j.validateSetScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference)SetTimeZone(val *string) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

func (g *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) ResetDestinationPhase() {
	_jsii_.InvokeVoid(
		g,
		"resetDestinationPhase",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) ResetDestinationTargetId() {
	_jsii_.InvokeVoid(
		g,
		"resetDestinationTargetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

