package decoder

import "bytes"

// #EXT-X-PRELOAD-HINT:<attribute-list>

// PreloadHintTag implements CustomTag and almost CustomDecoder (see customdecoder.PreloadHintTag)
// interfaces.
type PreloadHintTag struct {
	URI  string
	Type string
}

// TagName should return the full indentifier including the leading '#' and trailing ':'
// if the tag also contains a value or attribute list
func (tag *PreloadHintTag) TagName() string {
	return "#EXT-X-PRELOAD-HINT:"
}

// DecodeToStruct decodes the input string to the internal structure. The line
// will be the entire matched line, including the identifier.
func (tag *PreloadHintTag) DecodeToStruct(line string) (*PreloadHintTag, error) {
	var err error

	// Since this is a Segment tag, we want to create a new tag every time it is decoded
	// as there can be one for each segment with
	newTag := new(PreloadHintTag)

	for k, v := range DecodeAttributeList(line[20:]) {
		switch k {
		case "URI":
			newTag.URI = v
		case "TYPE":
			newTag.Type = v
		}
	}

	return newTag, err
}

// SegmentTag is a playlist tag example.
func (tag *PreloadHintTag) SegmentTag() bool {
	return true
}

// Encode encodes the structure to the text result.
func (tag *PreloadHintTag) Encode() *bytes.Buffer {
	buf := new(bytes.Buffer)

	if tag.URI != "" {
		buf.WriteString(tag.TagName())
		buf.WriteString("TYPE=")
		buf.WriteString(tag.Type)
		buf.WriteString(",URI=\"")
		buf.WriteString(tag.URI)
		buf.WriteString("\"")
	}

	return buf
}

// String implements Stringer interface.
func (tag *PreloadHintTag) String() string {
	return tag.Encode().String()
}
