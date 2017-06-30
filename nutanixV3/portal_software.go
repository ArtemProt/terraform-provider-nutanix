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

import (
	"time"
)

// Software information from Nutanix Portal
type PortalSoftware struct {

	// List of Prism Element compatible versions
	CompatiblePeVersionList []string `json:"compatible_pe_version_list,omitempty" bson:"compatible_pe_version_list,omitempty"`

	// List of software versions that this version can be upgraded from
	CompatibleVersionList []string `json:"compatible_version_list,omitempty" bson:"compatible_version_list,omitempty"`

	// MD5 checksum of the software file
	Md5sum string `json:"md5sum,omitempty" bson:"md5sum,omitempty"`

	// Release date of this software in RFC3339 format.
	ReleaseDate time.Time `json:"release_date,omitempty" bson:"release_date,omitempty"`

	// URL to point to the support portal release note of this software. Currently only set and used for NOS releases
	ReleaseNoteUrl string `json:"release_note_url,omitempty" bson:"release_note_url,omitempty"`

	// Total size of the software file in mebibytes
	SizeInMib int64 `json:"size_in_mib,omitempty" bson:"size_in_mib,omitempty"`

	// Software type
	SoftwareType string `json:"software_type,omitempty" bson:"software_type,omitempty"`

	UpgradeNotification UpgradeNotification `json:"upgrade_notification,omitempty" bson:"upgrade_notification,omitempty"`

	// Software version string
	Version string `json:"version,omitempty" bson:"version,omitempty"`
}
