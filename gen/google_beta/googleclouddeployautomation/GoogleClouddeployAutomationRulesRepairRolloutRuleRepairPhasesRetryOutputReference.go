package googleclouddeployautomation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleclouddeployautomation/internal"
)

type GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference interface {
	cdktf.ComplexObject
	Attempts() *string
	SetAttempts(val *string)
	AttemptsInput() *string
	BackoffMode() *string
	SetBackoffMode(val *string)
	BackoffModeInput() *string
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
	InternalValue() *GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetry
	SetInternalValue(val *GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetry)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Wait() *string
	SetWait(val *string)
	WaitInput() *string
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
	ResetBackoffMode()
	ResetWait()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference
type jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) Attempts() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) AttemptsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) BackoffMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backoffMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) BackoffModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backoffModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) InternalValue() *GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetry {
	var returns *GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetry
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) Wait() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wait",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) WaitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waitInput",
		&returns,
	)
	return returns
}


func NewGoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleClouddeployAutomation.GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference_Override(g GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleClouddeployAutomation.GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference)SetAttempts(val *string) {
	if err := j.validateSetAttemptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attempts",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference)SetBackoffMode(val *string) {
	if err := j.validateSetBackoffModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backoffMode",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference)SetInternalValue(val *GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetry) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference)SetWait(val *string) {
	if err := j.validateSetWaitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wait",
		val,
	)
}

func (g *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) ResetBackoffMode() {
	_jsii_.InvokeVoid(
		g,
		"resetBackoffMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) ResetWait() {
	_jsii_.InvokeVoid(
		g,
		"resetWait",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

