/*
 * Nutanix Intentful API
 *
 * Move programming from the user to the machine
 *
 * OpenAPI spec version: 3.0.0
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package nutanix

// Container options required to run the container.
type ContainerOptions struct {

	// Command to be executed in the container.
	Command string `json:"command,omitempty" bson:"command,omitempty"`

	// Additional options to be used for container.
	ContainerAdditionalOptions string `json:"container_additional_options,omitempty" bson:"container_additional_options,omitempty"`

	// Name of container runtime.
	ContainerRuntime string `json:"container_runtime,omitempty" bson:"container_runtime,omitempty"`

	// Memory in MiBs.
	MemoryMib int64 `json:"memory_mib,omitempty" bson:"memory_mib,omitempty"`

	// Number of cpu shares.
	NumCpuShares int64 `json:"num_cpu_shares,omitempty" bson:"num_cpu_shares,omitempty"`

	// Container restart policy.
	RestartPolicy string `json:"restart_policy,omitempty" bson:"restart_policy,omitempty"`
}
