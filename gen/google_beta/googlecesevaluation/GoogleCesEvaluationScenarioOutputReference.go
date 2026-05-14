package googlecesevaluation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecesevaluation/internal"
)

type GoogleCesEvaluationScenarioOutputReference interface {
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
	EvaluationExpectations() *[]*string
	SetEvaluationExpectations(val *[]*string)
	EvaluationExpectationsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleCesEvaluationScenario
	SetInternalValue(val *GoogleCesEvaluationScenario)
	MaxTurns() *float64
	SetMaxTurns(val *float64)
	MaxTurnsInput() *float64
	Rubrics() *[]*string
	SetRubrics(val *[]*string)
	RubricsInput() *[]*string
	ScenarioExpectations() GoogleCesEvaluationScenarioScenarioExpectationsList
	ScenarioExpectationsInput() interface{}
	Task() *string
	SetTask(val *string)
	TaskCompletionBehavior() *string
	SetTaskCompletionBehavior(val *string)
	TaskCompletionBehaviorInput() *string
	TaskInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserFacts() GoogleCesEvaluationScenarioUserFactsList
	UserFactsInput() interface{}
	UserGoalBehavior() *string
	SetUserGoalBehavior(val *string)
	UserGoalBehaviorInput() *string
	VariableOverrides() *map[string]*string
	SetVariableOverrides(val *map[string]*string)
	VariableOverridesInput() *map[string]*string
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
	PutScenarioExpectations(value interface{})
	PutUserFacts(value interface{})
	ResetEvaluationExpectations()
	ResetMaxTurns()
	ResetTaskCompletionBehavior()
	ResetUserFacts()
	ResetUserGoalBehavior()
	ResetVariableOverrides()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesEvaluationScenarioOutputReference
type jsiiProxy_GoogleCesEvaluationScenarioOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) EvaluationExpectations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"evaluationExpectations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) EvaluationExpectationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"evaluationExpectationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) InternalValue() *GoogleCesEvaluationScenario {
	var returns *GoogleCesEvaluationScenario
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) MaxTurns() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTurns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) MaxTurnsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTurnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) Rubrics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rubrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) RubricsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rubricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) ScenarioExpectations() GoogleCesEvaluationScenarioScenarioExpectationsList {
	var returns GoogleCesEvaluationScenarioScenarioExpectationsList
	_jsii_.Get(
		j,
		"scenarioExpectations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) ScenarioExpectationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scenarioExpectationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) Task() *string {
	var returns *string
	_jsii_.Get(
		j,
		"task",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) TaskCompletionBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskCompletionBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) TaskCompletionBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskCompletionBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) TaskInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) UserFacts() GoogleCesEvaluationScenarioUserFactsList {
	var returns GoogleCesEvaluationScenarioUserFactsList
	_jsii_.Get(
		j,
		"userFacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) UserFactsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userFactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) UserGoalBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userGoalBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) UserGoalBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userGoalBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) VariableOverrides() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"variableOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) VariableOverridesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"variableOverridesInput",
		&returns,
	)
	return returns
}


func NewGoogleCesEvaluationScenarioOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCesEvaluationScenarioOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesEvaluationScenarioOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesEvaluationScenarioOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCesEvaluation.GoogleCesEvaluationScenarioOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesEvaluationScenarioOutputReference_Override(g GoogleCesEvaluationScenarioOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCesEvaluation.GoogleCesEvaluationScenarioOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference)SetEvaluationExpectations(val *[]*string) {
	if err := j.validateSetEvaluationExpectationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluationExpectations",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference)SetInternalValue(val *GoogleCesEvaluationScenario) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference)SetMaxTurns(val *float64) {
	if err := j.validateSetMaxTurnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTurns",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference)SetRubrics(val *[]*string) {
	if err := j.validateSetRubricsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rubrics",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference)SetTask(val *string) {
	if err := j.validateSetTaskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"task",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference)SetTaskCompletionBehavior(val *string) {
	if err := j.validateSetTaskCompletionBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskCompletionBehavior",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference)SetUserGoalBehavior(val *string) {
	if err := j.validateSetUserGoalBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userGoalBehavior",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioOutputReference)SetVariableOverrides(val *map[string]*string) {
	if err := j.validateSetVariableOverridesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"variableOverrides",
		val,
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) PutScenarioExpectations(value interface{}) {
	if err := g.validatePutScenarioExpectationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putScenarioExpectations",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) PutUserFacts(value interface{}) {
	if err := g.validatePutUserFactsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUserFacts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) ResetEvaluationExpectations() {
	_jsii_.InvokeVoid(
		g,
		"resetEvaluationExpectations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) ResetMaxTurns() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxTurns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) ResetTaskCompletionBehavior() {
	_jsii_.InvokeVoid(
		g,
		"resetTaskCompletionBehavior",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) ResetUserFacts() {
	_jsii_.InvokeVoid(
		g,
		"resetUserFacts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) ResetUserGoalBehavior() {
	_jsii_.InvokeVoid(
		g,
		"resetUserGoalBehavior",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) ResetVariableOverrides() {
	_jsii_.InvokeVoid(
		g,
		"resetVariableOverrides",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

