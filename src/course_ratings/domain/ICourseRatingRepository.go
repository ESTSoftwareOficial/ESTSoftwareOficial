package domain

import "estsoftwareoficial/src/course_ratings/domain/entities"

type CourseRatingRepository interface {
	Save(courseRating *entities.CourseRating) (*entities.CourseRating, error)
	GetByID(id int) (*entities.CourseRating, error)
	GetAll() ([]*entities.CourseRating, error)
	GetByCourse(courseID int) ([]*entities.CourseRating, error)
	GetByUserAndCourse(userID int, courseID int) (*entities.CourseRating, error)
	Update(courseRating *entities.CourseRating) error
	Delete(id int) error
	CalculateAverageRating(courseID int) (float64, int, error)
}
