/*
Copyright 2019 HAProxy Technologies

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

package parsers

import (
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type ACL struct {
	data []types.ACL
}

func (h *ACL) parse(line string, parts []string, comment string) (*types.ACL, error) {
	if len(parts) >= 3 {
		data := &types.ACL{
			Name:      parts[1],
			Criterion: parts[2],
			Value:     strings.Join(parts[3:], " "),
			Comment:   comment,
		}
		return data, nil
	}
	return nil, &errors.ParseError{Parser: "ACLLines", Line: line}
}

func (h *ACL) Result() ([]common.ReturnResultLine, error) {
	if len(h.data) == 0 {
		return nil, errors.ErrFetch
	}
	result := make([]common.ReturnResultLine, len(h.data))
	for index, req := range h.data {
		var sb strings.Builder
		sb.WriteString("acl ")
		sb.WriteString(req.Name)
		sb.WriteString(" ")
		sb.WriteString(req.Criterion)
		sb.WriteString(" ")
		sb.WriteString(req.Value)
		result[index] = common.ReturnResultLine{
			Data:    sb.String(),
			Comment: req.Comment,
		}
	}
	return result, nil
}
