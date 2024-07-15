package googlesccmanagementfoldersecurityhealthanalyticscustommodule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlesccmanagementfoldersecurityhealthanalyticscustommodule/internal"
)

type GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference interface {
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
	CustomOutput() GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference
	CustomOutputInput() *GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutput
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfig
	SetInternalValue(val *GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfig)
	Predicate() GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigPredicateOutputReference
	PredicateInput() *GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigPredicate
	Recommendation() *string
	SetRecommendation(val *string)
	RecommendationInput() *string
	ResourceSelector() GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelectorOutputReference
	ResourceSelectorInput() *GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelector
	Severity() *string
	SetSeverity(val *string)
	SeverityInput() *string
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
	PutCustomOutput(value *GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutput)
	PutPredicate(value *GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigPredicate)
	PutResourceSelector(value *GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelector)
	ResetCustomOutput()
	ResetDescription()
	ResetPredicate()
	ResetRecommendation()
	ResetResourceSelector()
	ResetSeverity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference
type jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) CustomOutput() GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference {
	var returns GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference
	_jsii_.Get(
		j,
		"customOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) CustomOutputInput() *GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutput {
	var returns *GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutput
	_jsii_.Get(
		j,
		"customOutputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) InternalValue() *GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfig {
	var returns *GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) Predicate() GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigPredicateOutputReference {
	var returns GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigPredicateOutputReference
	_jsii_.Get(
		j,
		"predicate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) PredicateInput() *GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigPredicate {
	var returns *GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigPredicate
	_jsii_.Get(
		j,
		"predicateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) Recommendation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recommendation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) RecommendationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recommendationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) ResourceSelector() GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelectorOutputReference {
	var returns GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelectorOutputReference
	_jsii_.Get(
		j,
		"resourceSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) ResourceSelectorInput() *GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelector {
	var returns *GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelector
	_jsii_.Get(
		j,
		"resourceSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) Severity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) SeverityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleSccManagementFolderSecurityHealthAnalyticsCustomModule.GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference_Override(g GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleSccManagementFolderSecurityHealthAnalyticsCustomModule.GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference)SetInternalValue(val *GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference)SetRecommendation(val *string) {
	if err := j.validateSetRecommendationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recommendation",
		val,
	)
}

func (j *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference)SetSeverity(val *string) {
	if err := j.validateSetSeverityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"severity",
		val,
	)
}

func (j *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) PutCustomOutput(value *GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutput) {
	if err := g.validatePutCustomOutputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCustomOutput",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) PutPredicate(value *GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigPredicate) {
	if err := g.validatePutPredicateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPredicate",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) PutResourceSelector(value *GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelector) {
	if err := g.validatePutResourceSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putResourceSelector",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) ResetCustomOutput() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomOutput",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) ResetPredicate() {
	_jsii_.InvokeVoid(
		g,
		"resetPredicate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) ResetRecommendation() {
	_jsii_.InvokeVoid(
		g,
		"resetRecommendation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) ResetResourceSelector() {
	_jsii_.InvokeVoid(
		g,
		"resetResourceSelector",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) ResetSeverity() {
	_jsii_.InvokeVoid(
		g,
		"resetSeverity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

