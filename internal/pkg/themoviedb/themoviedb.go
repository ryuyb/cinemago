package themoviedb

import (
	"cinemago/internal/model/dto"
	"github.com/cyruzin/golang-tmdb"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jinzhu/copier"
	"strconv"
)

const (
	defaultLanguage = "zh-CN"
)

type Client struct {
	client *tmdb.Client
}

func NewClient() (*Client, error) {
	client, err := tmdb.Init("")
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func (c *Client) GetAllTrending(page int, mediaType string, timeWindow string) ([]dto.BasicMediaInfo, error) {
	trending, err := c.client.GetTrending(mediaType, timeWindow, map[string]string{
		"language": defaultLanguage,
		"page":     strconv.Itoa(page),
	})
	if err != nil {
		return nil, err
	}
	infos := make([]dto.BasicMediaInfo, 0, len(trending.Results))
	for _, result := range trending.Results {
		info := dto.BasicMediaInfo{}
		err := copier.Copy(&info, &result)
		if err != nil {
			log.Errorw("copier.Copy failed", "result", result, "err", err)
			continue
		}
		info.TmdbId = result.ID
		if result.MediaType == "tv" {
			info.Title = result.Name
			info.OriginalTitle = result.OriginalName
			info.ReleaseDate = result.FirstAirDate
		}
		infos = append(infos, info)
	}
	return infos, nil
}

func (c *Client) GetMovieDetail(movieId int) (*dto.MovieDetail, error) {
	details, err := c.client.GetMovieDetails(movieId, map[string]string{
		"language": defaultLanguage,
	})
	if err != nil {
		return nil, err
	}
	movieDetail := dto.MovieDetail{}
	err = copier.Copy(&movieDetail, &details)
	if err != nil {
		log.Errorw("copier.Copy failed", "details", details, "err", err)
	}
	movieDetail.TmdbId = details.ID
	movieDetail.ImdbId = details.IMDbID
	return &movieDetail, nil
}
