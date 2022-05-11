package service

import (
	"driverlocation/config"
	"driverlocation/dto"
	"driverlocation/exception"
	"driverlocation/model"
	"driverlocation/repository"
	"driverlocation/utils"
	"encoding/csv"
	"github.com/umahmood/haversine"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"math"
	"mime/multipart"
	"os"
	"strconv"
	"sync"
)

var lock = &sync.Mutex{}

var singleInstance ILocationService

type ILocationService interface {
	Create(*model.Point) error
	CreateBulkUpdate(*multipart.FileHeader) error
	GetNearestDriver(*model.Location, int) (*dto.NearestDriverResponse, error)
}

type LocationService struct {
	repo repository.ILocationRepository
}

func NewLocationService(repo repository.ILocationRepository) ILocationService {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &LocationService{
				repo: repo,
			}
		}
	}
	return singleInstance
}

func (l LocationService) Create(point *model.Point) error {
	point.Title = utils.RandSeq(10)
	point.ID = primitive.NewObjectID()
	return l.repo.Insert(point)
}

func (l LocationService) CreateBulkUpdate(file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(file.Filename)
	if err != nil {
		return exception.GetGenericException(err.Error())
	}

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return exception.GetGenericException(err.Error())
	}

	defer os.Remove("./" + dst.Name())
	defer dst.Close()

	csvReader, csvFile := utils.CreateReader(dst.Name())
	defer csvFile.Close()
	csvReader.Read()

	return readCoordinates(csvReader, l.repo)
}

func readCoordinates(csvReader *csv.Reader, repo repository.ILocationRepository) error {
	rawCoordinates, err := csvReader.ReadAll()
	if err != nil {
		return exception.GetGenericException(err.Error())
	}

	bulkResultChan := make(chan error)
	defer close(bulkResultChan)

	goRoutineCount := 0
	var wg sync.WaitGroup
	for i := 0; i < len(rawCoordinates); i = i + config.Get().RecordBatchCount {
		goRoutineCount++
		if i+config.Get().RecordBatchCount > len(rawCoordinates) {
			go convertAndInsert(rawCoordinates[i:len(rawCoordinates)], &wg, repo, bulkResultChan)
		} else {
			go convertAndInsert(rawCoordinates[i:i+config.Get().RecordBatchCount], &wg, repo, bulkResultChan)
		}
	}
	var errorOfGoRoutines error = nil
	wg.Wait()
	for i := 0; i < goRoutineCount; i++ {
		result := <-bulkResultChan
		if result != nil {
			errorOfGoRoutines = result
		}
	}

	return errorOfGoRoutines
}

func convertAndInsert(rawCoordinates [][]string, wg *sync.WaitGroup, repo repository.ILocationRepository, bulkResultChan chan error) {
	defer wg.Done()
	wg.Add(1)

	coordinates := make([]interface{}, 0)
	for _, coordinate := range rawCoordinates {
		latitude, err := strconv.ParseFloat(coordinate[0], 64)
		if err != nil {
			continue
		}
		longitude, err := strconv.ParseFloat(coordinate[1], 64)
		if err != nil {
			continue
		}
		point := model.Point{
			ID:    primitive.NewObjectID(),
			Title: utils.RandSeq(10),
			Location: model.Location{
				Type:        "Point",
				Coordinates: []float64{longitude, latitude},
			},
		}

		coordinates = append(coordinates, &point)
	}
	bulkResultChan <- repo.InsertBulk(coordinates)
}

func (l LocationService) GetNearestDriver(location *model.Location, radius int) (*dto.NearestDriverResponse, error) {
	points, err := l.repo.FindDriverInDistanceWithRadius(location, radius)
	if err != nil {
		return nil, err
	}

	if len(points) == 0 {
		return nil, exception.GetThereIsNoDriverException(radius)
	}

	userLocation := haversine.Coord{Lat: location.Coordinates[1], Lon: location.Coordinates[0]}
	response := &dto.NearestDriverResponse{
		Distance: math.MaxFloat64,
		Title:    "",
	}
	for _, val := range points {
		point := haversine.Coord{Lat: val.Location.Coordinates[1], Lon: val.Location.Coordinates[0]}

		_, km := haversine.Distance(userLocation, point)
		if km < response.Distance {
			response.Distance = km
			response.Title = val.Title
		}
	}
	return response, nil
}
