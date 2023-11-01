package googlesccorganizationcustommodule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlesccorganizationcustommodule/internal"
)

type GoogleSccOrganizationCustomModuleCustomConfigOutputReference interface {
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
	CustomOutput() GoogleSccOrganizationCustomModuleCustomConfigCustomOutputOutputReference
	CustomOutputInput() *GoogleSccOrganizationCustomModuleCustomConfigCustomOutput
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleSccOrganizationCustomModuleCustomConfig
	SetInternalValue(val *GoogleSccOrganizationCustomModuleCustomConfig)
	Predicate() GoogleSccOrganizationCustomModuleCustomConfigPredicateOutputReference
	PredicateInput() *GoogleSccOrganizationCustomModuleCustomConfigPredicate
	Recommendation() *string
	SetRecommendation(val *string)
	RecommendationInput() *string
	ResourceSelector() GoogleSccOrganizationCustomModuleCustomConfigResourceSelectorOutputReference
	ResourceSelectorInput() *GoogleSccOrganizationCustomModuleCustomConfigResourceSelector
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
	PutCustomOutput(value *GoogleSccOrganizationCustomModuleCustomConfigCustomOutput)
	PutPredicate(value *GoogleSccOrganizationCustomModuleCustomConfigPredicate)
	PutResourceSelector(value *GoogleSccOrganizationCustomModuleCustomConfigResourceSelector)
	ResetCustomOutput()
	ResetDescription()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleSccOrganizationCustomModuleCustomConfigOutputReference
type jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) CustomOutput() GoogleSccOrganizationCustomModuleCustomConfigCustomOutputOutputReference {
	var returns GoogleSccOrganizationCustomModuleCustomConfigCustomOutputOutputReference
	_jsii_.Get(
		j,
		"customOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) CustomOutputInput() *GoogleSccOrganizationCustomModuleCustomConfigCustomOutput {
	var returns *GoogleSccOrganizationCustomModuleCustomConfigCustomOutput
	_jsii_.Get(
		j,
		"customOutputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) InternalValue() *GoogleSccOrganizationCustomModuleCustomConfig {
	var returns *GoogleSccOrganizationCustomModuleCustomConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) Predicate() GoogleSccOrganizationCustomModuleCustomConfigPredicateOutputReference {
	var returns GoogleSccOrganizationCustomModuleCustomConfigPredicateOutputReference
	_jsii_.Get(
		j,
		"predicate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) PredicateInput() *GoogleSccOrganizationCustomModuleCustomConfigPredicate {
	var returns *GoogleSccOrganizationCustomModuleCustomConfigPredicate
	_jsii_.Get(
		j,
		"predicateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) Recommendation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recommendation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) RecommendationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recommendationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) ResourceSelector() GoogleSccOrganizationCustomModuleCustomConfigResourceSelectorOutputReference {
	var returns GoogleSccOrganizationCustomModuleCustomConfigResourceSelectorOutputReference
	_jsii_.Get(
		j,
		"resourceSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) ResourceSelectorInput() *GoogleSccOrganizationCustomModuleCustomConfigResourceSelector {
	var returns *GoogleSccOrganizationCustomModuleCustomConfigResourceSelector
	_jsii_.Get(
		j,
		"resourceSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) Severity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) SeverityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleSccOrganizationCustomModuleCustomConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleSccOrganizationCustomModuleCustomConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleSccOrganizationCustomModuleCustomConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleSccOrganizationCustomModule.GoogleSccOrganizationCustomModuleCustomConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleSccOrganizationCustomModuleCustomConfigOutputReference_Override(g GoogleSccOrganizationCustomModuleCustomConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleSccOrganizationCustomModule.GoogleSccOrganizationCustomModuleCustomConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference)SetInternalValue(val *GoogleSccOrganizationCustomModuleCustomConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference)SetRecommendation(val *string) {
	if err := j.validateSetRecommendationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recommendation",
		val,
	)
}

func (j *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference)SetSeverity(val *string) {
	if err := j.validateSetSeverityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"severity",
		val,
	)
}

func (j *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) PutCustomOutput(value *GoogleSccOrganizationCustomModuleCustomConfigCustomOutput) {
	if err := g.validatePutCustomOutputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCustomOutput",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) PutPredicate(value *GoogleSccOrganizationCustomModuleCustomConfigPredicate) {
	if err := g.validatePutPredicateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPredicate",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) PutResourceSelector(value *GoogleSccOrganizationCustomModuleCustomConfigResourceSelector) {
	if err := g.validatePutResourceSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putResourceSelector",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) ResetCustomOutput() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomOutput",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleSccOrganizationCustomModuleCustomConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

