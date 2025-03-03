/*
   Copyright (c) 2020 gingfrederik
   Copyright (c) 2021 Gonzalo Fernandez-Victorio
   Copyright (c) 2021 Basement Crowd Ltd (https://www.basementcrowd.com)
   Copyright (c) 2023 Fumiama Minamoto (源文雨)

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package docx

import (
	"errors"
	"strings"
)

// RangeRelationships goes through each doc relation
func (f *Docx) RangeRelationships(iter func(*Relationship) error) error {
	for _, r := range f.docRelation.Relationship {
		err := iter(&r)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *Docx) RangeRelationshipsPicture(blipEmbed string) ([]byte, string, error) {
	target := ""
	for _, r := range f.docRelation.Relationship {
		if r.ID == blipEmbed {
			target = r.Target
		}
	}
	if target == "" {
		return nil, "", errors.New("not found")
	}
	targetList := strings.Split(target, "/")
	if len(targetList) <= 1 {
		return nil, "", errors.New("target value error " + target)
	}
	idx, ok := f.mediaNameIdx[targetList[len(targetList)-1]]
	if !ok {
		return nil, "", errors.New("target not found " + target)
	}
	if idx < 0 || idx >= len(f.media) {
		return nil, "", errors.New("idx out of range " + target)
	}
	return f.media[idx].Data, targetList[len(targetList)-1], nil
}
