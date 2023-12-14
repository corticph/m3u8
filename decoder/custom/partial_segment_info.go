package customdecoder

import (
	"github.com/corticph/m3u8"
	"github.com/corticph/m3u8/decoder"
)

// #EXT-X-PART-INF:<attribute-list>

// PartialSegmentInfoTag implements both CustomTag and CustomDecoder
// interfaces.
type PartialSegmentInfoTag struct {
	*decoder.PartialSegmentInfoTag
}

// Decode decodes the input string to the internal structure. The line
// will be the entire matched line, including the identifier.
func (tag *PartialSegmentInfoTag) Decode(line string) (m3u8.CustomTag, error) {
	s, err := tag.DecodeToStruct(line)
	if err != nil {
		return nil, err
	}

	return m3u8.CustomTag(s), err
}
