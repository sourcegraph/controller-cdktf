package discoveryenginesearchengine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/discoveryenginesearchengine/internal"
)

type DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference interface {
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
	DisablePrivateKgAutoComplete() interface{}
	SetDisablePrivateKgAutoComplete(val interface{})
	DisablePrivateKgAutoCompleteInput() interface{}
	DisablePrivateKgEnrichment() interface{}
	SetDisablePrivateKgEnrichment(val interface{})
	DisablePrivateKgEnrichmentInput() interface{}
	DisablePrivateKgQueryUiChips() interface{}
	SetDisablePrivateKgQueryUiChips(val interface{})
	DisablePrivateKgQueryUiChipsInput() interface{}
	DisablePrivateKgQueryUnderstanding() interface{}
	SetDisablePrivateKgQueryUnderstanding(val interface{})
	DisablePrivateKgQueryUnderstandingInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfig
	SetInternalValue(val *DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfig)
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
	ResetDisablePrivateKgAutoComplete()
	ResetDisablePrivateKgEnrichment()
	ResetDisablePrivateKgQueryUiChips()
	ResetDisablePrivateKgQueryUnderstanding()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference
type jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) DisablePrivateKgAutoComplete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePrivateKgAutoComplete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) DisablePrivateKgAutoCompleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePrivateKgAutoCompleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) DisablePrivateKgEnrichment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePrivateKgEnrichment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) DisablePrivateKgEnrichmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePrivateKgEnrichmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) DisablePrivateKgQueryUiChips() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePrivateKgQueryUiChips",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) DisablePrivateKgQueryUiChipsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePrivateKgQueryUiChipsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) DisablePrivateKgQueryUnderstanding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePrivateKgQueryUnderstanding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) DisablePrivateKgQueryUnderstandingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePrivateKgQueryUnderstandingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) InternalValue() *DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfig {
	var returns *DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.discoveryEngineSearchEngine.DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference_Override(d DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.discoveryEngineSearchEngine.DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference)SetDisablePrivateKgAutoComplete(val interface{}) {
	if err := j.validateSetDisablePrivateKgAutoCompleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disablePrivateKgAutoComplete",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference)SetDisablePrivateKgEnrichment(val interface{}) {
	if err := j.validateSetDisablePrivateKgEnrichmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disablePrivateKgEnrichment",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference)SetDisablePrivateKgQueryUiChips(val interface{}) {
	if err := j.validateSetDisablePrivateKgQueryUiChipsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disablePrivateKgQueryUiChips",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference)SetDisablePrivateKgQueryUnderstanding(val interface{}) {
	if err := j.validateSetDisablePrivateKgQueryUnderstandingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disablePrivateKgQueryUnderstanding",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference)SetInternalValue(val *DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) ResetDisablePrivateKgAutoComplete() {
	_jsii_.InvokeVoid(
		d,
		"resetDisablePrivateKgAutoComplete",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) ResetDisablePrivateKgEnrichment() {
	_jsii_.InvokeVoid(
		d,
		"resetDisablePrivateKgEnrichment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) ResetDisablePrivateKgQueryUiChips() {
	_jsii_.InvokeVoid(
		d,
		"resetDisablePrivateKgQueryUiChips",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) ResetDisablePrivateKgQueryUnderstanding() {
	_jsii_.InvokeVoid(
		d,
		"resetDisablePrivateKgQueryUnderstanding",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

