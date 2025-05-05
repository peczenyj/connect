// Copyright 2024 Redpanda Data, Inc.
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

package changelog

import (
	"github.com/go-viper/mapstructure/v2"
	"github.com/r3labs/diff/v3"
	"go.uber.org/multierr"

	"github.com/redpanda-data/benthos/v4/public/bloblang"
)

func init() {
	diffSpec := bloblang.NewPluginSpec().
		Beta().
		Category("Object & Array Manipulation").
		Description(`Create a diff by comparing the current value with the given one. Wraps the github.com/r3labs/diff/v3 package. See its https://pkg.go.dev/github.com/r3labs/diff/v3[docs^] for more information.`).
		Version("4.25.0").
		Param(bloblang.NewAnyParam("other").Description("The value to compare against."))

	if err := bloblang.RegisterMethodV2("diff", diffSpec, func(args *bloblang.ParsedParams) (bloblang.Method, error) {
		other, err := args.Get("other")
		if err != nil {
			return nil, err
		}

		return func(v any) (any, error) {
			if v == nil {
				return nil, nil
			}
			cl, err := diff.Diff(v, other)
			if err != nil {
				return nil, err
			}

			var result []map[string]any
			if err := mapstructure.Decode(cl, &result); err != nil {
				return nil, err
			}

			return result, nil
		}, nil
	}); err != nil {
		panic(err)
	}

	patchSpec := bloblang.NewPluginSpec().
		Beta().
		Category("Object & Array Manipulation").
		Description(`Create a diff by comparing the current value with the given one. Wraps the github.com/r3labs/diff/v3 package. See its https://pkg.go.dev/github.com/r3labs/diff/v3[docs^] for more information.`).
		Version("4.25.0").
		Param(bloblang.NewAnyParam("changelog").Description("The changelog to apply."))

	if err := bloblang.RegisterMethodV2("patch", patchSpec, func(args *bloblang.ParsedParams) (bloblang.Method, error) {
		clog, err := args.Get("changelog")
		if err != nil {
			return nil, err
		}

		var cl diff.Changelog
		if err := mapstructure.Decode(clog, &cl); err != nil {
			return nil, err
		}

		return func(v any) (any, error) {
			if v == nil {
				return nil, nil
			}

			pl := diff.Patch(cl, &v)

			if pl.HasErrors() {
				var e error
				for _, ple := range pl {
					if ple.Errors != nil {
						if err := multierr.Append(e, ple.Errors); err != nil {
							return nil, err
						}
					}
				}

				return nil, e
			}

			return v, nil
		}, nil
	}); err != nil {
		panic(err)
	}
}
