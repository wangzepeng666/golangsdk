package versions

import (
	"strconv"

	"github.com/chnsz/golangsdk/pagination"
)

// Version is the structure that represents the specified function version details.
type Version struct {
	// Function URN.
	FuncUrn string `json:"func_urn"`
	// Function name.
	FuncName string `json:"func_name"`
	// Domain ID.
	DomainId string `json:"domain_id"`
	// Project ID.
	Namespace string `json:"namespace"`
	// Project name.
	ProjectName string `json:"project_name"`
	// Group package to which the function belongs. This field is defined to group functions.
	Package string `json:"package"`
	// Environment in which a FunctionGraph function is executed. Options:
	// Enumeration values:
	// + Java8
	// + Java11
	// + Node.js6.10
	// + Node.js8.10
	// + Node.js10.16
	// + Node.js12.13
	// + Node.js14.18
	// + Python2.7
	// + Python3.6
	// + Go1.8
	// + Go1.x
	// + C#(.NET Core 2.0)
	// + C#(.NET Core 2.1)
	// + C#(.NET Core 3.1)
	// + Custom
	// + PHP7.3
	// + Python3.9
	// + http
	Runtime string `json:"runtime"`
	// Maximum duration the function can be executed. Value range: 3s–900s.
	// The value can be up to 12 hours for whitelisted users.
	// For details, contact FunctionGraph personnel.
	Timeout int `json:"timeout"`
	// Handler of a function in the format of "xx.xx". It must contain a period (.).
	// For example, for Node.js function myfunction.handler, the file name is myfunction.js,
	// and the handler function is handler.
	Handle string `json:"handle"`
	// Memory consumed by a function.
	// Unit: MB.
	// The value can be 128, 256, 512, 768, 1024, 1280, 1536, 1792, 2048, 2560, 3072, 3584 or 4096.
	// The value ranges from 128 to 4096.
	MemorySize int `json:"memory_size"`
	// CPU resources of a function.
	// Unit: millicore (1 core = 1000 millicores).
	// The value of this field is proportional to that of MemorySize.
	// By default, 100 CPU millicores are required for 128 MB memory.
	// The value is calculated as follows: Memory/128 x 100 + 200 (basic CPU millicores)
	CPU int `json:"cpu"`
	// Function code type. Options:
	// + inline: inline code.
	// + zip: ZIP file.
	// + obs: function code stored in an OBS bucket.
	// + jar: JAR file, for Java functions.
	CodeType string `json:"code_type"`
	// If CodeType is set to obs, enter the OBS URL of the function code package.
	// If CodeType is not set to obs, leave this parameter blank.
	CodeURL string `json:"code_url"`
	// Name of a function file. This parameter is mandatory only when CodeType is set to jar or zip.
	CodeFileName string `json:"code_filename"`
	// Code size in bytes.
	CodeSize int `json:"code_size"`
	// Name/Value information defined for the function. These are parameters used in the function.
	// For example, if a function needs to access a host, define Host={host_ip}.
	// You can define a maximum of 20 such parameters, and their total length cannot exceed 4 KB.
	UserData string `json:"user_data"`
	// User-defined name/value to be encrypted.
	EncryptedUserData string `json:"encrypted_user_data"`
	// SHA512 hash value of function code, which is used to determine whether the function has changed.
	Digest string `json:"digest"`
	// Function version, which is automatically generated by the system.
	// The version name is in the format of "vYYYYMMDD-HHMMSS" (v+year/month/day-hour/minute/second).
	Version string `json:"version"`
	// Internal identifier of a function version.
	ImageName string `json:"image_name"`
	// Agency used by the function. You need to create an agency on the IAM console.
	// This field is mandatory when a function needs to access other services.
	Xrole string `json:"xrole"`
	// Agency used by the function app. You need to create an agency on the IAM console.
	// This field is mandatory when a function needs to access other services.
	AppXrole string `json:"app_xrole"`
	// Time when the function was last updated.
	LastModified string `json:"last_modified"`
	// VPC ID.
	FuncVpcId string `json:"func_vpc_id"`
	// Whether the function is disabled.
	// + 0: The function is disabled.
	// + -1: The function is enabled.
	Concurrency int `json:"concurrency"`
	// Number of concurrent instances.
	ConcurrentNum int `json:"concurrent_num"`
	// Function policy configuration.
	StrategyConfig StrategyConfig `json:"strategy_config"`
	// Initializer of the function. It is in the format of "xx.xx" and must contain a period (.).
	// For example, for Node.js function myfunction.initializer, the file name is myfunction.js,
	// and the initialization function is initializer.
	InitializerHandler string `json:"initializer_handler"`
	// Maximum duration the function can be initialized. Value range: 1s–300s.
	InitializerTimeout int `json:"initializer_timeout"`
	// Whether long-term running is supported.
	LongTime bool `json:"long_time"`
	// Return struct of the asynchronous execution notification settings.
	FunctionAsyncConfig FunctionAsyncConfig `json:"function_async_config"`
	// Function type.
	Type string `json:"type"`
	// Whether to enable cloud-based debugging.
	EnableCloudDebug string `json:"enable_cloud_debug"`
	// Whether to enable dynamic memory allocation.
	EnableDynamicMemory bool `json:"enable_dynamic_memory"`
	// Enterprise project ID. This parameter is mandatory if you create a function as an enterprise user.
	EnterpriseProjectId string `json:"enterprise_project_id"`
	// Whether stateful functions are supported. If they are supported, set this parameter to true.
	// This parameter is supported in FunctionGraph v2.
	IsStatefuleFunction bool `json:"is_stateful_function"`
	// Whether to allow authentication information in the request header.
	EnableAuthInHeader bool `json:"enable_auth_in_header"`
	// Container image.
	CustomImage CustomImage `json:"custom_image"`
	// Whether to enable idle mode for reserved instances.
	ReservedInstanceIdleMode bool `json:"reserved_instance_idle_mode"`
}

// StrategyConfig is the structure that represents the function policy definition.
type StrategyConfig struct {
	// Maximum number of instances for a single function.
	// + -1: The function has unlimited instances.
	// + 0: The function is disabled.
	// For v1, the value can be 0 or –1; for v2, it ranges from –1 to 1000.
	Concurrency int `json:"concurrency"`
	// Maximum number of concurrent requests per instance.
	// This parameter is supported only by v2.
	// The value ranges from –1 to 1000.
	ConcurrentNum int `json:"concurrent_num"`
}

// FunctionAsyncConfig is the structure that represents the asynchronous execution notification settings
type FunctionAsyncConfig struct {
	// Maximum validity period of a message. Value range: 60–86,400. Unit: second.
	MaxAsyncEventAgeInSecond int `json:"max_async_event_age_in_second"`
	// Maximum number of retry attempts to be made if asynchronous invocation fails.
	// Default value: 3. Value range: 0–8.
	MaxAsyncRetryAttempts int `json:"max_async_retry_attempts"`
	// Asynchronous invocation target.
	DestinationConfig FuncAsyncDestinationConfig `json:"destination_config"`
	// Time when asynchronous execution notification was configured.
	CreatedTime string `json:"created_time"`
	// Time when the asynchronous execution notification settings were last modified.
	LastModifed string `json:"last_modified"`
}

// FuncAsyncDestinationConfig is the structure that represents the asynchronous invocation target.
type FuncAsyncDestinationConfig struct {
	// Target to be invoked when a function is successfully executed.
	OnSuccess FuncDestinationConfig `json:"on_success"`
	// Target to be invoked when a function fails to be executed due to a system error or an internal error.
	OnFailure FuncDestinationConfig `json:"on_failure"`
}

// FuncDestinationConfig is the structure that represents the destination configuration details for function.
type FuncDestinationConfig struct {
	// Object type.
	// + OBS
	// + SMN
	// + DIS
	// + FunctionGraph
	Destination string `json:"destination"`
	// Parameters (in JSON format) corresponding to the target service.
	// + OBS: Parameters related to the bucket name, object directory prefix, and object expiration time are included.
	// The object expiration time ranges from 0 to 365 days. If the value is 0, the object will not expire.
	// + SMN: The topic_urn parameter is included.
	// + DIS: The stream_name parameter is included.
	// + FunctionGraph: The func_urn parameter is included.
	Param string `json:"param"`
}

// CustomImage is the structure that represents the container image details.
type CustomImage struct {
	// Whether to enable this feature.
	Enabled bool `json:"enabled"`
	// Image address.
	Image string `json:"image"`
	// Command for starting a container image.
	Command string `json:"command"`
	// Command line parameter for starting a container image.
	Args string `json:"args"`
	// Working directory of an image container.
	WorkingDir string `json:"working_dir"`
	// User ID of an image container.
	UID string `json:"uid"`
	// User group ID of an image container.
	GID string `json:"gid"`
}

type pageInfo struct {
	// Version list.
	Versions []Version `json:"versions"`
	// Next record location.
	NextMarker int `json:"next_marker"`
	// Total number of versions.
	Count int `json:"count"`
}

type VersionPage struct {
	pagination.MarkerPageBase
}

// IsEmpty returns true if a list result no version.
func (r VersionPage) IsEmpty() (bool, error) {
	resp, err := extractPageInfo(r)
	return len(resp.Versions) == 0, err
}

// LastMarker returns the last marker index in a pageInfo.
func (r VersionPage) LastMarker() (string, error) {
	resp, err := extractPageInfo(r)
	if err != nil {
		return "", err
	}
	if resp.NextMarker == resp.Count {
		return "", nil
	}
	return strconv.Itoa(resp.NextMarker), nil
}

// extractPageInfo is a method which to extract the response of the page information.
func extractPageInfo(r pagination.Page) (*pageInfo, error) {
	var s pageInfo
	err := r.(VersionPage).Result.ExtractInto(&s)
	return &s, err
}
