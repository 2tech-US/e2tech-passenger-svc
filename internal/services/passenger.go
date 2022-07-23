package services

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/lntvan166/e2tech-passenger-svc/internal/db"
	"github.com/lntvan166/e2tech-passenger-svc/internal/pb"
	"github.com/lntvan166/e2tech-passenger-svc/internal/utils"
)

func (s *Server) CreatePassenger(context context.Context, req *pb.CreatePassengerRequest) (*pb.CreatePassengerResponse, error) {
	_, err := s.DB.GetPassengerByPhone(context, req.Phone)
	if err != sql.ErrNoRows {
		return &pb.CreatePassengerResponse{
			Status: http.StatusBadRequest,
			Error:  "user already exists",
		}, nil
	}

	hashedPassword := utils.HashPassword(req.Password)
	parsedDate, err := utils.ParseStringToDate(req.DateOfBirth)
	if err != nil {
		return &pb.CreatePassengerResponse{
			Status: http.StatusBadRequest,
			Error:  "invalid date of birth",
		}, nil
	}

	arg := db.CreatePassengerParams{
		Phone:          req.Phone,
		HashedPassword: hashedPassword,
		Name:           req.Name,
		DateOfBirth:    parsedDate,
	}

	passenger, err := s.DB.CreatePassenger(context, arg)
	if err != nil {
		return &pb.CreatePassengerResponse{
			Status: http.StatusInternalServerError,
			Error:  fmt.Sprintf("create passenger error: %s", err),
		}, nil
	}

	dataRsp := &pb.Passenger{
		Id:          passenger.ID,
		Phone:       passenger.Phone,
		Name:        passenger.Name,
		Verified:    passenger.Verified,
		DateOfBirth: utils.ParsedDateToString(passenger.DateOfBirth),
	}

	return &pb.CreatePassengerResponse{
		Status:    http.StatusCreated,
		Passenger: dataRsp,
	}, nil
}

func (s *Server) GetPassengerByPhone(context context.Context, req *pb.GetPassengerByPhoneRequest) (*pb.GetPassengerByPhoneResponse, error) {
	passenger, err := s.DB.GetPassengerByPhone(context, req.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.GetPassengerByPhoneResponse{
				Status: http.StatusBadRequest,
				Error:  "user not found",
			}, nil
		}

		return &pb.GetPassengerByPhoneResponse{
			Status: http.StatusInternalServerError,
			Error:  "failed to get user",
		}, nil
	}

	dataRsp := &pb.Passenger{
		Id:          passenger.ID,
		Phone:       passenger.Phone,
		Name:        passenger.Name,
		Verified:    passenger.Verified,
		DateOfBirth: utils.ParsedDateToString(passenger.DateOfBirth),
	}

	return &pb.GetPassengerByPhoneResponse{
		Status:    http.StatusOK,
		Passenger: dataRsp,
	}, nil
}

func (s *Server) ListPassengers(context context.Context, req *pb.ListPassengersRequest) (*pb.ListPassengersResponse, error) {
	arg := db.ListPassengersParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	}

	passengers, err := s.DB.ListPassengers(context, arg)
	if err != nil {
		return &pb.ListPassengersResponse{
			Status: http.StatusInternalServerError,
			Error:  "failed to list passengers",
		}, nil
	}

	dataRsp := make([]*pb.Passenger, len(passengers))
	for i, passenger := range passengers {
		dataRsp[i] = &pb.Passenger{
			Id:          passenger.ID,
			Phone:       passenger.Phone,
			Name:        passenger.Name,
			Verified:    passenger.Verified,
			DateOfBirth: utils.ParsedDateToString(passenger.DateOfBirth),
		}
	}

	return &pb.ListPassengersResponse{
		Status:    http.StatusOK,
		Passenger: dataRsp,
	}, nil
}

func (s *Server) UpdatePassenger(context context.Context, req *pb.UpdatePassengerRequest) (*pb.UpdatePassengerResponse, error) {
	passenger, err := s.DB.GetPassengerByPhone(context, req.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.UpdatePassengerResponse{
				Status: http.StatusBadRequest,
				Error:  "user not found",
			}, nil
		}

		return &pb.UpdatePassengerResponse{
			Status: http.StatusInternalServerError,
			Error:  "failed to get user",
		}, nil
	}

	strDate, err := utils.ParseStringToDate(req.DateOfBirth)
	if err != nil {
		return &pb.UpdatePassengerResponse{
			Status: http.StatusBadRequest,
			Error:  "invalid date of birth",
		}, nil
	}

	arg := db.UpdatePassengerParams{
		ID:          req.Id,
		Phone:       req.Phone,
		Name:        req.Name,
		DateOfBirth: strDate,
	}

	passenger, err = s.DB.UpdatePassenger(context, arg)
	if err != nil {
		return &pb.UpdatePassengerResponse{
			Status: http.StatusInternalServerError,
			Error:  "failed to update passenger",
		}, nil
	}

	dataRsp := &pb.Passenger{
		Id:          passenger.ID,
		Phone:       passenger.Phone,
		Name:        passenger.Name,
		Verified:    passenger.Verified,
		DateOfBirth: utils.ParsedDateToString(passenger.DateOfBirth),
	}

	return &pb.UpdatePassengerResponse{
		Status:    http.StatusOK,
		Passenger: dataRsp,
	}, nil
}

func (s *Server) UpdatePassword(context context.Context, req *pb.UpdatePasswordRequest) (*pb.UpdatePasswordResponse, error) {
	passenger, err := s.DB.GetPassengerByPhone(context, req.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.UpdatePasswordResponse{
				Status: http.StatusBadRequest,
				Error:  "user not found",
			}, nil
		}

		return &pb.UpdatePasswordResponse{
			Status: http.StatusInternalServerError,
			Error:  "failed to get user",
		}, nil
	}

	hashedPassword := utils.HashPassword(req.Password)
	arg := db.UpdatePasswordParams{
		ID:             passenger.ID,
		HashedPassword: hashedPassword,
	}

	passenger, err = s.DB.UpdatePassword(context, arg)
	if err != nil {
		return &pb.UpdatePasswordResponse{
			Status: http.StatusInternalServerError,
			Error:  "failed to update password",
		}, nil
	}

	dataRsp := &pb.Passenger{
		Id:          passenger.ID,
		Phone:       passenger.Phone,
		Name:        passenger.Name,
		Verified:    passenger.Verified,
		DateOfBirth: utils.ParsedDateToString(passenger.DateOfBirth),
	}

	return &pb.UpdatePasswordResponse{
		Status:    http.StatusOK,
		Passenger: dataRsp,
	}, nil
}

func (s *Server) VerifyPassenger(context context.Context, req *pb.VerifyPassengerRequest) (*pb.VerifyPassengerResponse, error) {
	passenger, err := s.DB.GetPassengerByPhone(context, req.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.VerifyPassengerResponse{
				Status: http.StatusBadRequest,
				Error:  "user not found",
			}, nil
		}

		return &pb.VerifyPassengerResponse{
			Status: http.StatusInternalServerError,
			Error:  "failed to get user",
		}, nil
	}

	passenger, err = s.DB.Verify(context, req.Phone)
	if err != nil {
		return &pb.VerifyPassengerResponse{
			Status: http.StatusInternalServerError,
			Error:  "failed to verify passenger",
		}, nil
	}

	dataRsp := &pb.Passenger{
		Id:          passenger.ID,
		Phone:       passenger.Phone,
		Name:        passenger.Name,
		Verified:    passenger.Verified,
		DateOfBirth: utils.ParsedDateToString(passenger.DateOfBirth),
	}

	return &pb.VerifyPassengerResponse{
		Status:    http.StatusOK,
		Passenger: dataRsp,
	}, nil
}

func (s *Server) DeletePassenger(context context.Context, req *pb.DeletePassengerRequest) (*pb.DeletePassengerResponse, error) {
	passenger, err := s.DB.GetPassengerByPhone(context, req.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.DeletePassengerResponse{
				Status: http.StatusBadRequest,
				Error:  "user not found",
			}, nil
		}

		return &pb.DeletePassengerResponse{
			Status: http.StatusInternalServerError,
			Error:  "failed to get user",
		}, nil
	}

	err = s.DB.DeletePassenger(context, passenger.Phone)
	if err != nil {
		return &pb.DeletePassengerResponse{
			Status: http.StatusInternalServerError,
			Error:  "failed to delete passenger",
		}, nil
	}

	return &pb.DeletePassengerResponse{
		Status: http.StatusOK,
	}, nil
}