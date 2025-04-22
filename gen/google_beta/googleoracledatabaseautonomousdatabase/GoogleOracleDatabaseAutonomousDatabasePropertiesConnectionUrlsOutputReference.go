package googleoracledatabaseautonomousdatabase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleoracledatabaseautonomousdatabase/internal"
)

type GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference interface {
	cdktf.ComplexObject
	ApexUri() *string
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
	DatabaseTransformsUri() *string
	// Experimental.
	Fqn() *string
	GraphStudioUri() *string
	InternalValue() *GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrls
	SetInternalValue(val *GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrls)
	MachineLearningNotebookUri() *string
	MachineLearningUserManagementUri() *string
	MongoDbUri() *string
	OrdsUri() *string
	SqlDevWebUri() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference
type jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference) ApexUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apexUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference) DatabaseTransformsUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseTransformsUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference) GraphStudioUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"graphStudioUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference) InternalValue() *GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrls {
	var returns *GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrls
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference) MachineLearningNotebookUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineLearningNotebookUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference) MachineLearningUserManagementUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineLearningUserManagementUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference) MongoDbUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoDbUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference) OrdsUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ordsUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference) SqlDevWebUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlDevWebUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleOracleDatabaseAutonomousDatabase.GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference_Override(g GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleOracleDatabaseAutonomousDatabase.GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference)SetInternalValue(val *GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrls) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

