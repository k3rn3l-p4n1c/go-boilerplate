package provider_test

import (
	"context"
	"testing"

	"github.com/cafebazaar/go-boilerplate/pkg/postview"
	"github.com/stretchr/testify/suite"
	"golang.org/x/xerrors"
)

type MemoryProviderTestSuite struct {
	suite.Suite

	provider PostProvider
}

func TestMemoryProviderTestSuite(t *testing.T) {
	suite.Run(t, new(MemoryProviderTestSuite))
}

func (s *MemoryProviderTestSuite) TestShouldReturnNotFoundInitially() {
	_, err := s.provider.GetPost(context.Background(), "token")
	s.True(xerrors.Is(err, ErrNotFound))
}

func (s *MemoryProviderTestSuite) TestShouldReturnPostAfterAdd() {
	err := s.provider.AddPost(context.Background(), &postview.Post{
		Token: "abcd",
		Title: "a title",
	})
	s.Nil(err)
	if err != nil {
		return
	}

	post, err := s.provider.GetPost(context.Background(), "abcd")
	s.Nil(err)
	if err != nil {
		return
	}

	s.Equal("abcd", post.Token)
	s.Equal("a title", post.Title)
}

func (s *MemoryProviderTestSuite) SetupTest() {
	s.provider = NewMemory()
}
