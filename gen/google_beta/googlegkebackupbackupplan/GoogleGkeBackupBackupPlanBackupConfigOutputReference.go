package googlegkebackupbackupplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlegkebackupbackupplan/internal"
)

type GoogleGkeBackupBackupPlanBackupConfigOutputReference interface {
	cdktf.ComplexObject
	AllNamespaces() interface{}
	SetAllNamespaces(val interface{})
	AllNamespacesInput() interface{}
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
	EncryptionKey() GoogleGkeBackupBackupPlanBackupConfigEncryptionKeyOutputReference
	EncryptionKeyInput() *GoogleGkeBackupBackupPlanBackupConfigEncryptionKey
	// Experimental.
	Fqn() *string
	IncludeSecrets() interface{}
	SetIncludeSecrets(val interface{})
	IncludeSecretsInput() interface{}
	IncludeVolumeData() interface{}
	SetIncludeVolumeData(val interface{})
	IncludeVolumeDataInput() interface{}
	InternalValue() *GoogleGkeBackupBackupPlanBackupConfig
	SetInternalValue(val *GoogleGkeBackupBackupPlanBackupConfig)
	SelectedApplications() GoogleGkeBackupBackupPlanBackupConfigSelectedApplicationsOutputReference
	SelectedApplicationsInput() *GoogleGkeBackupBackupPlanBackupConfigSelectedApplications
	SelectedNamespaces() GoogleGkeBackupBackupPlanBackupConfigSelectedNamespacesOutputReference
	SelectedNamespacesInput() *GoogleGkeBackupBackupPlanBackupConfigSelectedNamespaces
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
	PutEncryptionKey(value *GoogleGkeBackupBackupPlanBackupConfigEncryptionKey)
	PutSelectedApplications(value *GoogleGkeBackupBackupPlanBackupConfigSelectedApplications)
	PutSelectedNamespaces(value *GoogleGkeBackupBackupPlanBackupConfigSelectedNamespaces)
	ResetAllNamespaces()
	ResetEncryptionKey()
	ResetIncludeSecrets()
	ResetIncludeVolumeData()
	ResetSelectedApplications()
	ResetSelectedNamespaces()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleGkeBackupBackupPlanBackupConfigOutputReference
type jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) AllNamespaces() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) AllNamespacesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allNamespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) EncryptionKey() GoogleGkeBackupBackupPlanBackupConfigEncryptionKeyOutputReference {
	var returns GoogleGkeBackupBackupPlanBackupConfigEncryptionKeyOutputReference
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) EncryptionKeyInput() *GoogleGkeBackupBackupPlanBackupConfigEncryptionKey {
	var returns *GoogleGkeBackupBackupPlanBackupConfigEncryptionKey
	_jsii_.Get(
		j,
		"encryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) IncludeSecrets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSecrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) IncludeSecretsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSecretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) IncludeVolumeData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeVolumeData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) IncludeVolumeDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeVolumeDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) InternalValue() *GoogleGkeBackupBackupPlanBackupConfig {
	var returns *GoogleGkeBackupBackupPlanBackupConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) SelectedApplications() GoogleGkeBackupBackupPlanBackupConfigSelectedApplicationsOutputReference {
	var returns GoogleGkeBackupBackupPlanBackupConfigSelectedApplicationsOutputReference
	_jsii_.Get(
		j,
		"selectedApplications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) SelectedApplicationsInput() *GoogleGkeBackupBackupPlanBackupConfigSelectedApplications {
	var returns *GoogleGkeBackupBackupPlanBackupConfigSelectedApplications
	_jsii_.Get(
		j,
		"selectedApplicationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) SelectedNamespaces() GoogleGkeBackupBackupPlanBackupConfigSelectedNamespacesOutputReference {
	var returns GoogleGkeBackupBackupPlanBackupConfigSelectedNamespacesOutputReference
	_jsii_.Get(
		j,
		"selectedNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) SelectedNamespacesInput() *GoogleGkeBackupBackupPlanBackupConfigSelectedNamespaces {
	var returns *GoogleGkeBackupBackupPlanBackupConfigSelectedNamespaces
	_jsii_.Get(
		j,
		"selectedNamespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleGkeBackupBackupPlanBackupConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleGkeBackupBackupPlanBackupConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleGkeBackupBackupPlanBackupConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleGkeBackupBackupPlan.GoogleGkeBackupBackupPlanBackupConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleGkeBackupBackupPlanBackupConfigOutputReference_Override(g GoogleGkeBackupBackupPlanBackupConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleGkeBackupBackupPlan.GoogleGkeBackupBackupPlanBackupConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference)SetAllNamespaces(val interface{}) {
	if err := j.validateSetAllNamespacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allNamespaces",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference)SetIncludeSecrets(val interface{}) {
	if err := j.validateSetIncludeSecretsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeSecrets",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference)SetIncludeVolumeData(val interface{}) {
	if err := j.validateSetIncludeVolumeDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeVolumeData",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference)SetInternalValue(val *GoogleGkeBackupBackupPlanBackupConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) PutEncryptionKey(value *GoogleGkeBackupBackupPlanBackupConfigEncryptionKey) {
	if err := g.validatePutEncryptionKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEncryptionKey",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) PutSelectedApplications(value *GoogleGkeBackupBackupPlanBackupConfigSelectedApplications) {
	if err := g.validatePutSelectedApplicationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSelectedApplications",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) PutSelectedNamespaces(value *GoogleGkeBackupBackupPlanBackupConfigSelectedNamespaces) {
	if err := g.validatePutSelectedNamespacesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSelectedNamespaces",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) ResetAllNamespaces() {
	_jsii_.InvokeVoid(
		g,
		"resetAllNamespaces",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) ResetEncryptionKey() {
	_jsii_.InvokeVoid(
		g,
		"resetEncryptionKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) ResetIncludeSecrets() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludeSecrets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) ResetIncludeVolumeData() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludeVolumeData",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) ResetSelectedApplications() {
	_jsii_.InvokeVoid(
		g,
		"resetSelectedApplications",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) ResetSelectedNamespaces() {
	_jsii_.InvokeVoid(
		g,
		"resetSelectedNamespaces",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

