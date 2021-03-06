// Copyright 2017 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package compare

import (
	"github.com/spf13/hugo/deps"
	"github.com/spf13/hugo/tpl/internal"
)

const name = "compare"

func init() {
	f := func(d *deps.Deps) *internal.TemplateFuncsNamespace {
		ctx := New()

		ns := &internal.TemplateFuncsNamespace{
			Name:    name,
			Context: func() interface{} { return ctx },
		}

		ns.AddMethodMapping(ctx.Default,
			[]string{"default"},
			[][2]string{
				{`{{ "Hugo Rocks!" | default "Hugo Rules!" }}`, `Hugo Rocks!`},
				{`{{ "" | default "Hugo Rules!" }}`, `Hugo Rules!`},
			},
		)

		ns.AddMethodMapping(ctx.Eq,
			[]string{"eq"},
			[][2]string{
				{`{{ if eq .Section "blog" }}current{{ end }}`, `current`},
			},
		)

		ns.AddMethodMapping(ctx.Ge,
			[]string{"ge"},
			[][2]string{},
		)

		ns.AddMethodMapping(ctx.Gt,
			[]string{"gt"},
			[][2]string{},
		)

		ns.AddMethodMapping(ctx.Le,
			[]string{"le"},
			[][2]string{},
		)

		ns.AddMethodMapping(ctx.Lt,
			[]string{"lt"},
			[][2]string{},
		)

		ns.AddMethodMapping(ctx.Ne,
			[]string{"ne"},
			[][2]string{},
		)

		return ns

	}

	internal.AddTemplateFuncsNamespace(f)
}
