package googledatastreamstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatastreamstream/internal"
)

type GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PostgresqlColumns() GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesPostgresqlColumnsList
	PostgresqlColumnsInput() interface{}
	Table() *string
	SetTable(val *string)
	TableInput() *string
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
	PutPostgresqlColumns(value interface{})
	ResetPostgresqlColumns()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference
type jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference) PostgresqlColumns() GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesPostgresqlColumnsList {
	var returns GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesPostgresqlColumnsList
	_jsii_.Get(
		j,
		"postgresqlColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference) PostgresqlColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"postgresqlColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference) Table() *string {
	var returns *string
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference) TableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference_Override(g GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference)SetTable(val *string) {
	if err := j.validateSetTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"table",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference) PutPostgresqlColumns(value interface{}) {
	if err := g.validatePutPostgresqlColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPostgresqlColumns",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference) ResetPostgresqlColumns() {
	_jsii_.InvokeVoid(
		g,
		"resetPostgresqlColumns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

