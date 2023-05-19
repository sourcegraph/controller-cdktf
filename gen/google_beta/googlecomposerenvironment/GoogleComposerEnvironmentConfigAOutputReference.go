package googlecomposerenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomposerenvironment/internal"
)

type GoogleComposerEnvironmentConfigAOutputReference interface {
	cdktf.ComplexObject
	AirflowUri() *string
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
	DagGcsPrefix() *string
	DatabaseConfig() GoogleComposerEnvironmentConfigDatabaseConfigOutputReference
	DatabaseConfigInput() *GoogleComposerEnvironmentConfigDatabaseConfig
	EncryptionConfig() GoogleComposerEnvironmentConfigEncryptionConfigOutputReference
	EncryptionConfigInput() *GoogleComposerEnvironmentConfigEncryptionConfig
	EnvironmentSize() *string
	SetEnvironmentSize(val *string)
	EnvironmentSizeInput() *string
	// Experimental.
	Fqn() *string
	GkeCluster() *string
	InternalValue() *GoogleComposerEnvironmentConfigA
	SetInternalValue(val *GoogleComposerEnvironmentConfigA)
	MaintenanceWindow() GoogleComposerEnvironmentConfigMaintenanceWindowOutputReference
	MaintenanceWindowInput() *GoogleComposerEnvironmentConfigMaintenanceWindow
	MasterAuthorizedNetworksConfig() GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigOutputReference
	MasterAuthorizedNetworksConfigInput() *GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfig
	NodeConfig() GoogleComposerEnvironmentConfigNodeConfigOutputReference
	NodeConfigInput() *GoogleComposerEnvironmentConfigNodeConfig
	NodeCount() *float64
	SetNodeCount(val *float64)
	NodeCountInput() *float64
	PrivateEnvironmentConfig() GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference
	PrivateEnvironmentConfigInput() *GoogleComposerEnvironmentConfigPrivateEnvironmentConfig
	RecoveryConfig() GoogleComposerEnvironmentConfigRecoveryConfigOutputReference
	RecoveryConfigInput() *GoogleComposerEnvironmentConfigRecoveryConfig
	SoftwareConfig() GoogleComposerEnvironmentConfigSoftwareConfigOutputReference
	SoftwareConfigInput() *GoogleComposerEnvironmentConfigSoftwareConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WebServerConfig() GoogleComposerEnvironmentConfigWebServerConfigOutputReference
	WebServerConfigInput() *GoogleComposerEnvironmentConfigWebServerConfig
	WebServerNetworkAccessControl() GoogleComposerEnvironmentConfigWebServerNetworkAccessControlOutputReference
	WebServerNetworkAccessControlInput() *GoogleComposerEnvironmentConfigWebServerNetworkAccessControl
	WorkloadsConfig() GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference
	WorkloadsConfigInput() *GoogleComposerEnvironmentConfigWorkloadsConfig
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
	PutDatabaseConfig(value *GoogleComposerEnvironmentConfigDatabaseConfig)
	PutEncryptionConfig(value *GoogleComposerEnvironmentConfigEncryptionConfig)
	PutMaintenanceWindow(value *GoogleComposerEnvironmentConfigMaintenanceWindow)
	PutMasterAuthorizedNetworksConfig(value *GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfig)
	PutNodeConfig(value *GoogleComposerEnvironmentConfigNodeConfig)
	PutPrivateEnvironmentConfig(value *GoogleComposerEnvironmentConfigPrivateEnvironmentConfig)
	PutRecoveryConfig(value *GoogleComposerEnvironmentConfigRecoveryConfig)
	PutSoftwareConfig(value *GoogleComposerEnvironmentConfigSoftwareConfig)
	PutWebServerConfig(value *GoogleComposerEnvironmentConfigWebServerConfig)
	PutWebServerNetworkAccessControl(value *GoogleComposerEnvironmentConfigWebServerNetworkAccessControl)
	PutWorkloadsConfig(value *GoogleComposerEnvironmentConfigWorkloadsConfig)
	ResetDatabaseConfig()
	ResetEncryptionConfig()
	ResetEnvironmentSize()
	ResetMaintenanceWindow()
	ResetMasterAuthorizedNetworksConfig()
	ResetNodeConfig()
	ResetNodeCount()
	ResetPrivateEnvironmentConfig()
	ResetRecoveryConfig()
	ResetSoftwareConfig()
	ResetWebServerConfig()
	ResetWebServerNetworkAccessControl()
	ResetWorkloadsConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComposerEnvironmentConfigAOutputReference
type jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) AirflowUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"airflowUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) DagGcsPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dagGcsPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) DatabaseConfig() GoogleComposerEnvironmentConfigDatabaseConfigOutputReference {
	var returns GoogleComposerEnvironmentConfigDatabaseConfigOutputReference
	_jsii_.Get(
		j,
		"databaseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) DatabaseConfigInput() *GoogleComposerEnvironmentConfigDatabaseConfig {
	var returns *GoogleComposerEnvironmentConfigDatabaseConfig
	_jsii_.Get(
		j,
		"databaseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) EncryptionConfig() GoogleComposerEnvironmentConfigEncryptionConfigOutputReference {
	var returns GoogleComposerEnvironmentConfigEncryptionConfigOutputReference
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) EncryptionConfigInput() *GoogleComposerEnvironmentConfigEncryptionConfig {
	var returns *GoogleComposerEnvironmentConfigEncryptionConfig
	_jsii_.Get(
		j,
		"encryptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) EnvironmentSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) EnvironmentSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) GkeCluster() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gkeCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) InternalValue() *GoogleComposerEnvironmentConfigA {
	var returns *GoogleComposerEnvironmentConfigA
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) MaintenanceWindow() GoogleComposerEnvironmentConfigMaintenanceWindowOutputReference {
	var returns GoogleComposerEnvironmentConfigMaintenanceWindowOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) MaintenanceWindowInput() *GoogleComposerEnvironmentConfigMaintenanceWindow {
	var returns *GoogleComposerEnvironmentConfigMaintenanceWindow
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) MasterAuthorizedNetworksConfig() GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigOutputReference {
	var returns GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigOutputReference
	_jsii_.Get(
		j,
		"masterAuthorizedNetworksConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) MasterAuthorizedNetworksConfigInput() *GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfig {
	var returns *GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfig
	_jsii_.Get(
		j,
		"masterAuthorizedNetworksConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) NodeConfig() GoogleComposerEnvironmentConfigNodeConfigOutputReference {
	var returns GoogleComposerEnvironmentConfigNodeConfigOutputReference
	_jsii_.Get(
		j,
		"nodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) NodeConfigInput() *GoogleComposerEnvironmentConfigNodeConfig {
	var returns *GoogleComposerEnvironmentConfigNodeConfig
	_jsii_.Get(
		j,
		"nodeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) NodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) PrivateEnvironmentConfig() GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference {
	var returns GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference
	_jsii_.Get(
		j,
		"privateEnvironmentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) PrivateEnvironmentConfigInput() *GoogleComposerEnvironmentConfigPrivateEnvironmentConfig {
	var returns *GoogleComposerEnvironmentConfigPrivateEnvironmentConfig
	_jsii_.Get(
		j,
		"privateEnvironmentConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) RecoveryConfig() GoogleComposerEnvironmentConfigRecoveryConfigOutputReference {
	var returns GoogleComposerEnvironmentConfigRecoveryConfigOutputReference
	_jsii_.Get(
		j,
		"recoveryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) RecoveryConfigInput() *GoogleComposerEnvironmentConfigRecoveryConfig {
	var returns *GoogleComposerEnvironmentConfigRecoveryConfig
	_jsii_.Get(
		j,
		"recoveryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) SoftwareConfig() GoogleComposerEnvironmentConfigSoftwareConfigOutputReference {
	var returns GoogleComposerEnvironmentConfigSoftwareConfigOutputReference
	_jsii_.Get(
		j,
		"softwareConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) SoftwareConfigInput() *GoogleComposerEnvironmentConfigSoftwareConfig {
	var returns *GoogleComposerEnvironmentConfigSoftwareConfig
	_jsii_.Get(
		j,
		"softwareConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) WebServerConfig() GoogleComposerEnvironmentConfigWebServerConfigOutputReference {
	var returns GoogleComposerEnvironmentConfigWebServerConfigOutputReference
	_jsii_.Get(
		j,
		"webServerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) WebServerConfigInput() *GoogleComposerEnvironmentConfigWebServerConfig {
	var returns *GoogleComposerEnvironmentConfigWebServerConfig
	_jsii_.Get(
		j,
		"webServerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) WebServerNetworkAccessControl() GoogleComposerEnvironmentConfigWebServerNetworkAccessControlOutputReference {
	var returns GoogleComposerEnvironmentConfigWebServerNetworkAccessControlOutputReference
	_jsii_.Get(
		j,
		"webServerNetworkAccessControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) WebServerNetworkAccessControlInput() *GoogleComposerEnvironmentConfigWebServerNetworkAccessControl {
	var returns *GoogleComposerEnvironmentConfigWebServerNetworkAccessControl
	_jsii_.Get(
		j,
		"webServerNetworkAccessControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) WorkloadsConfig() GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference {
	var returns GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference
	_jsii_.Get(
		j,
		"workloadsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) WorkloadsConfigInput() *GoogleComposerEnvironmentConfigWorkloadsConfig {
	var returns *GoogleComposerEnvironmentConfigWorkloadsConfig
	_jsii_.Get(
		j,
		"workloadsConfigInput",
		&returns,
	)
	return returns
}


func NewGoogleComposerEnvironmentConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComposerEnvironmentConfigAOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComposerEnvironmentConfigAOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference{}

	_jsii_.Create(
		"google-beta.googleComposerEnvironment.GoogleComposerEnvironmentConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComposerEnvironmentConfigAOutputReference_Override(g GoogleComposerEnvironmentConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleComposerEnvironment.GoogleComposerEnvironmentConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference)SetEnvironmentSize(val *string) {
	if err := j.validateSetEnvironmentSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentSize",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference)SetInternalValue(val *GoogleComposerEnvironmentConfigA) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference)SetNodeCount(val *float64) {
	if err := j.validateSetNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeCount",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) PutDatabaseConfig(value *GoogleComposerEnvironmentConfigDatabaseConfig) {
	if err := g.validatePutDatabaseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDatabaseConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) PutEncryptionConfig(value *GoogleComposerEnvironmentConfigEncryptionConfig) {
	if err := g.validatePutEncryptionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEncryptionConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) PutMaintenanceWindow(value *GoogleComposerEnvironmentConfigMaintenanceWindow) {
	if err := g.validatePutMaintenanceWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) PutMasterAuthorizedNetworksConfig(value *GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfig) {
	if err := g.validatePutMasterAuthorizedNetworksConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMasterAuthorizedNetworksConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) PutNodeConfig(value *GoogleComposerEnvironmentConfigNodeConfig) {
	if err := g.validatePutNodeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNodeConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) PutPrivateEnvironmentConfig(value *GoogleComposerEnvironmentConfigPrivateEnvironmentConfig) {
	if err := g.validatePutPrivateEnvironmentConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPrivateEnvironmentConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) PutRecoveryConfig(value *GoogleComposerEnvironmentConfigRecoveryConfig) {
	if err := g.validatePutRecoveryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRecoveryConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) PutSoftwareConfig(value *GoogleComposerEnvironmentConfigSoftwareConfig) {
	if err := g.validatePutSoftwareConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSoftwareConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) PutWebServerConfig(value *GoogleComposerEnvironmentConfigWebServerConfig) {
	if err := g.validatePutWebServerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWebServerConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) PutWebServerNetworkAccessControl(value *GoogleComposerEnvironmentConfigWebServerNetworkAccessControl) {
	if err := g.validatePutWebServerNetworkAccessControlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWebServerNetworkAccessControl",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) PutWorkloadsConfig(value *GoogleComposerEnvironmentConfigWorkloadsConfig) {
	if err := g.validatePutWorkloadsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWorkloadsConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) ResetDatabaseConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetDatabaseConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) ResetEncryptionConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetEncryptionConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) ResetEnvironmentSize() {
	_jsii_.InvokeVoid(
		g,
		"resetEnvironmentSize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		g,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) ResetMasterAuthorizedNetworksConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetMasterAuthorizedNetworksConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) ResetNodeConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) ResetNodeCount() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) ResetPrivateEnvironmentConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetPrivateEnvironmentConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) ResetRecoveryConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetRecoveryConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) ResetSoftwareConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSoftwareConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) ResetWebServerConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetWebServerConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) ResetWebServerNetworkAccessControl() {
	_jsii_.InvokeVoid(
		g,
		"resetWebServerNetworkAccessControl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) ResetWorkloadsConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkloadsConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

