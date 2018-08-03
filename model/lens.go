package model

type Lens struct {
	ID                     int    `json:"id"`
	Name                   string `json:"name"`
	LensType               string `json:"lens_type"`
	ReleaseDate            string `json:"release_date"`
	FocalLength            string `json:"focal_length"`
	EqFullFrameFocalLength string `json:"eq_full_frame_focal_length"`
	MaximumAperture        string `json:"maximum_aperture"`
	MinimumAperture        string `json:"minimum_aperture"`
	ViewingAngle           string `json:"viewing_angle"`
	LensStructure          string `json:"lens_structure"`
	AsphericalLens         string `json:"aspherical_lens"`
	LowDispersionLens      string `json:"low_dispersion_lens"`
	ApertureBlade          string `json:"aperture_blade"`
	RecentFocusDistance    string `json:"recent_focus_distance"`
	MaximumMagnification   string `json:"maximum_magnification"`
	FilterAperture         string `json:"filter_aperture"`
	LensMount              string `json:"lens_mount"`
	ImageStabilizer        string `json:"image_stabilizer"`
	Colour                 string `json:"colour"`
	Size                   string `json:"size"`
	Weight                 string `json:"weight"`
	Annex                  string `json:"annex"`
	Remarks                string `json:"remarks"`
}
