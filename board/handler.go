package board

import (
	"log"
	"net/http"
	"strconv"

	"github.com/harkce/beasiswakita"
	"github.com/harkce/beasiswakita/errors"
	"github.com/harkce/beasiswakita/utils"

	"github.com/harkce/beasiswakita/authentication"
	"github.com/harkce/beasiswakita/server/response"
	"github.com/julienschmidt/httprouter"
)

type BoardHandler struct{}

func (h *BoardHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	owner, err := authentication.Token(r.Header.Get("Authorization"), []string{})
	if err != nil {
		response.Error(w, errors.Unauthorized)
		log.Println(err)
		return
	}

	var board beasiswakita.StudentBoard
	err = utils.Decode(r, &board)
	if err != nil {
		response.Error(w, errors.InternalServerError)
		log.Println(err)
		return
	}

	board.UserID = owner.ID

	beasiswakita.Transaction, err = beasiswakita.DbMap.Begin()
	if err != nil {
		response.Error(w, errors.InternalServerError)
		log.Println(err)
		return
	}

	newBoard, err := CreateBoard(board)
	if err != nil {
		beasiswakita.Transaction.Rollback()
		response.Error(w, errors.InternalServerError)
		log.Println(err)
		return
	}

	err = beasiswakita.Transaction.Commit()
	if err != nil {
		beasiswakita.Transaction.Rollback()
		response.Error(w, errors.InternalServerError)
		log.Println(err)
		return
	}

	response.Created(w, newBoard)
	return
}

func (h *BoardHandler) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := authentication.Token(r.Header.Get("Authorization"), []string{})
	if err != nil {
		response.Error(w, errors.Unauthorized)
		log.Println(err)
		return
	}

	var board beasiswakita.StudentBoard
	err = utils.Decode(r, &board)
	if err != nil {
		response.Error(w, errors.InternalServerError)
		log.Println(err)
		return
	}

	beasiswakita.Transaction, err = beasiswakita.DbMap.Begin()
	if err != nil {
		beasiswakita.Transaction.Rollback()
		response.Error(w, errors.InternalServerError)
		log.Println(err)
		return
	}

	updatedBoard, err := UpdateBoard(board)
	if err != nil {
		beasiswakita.Transaction.Rollback()
		response.Error(w, errors.InternalServerError)
		log.Println(err)
		return
	}

	err = beasiswakita.Transaction.Commit()
	if err != nil {
		beasiswakita.Transaction.Rollback()
		response.Error(w, errors.InternalServerError)
		log.Println(err)
		return
	}

	response.OK(w, updatedBoard)
	return
}

func (h *BoardHandler) State(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_, err := authentication.Token(r.Header.Get("Authorization"), []string{})
	if err != nil {
		response.Error(w, errors.Unauthorized)
		log.Println(err)
		return
	}

	ID := ps.ByName("boardID")
	boardID, err := strconv.Atoi(ID)
	if err != nil {
		response.Error(w, errors.InternalServerError)
		log.Println(err)
		return
	}

	var state struct {
		State int `json:"state"`
	}

	err = utils.Decode(r, &state)
	if err != nil {
		response.Error(w, errors.InternalServerError)
		log.Println(err)
		return
	}

	beasiswakita.Transaction, err = beasiswakita.DbMap.Begin()
	if err != nil {
		beasiswakita.Transaction.Rollback()
		response.Error(w, errors.InternalServerError)
		log.Println(err)
		return
	}

	err = ChangeBoardState(boardID, state.State)
	if err != nil {
		beasiswakita.Transaction.Rollback()
		response.Error(w, errors.InternalServerError)
		log.Println(err)
		return
	}

	err = beasiswakita.Transaction.Commit()
	if err != nil {
		beasiswakita.Transaction.Rollback()
		response.Error(w, errors.InternalServerError)
		log.Println(err)
		return
	}

	response.OK(w, state)
	return
}

func (h *BoardHandler) Delete(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := authentication.Token(r.Header.Get("Authorization"), []string{})
	if err != nil {
		response.Error(w, errors.Unauthorized)
		log.Println(err)
		return
	}

	var board beasiswakita.StudentBoard
	err = utils.Decode(r, &board)
	if err != nil {
		response.Error(w, errors.InternalServerError)
		log.Println(err)
		return
	}

	beasiswakita.Transaction, err = beasiswakita.DbMap.Begin()
	if err != nil {
		beasiswakita.Transaction.Rollback()
		response.Error(w, errors.InternalServerError)
		log.Println(err)
		return
	}

	deletedBoard, err := DeleteBoard(board)
	if err != nil {
		beasiswakita.Transaction.Rollback()
		response.Error(w, errors.InternalServerError)
		log.Println(err)
		return
	}

	err = beasiswakita.Transaction.Commit()
	if err != nil {
		beasiswakita.Transaction.Rollback()
		response.Error(w, errors.InternalServerError)
		log.Println(err)
		return
	}

	response.OK(w, deletedBoard)
	return
}
