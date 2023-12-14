package customdecoder

import (
	"github.com/corticph/m3u8"
	"github.com/corticph/m3u8/decoder"
)

// #EXT-X-PART:<attribute-list>

// PartialSegmentTag implements both CustomTag and CustomDecoder
// interfaces.
type PartialSegmentTag struct {
	*decoder.PartialSegmentTag
}

// Decode decodes the input string to the internal structure. The line
// will be the entire matched line, including the identifier.
func (tag *PartialSegmentTag) Decode(line string) (m3u8.CustomTag, error) {
	s, err := tag.DecodeToStruct(line)
	if err != nil {
		return nil, err
	}

	return m3u8.CustomTag(s), err
}
