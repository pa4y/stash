// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"
)

type ConfigGeneralInput struct {
	// Array of file paths to content
	Stashes []string `json:"stashes"`
}

type ConfigGeneralResult struct {
	// Array of file paths to content
	Stashes []string `json:"stashes"`
}

type FindFilterType struct {
	Q         *string            `json:"q"`
	Page      *int               `json:"page"`
	PerPage   *int               `json:"per_page"`
	Sort      *string            `json:"sort"`
	Direction *SortDirectionEnum `json:"direction"`
}

type FindGalleriesResultType struct {
	Count     int       `json:"count"`
	Galleries []Gallery `json:"galleries"`
}

type FindPerformersResultType struct {
	Count      int         `json:"count"`
	Performers []Performer `json:"performers"`
}

type FindSceneMarkersResultType struct {
	Count        int           `json:"count"`
	SceneMarkers []SceneMarker `json:"scene_markers"`
}

type FindScenesResultType struct {
	Count  int     `json:"count"`
	Scenes []Scene `json:"scenes"`
}

type FindStudiosResultType struct {
	Count   int      `json:"count"`
	Studios []Studio `json:"studios"`
}

type GalleryFilesType struct {
	Index int     `json:"index"`
	Name  *string `json:"name"`
	Path  *string `json:"path"`
}

type MarkerStringsResultType struct {
	Count int    `json:"count"`
	ID    string `json:"id"`
	Title string `json:"title"`
}

type PerformerCreateInput struct {
	Name         *string `json:"name"`
	URL          *string `json:"url"`
	Birthdate    *string `json:"birthdate"`
	Ethnicity    *string `json:"ethnicity"`
	Country      *string `json:"country"`
	EyeColor     *string `json:"eye_color"`
	Height       *string `json:"height"`
	Measurements *string `json:"measurements"`
	FakeTits     *string `json:"fake_tits"`
	CareerLength *string `json:"career_length"`
	Tattoos      *string `json:"tattoos"`
	Piercings    *string `json:"piercings"`
	Aliases      *string `json:"aliases"`
	Twitter      *string `json:"twitter"`
	Instagram    *string `json:"instagram"`
	Favorite     *bool   `json:"favorite"`
	// This should be base64 encoded
	Image string `json:"image"`
}

type PerformerFilterType struct {
	// Filter by favorite
	FilterFavorites *bool `json:"filter_favorites"`
}

type PerformerUpdateInput struct {
	ID           string  `json:"id"`
	Name         *string `json:"name"`
	URL          *string `json:"url"`
	Birthdate    *string `json:"birthdate"`
	Ethnicity    *string `json:"ethnicity"`
	Country      *string `json:"country"`
	EyeColor     *string `json:"eye_color"`
	Height       *string `json:"height"`
	Measurements *string `json:"measurements"`
	FakeTits     *string `json:"fake_tits"`
	CareerLength *string `json:"career_length"`
	Tattoos      *string `json:"tattoos"`
	Piercings    *string `json:"piercings"`
	Aliases      *string `json:"aliases"`
	Twitter      *string `json:"twitter"`
	Instagram    *string `json:"instagram"`
	Favorite     *bool   `json:"favorite"`
	// This should be base64 encoded
	Image *string `json:"image"`
}

type SceneFileType struct {
	Size       *string  `json:"size"`
	Duration   *float64 `json:"duration"`
	VideoCodec *string  `json:"video_codec"`
	AudioCodec *string  `json:"audio_codec"`
	Width      *int     `json:"width"`
	Height     *int     `json:"height"`
	Framerate  *float64 `json:"framerate"`
	Bitrate    *int     `json:"bitrate"`
}

type SceneFilterType struct {
	// Filter by rating
	Rating *int `json:"rating"`
	// Filter by resolution
	Resolution *ResolutionEnum `json:"resolution"`
	// Filter to only include scenes which have markers. `true` or `false`
	HasMarkers *string `json:"has_markers"`
	// Filter to only include scenes missing this property
	IsMissing *string `json:"is_missing"`
	// Filter to only include scenes with this studio
	StudioID *string `json:"studio_id"`
	// Filter to only include scenes with these tags
	Tags []string `json:"tags"`
	// Filter to only include scenes with this performer
	PerformerID *string `json:"performer_id"`
}

type SceneMarkerCreateInput struct {
	Title        string   `json:"title"`
	Seconds      float64  `json:"seconds"`
	SceneID      string   `json:"scene_id"`
	PrimaryTagID string   `json:"primary_tag_id"`
	TagIds       []string `json:"tag_ids"`
}

type SceneMarkerFilterType struct {
	// Filter to only include scene markers with this tag
	TagID *string `json:"tag_id"`
	// Filter to only include scene markers with these tags
	Tags []string `json:"tags"`
	// Filter to only include scene markers attached to a scene with these tags
	SceneTags []string `json:"scene_tags"`
	// Filter to only include scene markers with these performers
	Performers []string `json:"performers"`
}

type SceneMarkerTag struct {
	Tag          Tag           `json:"tag"`
	SceneMarkers []SceneMarker `json:"scene_markers"`
}

type SceneMarkerUpdateInput struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	Seconds      float64  `json:"seconds"`
	SceneID      string   `json:"scene_id"`
	PrimaryTagID string   `json:"primary_tag_id"`
	TagIds       []string `json:"tag_ids"`
}

type ScenePathsType struct {
	Screenshot  *string `json:"screenshot"`
	Preview     *string `json:"preview"`
	Stream      *string `json:"stream"`
	Webp        *string `json:"webp"`
	Vtt         *string `json:"vtt"`
	ChaptersVtt *string `json:"chapters_vtt"`
}

type SceneUpdateInput struct {
	ClientMutationID *string  `json:"clientMutationId"`
	ID               string   `json:"id"`
	Title            *string  `json:"title"`
	Details          *string  `json:"details"`
	URL              *string  `json:"url"`
	Date             *string  `json:"date"`
	Rating           *int     `json:"rating"`
	StudioID         *string  `json:"studio_id"`
	GalleryID        *string  `json:"gallery_id"`
	PerformerIds     []string `json:"performer_ids"`
	TagIds           []string `json:"tag_ids"`
}

// A performer from a scraping operation...
type ScrapedPerformer struct {
	Name         *string `json:"name"`
	URL          *string `json:"url"`
	Twitter      *string `json:"twitter"`
	Instagram    *string `json:"instagram"`
	Birthdate    *string `json:"birthdate"`
	Ethnicity    *string `json:"ethnicity"`
	Country      *string `json:"country"`
	EyeColor     *string `json:"eye_color"`
	Height       *string `json:"height"`
	Measurements *string `json:"measurements"`
	FakeTits     *string `json:"fake_tits"`
	CareerLength *string `json:"career_length"`
	Tattoos      *string `json:"tattoos"`
	Piercings    *string `json:"piercings"`
	Aliases      *string `json:"aliases"`
}

type StatsResultType struct {
	SceneCount     int `json:"scene_count"`
	GalleryCount   int `json:"gallery_count"`
	PerformerCount int `json:"performer_count"`
	StudioCount    int `json:"studio_count"`
	TagCount       int `json:"tag_count"`
}

type StudioCreateInput struct {
	Name string  `json:"name"`
	URL  *string `json:"url"`
	// This should be base64 encoded
	Image string `json:"image"`
}

type StudioUpdateInput struct {
	ID   string  `json:"id"`
	Name *string `json:"name"`
	URL  *string `json:"url"`
	// This should be base64 encoded
	Image *string `json:"image"`
}

type TagCreateInput struct {
	Name string `json:"name"`
}

type TagDestroyInput struct {
	ID string `json:"id"`
}

type TagUpdateInput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ResolutionEnum string

const (
	// 240p
	ResolutionEnumLow ResolutionEnum = "LOW"
	// 480p
	ResolutionEnumStandard ResolutionEnum = "STANDARD"
	// 720p
	ResolutionEnumStandardHd ResolutionEnum = "STANDARD_HD"
	// 1080p
	ResolutionEnumFullHd ResolutionEnum = "FULL_HD"
	// 4k
	ResolutionEnumFourK ResolutionEnum = "FOUR_K"
)

var AllResolutionEnum = []ResolutionEnum{
	ResolutionEnumLow,
	ResolutionEnumStandard,
	ResolutionEnumStandardHd,
	ResolutionEnumFullHd,
	ResolutionEnumFourK,
}

func (e ResolutionEnum) IsValid() bool {
	switch e {
	case ResolutionEnumLow, ResolutionEnumStandard, ResolutionEnumStandardHd, ResolutionEnumFullHd, ResolutionEnumFourK:
		return true
	}
	return false
}

func (e ResolutionEnum) String() string {
	return string(e)
}

func (e *ResolutionEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ResolutionEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ResolutionEnum", str)
	}
	return nil
}

func (e ResolutionEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SortDirectionEnum string

const (
	SortDirectionEnumAsc  SortDirectionEnum = "ASC"
	SortDirectionEnumDesc SortDirectionEnum = "DESC"
)

var AllSortDirectionEnum = []SortDirectionEnum{
	SortDirectionEnumAsc,
	SortDirectionEnumDesc,
}

func (e SortDirectionEnum) IsValid() bool {
	switch e {
	case SortDirectionEnumAsc, SortDirectionEnumDesc:
		return true
	}
	return false
}

func (e SortDirectionEnum) String() string {
	return string(e)
}

func (e *SortDirectionEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortDirectionEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortDirectionEnum", str)
	}
	return nil
}

func (e SortDirectionEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
