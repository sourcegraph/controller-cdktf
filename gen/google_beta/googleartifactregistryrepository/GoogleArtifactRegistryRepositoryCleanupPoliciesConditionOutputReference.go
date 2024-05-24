package googleartifactregistryrepository

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleartifactregistryrepository/internal"
)

type GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference interface {
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
	InternalValue() *GoogleArtifactRegistryRepositoryCleanupPoliciesCondition
	SetInternalValue(val *GoogleArtifactRegistryRepositoryCleanupPoliciesCondition)
	NewerThan() *string
	SetNewerThan(val *string)
	NewerThanInput() *string
	OlderThan() *string
	SetOlderThan(val *string)
	OlderThanInput() *string
	PackageNamePrefixes() *[]*string
	SetPackageNamePrefixes(val *[]*string)
	PackageNamePrefixesInput() *[]*string
	TagPrefixes() *[]*string
	SetTagPrefixes(val *[]*string)
	TagPrefixesInput() *[]*string
	TagState() *string
	SetTagState(val *string)
	TagStateInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VersionNamePrefixes() *[]*string
	SetVersionNamePrefixes(val *[]*string)
	VersionNamePrefixesInput() *[]*string
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
	ResetNewerThan()
	ResetOlderThan()
	ResetPackageNamePrefixes()
	ResetTagPrefixes()
	ResetTagState()
	ResetVersionNamePrefixes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference
type jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) InternalValue() *GoogleArtifactRegistryRepositoryCleanupPoliciesCondition {
	var returns *GoogleArtifactRegistryRepositoryCleanupPoliciesCondition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) NewerThan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newerThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) NewerThanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newerThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) OlderThan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"olderThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) OlderThanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"olderThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) PackageNamePrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"packageNamePrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) PackageNamePrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"packageNamePrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) TagPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) TagPrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagPrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) TagState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) TagStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) VersionNamePrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"versionNamePrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) VersionNamePrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"versionNamePrefixesInput",
		&returns,
	)
	return returns
}


func NewGoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleArtifactRegistryRepository.GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference_Override(g GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleArtifactRegistryRepository.GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference)SetInternalValue(val *GoogleArtifactRegistryRepositoryCleanupPoliciesCondition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference)SetNewerThan(val *string) {
	if err := j.validateSetNewerThanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newerThan",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference)SetOlderThan(val *string) {
	if err := j.validateSetOlderThanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"olderThan",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference)SetPackageNamePrefixes(val *[]*string) {
	if err := j.validateSetPackageNamePrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packageNamePrefixes",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference)SetTagPrefixes(val *[]*string) {
	if err := j.validateSetTagPrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagPrefixes",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference)SetTagState(val *string) {
	if err := j.validateSetTagStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagState",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference)SetVersionNamePrefixes(val *[]*string) {
	if err := j.validateSetVersionNamePrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionNamePrefixes",
		val,
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) ResetNewerThan() {
	_jsii_.InvokeVoid(
		g,
		"resetNewerThan",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) ResetOlderThan() {
	_jsii_.InvokeVoid(
		g,
		"resetOlderThan",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) ResetPackageNamePrefixes() {
	_jsii_.InvokeVoid(
		g,
		"resetPackageNamePrefixes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) ResetTagPrefixes() {
	_jsii_.InvokeVoid(
		g,
		"resetTagPrefixes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) ResetTagState() {
	_jsii_.InvokeVoid(
		g,
		"resetTagState",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) ResetVersionNamePrefixes() {
	_jsii_.InvokeVoid(
		g,
		"resetVersionNamePrefixes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

