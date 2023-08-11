package googledataplexasset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataplexasset/internal"
)

type GoogleDataplexAssetDiscoverySpecOutputReference interface {
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
	CsvOptions() GoogleDataplexAssetDiscoverySpecCsvOptionsOutputReference
	CsvOptionsInput() *GoogleDataplexAssetDiscoverySpecCsvOptions
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	ExcludePatterns() *[]*string
	SetExcludePatterns(val *[]*string)
	ExcludePatternsInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludePatterns() *[]*string
	SetIncludePatterns(val *[]*string)
	IncludePatternsInput() *[]*string
	InternalValue() *GoogleDataplexAssetDiscoverySpec
	SetInternalValue(val *GoogleDataplexAssetDiscoverySpec)
	JsonOptions() GoogleDataplexAssetDiscoverySpecJsonOptionsOutputReference
	JsonOptionsInput() *GoogleDataplexAssetDiscoverySpecJsonOptions
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
	PutCsvOptions(value *GoogleDataplexAssetDiscoverySpecCsvOptions)
	PutJsonOptions(value *GoogleDataplexAssetDiscoverySpecJsonOptions)
	ResetCsvOptions()
	ResetExcludePatterns()
	ResetIncludePatterns()
	ResetJsonOptions()
	ResetSchedule()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataplexAssetDiscoverySpecOutputReference
type jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) CsvOptions() GoogleDataplexAssetDiscoverySpecCsvOptionsOutputReference {
	var returns GoogleDataplexAssetDiscoverySpecCsvOptionsOutputReference
	_jsii_.Get(
		j,
		"csvOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) CsvOptionsInput() *GoogleDataplexAssetDiscoverySpecCsvOptions {
	var returns *GoogleDataplexAssetDiscoverySpecCsvOptions
	_jsii_.Get(
		j,
		"csvOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) ExcludePatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludePatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) ExcludePatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludePatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) IncludePatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includePatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) IncludePatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includePatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) InternalValue() *GoogleDataplexAssetDiscoverySpec {
	var returns *GoogleDataplexAssetDiscoverySpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) JsonOptions() GoogleDataplexAssetDiscoverySpecJsonOptionsOutputReference {
	var returns GoogleDataplexAssetDiscoverySpecJsonOptionsOutputReference
	_jsii_.Get(
		j,
		"jsonOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) JsonOptionsInput() *GoogleDataplexAssetDiscoverySpecJsonOptions {
	var returns *GoogleDataplexAssetDiscoverySpecJsonOptions
	_jsii_.Get(
		j,
		"jsonOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) Schedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) ScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataplexAssetDiscoverySpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataplexAssetDiscoverySpecOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataplexAssetDiscoverySpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataplexAsset.GoogleDataplexAssetDiscoverySpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataplexAssetDiscoverySpecOutputReference_Override(g GoogleDataplexAssetDiscoverySpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataplexAsset.GoogleDataplexAssetDiscoverySpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference)SetExcludePatterns(val *[]*string) {
	if err := j.validateSetExcludePatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludePatterns",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference)SetIncludePatterns(val *[]*string) {
	if err := j.validateSetIncludePatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includePatterns",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference)SetInternalValue(val *GoogleDataplexAssetDiscoverySpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference)SetSchedule(val *string) {
	if err := j.validateSetScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) PutCsvOptions(value *GoogleDataplexAssetDiscoverySpecCsvOptions) {
	if err := g.validatePutCsvOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCsvOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) PutJsonOptions(value *GoogleDataplexAssetDiscoverySpecJsonOptions) {
	if err := g.validatePutJsonOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putJsonOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) ResetCsvOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetCsvOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) ResetExcludePatterns() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludePatterns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) ResetIncludePatterns() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludePatterns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) ResetJsonOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetJsonOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) ResetSchedule() {
	_jsii_.InvokeVoid(
		g,
		"resetSchedule",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataplexAssetDiscoverySpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

