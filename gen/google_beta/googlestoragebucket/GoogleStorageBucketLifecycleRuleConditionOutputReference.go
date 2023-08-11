package googlestoragebucket

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlestoragebucket/internal"
)

type GoogleStorageBucketLifecycleRuleConditionOutputReference interface {
	cdktf.ComplexObject
	Age() *float64
	SetAge(val *float64)
	AgeInput() *float64
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
	CreatedBefore() *string
	SetCreatedBefore(val *string)
	CreatedBeforeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomTimeBefore() *string
	SetCustomTimeBefore(val *string)
	CustomTimeBeforeInput() *string
	DaysSinceCustomTime() *float64
	SetDaysSinceCustomTime(val *float64)
	DaysSinceCustomTimeInput() *float64
	DaysSinceNoncurrentTime() *float64
	SetDaysSinceNoncurrentTime(val *float64)
	DaysSinceNoncurrentTimeInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleStorageBucketLifecycleRuleCondition
	SetInternalValue(val *GoogleStorageBucketLifecycleRuleCondition)
	MatchesPrefix() *[]*string
	SetMatchesPrefix(val *[]*string)
	MatchesPrefixInput() *[]*string
	MatchesStorageClass() *[]*string
	SetMatchesStorageClass(val *[]*string)
	MatchesStorageClassInput() *[]*string
	MatchesSuffix() *[]*string
	SetMatchesSuffix(val *[]*string)
	MatchesSuffixInput() *[]*string
	NoncurrentTimeBefore() *string
	SetNoncurrentTimeBefore(val *string)
	NoncurrentTimeBeforeInput() *string
	NumNewerVersions() *float64
	SetNumNewerVersions(val *float64)
	NumNewerVersionsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WithState() *string
	SetWithState(val *string)
	WithStateInput() *string
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
	ResetAge()
	ResetCreatedBefore()
	ResetCustomTimeBefore()
	ResetDaysSinceCustomTime()
	ResetDaysSinceNoncurrentTime()
	ResetMatchesPrefix()
	ResetMatchesStorageClass()
	ResetMatchesSuffix()
	ResetNoncurrentTimeBefore()
	ResetNumNewerVersions()
	ResetWithState()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleStorageBucketLifecycleRuleConditionOutputReference
type jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) Age() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"age",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) AgeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) CreatedBefore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBefore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) CreatedBeforeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBeforeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) CustomTimeBefore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customTimeBefore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) CustomTimeBeforeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customTimeBeforeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) DaysSinceCustomTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"daysSinceCustomTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) DaysSinceCustomTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"daysSinceCustomTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) DaysSinceNoncurrentTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"daysSinceNoncurrentTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) DaysSinceNoncurrentTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"daysSinceNoncurrentTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) InternalValue() *GoogleStorageBucketLifecycleRuleCondition {
	var returns *GoogleStorageBucketLifecycleRuleCondition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) MatchesPrefix() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchesPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) MatchesPrefixInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchesPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) MatchesStorageClass() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchesStorageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) MatchesStorageClassInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchesStorageClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) MatchesSuffix() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchesSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) MatchesSuffixInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchesSuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) NoncurrentTimeBefore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noncurrentTimeBefore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) NoncurrentTimeBeforeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noncurrentTimeBeforeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) NumNewerVersions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numNewerVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) NumNewerVersionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numNewerVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) WithState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"withState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) WithStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"withStateInput",
		&returns,
	)
	return returns
}


func NewGoogleStorageBucketLifecycleRuleConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleStorageBucketLifecycleRuleConditionOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleStorageBucketLifecycleRuleConditionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleStorageBucket.GoogleStorageBucketLifecycleRuleConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleStorageBucketLifecycleRuleConditionOutputReference_Override(g GoogleStorageBucketLifecycleRuleConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleStorageBucket.GoogleStorageBucketLifecycleRuleConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference)SetAge(val *float64) {
	if err := j.validateSetAgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"age",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference)SetCreatedBefore(val *string) {
	if err := j.validateSetCreatedBeforeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdBefore",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference)SetCustomTimeBefore(val *string) {
	if err := j.validateSetCustomTimeBeforeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customTimeBefore",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference)SetDaysSinceCustomTime(val *float64) {
	if err := j.validateSetDaysSinceCustomTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"daysSinceCustomTime",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference)SetDaysSinceNoncurrentTime(val *float64) {
	if err := j.validateSetDaysSinceNoncurrentTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"daysSinceNoncurrentTime",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference)SetInternalValue(val *GoogleStorageBucketLifecycleRuleCondition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference)SetMatchesPrefix(val *[]*string) {
	if err := j.validateSetMatchesPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchesPrefix",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference)SetMatchesStorageClass(val *[]*string) {
	if err := j.validateSetMatchesStorageClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchesStorageClass",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference)SetMatchesSuffix(val *[]*string) {
	if err := j.validateSetMatchesSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchesSuffix",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference)SetNoncurrentTimeBefore(val *string) {
	if err := j.validateSetNoncurrentTimeBeforeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noncurrentTimeBefore",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference)SetNumNewerVersions(val *float64) {
	if err := j.validateSetNumNewerVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numNewerVersions",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference)SetWithState(val *string) {
	if err := j.validateSetWithStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"withState",
		val,
	)
}

func (g *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) ResetAge() {
	_jsii_.InvokeVoid(
		g,
		"resetAge",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) ResetCreatedBefore() {
	_jsii_.InvokeVoid(
		g,
		"resetCreatedBefore",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) ResetCustomTimeBefore() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomTimeBefore",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) ResetDaysSinceCustomTime() {
	_jsii_.InvokeVoid(
		g,
		"resetDaysSinceCustomTime",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) ResetDaysSinceNoncurrentTime() {
	_jsii_.InvokeVoid(
		g,
		"resetDaysSinceNoncurrentTime",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) ResetMatchesPrefix() {
	_jsii_.InvokeVoid(
		g,
		"resetMatchesPrefix",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) ResetMatchesStorageClass() {
	_jsii_.InvokeVoid(
		g,
		"resetMatchesStorageClass",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) ResetMatchesSuffix() {
	_jsii_.InvokeVoid(
		g,
		"resetMatchesSuffix",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) ResetNoncurrentTimeBefore() {
	_jsii_.InvokeVoid(
		g,
		"resetNoncurrentTimeBefore",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) ResetNumNewerVersions() {
	_jsii_.InvokeVoid(
		g,
		"resetNumNewerVersions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) ResetWithState() {
	_jsii_.InvokeVoid(
		g,
		"resetWithState",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleStorageBucketLifecycleRuleConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

