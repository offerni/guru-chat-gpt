package guruchatgpt

import (
	"strings"
	"time"
)

type GuruRepository interface {
	ListCards() (*ListGuruCards, error)
}

type ListGuruCards struct {
	Cards []Card
}

type Card struct {
	ID                    string
	Content               string
	Owner                 Owner
	LastModified          CustomDate
	TeamID                string
	Collection            Collection
	LastVerified          CustomDate
	LastVerifiedBy        Owner
	LastModifiedBy        Owner
	PreferredPhrase       string
	ShareStatus           string
	HTMLContent           bool
	VerificationInterval  int
	VerificationType      string
	Slug                  string
	DateCreated           CustomDate
	CardType              string
	ContentSchemaVersion  string
	Followed              bool
	VerificationState     string
	OriginalOwner         Owner
	NextVerificationDate  CustomDate
	GuruSlateToolsVersion string
	VerificationReasons   []string
}

type Owner struct {
	ID            string
	Status        string
	Email         string
	LastName      string
	FirstName     string
	ProfilePicURL string
}

type Collection struct {
	Name                 string
	ID                   string
	Color                string
	HomeBoardSlug        string
	Description          string
	ROIEnabled           bool
	AssistEnabled        bool
	PublicCardsEnabled   bool
	CollectionType       string
	CollectionTypeDetail string
	Slug                 string
	DateCreated          CustomDate
}

type CustomDate struct {
	time.Time
}

func (t *CustomDate) UnmarshalJSON(b []byte) error {
	var err error
	timeString := strings.Trim(string(b), "\"")
	if timeString == "null" {
		t.Time = time.Time{}
		return nil
	}
	t.Time, err = time.Parse("2006-01-02T15:04:05.999Z0700", timeString)
	return err
}

func (t CustomDate) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + t.Time.Format("2006-01-02T15:04:05.999Z0700") + `"`), nil
}
