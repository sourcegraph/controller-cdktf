package googledatastreamstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatastreamstream/internal"
)

type GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasList
type jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewGoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasList {
	_init_.Initialize()

	if err := validateNewGoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasList{}

	_jsii_.Create(
		"google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewGoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasList_Override(g GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		g,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasList) Get(index *float64) GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOutputReference {
	if err := g.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOutputReference

	_jsii_.Invoke(
		g,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

