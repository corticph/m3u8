package decoder

import (
	"bytes"
	"strconv"
)

// #EXT-X-PART-INF:<attribute-list>

// PartialSegmentInfoTag implements CustomTag and almost CustomDecoder (see customdecoder.PartialSegmentInfoTag)
// interfaces.
type PartialSegmentInfoTag struct {
	TargetDuration float64
}

// TagName should return the full indentifier including the leading '#' and trailing ':'
// if the tag also contains a value or attribute list
func (tag *PartialSegmentInfoTag) TagName() string {
	return "#EXT-X-PART-INF:"
}

// DecodeToStruct decodes the input string to the internal structure. The line
// will be the entire matched line, including the identifier.
func (tag *PartialSegmentInfoTag) DecodeToStruct(line string) (*PartialSegmentInfoTag, error) {
	var err error

	// Since this is a Segment tag, we want to create a new tag every time it is decoded
	// as there can be one for each segment with
	newTag := new(PartialSegmentInfoTag)

	for k, v := range DecodeAttributeList(line[16:]) {
		switch k {
		case "PART-TARGET":
			duration, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return newTag, nil
			}
			newTag.TargetDuration = duration
		}
	}

	return newTag, err
}

// SegmentTag is a playlist tag example.
func (tag *PartialSegmentInfoTag) SegmentTag() bool {
	return true
}

// Encode encodes the structure to the text result.
func (tag *PartialSegmentInfoTag) Encode() *bytes.Buffer {
	buf := new(bytes.Buffer)

	if tag.TargetDuration != 0 {
		buf.WriteString(tag.TagName())
		buf.WriteString("PART-TARGET=")
		buf.WriteString(strconv.FormatFloat(tag.TargetDuration, 'f', 3, 32))
	}

	return buf
}

// String implements Stringer interface.
func (tag *PartialSegmentInfoTag) String() string {
	return tag.Encode().String()
}
