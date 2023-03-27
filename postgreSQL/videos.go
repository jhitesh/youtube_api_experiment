package postgreSQL

import (
	"fmt"
	"log"
	"strings"
)

type DBVideoData struct {
	ID          string
	Title       string
	Description string
	PublishedAt string
}

func (dbConn *DBConnection) SaveVideoData(videoData *DBVideoData) error {
	videoData.Title = strings.ReplaceAll(videoData.Title, "'", "\"")
	videoData.ID = strings.ReplaceAll(videoData.ID, "'", "\"")
	videoData.Description = strings.ReplaceAll(videoData.Description, "'", "\"")

	insertQuery := fmt.Sprintf("INSERT INTO videos (id, title, description, published_at) VALUES ('%s', '%s', '%s', '%s')",
		videoData.ID, videoData.Title, videoData.Description, videoData.PublishedAt)
	_, err := dbConn.DBConn.Query(insertQuery)
	if err != nil {
		log.Println("ERROR in inserting data in the DB for videoID ", videoData.ID, err.Error())
	}

	return err
}

func (dbConn *DBConnection) GetAllVideos(pageNumber, entriesPerPage int) ([]DBVideoData, error) {
	offset := entriesPerPage * (pageNumber - 1)
	limit := entriesPerPage

	selectQuery := fmt.Sprintf("SELECT * FROM videos ORDER BY published_at DESC OFFSET %d LIMIT %d", offset, limit)
	rows, err := dbConn.DBConn.Query(selectQuery)
	if err != nil {
		log.Println("ERROR in getting stored videos from database for page ", pageNumber, err.Error())
		return nil, err
	}

	ret := make([]DBVideoData, 0)
	for rows.Next() {
		nextVideo := DBVideoData{}
		_ = rows.Scan(&nextVideo.ID, &nextVideo.Title, &nextVideo.Description, &nextVideo.PublishedAt)

		ret = append(ret, nextVideo)
	}

	return ret, nil
}

func (dbConn *DBConnection) SearchVideos(query string, pageNumber, entriesPerPage int) []DBVideoData {
	skip := entriesPerPage * (pageNumber - 1)
	limit := entriesPerPage

	query = "%" + query + "%"

	searchQuery := fmt.Sprintf("SELECT * FROM videos WHERE title ILIKE '%s' OR description ILIKE '%s' ORDER BY published_at DESC OFFSET %d LIMIT %d", query, query, skip, limit)
	rows, err := dbConn.DBConn.Query(searchQuery)
	if err != nil {
		log.Println("ERROR in getting stored videos from database for page ", pageNumber, err.Error())
		return nil
	}

	ret := make([]DBVideoData, 0)
	for rows.Next() {
		nextVideo := DBVideoData{}
		_ = rows.Scan(&nextVideo.ID, &nextVideo.Title, &nextVideo.Description, &nextVideo.PublishedAt)
		ret = append(ret, nextVideo)
	}

	return ret
}
