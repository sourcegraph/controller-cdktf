package googleosconfigpatchdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleosconfigpatchdeployment/internal"
)

type GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference interface {
	cdktf.ComplexObject
	Categories() *[]*string
	SetCategories(val *[]*string)
	CategoriesInput() *[]*string
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
	Excludes() *[]*string
	SetExcludes(val *[]*string)
	ExcludesInput() *[]*string
	ExclusivePatches() *[]*string
	SetExclusivePatches(val *[]*string)
	ExclusivePatchesInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleOsConfigPatchDeploymentPatchConfigZypper
	SetInternalValue(val *GoogleOsConfigPatchDeploymentPatchConfigZypper)
	Severities() *[]*string
	SetSeverities(val *[]*string)
	SeveritiesInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WithOptional() interface{}
	SetWithOptional(val interface{})
	WithOptionalInput() interface{}
	WithUpdate() interface{}
	SetWithUpdate(val interface{})
	WithUpdateInput() interface{}
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
	ResetCategories()
	ResetExcludes()
	ResetExclusivePatches()
	ResetSeverities()
	ResetWithOptional()
	ResetWithUpdate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference
type jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) Categories() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"categories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) CategoriesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"categoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) Excludes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) ExcludesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) ExclusivePatches() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclusivePatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) ExclusivePatchesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclusivePatchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) InternalValue() *GoogleOsConfigPatchDeploymentPatchConfigZypper {
	var returns *GoogleOsConfigPatchDeploymentPatchConfigZypper
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) Severities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"severities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) SeveritiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"severitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) WithOptional() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withOptional",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) WithOptionalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withOptionalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) WithUpdate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) WithUpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withUpdateInput",
		&returns,
	)
	return returns
}


func NewGoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOsConfigPatchDeploymentPatchConfigZypperOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleOsConfigPatchDeployment.GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference_Override(g GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleOsConfigPatchDeployment.GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference)SetCategories(val *[]*string) {
	if err := j.validateSetCategoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"categories",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference)SetExcludes(val *[]*string) {
	if err := j.validateSetExcludesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludes",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference)SetExclusivePatches(val *[]*string) {
	if err := j.validateSetExclusivePatchesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exclusivePatches",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference)SetInternalValue(val *GoogleOsConfigPatchDeploymentPatchConfigZypper) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference)SetSeverities(val *[]*string) {
	if err := j.validateSetSeveritiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"severities",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference)SetWithOptional(val interface{}) {
	if err := j.validateSetWithOptionalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"withOptional",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference)SetWithUpdate(val interface{}) {
	if err := j.validateSetWithUpdateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"withUpdate",
		val,
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) ResetCategories() {
	_jsii_.InvokeVoid(
		g,
		"resetCategories",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) ResetExcludes() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) ResetExclusivePatches() {
	_jsii_.InvokeVoid(
		g,
		"resetExclusivePatches",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) ResetSeverities() {
	_jsii_.InvokeVoid(
		g,
		"resetSeverities",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) ResetWithOptional() {
	_jsii_.InvokeVoid(
		g,
		"resetWithOptional",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) ResetWithUpdate() {
	_jsii_.InvokeVoid(
		g,
		"resetWithUpdate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

