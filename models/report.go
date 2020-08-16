package models

import (
	"github.com/jinzhu/gorm"

)

type Report struct {

	gorm.Model
	Title string `gorm:"type:varchar(100);" json:"title"`
	Description string `gorm:"type:varchar(250);" json:"description"`
	Screenshot string `gorm:"size:255;" json:"screenshot"`
}


func (r *Report) createReport(db *gorm.DB) (*Report, error) {
	var err error

	err = db.Debug().Create(&r).Error

	if err != nil {
		return &Report{}, err
	}
	return r, nil
}

func (r *Report) updateReport(title string, description string, screenshot string, db *gorm.DB) (*Report, error) {
	var err error
	report := &Report{
		Title: title,
		Description: description,
		Screenshot: screenshot,
	}
	err = db.Debug().Table("reports").Where("id = ?", r.ID).Update(report).Error

	if err != nil {
		return nil, err
	}

	return report, nil
}

func (r *Report) getReport(db *gorm.DB) (*Report, error) {
	report := &Report{}
	err := db.Debug().Table("reports").Where("id = ?", r.ID).First(report).Error

	if err != nil {
		return nil, err
	}
	return report, nil
}

func (r *Report) deleteReport(db *gorm.DB) (string, error) {
	report := &Report{}
	err := db.Debug().Table("reports").Where("id = ?", r.ID).Delete(report).Error

	if err != nil {
		return "", err
	}
	return "report deleted successfully", nil

}