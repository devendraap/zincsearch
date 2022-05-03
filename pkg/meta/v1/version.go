/* Copyright 2022 Zinc Labs Inc. and Contributors
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// export VERSION=`git describe --tags --always`
// export DATE=`date -u '+%Y-%m-%d_%I:%M:%S%p-GMT'`
// export COMMIT_HASH=`git rev-parse HEAD`

// go build -a -ldflags "-X github.com/zinclabs/zinc/pkg/meta/v1.Version=$VERSION" -o zinc
var (
	// Version is the version of the program - Git tag - git describe --tags --always
	Version = "v0.0.0"
	// Build is the build number of the program as generated by the CI system
	Build = "0"
	// CommitHash is the commit hash of the program - git rev-parse HEAD
	CommitHash = "0"
	// Branch is the branch of the program
	Branch = "0"
	// Date is the date of the build - date -u '+%Y-%m-%d_%I:%M:%S%p-GMT'
	BuildDate = "0"
)

// GetVersion returns the version of the program
func GetVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Version":    Version,
		"Build":      Build,
		"CommitHash": CommitHash,
		"Branch":     Branch,
		"BuildDate":  BuildDate,
	})
}
