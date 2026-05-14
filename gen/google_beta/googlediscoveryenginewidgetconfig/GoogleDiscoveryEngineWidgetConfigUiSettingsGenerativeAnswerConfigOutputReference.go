package googlediscoveryenginewidgetconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlediscoveryenginewidgetconfig/internal"
)

type GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference interface {
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
	DisableRelatedQuestions() interface{}
	SetDisableRelatedQuestions(val interface{})
	DisableRelatedQuestionsInput() interface{}
	// Experimental.
	Fqn() *string
	IgnoreAdversarialQuery() interface{}
	SetIgnoreAdversarialQuery(val interface{})
	IgnoreAdversarialQueryInput() interface{}
	IgnoreLowRelevantContent() interface{}
	SetIgnoreLowRelevantContent(val interface{})
	IgnoreLowRelevantContentInput() interface{}
	IgnoreNonAnswerSeekingQuery() interface{}
	SetIgnoreNonAnswerSeekingQuery(val interface{})
	IgnoreNonAnswerSeekingQueryInput() interface{}
	ImageSource() *string
	SetImageSource(val *string)
	ImageSourceInput() *string
	InternalValue() *GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfig
	SetInternalValue(val *GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfig)
	LanguageCode() *string
	SetLanguageCode(val *string)
	LanguageCodeInput() *string
	MaxRephraseSteps() *float64
	SetMaxRephraseSteps(val *float64)
	MaxRephraseStepsInput() *float64
	ModelPromptPreamble() *string
	SetModelPromptPreamble(val *string)
	ModelPromptPreambleInput() *string
	ModelVersion() *string
	SetModelVersion(val *string)
	ModelVersionInput() *string
	ResultCount() *float64
	SetResultCount(val *float64)
	ResultCountInput() *float64
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
	ResetDisableRelatedQuestions()
	ResetIgnoreAdversarialQuery()
	ResetIgnoreLowRelevantContent()
	ResetIgnoreNonAnswerSeekingQuery()
	ResetImageSource()
	ResetLanguageCode()
	ResetMaxRephraseSteps()
	ResetModelPromptPreamble()
	ResetModelVersion()
	ResetResultCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference
type jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) DisableRelatedQuestions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableRelatedQuestions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) DisableRelatedQuestionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableRelatedQuestionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) IgnoreAdversarialQuery() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreAdversarialQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) IgnoreAdversarialQueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreAdversarialQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) IgnoreLowRelevantContent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreLowRelevantContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) IgnoreLowRelevantContentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreLowRelevantContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) IgnoreNonAnswerSeekingQuery() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreNonAnswerSeekingQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) IgnoreNonAnswerSeekingQueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreNonAnswerSeekingQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ImageSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ImageSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) InternalValue() *GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfig {
	var returns *GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) LanguageCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) LanguageCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) MaxRephraseSteps() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRephraseSteps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) MaxRephraseStepsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRephraseStepsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ModelPromptPreamble() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelPromptPreamble",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ModelPromptPreambleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelPromptPreambleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ModelVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ModelVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ResultCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"resultCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ResultCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"resultCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDiscoveryEngineWidgetConfig.GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference_Override(g GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDiscoveryEngineWidgetConfig.GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference)SetDisableRelatedQuestions(val interface{}) {
	if err := j.validateSetDisableRelatedQuestionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableRelatedQuestions",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference)SetIgnoreAdversarialQuery(val interface{}) {
	if err := j.validateSetIgnoreAdversarialQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreAdversarialQuery",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference)SetIgnoreLowRelevantContent(val interface{}) {
	if err := j.validateSetIgnoreLowRelevantContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreLowRelevantContent",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference)SetIgnoreNonAnswerSeekingQuery(val interface{}) {
	if err := j.validateSetIgnoreNonAnswerSeekingQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreNonAnswerSeekingQuery",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference)SetImageSource(val *string) {
	if err := j.validateSetImageSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageSource",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference)SetInternalValue(val *GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference)SetLanguageCode(val *string) {
	if err := j.validateSetLanguageCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"languageCode",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference)SetMaxRephraseSteps(val *float64) {
	if err := j.validateSetMaxRephraseStepsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRephraseSteps",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference)SetModelPromptPreamble(val *string) {
	if err := j.validateSetModelPromptPreambleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelPromptPreamble",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference)SetModelVersion(val *string) {
	if err := j.validateSetModelVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference)SetResultCount(val *float64) {
	if err := j.validateSetResultCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resultCount",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ResetDisableRelatedQuestions() {
	_jsii_.InvokeVoid(
		g,
		"resetDisableRelatedQuestions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ResetIgnoreAdversarialQuery() {
	_jsii_.InvokeVoid(
		g,
		"resetIgnoreAdversarialQuery",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ResetIgnoreLowRelevantContent() {
	_jsii_.InvokeVoid(
		g,
		"resetIgnoreLowRelevantContent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ResetIgnoreNonAnswerSeekingQuery() {
	_jsii_.InvokeVoid(
		g,
		"resetIgnoreNonAnswerSeekingQuery",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ResetImageSource() {
	_jsii_.InvokeVoid(
		g,
		"resetImageSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ResetLanguageCode() {
	_jsii_.InvokeVoid(
		g,
		"resetLanguageCode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ResetMaxRephraseSteps() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxRephraseSteps",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ResetModelPromptPreamble() {
	_jsii_.InvokeVoid(
		g,
		"resetModelPromptPreamble",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ResetModelVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetModelVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ResetResultCount() {
	_jsii_.InvokeVoid(
		g,
		"resetResultCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

