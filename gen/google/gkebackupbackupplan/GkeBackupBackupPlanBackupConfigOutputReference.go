package gkebackupbackupplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/gkebackupbackupplan/internal"
)

type GkeBackupBackupPlanBackupConfigOutputReference interface {
	cdktf.ComplexObject
	AllNamespaces() any
	SetAllNamespaces(val any)
	AllNamespacesInput() any
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
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
	EncryptionKey() GkeBackupBackupPlanBackupConfigEncryptionKeyOutputReference
	EncryptionKeyInput() *GkeBackupBackupPlanBackupConfigEncryptionKey
	// Experimental.
	Fqn() *string
	IncludeSecrets() any
	SetIncludeSecrets(val any)
	IncludeSecretsInput() any
	IncludeVolumeData() any
	SetIncludeVolumeData(val any)
	IncludeVolumeDataInput() any
	InternalValue() *GkeBackupBackupPlanBackupConfig
	SetInternalValue(val *GkeBackupBackupPlanBackupConfig)
	PermissiveMode() any
	SetPermissiveMode(val any)
	PermissiveModeInput() any
	SelectedApplications() GkeBackupBackupPlanBackupConfigSelectedApplicationsOutputReference
	SelectedApplicationsInput() *GkeBackupBackupPlanBackupConfigSelectedApplications
	SelectedNamespaceLabels() GkeBackupBackupPlanBackupConfigSelectedNamespaceLabelsOutputReference
	SelectedNamespaceLabelsInput() *GkeBackupBackupPlanBackupConfigSelectedNamespaceLabels
	SelectedNamespaces() GkeBackupBackupPlanBackupConfigSelectedNamespacesOutputReference
	SelectedNamespacesInput() *GkeBackupBackupPlanBackupConfigSelectedNamespaces
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
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	PutEncryptionKey(value *GkeBackupBackupPlanBackupConfigEncryptionKey)
	PutSelectedApplications(value *GkeBackupBackupPlanBackupConfigSelectedApplications)
	PutSelectedNamespaceLabels(value *GkeBackupBackupPlanBackupConfigSelectedNamespaceLabels)
	PutSelectedNamespaces(value *GkeBackupBackupPlanBackupConfigSelectedNamespaces)
	ResetAllNamespaces()
	ResetEncryptionKey()
	ResetIncludeSecrets()
	ResetIncludeVolumeData()
	ResetPermissiveMode()
	ResetSelectedApplications()
	ResetSelectedNamespaceLabels()
	ResetSelectedNamespaces()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GkeBackupBackupPlanBackupConfigOutputReference
type jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) AllNamespaces() any {
	var returns any
	_jsii_.Get(
		j,
		"allNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) AllNamespacesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"allNamespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) EncryptionKey() GkeBackupBackupPlanBackupConfigEncryptionKeyOutputReference {
	var returns GkeBackupBackupPlanBackupConfigEncryptionKeyOutputReference
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) EncryptionKeyInput() *GkeBackupBackupPlanBackupConfigEncryptionKey {
	var returns *GkeBackupBackupPlanBackupConfigEncryptionKey
	_jsii_.Get(
		j,
		"encryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) IncludeSecrets() any {
	var returns any
	_jsii_.Get(
		j,
		"includeSecrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) IncludeSecretsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"includeSecretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) IncludeVolumeData() any {
	var returns any
	_jsii_.Get(
		j,
		"includeVolumeData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) IncludeVolumeDataInput() any {
	var returns any
	_jsii_.Get(
		j,
		"includeVolumeDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) InternalValue() *GkeBackupBackupPlanBackupConfig {
	var returns *GkeBackupBackupPlanBackupConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) PermissiveMode() any {
	var returns any
	_jsii_.Get(
		j,
		"permissiveMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) PermissiveModeInput() any {
	var returns any
	_jsii_.Get(
		j,
		"permissiveModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) SelectedApplications() GkeBackupBackupPlanBackupConfigSelectedApplicationsOutputReference {
	var returns GkeBackupBackupPlanBackupConfigSelectedApplicationsOutputReference
	_jsii_.Get(
		j,
		"selectedApplications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) SelectedApplicationsInput() *GkeBackupBackupPlanBackupConfigSelectedApplications {
	var returns *GkeBackupBackupPlanBackupConfigSelectedApplications
	_jsii_.Get(
		j,
		"selectedApplicationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) SelectedNamespaceLabels() GkeBackupBackupPlanBackupConfigSelectedNamespaceLabelsOutputReference {
	var returns GkeBackupBackupPlanBackupConfigSelectedNamespaceLabelsOutputReference
	_jsii_.Get(
		j,
		"selectedNamespaceLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) SelectedNamespaceLabelsInput() *GkeBackupBackupPlanBackupConfigSelectedNamespaceLabels {
	var returns *GkeBackupBackupPlanBackupConfigSelectedNamespaceLabels
	_jsii_.Get(
		j,
		"selectedNamespaceLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) SelectedNamespaces() GkeBackupBackupPlanBackupConfigSelectedNamespacesOutputReference {
	var returns GkeBackupBackupPlanBackupConfigSelectedNamespacesOutputReference
	_jsii_.Get(
		j,
		"selectedNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) SelectedNamespacesInput() *GkeBackupBackupPlanBackupConfigSelectedNamespaces {
	var returns *GkeBackupBackupPlanBackupConfigSelectedNamespaces
	_jsii_.Get(
		j,
		"selectedNamespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGkeBackupBackupPlanBackupConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GkeBackupBackupPlanBackupConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGkeBackupBackupPlanBackupConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.gkeBackupBackupPlan.GkeBackupBackupPlanBackupConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGkeBackupBackupPlanBackupConfigOutputReference_Override(g GkeBackupBackupPlanBackupConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.gkeBackupBackupPlan.GkeBackupBackupPlanBackupConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) SetAllNamespaces(val any) {
	if err := j.validateSetAllNamespacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allNamespaces",
		val,
	)
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) SetIncludeSecrets(val any) {
	if err := j.validateSetIncludeSecretsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeSecrets",
		val,
	)
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) SetIncludeVolumeData(val any) {
	if err := j.validateSetIncludeVolumeDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeVolumeData",
		val,
	)
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) SetInternalValue(val *GkeBackupBackupPlanBackupConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) SetPermissiveMode(val any) {
	if err := j.validateSetPermissiveModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permissiveMode",
		val,
	)
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) PutEncryptionKey(value *GkeBackupBackupPlanBackupConfigEncryptionKey) {
	if err := g.validatePutEncryptionKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEncryptionKey",
		[]any{value},
	)
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) PutSelectedApplications(value *GkeBackupBackupPlanBackupConfigSelectedApplications) {
	if err := g.validatePutSelectedApplicationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSelectedApplications",
		[]any{value},
	)
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) PutSelectedNamespaceLabels(value *GkeBackupBackupPlanBackupConfigSelectedNamespaceLabels) {
	if err := g.validatePutSelectedNamespaceLabelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSelectedNamespaceLabels",
		[]any{value},
	)
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) PutSelectedNamespaces(value *GkeBackupBackupPlanBackupConfigSelectedNamespaces) {
	if err := g.validatePutSelectedNamespacesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSelectedNamespaces",
		[]any{value},
	)
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) ResetAllNamespaces() {
	_jsii_.InvokeVoid(
		g,
		"resetAllNamespaces",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) ResetEncryptionKey() {
	_jsii_.InvokeVoid(
		g,
		"resetEncryptionKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) ResetIncludeSecrets() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludeSecrets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) ResetIncludeVolumeData() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludeVolumeData",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) ResetPermissiveMode() {
	_jsii_.InvokeVoid(
		g,
		"resetPermissiveMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) ResetSelectedApplications() {
	_jsii_.InvokeVoid(
		g,
		"resetSelectedApplications",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) ResetSelectedNamespaceLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetSelectedNamespaceLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) ResetSelectedNamespaces() {
	_jsii_.InvokeVoid(
		g,
		"resetSelectedNamespaces",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		g,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
