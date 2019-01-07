package scoutred

import (
	"time"

	"github.com/scoutred/app/geojson"
)

type Permit struct {
	ID                   *int32                  `json:"id"`
	ParcelID             *int64                  `json:"parcelId"`
	Jurisdiction         *Jurisdiction           `json:"jurisdiction"`
	JurisdictionPermitID *string                 `json:"jurisdictionPermitId"`
	Description          *string                 `json:"description"`
	AppliedDate          *time.Time              `json:"appliedDate"`
	IssuedDate           *time.Time              `json:"issueDate"`
	CompletedDate        *time.Time              `json:"completedDate"`
	OriginalAddress1     *string                 `json:"originalAddress1"`
	OriginalAddress2     *string                 `json:"originalAddress2"`
	OriginalCity         *string                 `json:"originalCity"`
	OriginalState        *string                 `json:"originalState"`
	OriginalZip          *string                 `json:"originalZip"`
	ClassRaw             *string                 `json:"classRaw"`
	ClassMapped          *string                 `json:"classMapped"`
	StatusCurrentRaw     *string                 `json:"statusCurrentRaw"`
	StatusCurrentID      *int32                  `json:"statusCurrentId"`
	WorkClass            *string                 `json:"workClass"`
	WorkClassMapped      *string                 `json:"workClassMapped"`
	TypeRaw              *string                 `json:"typeRaw"`
	TypeID               *int32                  `json:"typeId"`
	TypeDescription      *string                 `json:"typeDescription"`
	StatusDate           *time.Time              `json:"statusDate"`
	TotalSF              *int32                  `json:"totalSf"`
	Link                 *string                 `json:"link"`
	EstimatedProjectCost *float64                `json:"estimatedProjectCost"`
	HousingUnits         *int32                  `json:"housingUnits"`
	APN                  *string                 `json:"apn"`
	ProposedUse          *string                 `json:"proposedUse"`
	AddedSF              *int32                  `json:"addedSf"`
	RemovedSF            *int32                  `json:"removedSf"`
	MasterPermitNumber   *string                 `json:"masterPermitNumber"`
	ExpiredDate          *time.Time              `json:"expiredDate"`
	COIssuedDate         *time.Time              `json:"coIssueDate"`
	HoldDate             *time.Time              `json:"holdDate"`
	VoidDate             *time.Time              `json:"voidDate"`
	ProjectName          *string                 `json:"projectName"`
	ProjectID            *string                 `json:"projectId"`
	TotalFinishedSF      *int32                  `json:"totalFinishedSf"`
	TotalUnfinishedSF    *int32                  `json:"totalUnfinishedSf"`
	TotalHeatedSF        *int32                  `json:"totalHeatedSf"`
	TotalUnheatedSF      *int32                  `json:"totalUnheatedSf"`
	TotalAccessorySF     *int32                  `json:"totalAccessorySf"`
	TotalSprinkledSF     *int32                  `json:"totalSprinkledSf"`
	ExtraFileds          *map[string]interface{} `json:"extraFields"`
	Publisher            *string                 `json:"publisher"`
	Fee                  *int32                  `json:"fee"`
	Contractor           *Contractor             `json:"contractor"`
	Geom                 *geojson.Point          `json:"geom"`
	Created              *time.Time              `json:"created"`
	Updated              *time.Time              `json:"updated"`
}
