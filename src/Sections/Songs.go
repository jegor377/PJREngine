package Sections

import "bytes"

type SongsSection struct {
	Songs []string
	IsEmpty bool
}

func NewSongsSection() *SongsSection {
	return &SongsSection{
		make([]string, 0),
		false,
	}
}

func (songs *SongsSection) DecodeSongs(byteCode []byte) {
	if len(byteCode) == 0 {
		songs.IsEmpty = true
		return
	}
	songs.Songs = append(songs.Songs, "none") // first must be none
	paths := bytes.Split(byteCode, []byte{0})
	for _, e := range(paths) {
		if len(e) > 0 {
			songs.Songs = append(songs.Songs, string(e))
		}
	}
}