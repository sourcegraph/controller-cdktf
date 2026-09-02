package googlediscoveryenginewidgetconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlediscoveryenginewidgetconfig/internal"
)

type GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference interface {
	cdktn.ComplexObject
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
	DataStoreUiConfigs() GoogleDiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigsList
	DataStoreUiConfigsInput() interface{}
	DefaultSearchRequestOrderBy() *string
	SetDefaultSearchRequestOrderBy(val *string)
	DefaultSearchRequestOrderByInput() *string
	DisableUserEventsCollection() interface{}
	SetDisableUserEventsCollection(val interface{})
	DisableUserEventsCollectionInput() interface{}
	EnableAutocomplete() interface{}
	SetEnableAutocomplete(val interface{})
	EnableAutocompleteInput() interface{}
	EnableCreateAgentButton() interface{}
	SetEnableCreateAgentButton(val interface{})
	EnableCreateAgentButtonInput() interface{}
	EnablePeopleSearch() interface{}
	SetEnablePeopleSearch(val interface{})
	EnablePeopleSearchInput() interface{}
	EnableQualityFeedback() interface{}
	SetEnableQualityFeedback(val interface{})
	EnableQualityFeedbackInput() interface{}
	EnableSafeSearch() interface{}
	SetEnableSafeSearch(val interface{})
	EnableSafeSearchInput() interface{}
	EnableSearchAsYouType() interface{}
	SetEnableSearchAsYouType(val interface{})
	EnableSearchAsYouTypeInput() interface{}
	EnableVisualContentSummary() interface{}
	SetEnableVisualContentSummary(val interface{})
	EnableVisualContentSummaryInput() interface{}
	// Experimental.
	Fqn() *string
	GenerativeAnswerConfig() GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference
	GenerativeAnswerConfigInput() *GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfig
	InteractionType() *string
	SetInteractionType(val *string)
	InteractionTypeInput() *string
	InternalValue() *GoogleDiscoveryEngineWidgetConfigUiSettings
	SetInternalValue(val *GoogleDiscoveryEngineWidgetConfigUiSettings)
	ResultDescriptionType() *string
	SetResultDescriptionType(val *string)
	ResultDescriptionTypeInput() *string
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
	PutDataStoreUiConfigs(value interface{})
	PutGenerativeAnswerConfig(value *GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfig)
	ResetDataStoreUiConfigs()
	ResetDefaultSearchRequestOrderBy()
	ResetDisableUserEventsCollection()
	ResetEnableAutocomplete()
	ResetEnableCreateAgentButton()
	ResetEnablePeopleSearch()
	ResetEnableQualityFeedback()
	ResetEnableSafeSearch()
	ResetEnableSearchAsYouType()
	ResetEnableVisualContentSummary()
	ResetGenerativeAnswerConfig()
	ResetInteractionType()
	ResetResultDescriptionType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference
type jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) DataStoreUiConfigs() GoogleDiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigsList {
	var returns GoogleDiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigsList
	_jsii_.Get(
		j,
		"dataStoreUiConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) DataStoreUiConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataStoreUiConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) DefaultSearchRequestOrderBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSearchRequestOrderBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) DefaultSearchRequestOrderByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSearchRequestOrderByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) DisableUserEventsCollection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableUserEventsCollection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) DisableUserEventsCollectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableUserEventsCollectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) EnableAutocomplete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutocomplete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) EnableAutocompleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutocompleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) EnableCreateAgentButton() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCreateAgentButton",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) EnableCreateAgentButtonInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCreateAgentButtonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) EnablePeopleSearch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePeopleSearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) EnablePeopleSearchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePeopleSearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) EnableQualityFeedback() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableQualityFeedback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) EnableQualityFeedbackInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableQualityFeedbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) EnableSafeSearch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSafeSearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) EnableSafeSearchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSafeSearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) EnableSearchAsYouType() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSearchAsYouType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) EnableSearchAsYouTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSearchAsYouTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) EnableVisualContentSummary() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableVisualContentSummary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) EnableVisualContentSummaryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableVisualContentSummaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) GenerativeAnswerConfig() GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference {
	var returns GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference
	_jsii_.Get(
		j,
		"generativeAnswerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) GenerativeAnswerConfigInput() *GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfig {
	var returns *GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfig
	_jsii_.Get(
		j,
		"generativeAnswerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) InteractionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interactionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) InteractionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interactionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) InternalValue() *GoogleDiscoveryEngineWidgetConfigUiSettings {
	var returns *GoogleDiscoveryEngineWidgetConfigUiSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) ResultDescriptionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultDescriptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) ResultDescriptionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultDescriptionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDiscoveryEngineWidgetConfigUiSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDiscoveryEngineWidgetConfig.GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference_Override(g GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDiscoveryEngineWidgetConfig.GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference)SetDefaultSearchRequestOrderBy(val *string) {
	if err := j.validateSetDefaultSearchRequestOrderByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultSearchRequestOrderBy",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference)SetDisableUserEventsCollection(val interface{}) {
	if err := j.validateSetDisableUserEventsCollectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableUserEventsCollection",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference)SetEnableAutocomplete(val interface{}) {
	if err := j.validateSetEnableAutocompleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAutocomplete",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference)SetEnableCreateAgentButton(val interface{}) {
	if err := j.validateSetEnableCreateAgentButtonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableCreateAgentButton",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference)SetEnablePeopleSearch(val interface{}) {
	if err := j.validateSetEnablePeopleSearchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePeopleSearch",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference)SetEnableQualityFeedback(val interface{}) {
	if err := j.validateSetEnableQualityFeedbackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableQualityFeedback",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference)SetEnableSafeSearch(val interface{}) {
	if err := j.validateSetEnableSafeSearchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSafeSearch",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference)SetEnableSearchAsYouType(val interface{}) {
	if err := j.validateSetEnableSearchAsYouTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSearchAsYouType",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference)SetEnableVisualContentSummary(val interface{}) {
	if err := j.validateSetEnableVisualContentSummaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableVisualContentSummary",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference)SetInteractionType(val *string) {
	if err := j.validateSetInteractionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interactionType",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference)SetInternalValue(val *GoogleDiscoveryEngineWidgetConfigUiSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference)SetResultDescriptionType(val *string) {
	if err := j.validateSetResultDescriptionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resultDescriptionType",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) PutDataStoreUiConfigs(value interface{}) {
	if err := g.validatePutDataStoreUiConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDataStoreUiConfigs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) PutGenerativeAnswerConfig(value *GoogleDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfig) {
	if err := g.validatePutGenerativeAnswerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGenerativeAnswerConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) ResetDataStoreUiConfigs() {
	_jsii_.InvokeVoid(
		g,
		"resetDataStoreUiConfigs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) ResetDefaultSearchRequestOrderBy() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultSearchRequestOrderBy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) ResetDisableUserEventsCollection() {
	_jsii_.InvokeVoid(
		g,
		"resetDisableUserEventsCollection",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) ResetEnableAutocomplete() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableAutocomplete",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) ResetEnableCreateAgentButton() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableCreateAgentButton",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) ResetEnablePeopleSearch() {
	_jsii_.InvokeVoid(
		g,
		"resetEnablePeopleSearch",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) ResetEnableQualityFeedback() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableQualityFeedback",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) ResetEnableSafeSearch() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableSafeSearch",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) ResetEnableSearchAsYouType() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableSearchAsYouType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) ResetEnableVisualContentSummary() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableVisualContentSummary",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) ResetGenerativeAnswerConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetGenerativeAnswerConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) ResetInteractionType() {
	_jsii_.InvokeVoid(
		g,
		"resetInteractionType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) ResetResultDescriptionType() {
	_jsii_.InvokeVoid(
		g,
		"resetResultDescriptionType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

