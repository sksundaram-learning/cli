// Copyright 2017 Monax Industries Limited
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

// ConvertMapStringIntToPermissions converts a map of string-integer pairs and a slice of
// strings for the roles to an AccountPermissions type.  The integer needs to be greater
// than zero to set the permission.  For all unmentioned permissions the ZeroBasePermissions
// is defaulted to.
// TODO: [ben] re-evaluate the use of int for setting the permission.
func ConvertPermissionsMapAndRolesToAccountPermissions(permissions map[string]int64, roles []string) (*AccountPermissions, error) {
	var err error
	accountPermissions := &AccountPermissions{}
	accountPermissions.Base, err = convertPermissionsMapStringIntToBasePermissions(permissions)
	if err != nil {
		return nil, err
	}
	accountPermissions.Roles = roles
	return accountPermissions, nil
}

// convertPermissionsMapStringIntToBasePermissions converts a map of string-integer pairs to
// BasePermissions.
func convertPermissionsMapStringIntToBasePermissions(permissions map[string]int64) (BasePermissions, error) {
	// initialise basePermissions as ZeroBasePermissions
	basePermissions := ZeroBasePermissions

	for permissionName, value := range permissions {
		permissionsFlag, err := PermStringToFlag(permissionName)
		if err != nil {
			return basePermissions, err
		}
		// sets the permissions bitflag and the setbit flag for the permission.
		basePermissions.Set(permissionsFlag, value > 0)
	}

	return basePermissions, nil
}
