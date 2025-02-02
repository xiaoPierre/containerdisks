/*
 * HCS API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

// Passed to HcsRpc_ModifyProcess
type ProcessModifyRequest struct {
	Operation string `json:"Operation,omitempty"`

	ConsoleSize *ConsoleSize `json:"ConsoleSize,omitempty"`

	CloseHandle *CloseHandle `json:"CloseHandle,omitempty"`
}
