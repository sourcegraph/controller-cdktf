package googlealloydbinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlealloydbinstance/internal"
)

type GoogleAlloydbInstanceObservabilityConfigOutputReference interface {
	cdktf.ComplexObject
	AssistiveExperiencesEnabled() interface{}
	SetAssistiveExperiencesEnabled(val interface{})
	AssistiveExperiencesEnabledInput() interface{}
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleAlloydbInstanceObservabilityConfig
	SetInternalValue(val *GoogleAlloydbInstanceObservabilityConfig)
	MaxQueryStringLength() *float64
	SetMaxQueryStringLength(val *float64)
	MaxQueryStringLengthInput() *float64
	PreserveComments() interface{}
	SetPreserveComments(val interface{})
	PreserveCommentsInput() interface{}
	QueryPlansPerMinute() *float64
	SetQueryPlansPerMinute(val *float64)
	QueryPlansPerMinuteInput() *float64
	RecordApplicationTags() interface{}
	SetRecordApplicationTags(val interface{})
	RecordApplicationTagsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrackActiveQueries() interface{}
	SetTrackActiveQueries(val interface{})
	TrackActiveQueriesInput() interface{}
	TrackWaitEvents() interface{}
	SetTrackWaitEvents(val interface{})
	TrackWaitEventsInput() interface{}
	TrackWaitEventTypes() interface{}
	SetTrackWaitEventTypes(val interface{})
	TrackWaitEventTypesInput() interface{}
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
	ResetAssistiveExperiencesEnabled()
	ResetEnabled()
	ResetMaxQueryStringLength()
	ResetPreserveComments()
	ResetQueryPlansPerMinute()
	ResetRecordApplicationTags()
	ResetTrackActiveQueries()
	ResetTrackWaitEvents()
	ResetTrackWaitEventTypes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleAlloydbInstanceObservabilityConfigOutputReference
type jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) AssistiveExperiencesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assistiveExperiencesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) AssistiveExperiencesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assistiveExperiencesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) InternalValue() *GoogleAlloydbInstanceObservabilityConfig {
	var returns *GoogleAlloydbInstanceObservabilityConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) MaxQueryStringLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxQueryStringLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) MaxQueryStringLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxQueryStringLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) PreserveComments() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveComments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) PreserveCommentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveCommentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) QueryPlansPerMinute() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryPlansPerMinute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) QueryPlansPerMinuteInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryPlansPerMinuteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) RecordApplicationTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recordApplicationTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) RecordApplicationTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recordApplicationTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) TrackActiveQueries() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trackActiveQueries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) TrackActiveQueriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trackActiveQueriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) TrackWaitEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trackWaitEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) TrackWaitEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trackWaitEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) TrackWaitEventTypes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trackWaitEventTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) TrackWaitEventTypesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trackWaitEventTypesInput",
		&returns,
	)
	return returns
}


func NewGoogleAlloydbInstanceObservabilityConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleAlloydbInstanceObservabilityConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleAlloydbInstanceObservabilityConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleAlloydbInstance.GoogleAlloydbInstanceObservabilityConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleAlloydbInstanceObservabilityConfigOutputReference_Override(g GoogleAlloydbInstanceObservabilityConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleAlloydbInstance.GoogleAlloydbInstanceObservabilityConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference)SetAssistiveExperiencesEnabled(val interface{}) {
	if err := j.validateSetAssistiveExperiencesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assistiveExperiencesEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference)SetInternalValue(val *GoogleAlloydbInstanceObservabilityConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference)SetMaxQueryStringLength(val *float64) {
	if err := j.validateSetMaxQueryStringLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxQueryStringLength",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference)SetPreserveComments(val interface{}) {
	if err := j.validateSetPreserveCommentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveComments",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference)SetQueryPlansPerMinute(val *float64) {
	if err := j.validateSetQueryPlansPerMinuteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryPlansPerMinute",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference)SetRecordApplicationTags(val interface{}) {
	if err := j.validateSetRecordApplicationTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordApplicationTags",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference)SetTrackActiveQueries(val interface{}) {
	if err := j.validateSetTrackActiveQueriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trackActiveQueries",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference)SetTrackWaitEvents(val interface{}) {
	if err := j.validateSetTrackWaitEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trackWaitEvents",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference)SetTrackWaitEventTypes(val interface{}) {
	if err := j.validateSetTrackWaitEventTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trackWaitEventTypes",
		val,
	)
}

func (g *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) ResetAssistiveExperiencesEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetAssistiveExperiencesEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) ResetMaxQueryStringLength() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxQueryStringLength",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) ResetPreserveComments() {
	_jsii_.InvokeVoid(
		g,
		"resetPreserveComments",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) ResetQueryPlansPerMinute() {
	_jsii_.InvokeVoid(
		g,
		"resetQueryPlansPerMinute",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) ResetRecordApplicationTags() {
	_jsii_.InvokeVoid(
		g,
		"resetRecordApplicationTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) ResetTrackActiveQueries() {
	_jsii_.InvokeVoid(
		g,
		"resetTrackActiveQueries",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) ResetTrackWaitEvents() {
	_jsii_.InvokeVoid(
		g,
		"resetTrackWaitEvents",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) ResetTrackWaitEventTypes() {
	_jsii_.InvokeVoid(
		g,
		"resetTrackWaitEventTypes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleAlloydbInstanceObservabilityConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

