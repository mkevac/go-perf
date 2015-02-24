// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// generated by stringer -type=RecordType; DO NOT EDIT

package perffile

import "fmt"

const _RecordType_name = "RecordTypeMmapRecordTypeLostRecordTypeCommRecordTypeExitRecordTypeThrottleRecordTypeUnthrottleRecordTypeForkRecordTypeReadRecordTypeSamplerecordTypeMmap2"

var _RecordType_index = [...]uint8{0, 14, 28, 42, 56, 74, 94, 108, 122, 138, 153}

func (i RecordType) String() string {
	i -= 1
	if i+1 >= RecordType(len(_RecordType_index)) {
		return fmt.Sprintf("RecordType(%d)", i+1)
	}
	return _RecordType_name[_RecordType_index[i]:_RecordType_index[i+1]]
}
