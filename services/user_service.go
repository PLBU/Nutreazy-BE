package services

import (
    "be-nutreazy/models"
    "gorm.io/gorm"
)

type UserService struct {
    db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
    return &UserService{
        db: db,
    }
}

func (s *UserService) GetUsers() ([]models.User, error) {
    var users []models.User
    if err := s.db.Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}

func (s *UserService) GetUserByID(userID int) (models.User, error) {
    var user models.User
    if err := s.db.First(&user, "id = ?", userID).Error; err != nil {
        return models.User{}, err
    }
    return user, nil
}


func (s *UserService) CreateUser(user models.User) (models.User, error) {
    if err := s.db.Create(&user).Error; err != nil {
        return models.User{}, err
    }
    return user, nil
}

func (s *UserService) UpdateUser(user models.User) (models.User, error) {
    if err := s.db.First(&user, "id = ?", user.ID).Error; err != nil {
        return models.User{}, err
    }
    if err := s.db.Save(&user).Error; err != nil {
        return models.User{}, err
    }
    return user, nil
}

func (s *UserService) DeleteUser(userID int) error {
    return s.db.Delete(&models.User{}, userID).Error
}