package service

import (
	"log"

	"github.com/danmory/web-hackers-service/internal/core"
	"github.com/danmory/web-hackers-service/internal/storage/redis"
)

func RetrieveHackers() []core.Hacker {
	rowHackers, err := redis.RetrieveHackers()
	if err != nil {
		log.Printf("Error while retrieving hackers: %v", err)
		return nil
	}
	hackers := make([]core.Hacker, len(rowHackers))
	for i := range rowHackers {
		name, ok := rowHackers[i].Member.(string)
		if !ok {
			log.Printf("Incorrect database entry: %v", name)
			return nil
		}
		hackers[i] = core.Hacker{Name: rowHackers[i].Member.(string), Score: rowHackers[i].Score}
	}
	return hackers
}
