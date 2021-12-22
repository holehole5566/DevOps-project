package article

import (
	"database/sql"
	"strconv"
	"log"
	"github.com/holehole5566/goproject/model"
	"github.com/holehole5566/goproject/repo"
	C "github.com/holehole5566/goproject/pkg/constant"

)

func (srv *Service) GetTour(param string) (*model.Tour, error) {

	id, err := strconv.Atoi(param)
	if err != nil {
		log.Println("Get Tour: param id %s is not a number", param)
		return nil, C.ErrTourIDNotNumber
	}

	tour, err := repo.Tour.Get(id)
	if err == sql.ErrNoRows {
		log.Println("Get Tour: tour id %d record not found", id)
		return nil, C.ErrTourNotFound
	} else if err != nil {
		log.Println("Get Tour: unknown database error", err.Error())
		return nil, C.ErrDatabase
	}

	return tour, nil
}

func (srv *Service) GetTours() ([]*model.Tour, error) {
	// TODO: cache
	total, err := repo.Tour.Gets()
	if err != nil {
		log.Println("Get Tours: unknown database error, ", err.Error())
		return nil, C.ErrDatabase
	}
	return total, nil
}

func (srv *Service) GetTotalTours() (int, error) {
	// TODO: cache
	total, err := repo.Tour.GetTotal()
	if err != nil {
		log.Println("Get Tours: unknown database error, ", err.Error())
		return 0, C.ErrDatabase
	}
	return total, nil
}

func (srv *Service) AddTour(collectsID []int, title string) (int, error) {

	if len(collectsID) == 0 || title == "" {
		return 0, C.ErrTourAddFormatIncorrect
	}

	collectNum, err := repo.Collect.CheckManyExist(collectsID)
	if err != nil {
		log.Println("Check Collects Exist: ", err)
		return 0, C.ErrDatabase
	} else if len(collectsID) != int(collectNum) {
		return 0, C.ErrTourAddCollectsRecordNotFound
	}
	id, err := repo.Tour.Add(collectsID, title)
	if err != nil {
		log.Println("Add Tour: ", err)
		return 0, C.ErrDatabase
	}

	return id, nil
}

func (srv *Service) DelTour(param string) error {

	id, err := strconv.Atoi(param)
	if err != nil {
		log.Println("Del Tour: param id %s is not a number", param)
		return C.ErrTourIDNotNumber
	}

	if id < 0 {
		return C.ErrTourDelIDIncorrect
	}

	if tour, _ := repo.Tour.Get(id); tour == nil {
		return C.ErrTourDelDeleted
	}

	if err := repo.Tour.Del(id); err != nil {
		log.Println("Del Tour: ", err)
		return C.ErrDatabase
	}

	return nil
}