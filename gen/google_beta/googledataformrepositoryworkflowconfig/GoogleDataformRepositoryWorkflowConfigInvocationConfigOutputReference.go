package googledataformrepositoryworkflowconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataformrepositoryworkflowconfig/internal"
)

type GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference interface {
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
	FullyRefreshIncrementalTablesEnabled() interface{}
	SetFullyRefreshIncrementalTablesEnabled(val interface{})
	FullyRefreshIncrementalTablesEnabledInput() interface{}
	IncludedTags() *[]*string
	SetIncludedTags(val *[]*string)
	IncludedTagsInput() *[]*string
	IncludedTargets() GoogleDataformRepositoryWorkflowConfigInvocationConfigIncludedTargetsList
	IncludedTargetsInput() interface{}
	InternalValue() *GoogleDataformRepositoryWorkflowConfigInvocationConfig
	SetInternalValue(val *GoogleDataformRepositoryWorkflowConfigInvocationConfig)
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransitiveDependenciesIncluded() interface{}
	SetTransitiveDependenciesIncluded(val interface{})
	TransitiveDependenciesIncludedInput() interface{}
	TransitiveDependentsIncluded() interface{}
	SetTransitiveDependentsIncluded(val interface{})
	TransitiveDependentsIncludedInput() interface{}
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
	PutIncludedTargets(value interface{})
	ResetFullyRefreshIncrementalTablesEnabled()
	ResetIncludedTags()
	ResetIncludedTargets()
	ResetServiceAccount()
	ResetTransitiveDependenciesIncluded()
	ResetTransitiveDependentsIncluded()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference
type jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) FullyRefreshIncrementalTablesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fullyRefreshIncrementalTablesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) FullyRefreshIncrementalTablesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fullyRefreshIncrementalTablesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) IncludedTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) IncludedTagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) IncludedTargets() GoogleDataformRepositoryWorkflowConfigInvocationConfigIncludedTargetsList {
	var returns GoogleDataformRepositoryWorkflowConfigInvocationConfigIncludedTargetsList
	_jsii_.Get(
		j,
		"includedTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) IncludedTargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includedTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) InternalValue() *GoogleDataformRepositoryWorkflowConfigInvocationConfig {
	var returns *GoogleDataformRepositoryWorkflowConfigInvocationConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) TransitiveDependenciesIncluded() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transitiveDependenciesIncluded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) TransitiveDependenciesIncludedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transitiveDependenciesIncludedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) TransitiveDependentsIncluded() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transitiveDependentsIncluded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) TransitiveDependentsIncludedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transitiveDependentsIncludedInput",
		&returns,
	)
	return returns
}


func NewGoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataformRepositoryWorkflowConfig.GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference_Override(g GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataformRepositoryWorkflowConfig.GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference)SetFullyRefreshIncrementalTablesEnabled(val interface{}) {
	if err := j.validateSetFullyRefreshIncrementalTablesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullyRefreshIncrementalTablesEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference)SetIncludedTags(val *[]*string) {
	if err := j.validateSetIncludedTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedTags",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference)SetInternalValue(val *GoogleDataformRepositoryWorkflowConfigInvocationConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference)SetTransitiveDependenciesIncluded(val interface{}) {
	if err := j.validateSetTransitiveDependenciesIncludedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitiveDependenciesIncluded",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference)SetTransitiveDependentsIncluded(val interface{}) {
	if err := j.validateSetTransitiveDependentsIncludedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitiveDependentsIncluded",
		val,
	)
}

func (g *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) PutIncludedTargets(value interface{}) {
	if err := g.validatePutIncludedTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIncludedTargets",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) ResetFullyRefreshIncrementalTablesEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetFullyRefreshIncrementalTablesEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) ResetIncludedTags() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludedTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) ResetIncludedTargets() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludedTargets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) ResetTransitiveDependenciesIncluded() {
	_jsii_.InvokeVoid(
		g,
		"resetTransitiveDependenciesIncluded",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) ResetTransitiveDependentsIncluded() {
	_jsii_.InvokeVoid(
		g,
		"resetTransitiveDependentsIncluded",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataformRepositoryWorkflowConfigInvocationConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

