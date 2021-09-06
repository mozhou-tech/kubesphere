/*
 * HCS API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

type RegistryValue struct {

	Key *RegistryKey `json:"Key,omitempty"`

	Name string `json:"Name,omitempty"`

	Type_ string `json:"Type,omitempty"`

	//  One and only one value type must be set.
	StringValue string `json:"StringValue,omitempty"`

	BinaryValue string `json:"BinaryValue,omitempty"`

	DWordValue int32 `json:"DWordValue,omitempty"`

	QWordValue int32 `json:"QWordValue,omitempty"`

	//  Only used if RegistryValueType is CustomType  The data is in BinaryValue
	CustomType int32 `json:"CustomType,omitempty"`
}