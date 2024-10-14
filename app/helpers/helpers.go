package helpers

import (
	"fmt"
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

/*
	func IncrementRateLimit(ip string) (int, error) {
		redisClient := database.GetRedisClient()
		limitKey := "rate_limit:" + ip
		ctx := context.Background()
		current, err := redisClient.Get(ctx, limitKey).Int()
		if err != nil && err != redis.Nil {
			return 0, err
		}

		if current >= 5 {
			return current, nil
		}

		redisClient.Incr(ctx, limitKey)
		redisClient.Expire(ctx, limitKey, 60*time.Second)
		return current, nil
	}
*/
func AllShortUrls() ([]entities.Url, error) {
	items := []entities.Url{}

	if err := database.Connection().Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}
func IncClickCount(url entities.Url) {
	count := url.UsageCount + 1
	fmt.Println("usage_count : ", count)

	database.Connection().Model(&url).Where("shortened_url = ?", url.ShortenedUrl).Updates(entities.Url{
		UsageCount: count,
	})
}
func SaveAccessDetails(url entities.Url, ip string) {
	fmt.Println("accessed ip : ", ip)
	log := &entities.Log{}

	result := database.Connection().Model(&log).Where("shortened_url = ? AND accessed_ip = ?", url.ShortenedUrl, ip).First(&log)

	// Check if no record was found
	if result.RowsAffected == 0 {
		fmt.Println("Yeni Kayıt")
		database.Connection().Model(&entities.Log{}).Create(&entities.Log{
			ShortenedUrl: url.ShortenedUrl,
			AccessedIp:   ip,
			AccessedAt:   time.Now(),
			AccessCount:  1,
		})
	} else if result.Error == nil {
		UpdateAccessDetails(*log, ip)
	} else {
		fmt.Println("Error: ", result.Error)
	}
}

func UpdateAccessDetails(log entities.Log, ip string) {
	accessCount := log.AccessCount + 1
	database.Connection().Model(&log).Where("shortened_url = ? AND accessed_ip = ?", log.ShortenedUrl, ip).Updates(entities.Log{
		AccessCount: accessCount,
	})
}
