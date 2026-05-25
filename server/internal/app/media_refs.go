package app

import "gorm.io/gorm"

func mediaURLInUse(db *gorm.DB, url string) bool {
	var count int64
	db.Model(&Person{}).Where("avatar_url = ?", url).Count(&count)
	if count > 0 {
		return true
	}
	db.Model(&Publication{}).Where("image_url = ?", url).Count(&count)
	return count > 0
}
