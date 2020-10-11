// Copyright 2015 The etcd Authors
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

package flags

import (
	"strings"

	"github.com/mosheya/etcd/pkg/types"
)

type URLsValue types.URLs

// Set parses a command line set of URLs formatted like:
// http://127.0.0.1:2380,http://10.1.1.2:80
func (us *URLsValue) Set(s string) error {
	strs := strings.Split(s, ",")
	nus, err := types.NewURLs(strs)
	if err != nil {
		return err
	}

	*us = URLsValue(nus)
	return nil
}

func (us *URLsValue) String() string {
	all := make([]string, len(*us))
	for i, u := range *us {
		all[i] = u.String()
	}
	return strings.Join(all, ",")
}

func NewURLsValue(init string) *URLsValue {
	v := &URLsValue{}
	if err := v.Set(init); err != nil {
		plog.Panicf("new URLsValue should never fail: %v", err)
	}
	return v
}
