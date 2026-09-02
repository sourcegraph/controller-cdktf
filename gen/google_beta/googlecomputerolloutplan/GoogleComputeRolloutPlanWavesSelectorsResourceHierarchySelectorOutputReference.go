package googlecomputerolloutplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputerolloutplan/internal"
)

type GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference interface {
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
	// Experimental.
	Fqn() *string
	IncludedFolders() *[]*string
	SetIncludedFolders(val *[]*string)
	IncludedFoldersInput() *[]*string
	IncludedOrganizations() *[]*string
	SetIncludedOrganizations(val *[]*string)
	IncludedOrganizationsInput() *[]*string
	IncludedProjects() *[]*string
	SetIncludedProjects(val *[]*string)
	IncludedProjectsInput() *[]*string
	InternalValue() *GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelector
	SetInternalValue(val *GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelector)
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
	ResetIncludedFolders()
	ResetIncludedOrganizations()
	ResetIncludedProjects()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference
type jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference) IncludedFolders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedFolders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference) IncludedFoldersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedFoldersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference) IncludedOrganizations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedOrganizations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference) IncludedOrganizationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedOrganizationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference) IncludedProjects() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedProjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference) IncludedProjectsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedProjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference) InternalValue() *GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelector {
	var returns *GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelector
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeRolloutPlan.GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference_Override(g GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeRolloutPlan.GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference)SetIncludedFolders(val *[]*string) {
	if err := j.validateSetIncludedFoldersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedFolders",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference)SetIncludedOrganizations(val *[]*string) {
	if err := j.validateSetIncludedOrganizationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedOrganizations",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference)SetIncludedProjects(val *[]*string) {
	if err := j.validateSetIncludedProjectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedProjects",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference)SetInternalValue(val *GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelector) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference) ResetIncludedFolders() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludedFolders",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference) ResetIncludedOrganizations() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludedOrganizations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference) ResetIncludedProjects() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludedProjects",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelectorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

