package googledataplextask

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataplextask/internal"
)

type GoogleDataplexTaskNotebookOutputReference interface {
	cdktf.ComplexObject
	ArchiveUris() *[]*string
	SetArchiveUris(val *[]*string)
	ArchiveUrisInput() *[]*string
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
	FileUris() *[]*string
	SetFileUris(val *[]*string)
	FileUrisInput() *[]*string
	// Experimental.
	Fqn() *string
	InfrastructureSpec() GoogleDataplexTaskNotebookInfrastructureSpecOutputReference
	InfrastructureSpecInput() *GoogleDataplexTaskNotebookInfrastructureSpec
	InternalValue() *GoogleDataplexTaskNotebook
	SetInternalValue(val *GoogleDataplexTaskNotebook)
	Notebook() *string
	SetNotebook(val *string)
	NotebookInput() *string
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
	PutInfrastructureSpec(value *GoogleDataplexTaskNotebookInfrastructureSpec)
	ResetArchiveUris()
	ResetFileUris()
	ResetInfrastructureSpec()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataplexTaskNotebookOutputReference
type jsiiProxy_GoogleDataplexTaskNotebookOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) ArchiveUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"archiveUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) ArchiveUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"archiveUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) FileUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fileUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) FileUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fileUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) InfrastructureSpec() GoogleDataplexTaskNotebookInfrastructureSpecOutputReference {
	var returns GoogleDataplexTaskNotebookInfrastructureSpecOutputReference
	_jsii_.Get(
		j,
		"infrastructureSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) InfrastructureSpecInput() *GoogleDataplexTaskNotebookInfrastructureSpec {
	var returns *GoogleDataplexTaskNotebookInfrastructureSpec
	_jsii_.Get(
		j,
		"infrastructureSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) InternalValue() *GoogleDataplexTaskNotebook {
	var returns *GoogleDataplexTaskNotebook
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) Notebook() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) NotebookInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataplexTaskNotebookOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataplexTaskNotebookOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataplexTaskNotebookOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataplexTaskNotebookOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataplexTask.GoogleDataplexTaskNotebookOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataplexTaskNotebookOutputReference_Override(g GoogleDataplexTaskNotebookOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataplexTask.GoogleDataplexTaskNotebookOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookOutputReference)SetArchiveUris(val *[]*string) {
	if err := j.validateSetArchiveUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"archiveUris",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookOutputReference)SetFileUris(val *[]*string) {
	if err := j.validateSetFileUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileUris",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookOutputReference)SetInternalValue(val *GoogleDataplexTaskNotebook) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookOutputReference)SetNotebook(val *string) {
	if err := j.validateSetNotebookParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notebook",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) PutInfrastructureSpec(value *GoogleDataplexTaskNotebookInfrastructureSpec) {
	if err := g.validatePutInfrastructureSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInfrastructureSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) ResetArchiveUris() {
	_jsii_.InvokeVoid(
		g,
		"resetArchiveUris",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) ResetFileUris() {
	_jsii_.InvokeVoid(
		g,
		"resetFileUris",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) ResetInfrastructureSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetInfrastructureSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataplexTaskNotebookOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

