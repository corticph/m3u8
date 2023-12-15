package decoder

import (
	"bytes"
	"errors"
	"strconv"
)

// #EXT-X-PART:<attribute-list>

// PartialSegmentTag implements CustomTag and almost CustomDecoder (see customdecoder.PartialSegmentTag)
// interfaces.
type PartialSegmentTag struct {
	URI         string
	Duration    float64
	Independent bool
}

// TagName should return the full indentifier including the leading '#' and trailing ':'
// if the tag also contains a value or attribute list
func (tag *PartialSegmentTag) TagName() string {
	return "#EXT-X-PART:"
}

// DecodeToStruct decodes the input string to the internal structure. The line
// will be the entire matched line, including the identifier.
func (tag *PartialSegmentTag) DecodeToStruct(line string) (*PartialSegmentTag, error) {
	var err error

	// Since this is a Segment tag, we want to create a new tag every time it is decoded
	// as there can be one for each segment with
	newTag := new(PartialSegmentTag)

	for k, v := range DecodeAttributeList(line[12:]) {
		switch k {
		case "URI":
			newTag.URI = v
		case "DURATION":
			duration, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return newTag, nil
			}
			newTag.Duration = duration
		case "INDEPENDENT":
			if v == "YES" {
				newTag.Independent = true
			} else if v == "NO" {
				newTag.Independent = false
			} else {
				err = errors.New("Valid strings for INDEPENDENT attribute are YES and NO.")
			}
		}
	}

	return newTag, err
}

// SegmentTag is a playlist tag example.
func (tag *PartialSegmentTag) SegmentTag() bool {
	return true
}

// Encode encodes the structure to the text result.
func (tag *PartialSegmentTag) Encode() *bytes.Buffer {
	buf := new(bytes.Buffer)

	if tag.URI != "" {
		buf.WriteString(tag.TagName())
		buf.WriteString("DURATION=")
		buf.WriteString(strconv.FormatFloat(tag.Duration, 'f', 3, 32))
		buf.WriteString(",URI=\"")
		buf.WriteString(tag.URI)
		buf.WriteString("\",INDEPENDENT=")
		if tag.Independent {
			buf.WriteString("YES")
		} else {
			buf.WriteString("NO")
		}
	}

	return buf
}

// String implements Stringer interface.
func (tag *PartialSegmentTag) String() string {
	return tag.Encode().String()
}
