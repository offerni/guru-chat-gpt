package models

import "time"

type Cards struct {
	data []Card
}

type Card struct {
	ID                    string     `json:"id"`
	Content               string     `json:"content"`
	Owner                 Owner      `json:"owner"`
	LastModified          time.Time  `json:"lastModified"`
	TeamID                string     `json:"teamId"`
	Collection            Collection `json:"collection"`
	LastVerified          time.Time  `json:"lastVerified"`
	LastVerifiedBy        Owner      `json:"lastVerifiedBy"`
	LastModifiedBy        Owner      `json:"lastModifiedBy"`
	PreferredPhrase       string     `json:"preferredPhrase"`
	ShareStatus           string     `json:"shareStatus"`
	HTMLContent           bool       `json:"htmlContent"`
	VerificationInterval  int        `json:"verificationInterval"`
	VerificationType      string     `json:"verificationType"`
	Slug                  string     `json:"slug"`
	DateCreated           time.Time  `json:"dateCreated"`
	CardType              string     `json:"cardType"`
	ContentSchemaVersion  string     `json:"contentSchemaVersion"`
	Followed              bool       `json:"followed"`
	VerificationState     string     `json:"verificationState"`
	OriginalOwner         Owner      `json:"originalOwner"`
	NextVerificationDate  time.Time  `json:"nextVerificationDate"`
	GuruSlateToolsVersion string     `json:"guruSlateToolsVersion"`
	VerificationReasons   []string   `json:"verificationReasons"`
}

type Owner struct {
	ID            string `json:"id"`
	Status        string `json:"status"`
	Email         string `json:"email"`
	LastName      string `json:"lastName"`
	FirstName     string `json:"firstName"`
	ProfilePicURL string `json:"profilePicUrl"`
}

type Collection struct {
	Name                 string    `json:"name"`
	ID                   string    `json:"id"`
	Color                string    `json:"color"`
	HomeBoardSlug        string    `json:"homeBoardSlug"`
	Description          string    `json:"description"`
	ROIEnabled           bool      `json:"roiEnabled"`
	AssistEnabled        bool      `json:"assistEnabled"`
	PublicCardsEnabled   bool      `json:"publicCardsEnabled"`
	CollectionType       string    `json:"collectionType"`
	CollectionTypeDetail string    `json:"collectionTypeDetail"`
	Slug                 string    `json:"slug"`
	DateCreated          time.Time `json:"dateCreated"`
}
