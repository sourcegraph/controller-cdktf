package googlecontactcenterinsightsanalysisrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecontactcenterinsightsanalysisrule/internal"
)

type GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference interface {
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
	InternalValue() *GoogleContactCenterInsightsAnalysisRuleAnnotatorSelector
	SetInternalValue(val *GoogleContactCenterInsightsAnalysisRuleAnnotatorSelector)
	IssueModels() *[]*string
	SetIssueModels(val *[]*string)
	IssueModelsInput() *[]*string
	PhraseMatchers() *[]*string
	SetPhraseMatchers(val *[]*string)
	PhraseMatchersInput() *[]*string
	QaConfig() GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorQaConfigOutputReference
	QaConfigInput() *GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorQaConfig
	RunEntityAnnotator() interface{}
	SetRunEntityAnnotator(val interface{})
	RunEntityAnnotatorInput() interface{}
	RunIntentAnnotator() interface{}
	SetRunIntentAnnotator(val interface{})
	RunIntentAnnotatorInput() interface{}
	RunInterruptionAnnotator() interface{}
	SetRunInterruptionAnnotator(val interface{})
	RunInterruptionAnnotatorInput() interface{}
	RunIssueModelAnnotator() interface{}
	SetRunIssueModelAnnotator(val interface{})
	RunIssueModelAnnotatorInput() interface{}
	RunPhraseMatcherAnnotator() interface{}
	SetRunPhraseMatcherAnnotator(val interface{})
	RunPhraseMatcherAnnotatorInput() interface{}
	RunQaAnnotator() interface{}
	SetRunQaAnnotator(val interface{})
	RunQaAnnotatorInput() interface{}
	RunSentimentAnnotator() interface{}
	SetRunSentimentAnnotator(val interface{})
	RunSentimentAnnotatorInput() interface{}
	RunSilenceAnnotator() interface{}
	SetRunSilenceAnnotator(val interface{})
	RunSilenceAnnotatorInput() interface{}
	RunSummarizationAnnotator() interface{}
	SetRunSummarizationAnnotator(val interface{})
	RunSummarizationAnnotatorInput() interface{}
	SummarizationConfig() GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorSummarizationConfigOutputReference
	SummarizationConfigInput() *GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorSummarizationConfig
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
	PutQaConfig(value *GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorQaConfig)
	PutSummarizationConfig(value *GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorSummarizationConfig)
	ResetIssueModels()
	ResetPhraseMatchers()
	ResetQaConfig()
	ResetRunEntityAnnotator()
	ResetRunIntentAnnotator()
	ResetRunInterruptionAnnotator()
	ResetRunIssueModelAnnotator()
	ResetRunPhraseMatcherAnnotator()
	ResetRunQaAnnotator()
	ResetRunSentimentAnnotator()
	ResetRunSilenceAnnotator()
	ResetRunSummarizationAnnotator()
	ResetSummarizationConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference
type jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) InternalValue() *GoogleContactCenterInsightsAnalysisRuleAnnotatorSelector {
	var returns *GoogleContactCenterInsightsAnalysisRuleAnnotatorSelector
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) IssueModels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"issueModels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) IssueModelsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"issueModelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) PhraseMatchers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"phraseMatchers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) PhraseMatchersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"phraseMatchersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) QaConfig() GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorQaConfigOutputReference {
	var returns GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorQaConfigOutputReference
	_jsii_.Get(
		j,
		"qaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) QaConfigInput() *GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorQaConfig {
	var returns *GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorQaConfig
	_jsii_.Get(
		j,
		"qaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunEntityAnnotator() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runEntityAnnotator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunEntityAnnotatorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runEntityAnnotatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunIntentAnnotator() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runIntentAnnotator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunIntentAnnotatorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runIntentAnnotatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunInterruptionAnnotator() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runInterruptionAnnotator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunInterruptionAnnotatorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runInterruptionAnnotatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunIssueModelAnnotator() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runIssueModelAnnotator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunIssueModelAnnotatorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runIssueModelAnnotatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunPhraseMatcherAnnotator() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runPhraseMatcherAnnotator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunPhraseMatcherAnnotatorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runPhraseMatcherAnnotatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunQaAnnotator() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runQaAnnotator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunQaAnnotatorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runQaAnnotatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunSentimentAnnotator() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runSentimentAnnotator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunSentimentAnnotatorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runSentimentAnnotatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunSilenceAnnotator() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runSilenceAnnotator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunSilenceAnnotatorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runSilenceAnnotatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunSummarizationAnnotator() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runSummarizationAnnotator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunSummarizationAnnotatorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runSummarizationAnnotatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) SummarizationConfig() GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorSummarizationConfigOutputReference {
	var returns GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorSummarizationConfigOutputReference
	_jsii_.Get(
		j,
		"summarizationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) SummarizationConfigInput() *GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorSummarizationConfig {
	var returns *GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorSummarizationConfig
	_jsii_.Get(
		j,
		"summarizationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleContactCenterInsightsAnalysisRule.GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference_Override(g GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleContactCenterInsightsAnalysisRule.GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetInternalValue(val *GoogleContactCenterInsightsAnalysisRuleAnnotatorSelector) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetIssueModels(val *[]*string) {
	if err := j.validateSetIssueModelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issueModels",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetPhraseMatchers(val *[]*string) {
	if err := j.validateSetPhraseMatchersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phraseMatchers",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetRunEntityAnnotator(val interface{}) {
	if err := j.validateSetRunEntityAnnotatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runEntityAnnotator",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetRunIntentAnnotator(val interface{}) {
	if err := j.validateSetRunIntentAnnotatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runIntentAnnotator",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetRunInterruptionAnnotator(val interface{}) {
	if err := j.validateSetRunInterruptionAnnotatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runInterruptionAnnotator",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetRunIssueModelAnnotator(val interface{}) {
	if err := j.validateSetRunIssueModelAnnotatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runIssueModelAnnotator",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetRunPhraseMatcherAnnotator(val interface{}) {
	if err := j.validateSetRunPhraseMatcherAnnotatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runPhraseMatcherAnnotator",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetRunQaAnnotator(val interface{}) {
	if err := j.validateSetRunQaAnnotatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runQaAnnotator",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetRunSentimentAnnotator(val interface{}) {
	if err := j.validateSetRunSentimentAnnotatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runSentimentAnnotator",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetRunSilenceAnnotator(val interface{}) {
	if err := j.validateSetRunSilenceAnnotatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runSilenceAnnotator",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetRunSummarizationAnnotator(val interface{}) {
	if err := j.validateSetRunSummarizationAnnotatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runSummarizationAnnotator",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) PutQaConfig(value *GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorQaConfig) {
	if err := g.validatePutQaConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putQaConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) PutSummarizationConfig(value *GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorSummarizationConfig) {
	if err := g.validatePutSummarizationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSummarizationConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ResetIssueModels() {
	_jsii_.InvokeVoid(
		g,
		"resetIssueModels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ResetPhraseMatchers() {
	_jsii_.InvokeVoid(
		g,
		"resetPhraseMatchers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ResetQaConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetQaConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ResetRunEntityAnnotator() {
	_jsii_.InvokeVoid(
		g,
		"resetRunEntityAnnotator",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ResetRunIntentAnnotator() {
	_jsii_.InvokeVoid(
		g,
		"resetRunIntentAnnotator",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ResetRunInterruptionAnnotator() {
	_jsii_.InvokeVoid(
		g,
		"resetRunInterruptionAnnotator",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ResetRunIssueModelAnnotator() {
	_jsii_.InvokeVoid(
		g,
		"resetRunIssueModelAnnotator",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ResetRunPhraseMatcherAnnotator() {
	_jsii_.InvokeVoid(
		g,
		"resetRunPhraseMatcherAnnotator",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ResetRunQaAnnotator() {
	_jsii_.InvokeVoid(
		g,
		"resetRunQaAnnotator",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ResetRunSentimentAnnotator() {
	_jsii_.InvokeVoid(
		g,
		"resetRunSentimentAnnotator",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ResetRunSilenceAnnotator() {
	_jsii_.InvokeVoid(
		g,
		"resetRunSilenceAnnotator",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ResetRunSummarizationAnnotator() {
	_jsii_.InvokeVoid(
		g,
		"resetRunSummarizationAnnotator",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ResetSummarizationConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSummarizationConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

