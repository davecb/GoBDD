package services

import (
	"github.com/Golang-Coach/Lessons/GoBDD/mocks"
	"context"
	. "github.com/smartystreets/goconvey/convey"
	. "github.com/google/go-github/github"
	"testing"
	"errors"
)

func TestGithubAPI(t *testing.T) {
	Convey("Should return repository information", t, func() {
		backgroundContext := context.Background()
		// create mock of IRepositoryServices interface
		repositoryServices := new(mocks.IRepositoryServices)
		// pass mocked object in NewGithub constructor/func
		github := NewGithub(backgroundContext, repositoryServices)

		fullName := "ABC"
		starCount := 10
		repo := &Repository{
			FullName:        &fullName,
			Description:     &fullName,
			ForksCount:      &starCount,
			StargazersCount: &starCount,

		}
        // when code calls Get method of IRepositoryServices, it will return repo mocked object
		repositoryServices.On("Get", backgroundContext, "golang-coach", "Lessons").Return(repo, nil, nil)
		pack, _ := github.GetPackageRepoInfo("golang-coach", "Lessons")
        // assert
		So(pack.ForksCount, ShouldEqual, starCount)
	})

	Convey("Should return error when failed to retrieve  repository information", t, func() {
		backgroundContext := context.Background()
		repositoryServices := new(mocks.IRepositoryServices)
		github := NewGithub(backgroundContext, repositoryServices)
		repositoryServices.On("Get", backgroundContext, "golang-coach", "Lessons").Return(nil, nil, errors.New("Error has been occurred"))
		_, err := github.GetPackageRepoInfo("golang-coach", "Lessons")
		So(err, ShouldNotBeEmpty)
	})
}
