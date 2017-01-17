//
// Copyright 2016 Alsanium, SAS. or its affiliates. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package cloudwatchschedevt

import (
	"bytes"
	"time"
)

// UnmarshalJSON interprets the data as a RFC3339 time. It then sets *t to a
// copy of the interpreted time.
func (t *Timestamp) UnmarshalJSON(data []byte) error {
	v, err := time.Parse(time.RFC3339, string(bytes.Trim(data, `"`)))
	if err != nil {
		return err
	}

	t.Time = v
	return nil
}