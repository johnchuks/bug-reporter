package models

import (
	"strings"
	"github.com/jinzhu/gorm"

)
// Report struct that holds the report data
type Report struct {

	gorm.Model
	Title string `gorm:"type:varchar(100);not null;unique" json:"title"`
	Description string `gorm:"not null" json:"description"`
	Screenshot string `gorm:"size:255;" json:"screenshot"`
	UserID uint `gorm:"not null"  json:"user_id"`
}

// Strip removes any whitespaces from the user inputs
func (r *Report) Strip() {
	r.Title = strings.TrimSpace(r.Title)
	r.Description = strings.TrimSpace(r.Description)
	r.Screenshot = strings.TrimSpace(r.Screenshot)
}


// Create adds a new bug report
func (r *Report) Create(db *gorm.DB) (*Report, error) {
	var err error

	err = db.Debug().Create(&r).Error

	if err != nil {
		return &Report{}, err
	}
	return r, nil
}

// Update updates an existing bug report
func (r *Report) Update(id int, db *gorm.DB) (*Report, error) {
	var err error
	report := &Report{
		Title: r.Title,
		Description: r.Description,
		Screenshot: r.Screenshot,
	}
	err = db.Debug().Table("reports").Where("id = ?", id).Updates(report).Error

	if err != nil {
		return &Report{}, err
	}

	return r, nil
}

// Get gets a bug report with the id specified
func (r *Report) Get(id int, db *gorm.DB) (*Report, error) {
	report := &Report{}
	err := db.Debug().Table("reports").Where("id = ?", id).First(report).Error

	if err != nil {
		return nil, err
	}
	return report, nil
}

// Delete deletes an existing bug report for a user
func (r *Report) Delete(id int, db *gorm.DB) (string, error) {
	report := &Report{}
	err := db.Debug().Table("reports").Where("id = ?", id).Delete(report).Error

	if err != nil {
		return "", err
	}
	return "report deleted successfully", nil

}