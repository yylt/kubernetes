/*
Copyright 2015 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package hash

import (
	"hash"
	"regexp"

	"github.com/davecgh/go-spew/spew"
)

// DeepHashObject writes specified object to hash using the spew library
// which follows pointers and prints actual values of the nested objects
// ensuring the hash does not change when a pointer changes.
func DeepHashObject(hasher hash.Hash, objectToWrite interface{}) {
	hasher.Reset()
	printer := spew.ConfigState{
		Indent:         " ",
		SortKeys:       true,
		DisableMethods: true,
		SpewKeys:       true,
	}
	printer.Fprintf(hasher, "%#v", objectToWrite)
}

// LegacyDeepHashObject writes specified object to hash using the spew library
// which follows pointers and prints actual values of the nested objects
// ensuring the hash does not change when a pointer changes.
// LegacyDeepHashObject differs from DeepHashObject in this:
// It use spew's formatter to get the LegacyContainer string, replace Legacy
// strings and then write the result to hasher.
func LegacyDeepHashObject(hasher hash.Hash, objectToWrite interface{}) {
	hasher.Reset()
	printer := spew.ConfigState{
		Indent:         " ",
		SortKeys:       true,
		DisableMethods: true,
		SpewKeys:       true,
	}
	s := printer.Sprintf("%#v", objectToWrite)
	re, _ := regexp.Compile("LegacyContainer")
	s = re.ReplaceAllString(s, "Container")
	re, _ = regexp.Compile("LegacySecurityContext")
	s = re.ReplaceAllString(s, "SecurityContext")
	re, _ = regexp.Compile("LegacyVolumeMount")
	s = re.ReplaceAllString(s, "VolumeMount")
	hasher.Write([]byte(s))
}
