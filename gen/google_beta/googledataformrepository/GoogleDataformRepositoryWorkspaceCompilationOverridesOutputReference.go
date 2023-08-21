package googledataformrepository

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataformrepository/internal"
)

type GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference interface {
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
	DefaultDatabase() *string
	SetDefaultDatabase(val *string)
	DefaultDatabaseInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDataformRepositoryWorkspaceCompilationOverrides
	SetInternalValue(val *GoogleDataformRepositoryWorkspaceCompilationOverrides)
	SchemaSuffix() *string
	SetSchemaSuffix(val *string)
	SchemaSuffixInput() *string
	TablePrefix() *string
	SetTablePrefix(val *string)
	TablePrefixInput() *string
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
	ResetDefaultDatabase()
	ResetSchemaSuffix()
	ResetTablePrefix()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference
type jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference) DefaultDatabase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference) DefaultDatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDatabaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference) InternalValue() *GoogleDataformRepositoryWorkspaceCompilationOverrides {
	var returns *GoogleDataformRepositoryWorkspaceCompilationOverrides
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference) SchemaSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference) SchemaSuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaSuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference) TablePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tablePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference) TablePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tablePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataformRepositoryWorkspaceCompilationOverridesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataformRepository.GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference_Override(g GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataformRepository.GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference)SetDefaultDatabase(val *string) {
	if err := j.validateSetDefaultDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultDatabase",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference)SetInternalValue(val *GoogleDataformRepositoryWorkspaceCompilationOverrides) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference)SetSchemaSuffix(val *string) {
	if err := j.validateSetSchemaSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaSuffix",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference)SetTablePrefix(val *string) {
	if err := j.validateSetTablePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tablePrefix",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference) ResetDefaultDatabase() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultDatabase",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference) ResetSchemaSuffix() {
	_jsii_.InvokeVoid(
		g,
		"resetSchemaSuffix",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference) ResetTablePrefix() {
	_jsii_.InvokeVoid(
		g,
		"resetTablePrefix",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataformRepositoryWorkspaceCompilationOverridesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

