package scoutred

type ZoningRegulation struct {
	ID       *int32 `json:"id"`
	ZoningID *int32 `json:"zoningId"`
	LotSize  struct {
		MinArea        *int32 `json:"minArea"`
		MaxArea        *int32 `json:"maxArea"`
		MinWidth       *int32 `json:"minWidth"`
		MinWidthCorner *int32 `json:"minWidthCorner"`
		MinFrontage    *int32 `json:"minFrontage"`
		MinDepth       *int32 `json:"minDepth"`
	} `json:"lotSize"`
	Density struct {
		SfDu     *int32  `json:"sfDu"`
		SfDuNote *string `json:"sfDuNote"`
	} `json:"density"`
	HeightLimit struct {
		Max                  *int32  `json:"max"`
		AboveEnclosedParking *int32  `json:"aboveEnclosedParking"`
		RoofFlat             *int32  `json:"roofFlat"`
		RoofPitched          *int32  `json:"roofPitched"`
		Note                 *string `json:"note"`
	} `json:"heightLimit"`
	FAR struct {
		Base        *float32 `json:"base"`
		Min         *float32 `json:"min"`
		Max         *float32 `json:"max"`
		Residential *float32 `json:"residential"`
		Commercial  *float32 `json:"commercial"`
		Mixed       *float32 `json:"mixed"`
		Note        *string  `json:"note"`
	} `json:"far"`
	Setbacks struct {
		FrontMin         *int32  `json:"frontMin"`
		FrontMinNote     *string `json:"frontMinNote"`
		FrontMax         *int32  `json:"frontMax"`
		FrontMaxNote     *string `json:"frontMaxNote"`
		InteriorSide     *int32  `json:"interiorSide"`
		InteriorSideNote *string `json:"interiorSideNote"`
		StreetSide       *int32  `json:"streetSide"`
		StreetSideNote   *string `json:"streetSideNote"`
		Rear             *int32  `json:"rear"`
		RearAlley        *int32  `json:"rearAlley"`
		RearNote         *string `json:"rearNote"`
		GeneralNote      *string `json:"generalNote"`
	} `json:"setbacks"`
	LotCoverage struct {
		Min  *float32 `json:"min"`
		Max  *float32 `json:"max"`
		Note *string  `json:"note"`
	} `json:"lotCoverage"`
}
