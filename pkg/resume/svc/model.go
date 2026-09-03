package svc

import (
	"encoding/json"
	"strconv"
	"strings"
)

// FlexInt 吸收同一欄位在不同端點上時而為數字、時而為字串的情形
// (例如 status.id 在 /v2/search 是 "102"，在 /v2/jobs 是 102)。
type FlexInt int

// UnmarshalJSON 同時接受 JSON number 與帶引號的十進位字串。
func (f *FlexInt) UnmarshalJSON(data []byte) error {
	text := strings.Trim(string(data), `"`)
	if text == "" || text == "null" {
		*f = 0
		return nil
	}
	value, err := strconv.Atoi(text)
	if err != nil {
		return err
	}
	*f = FlexInt(value)
	return nil
}

// MarshalJSON 一律輸出成數字，讓落檔後的資料形狀一致。
func (f FlexInt) MarshalJSON() ([]byte, error) { return json.Marshal(int(f)) }

// Link 是 HATEOAS 連結。這組 _links 是探索 API 表面的主要線索。
type Link struct {
	Href string `json:"href"`
}

// Links 是回應中的 _links 區塊。
type Links struct {
	Self               *Link `json:"self,omitempty"`
	First              *Link `json:"first,omitempty"`
	Next               *Link `json:"next,omitempty"`
	Last               *Link `json:"last,omitempty"`
	Jobs               *Link `json:"jobs,omitempty"`
	Addresses          *Link `json:"addresses,omitempty"`
	Schemes            *Link `json:"schemes,omitempty"`
	ScreeningQuestions *Link `json:"screeningQuestions,omitempty"`
}

// Job 是職缺。/v2/search 的 results 元素與 /v2/jobs/{uuid} 共用同一形狀，
// 差別只在 search 不帶 description —— 要 JD 全文一定得再打一次 detail。
type Job struct {
	UUID                     string              `json:"uuid"`
	Title                    string              `json:"title"`
	Description              string              `json:"description,omitempty"`
	SourceCode               string              `json:"sourceCode,omitempty"`
	MinimumYearsExperience   int                 `json:"minimumYearsExperience,omitempty"`
	NumberOfVacancies        int                 `json:"numberOfVacancies,omitempty"`
	OtherRequirements        string              `json:"otherRequirements,omitempty"`
	WorkingHours             string              `json:"workingHours,omitempty"`
	ShiftPattern             *ShiftPattern       `json:"shiftPattern,omitempty"`
	SSOCCode                 string              `json:"ssocCode,omitempty"`
	SSOCVersion              string              `json:"ssocVersion,omitempty"`
	OccupationID             string              `json:"occupationId,omitempty"`
	SSECEQA                  string              `json:"ssecEqa,omitempty"`
	SSECFOS                  string              `json:"ssecFos,omitempty"`
	PSDURL                   string              `json:"psdUrl,omitempty"`
	Salary                   *Salary             `json:"salary,omitempty"`
	Status                   *Status             `json:"status,omitempty"`
	Address                  *Address            `json:"address,omitempty"`
	Metadata                 *Metadata           `json:"metadata,omitempty"`
	PostedCompany            *Company            `json:"postedCompany,omitempty"`
	HiringCompany            *Company            `json:"hiringCompany,omitempty"`
	Categories               []Category          `json:"categories,omitempty"`
	EmploymentTypes          []EmploymentType    `json:"employmentTypes,omitempty"`
	PositionLevels           []PositionLevel     `json:"positionLevels,omitempty"`
	Skills                   []Skill             `json:"skills,omitempty"`
	Schemes                  []Scheme            `json:"schemes,omitempty"`
	FlexibleWorkArrangements []FlexibleWork      `json:"flexibleWorkArrangements,omitempty"`
	ScreeningQuestions       []ScreeningQuestion `json:"screeningQuestions,omitempty"`
	Links                    Links               `json:"_links,omitempty"`

	// 以下三個分數只出現在 /v2/search 的結果中，用於檢視排序品質。
	Score            float64 `json:"score,omitempty"`
	TitleMatchScore  float64 `json:"title_match_score,omitempty"`
	SkillsMatchScore float64 `json:"skills_match_score,omitempty"`
	RecencyScore     float64 `json:"recency_score,omitempty"`
}

// Salary 是薪酬區間。isHideSalary 為真的職缺仍會帶推估值，不是雇主公開數字。
type Salary struct {
	Minimum int         `json:"minimum"`
	Maximum int         `json:"maximum"`
	Type    *SalaryType `json:"type,omitempty"`
}

// SalaryType 是薪酬週期，例如 Monthly、Annually。
type SalaryType struct {
	ID         FlexInt `json:"id,omitempty"`
	SalaryType string  `json:"salaryType"`
}

// Status 是職缺狀態，例如 Open、Closed。
type Status struct {
	ID        FlexInt `json:"id"`
	JobStatus string  `json:"jobStatus"`
}

// Category 是職務類別，也是 /v2/search 的 categories 篩選值來源。
type Category struct {
	ID       int    `json:"id"`
	Category string `json:"category"`
}

// EmploymentType 是聘僱型態，例如 Full Time、Contract。
type EmploymentType struct {
	ID             int    `json:"id"`
	EmploymentType string `json:"employmentType"`
}

// PositionLevel 是職級，例如 Manager、Senior Executive。
type PositionLevel struct {
	ID       int    `json:"id"`
	Position string `json:"position"`
}

// Skill 是系統萃取出的技能標籤，IsKeySkill 標記雇主指定的關鍵技能。
type Skill struct {
	UUID       string   `json:"uuid"`
	Skill      string   `json:"skill"`
	Confidence *float64 `json:"confidence,omitempty"`
	IsKeySkill bool     `json:"isKeySkill,omitempty"`
}

// Scheme 是政府補助計畫，例如 Career Conversion Programme。
type Scheme struct {
	ID        int    `json:"id,omitempty"`
	Scheme    string `json:"scheme,omitempty"`
	SubScheme string `json:"subScheme,omitempty"`
}

// FlexibleWork 是彈性工作安排，例如 Flexi-place、Flexi-time。
type FlexibleWork struct {
	ID                      int    `json:"id"`
	FlexibleWorkArrangement string `json:"flexibleWorkArrangement"`
}

// ShiftPattern 是輪班型態，多數職缺為空。
type ShiftPattern struct {
	ID           int    `json:"id,omitempty"`
	ShiftPattern string `json:"shiftPattern,omitempty"`
}

// Address 是工作地點。海外職缺走 foreignAddress 欄位，本地職缺走 block/street。
type Address struct {
	IsOverseas      bool       `json:"isOverseas"`
	OverseasCountry string     `json:"overseasCountry,omitempty"`
	ForeignAddress1 string     `json:"foreignAddress1,omitempty"`
	ForeignAddress2 string     `json:"foreignAddress2,omitempty"`
	Block           string     `json:"block,omitempty"`
	Street          string     `json:"street,omitempty"`
	Floor           string     `json:"floor,omitempty"`
	Unit            string     `json:"unit,omitempty"`
	Building        string     `json:"building,omitempty"`
	PostalCode      string     `json:"postalCode,omitempty"`
	Lat             float64    `json:"lat,omitempty"`
	Lng             float64    `json:"lng,omitempty"`
	Districts       []District `json:"districts,omitempty"`
}

// District 是行政區與其所屬區域。
type District struct {
	ID       int      `json:"id"`
	Location string   `json:"location"`
	Region   string   `json:"region"`
	RegionID string   `json:"regionId"`
	Sectors  []string `json:"sectors,omitempty"`
}

// Metadata 是張貼與統計資訊，JobDetailsUrl 是給人看的原始職缺連結。
type Metadata struct {
	JobPostID                 string `json:"jobPostId"`
	JobDetailsURL             string `json:"jobDetailsUrl"`
	CreatedAt                 string `json:"createdAt,omitempty"`
	UpdatedAt                 string `json:"updatedAt,omitempty"`
	NewPostingDate            string `json:"newPostingDate,omitempty"`
	OriginalPostingDate       string `json:"originalPostingDate,omitempty"`
	ExpiryDate                string `json:"expiryDate,omitempty"`
	EditCount                 int    `json:"editCount,omitempty"`
	RepostCount               int    `json:"repostCount,omitempty"`
	TotalNumberOfView         int    `json:"totalNumberOfView,omitempty"`
	TotalNumberJobApplication int    `json:"totalNumberJobApplication,omitempty"`
	IsHideSalary              bool   `json:"isHideSalary,omitempty"`
	IsHideEmployerName        bool   `json:"isHideEmployerName,omitempty"`
	IsHideHiringEmployerName  bool   `json:"isHideHiringEmployerName,omitempty"`
	IsPostedOnBehalf          bool   `json:"isPostedOnBehalf,omitempty"`
	IsJdEnhanced              bool   `json:"isJdEnhanced,omitempty"`
}

// Company 是雇主。UEN 是新加坡的統一企業號，也是 /v2/companies/{uen} 的主鍵。
type Company struct {
	UEN                 string              `json:"uen"`
	Name                string              `json:"name"`
	Description         string              `json:"description,omitempty"`
	CompanyURL          string              `json:"companyUrl,omitempty"`
	SSICCode            string              `json:"ssicCode,omitempty"`
	SSICCode2020        string              `json:"ssicCode2020,omitempty"`
	SSICDescription2020 string              `json:"ssicDescription2020,omitempty"`
	EmployeeCount       int                 `json:"employeeCount,omitempty"`
	LastSyncDate        string              `json:"lastSyncDate,omitempty"`
	LogoFileName        string              `json:"logoFileName,omitempty"`
	LogoUploadPath      string              `json:"logoUploadPath,omitempty"`
	ResponsiveEmployer  *ResponsiveEmployer `json:"responsiveEmployer,omitempty"`
	Links               Links               `json:"_links,omitempty"`
}

// ResponsiveEmployer 標記雇主是否被平台認定為「會回覆求職者」。
type ResponsiveEmployer struct {
	IsResponsive bool `json:"isResponsive"`
}

// ScreeningQuestion 是申請時的篩選題。多數職缺為空陣列。
type ScreeningQuestion struct {
	ID       FlexInt `json:"id,omitempty"`
	Question string  `json:"question,omitempty"`
	Type     string  `json:"type,omitempty"`
}
