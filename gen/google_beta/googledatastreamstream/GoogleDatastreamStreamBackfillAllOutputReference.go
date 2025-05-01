package googledatastreamstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatastreamstream/internal"
)

type GoogleDatastreamStreamBackfillAllOutputReference interface {
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
	InternalValue() *GoogleDatastreamStreamBackfillAll
	SetInternalValue(val *GoogleDatastreamStreamBackfillAll)
	MysqlExcludedObjects() GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsOutputReference
	MysqlExcludedObjectsInput() *GoogleDatastreamStreamBackfillAllMysqlExcludedObjects
	OracleExcludedObjects() GoogleDatastreamStreamBackfillAllOracleExcludedObjectsOutputReference
	OracleExcludedObjectsInput() *GoogleDatastreamStreamBackfillAllOracleExcludedObjects
	PostgresqlExcludedObjects() GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsOutputReference
	PostgresqlExcludedObjectsInput() *GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjects
	SalesforceExcludedObjects() GoogleDatastreamStreamBackfillAllSalesforceExcludedObjectsOutputReference
	SalesforceExcludedObjectsInput() *GoogleDatastreamStreamBackfillAllSalesforceExcludedObjects
	SqlServerExcludedObjects() GoogleDatastreamStreamBackfillAllSqlServerExcludedObjectsOutputReference
	SqlServerExcludedObjectsInput() *GoogleDatastreamStreamBackfillAllSqlServerExcludedObjects
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
	PutMysqlExcludedObjects(value *GoogleDatastreamStreamBackfillAllMysqlExcludedObjects)
	PutOracleExcludedObjects(value *GoogleDatastreamStreamBackfillAllOracleExcludedObjects)
	PutPostgresqlExcludedObjects(value *GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjects)
	PutSalesforceExcludedObjects(value *GoogleDatastreamStreamBackfillAllSalesforceExcludedObjects)
	PutSqlServerExcludedObjects(value *GoogleDatastreamStreamBackfillAllSqlServerExcludedObjects)
	ResetMysqlExcludedObjects()
	ResetOracleExcludedObjects()
	ResetPostgresqlExcludedObjects()
	ResetSalesforceExcludedObjects()
	ResetSqlServerExcludedObjects()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDatastreamStreamBackfillAllOutputReference
type jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) InternalValue() *GoogleDatastreamStreamBackfillAll {
	var returns *GoogleDatastreamStreamBackfillAll
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) MysqlExcludedObjects() GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsOutputReference {
	var returns GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsOutputReference
	_jsii_.Get(
		j,
		"mysqlExcludedObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) MysqlExcludedObjectsInput() *GoogleDatastreamStreamBackfillAllMysqlExcludedObjects {
	var returns *GoogleDatastreamStreamBackfillAllMysqlExcludedObjects
	_jsii_.Get(
		j,
		"mysqlExcludedObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) OracleExcludedObjects() GoogleDatastreamStreamBackfillAllOracleExcludedObjectsOutputReference {
	var returns GoogleDatastreamStreamBackfillAllOracleExcludedObjectsOutputReference
	_jsii_.Get(
		j,
		"oracleExcludedObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) OracleExcludedObjectsInput() *GoogleDatastreamStreamBackfillAllOracleExcludedObjects {
	var returns *GoogleDatastreamStreamBackfillAllOracleExcludedObjects
	_jsii_.Get(
		j,
		"oracleExcludedObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) PostgresqlExcludedObjects() GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsOutputReference {
	var returns GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsOutputReference
	_jsii_.Get(
		j,
		"postgresqlExcludedObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) PostgresqlExcludedObjectsInput() *GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjects {
	var returns *GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjects
	_jsii_.Get(
		j,
		"postgresqlExcludedObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) SalesforceExcludedObjects() GoogleDatastreamStreamBackfillAllSalesforceExcludedObjectsOutputReference {
	var returns GoogleDatastreamStreamBackfillAllSalesforceExcludedObjectsOutputReference
	_jsii_.Get(
		j,
		"salesforceExcludedObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) SalesforceExcludedObjectsInput() *GoogleDatastreamStreamBackfillAllSalesforceExcludedObjects {
	var returns *GoogleDatastreamStreamBackfillAllSalesforceExcludedObjects
	_jsii_.Get(
		j,
		"salesforceExcludedObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) SqlServerExcludedObjects() GoogleDatastreamStreamBackfillAllSqlServerExcludedObjectsOutputReference {
	var returns GoogleDatastreamStreamBackfillAllSqlServerExcludedObjectsOutputReference
	_jsii_.Get(
		j,
		"sqlServerExcludedObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) SqlServerExcludedObjectsInput() *GoogleDatastreamStreamBackfillAllSqlServerExcludedObjects {
	var returns *GoogleDatastreamStreamBackfillAllSqlServerExcludedObjects
	_jsii_.Get(
		j,
		"sqlServerExcludedObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDatastreamStreamBackfillAllOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDatastreamStreamBackfillAllOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDatastreamStreamBackfillAllOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDatastreamStreamBackfillAllOutputReference_Override(g GoogleDatastreamStreamBackfillAllOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference)SetInternalValue(val *GoogleDatastreamStreamBackfillAll) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) PutMysqlExcludedObjects(value *GoogleDatastreamStreamBackfillAllMysqlExcludedObjects) {
	if err := g.validatePutMysqlExcludedObjectsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMysqlExcludedObjects",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) PutOracleExcludedObjects(value *GoogleDatastreamStreamBackfillAllOracleExcludedObjects) {
	if err := g.validatePutOracleExcludedObjectsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOracleExcludedObjects",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) PutPostgresqlExcludedObjects(value *GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjects) {
	if err := g.validatePutPostgresqlExcludedObjectsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPostgresqlExcludedObjects",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) PutSalesforceExcludedObjects(value *GoogleDatastreamStreamBackfillAllSalesforceExcludedObjects) {
	if err := g.validatePutSalesforceExcludedObjectsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSalesforceExcludedObjects",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) PutSqlServerExcludedObjects(value *GoogleDatastreamStreamBackfillAllSqlServerExcludedObjects) {
	if err := g.validatePutSqlServerExcludedObjectsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSqlServerExcludedObjects",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) ResetMysqlExcludedObjects() {
	_jsii_.InvokeVoid(
		g,
		"resetMysqlExcludedObjects",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) ResetOracleExcludedObjects() {
	_jsii_.InvokeVoid(
		g,
		"resetOracleExcludedObjects",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) ResetPostgresqlExcludedObjects() {
	_jsii_.InvokeVoid(
		g,
		"resetPostgresqlExcludedObjects",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) ResetSalesforceExcludedObjects() {
	_jsii_.InvokeVoid(
		g,
		"resetSalesforceExcludedObjects",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) ResetSqlServerExcludedObjects() {
	_jsii_.InvokeVoid(
		g,
		"resetSqlServerExcludedObjects",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

