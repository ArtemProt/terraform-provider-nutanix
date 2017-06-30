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

type VmSnapshotDefStatusSnapshotFileList struct {

	// Pathname of the file that participated in the snapshot.
	FilePath string `json:"file_path,omitempty" bson:"file_path,omitempty"`

	// Pathname of the snapshot of the file.
	SnapshotFilePath string `json:"snapshot_file_path,omitempty" bson:"snapshot_file_path,omitempty"`
}
