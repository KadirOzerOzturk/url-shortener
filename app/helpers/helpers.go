package helpers

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/KadirOzerOzturk/url-shortener/app/entities"
	"github.com/KadirOzerOzturk/url-shortener/internal/database"
)

func GenerateShortUrl() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	randStr := make([]rune, rand.Intn(10-6)+6)

	for i := range randStr {
		randStr[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	allUrls, err := AllShortUrls()
	if err != nil {
		return ""
	}

	isUnique := false
	for !isUnique {
		isUnique = true
		for _, url := range allUrls {
			if url.ShortenedUrl == string(randStr) {
				for i := range randStr {
					randStr[i] = letterRunes[rand.Intn(len(letterRunes))]
				}
				isUnique = false
				break
			}
		}
	}

	return string(randStr)
}

func AllShortUrls() ([]entities.Url, error) {
	items := []entities.Url{}

	if err := database.Connection().Table("urls").Select("*").Find(&items).Error; err != nil {
		log.Println("Error fetching URLs:", err) // Log the error for better insight
		return nil, err
	}

	return items, nil
}
func IncClickCount(url entities.Url) {
	count := url.UsageCount + 1
	fmt.Println("usage_count : ", count)

	if err := database.Connection().Model(&url).Where("shortened_url = ?", url.ShortenedUrl).Updates(entities.Url{UsageCount: count}).Error; err != nil {
		log.Fatalf("Update failed: %v", err)
	}

}
func SaveAccessDetails(url entities.Url, ip string) {
	fmt.Println("accessed ip : ", ip)
	ipLog := &entities.Log{}

	result := database.Connection().Model(&ipLog).Where("shortened_url = ? AND accessed_ip = ?", url.ShortenedUrl, ip).First(&ipLog)
	if result.RowsAffected == 0 {
		fmt.Println("Yeni Kayıt")
		if err := database.Connection().Model(&entities.Log{}).Create(&entities.Log{
			ShortenedUrl: url.ShortenedUrl,
			AccessedIp:   ip,
			AccessedAt:   time.Now(),
			AccessCount:  1,
		}).Error; err != nil {
			log.Fatalf("Process failed: %v", err)
		}

	} else if result.Error == nil {
		UpdateAccessDetails(*ipLog, ip)
	} else {
		fmt.Println("Error: ", result.Error)
	}
}

func UpdateAccessDetails(ipLog entities.Log, ip string) {
	accessCount := ipLog.AccessCount + 1
	if err := database.Connection().Model(&ipLog).Where("shortened_url = ? AND accessed_ip = ?", ipLog.ShortenedUrl, ip).Updates(entities.Log{
		AccessCount: accessCount,
	}).Error; err != nil {
		log.Fatalf("Update Failed : %v", err)
	}
}

func SendMail(mail entities.Mail) error {
	if err := database.Connection().Model(&entities.Mail{}).Create(&mail).Error; err != nil {
		return err
	}
	return nil
}

func Login(user entities.User) (entities.User, error) {
	result := database.Connection().Model(&entities.User{}).Where("email = ?", user.Email).First(&user)
	if result.Error != nil {
		return entities.User{}, result.Error
	}
	return user, nil
}
func Register(user entities.User) (entities.User, error) {
	if err := database.Connection().Model(&entities.User{}).Create(&user).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}
