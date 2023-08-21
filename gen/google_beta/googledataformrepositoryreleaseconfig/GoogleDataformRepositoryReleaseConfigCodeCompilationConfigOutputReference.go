package googledataformrepositoryreleaseconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataformrepositoryreleaseconfig/internal"
)

type GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference interface {
	cdktf.ComplexObject
	AssertionSchema() *string
	SetAssertionSchema(val *string)
	AssertionSchemaInput() *string
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
	DatabaseSuffix() *string
	SetDatabaseSuffix(val *string)
	DatabaseSuffixInput() *string
	DefaultDatabase() *string
	SetDefaultDatabase(val *string)
	DefaultDatabaseInput() *string
	DefaultLocation() *string
	SetDefaultLocation(val *string)
	DefaultLocationInput() *string
	DefaultSchema() *string
	SetDefaultSchema(val *string)
	DefaultSchemaInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDataformRepositoryReleaseConfigCodeCompilationConfig
	SetInternalValue(val *GoogleDataformRepositoryReleaseConfigCodeCompilationConfig)
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
	Vars() *map[string]*string
	SetVars(val *map[string]*string)
	VarsInput() *map[string]*string
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
	ResetAssertionSchema()
	ResetDatabaseSuffix()
	ResetDefaultDatabase()
	ResetDefaultLocation()
	ResetDefaultSchema()
	ResetSchemaSuffix()
	ResetTablePrefix()
	ResetVars()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference
type jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) AssertionSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assertionSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) AssertionSchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assertionSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) DatabaseSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) DatabaseSuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseSuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) DefaultDatabase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) DefaultDatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDatabaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) DefaultLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) DefaultLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) DefaultSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) DefaultSchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) InternalValue() *GoogleDataformRepositoryReleaseConfigCodeCompilationConfig {
	var returns *GoogleDataformRepositoryReleaseConfigCodeCompilationConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) SchemaSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) SchemaSuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaSuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) TablePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tablePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) TablePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tablePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) Vars() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"vars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) VarsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"varsInput",
		&returns,
	)
	return returns
}


func NewGoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataformRepositoryReleaseConfig.GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference_Override(g GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataformRepositoryReleaseConfig.GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference)SetAssertionSchema(val *string) {
	if err := j.validateSetAssertionSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assertionSchema",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference)SetDatabaseSuffix(val *string) {
	if err := j.validateSetDatabaseSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseSuffix",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference)SetDefaultDatabase(val *string) {
	if err := j.validateSetDefaultDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultDatabase",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference)SetDefaultLocation(val *string) {
	if err := j.validateSetDefaultLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultLocation",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference)SetDefaultSchema(val *string) {
	if err := j.validateSetDefaultSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultSchema",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference)SetInternalValue(val *GoogleDataformRepositoryReleaseConfigCodeCompilationConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference)SetSchemaSuffix(val *string) {
	if err := j.validateSetSchemaSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaSuffix",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference)SetTablePrefix(val *string) {
	if err := j.validateSetTablePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tablePrefix",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference)SetVars(val *map[string]*string) {
	if err := j.validateSetVarsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vars",
		val,
	)
}

func (g *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) ResetAssertionSchema() {
	_jsii_.InvokeVoid(
		g,
		"resetAssertionSchema",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) ResetDatabaseSuffix() {
	_jsii_.InvokeVoid(
		g,
		"resetDatabaseSuffix",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) ResetDefaultDatabase() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultDatabase",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) ResetDefaultLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) ResetDefaultSchema() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultSchema",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) ResetSchemaSuffix() {
	_jsii_.InvokeVoid(
		g,
		"resetSchemaSuffix",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) ResetTablePrefix() {
	_jsii_.InvokeVoid(
		g,
		"resetTablePrefix",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) ResetVars() {
	_jsii_.InvokeVoid(
		g,
		"resetVars",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

