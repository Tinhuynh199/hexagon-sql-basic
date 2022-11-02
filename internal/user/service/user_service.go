package service

import (
	"context"
	"database/sql"
	"log"

	. "hexrestapi/internal/user/domain"
	. "hexrestapi/internal/user/port"
)

type UserService interface {
	Load(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, user *User) (int64, error)
	Update(ctx context.Context, user *User) (int64, error)
	Patch(ctx context.Context, user map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

type userService struct {
	db         *sql.DB
	repository UserRepository
}

func NewUserService(db *sql.DB, repos UserRepository) UserService {
	return &userService{db: db, repository: repos}
}

// Create implements UserService
func (s *userService) Create(ctx context.Context, user *User) (int64, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Fatal(err)
		return -1, nil
	}

	ctx = context.WithValue(ctx, "tx", tx)
	res, err := s.repository.Create(ctx, user)
	if err != nil {
		er := tx.Rollback()
		if er != nil {
			log.Fatal(er)
			return -1, nil
		}
		log.Fatal(err)
		return -1, err
	}
	err = tx.Commit()
	return res, err
}

// Delete implements UserService
func (s *userService) Delete(ctx context.Context, id string) (int64, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Fatal(err)
		return -1, nil
	}

	ctx = context.WithValue(ctx, "tx", tx)
	res, err := s.repository.Delete(ctx, id)
	if err != nil {
		er := tx.Rollback()
		if er != nil {
			log.Fatal(er)
			return -1, nil
		}
		log.Fatal(err)
		return -1, err
	}
	err = tx.Commit()
	return res, err
}

// Load implements UserService
func (s *userService) Load(ctx context.Context, id string) (*User, error) {
	return s.repository.Load(ctx, id)
}

// Patch implements UserService
func (s *userService) Patch(ctx context.Context, user map[string]interface{}) (int64, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Fatal(err)
		return -1, nil
	}

	ctx = context.WithValue(ctx, "tx", tx)
	res, err := s.repository.Patch(ctx, user)
	if err != nil {
		er := tx.Rollback()
		if er != nil {
			log.Fatal(er)
			return -1, nil
		}
		log.Fatal(err)
		return -1, err
	}
	err = tx.Commit()
	return res, err

}

// Update implements UserService
func (s *userService) Update(ctx context.Context, user *User) (int64, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Fatal(err)
		return -1, nil
	}

	ctx = context.WithValue(ctx, "tx", tx)
	res, err := s.repository.Update(ctx, user)
	if err != nil {
		er := tx.Rollback()
		if er != nil {
			log.Fatal(er)
			return -1, nil
		}
		log.Fatal(err)
		return -1, err
	}
	err = tx.Commit()
	return res, err
}
