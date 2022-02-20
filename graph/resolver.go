package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import "github.com/isadoramsouza/graphql-go/model"

type Resolver struct {
	Categories []*model.Category
	Courses    []*model.Course
	Chapter    []*model.Chapter
}
