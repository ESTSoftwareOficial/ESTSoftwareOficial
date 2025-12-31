package domain

import (
	"estsoftwareoficial/src/course_ratings/domain/entities"
	"estsoftwareoficial/src/course_ratings/domain/dto"
)

type CourseRatingRepository interface {
	Save(courseRating *entities.CourseRating) (*entities.CourseRating, error)
	GetByID(id int) (*dto.CourseRatingResponse, error)                    
	GetAll() ([]*dto.CourseRatingResponse, error)                         
	GetByCourse(courseID int) ([]*dto.CourseRatingResponse, error)        
	GetByUserAndCourse(userID int, courseID int) (*entities.CourseRating, error)
	Update(courseRating *entities.CourseRating) error
	Delete(id int) error
	CalculateAverageRating(courseID int) (float64, int, error)
}