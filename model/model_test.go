package model_test

import (
	"errors"
	"net/url"
	"testing"

	"github.com/axelspringer/swerve/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDBDomain(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DB domain suite")
}

var _ = Describe("type Domain", func() {
	It("Domain struct validating", func() {
		redirect := &model.Redirect{}
		errList := redirect.Validate()
		Expect(errList).To(Equal(
			errors.New(model.ErrInvalidHTTPCode)))
	})

	Context("type Domain struct", func() {
		It("Domain struct redirect", func() {
			redirect := &model.Redirect{
				RedirectFrom: "example.com",
				RedirectTo:   "https://www.example.com",
				Code:         301,
			}

			url, err := url.Parse("https://example.com/path/to/promote?with=query")
			Expect(err).To(BeNil())

			redirectURL, redirectCode := redirect.GetRedirect(url)
			Expect("https://www.example.com").To(Equal(redirectURL))
			Expect(redirectCode).To(Equal(301))
		})

		It("Domain struct redirect with query string not promotable and trailing slash", func() {
			redirect := &model.Redirect{
				RedirectFrom: "example.com",
				RedirectTo:   "https://www.example.com/",
				Code:         301,
				Promotable:   false,
			}

			url, err := url.Parse("https://example.com/path/to/promote?with=query")
			Expect(err).To(BeNil())

			redirectURL, redirectCode := redirect.GetRedirect(url)
			Expect("https://www.example.com/").To(Equal(redirectURL))
			Expect(redirectCode).To(Equal(301))
		})

		It("Domain struct redirect without query string and not promotable", func() {
			redirect := &model.Redirect{
				RedirectFrom: "example.com",
				RedirectTo:   "https://www.example.com/path/to/promote/",
				Code:         301,
				Promotable:   false,
			}

			url, err := url.Parse("https://example.com/path/to/promote/")
			Expect(err).To(BeNil())

			redirectURL, redirectCode := redirect.GetRedirect(url)
			Expect("https://www.example.com/path/to/promote/").To(Equal(redirectURL))
			Expect(redirectCode).To(Equal(301))
		})

		It("Domain struct redirect promotable without query string", func() {
			redirect := &model.Redirect{
				RedirectFrom: "example.com",
				RedirectTo:   "https://www.example.com",
				Code:         301,
				Promotable:   true,
			}

			url, err := url.Parse("https://example.com/path/to/promote/")
			Expect(err).To(BeNil())

			redirectURL, redirectCode := redirect.GetRedirect(url)
			Expect("https://www.example.com/path/to/promote/").To(Equal(redirectURL))
			Expect(redirectCode).To(Equal(301))
		})

		It("Domain struct redirect promotable with query string", func() {
			redirect := &model.Redirect{
				RedirectFrom: "example.com",
				RedirectTo:   "https://www.example.com",
				Code:         301,
				Promotable:   true,
			}

			url, err := url.Parse("https://example.com/path/to/promote?with=query")
			Expect(err).To(BeNil())

			redirectURL, redirectCode := redirect.GetRedirect(url)
			Expect("https://www.example.com/path/to/promote?with=query").To(Equal(redirectURL))
			Expect(redirectCode).To(Equal(301))
		})

		It("Domain struct redirect promotable path mapping", func() {
			redirect := &model.Redirect{
				RedirectFrom: "example.com",
				RedirectTo:   "https://www.example.com",
				Code:         301,
				Promotable:   true,
				PathMaps: []model.PathMap{
					model.PathMap{From: "/old/path", To: "/new/target"},
				},
			}

			url, err := url.Parse("https://example.com/path/to/promote?with=query")
			Expect(err).To(BeNil())
			redirectURL, redirectCode := redirect.GetRedirect(url)
			Expect("https://www.example.com/path/to/promote?with=query").To(Equal(redirectURL))
			Expect(redirectCode).To(Equal(301))

			url, err = url.Parse("https://example.com/old/path/to/promote?with=query")
			Expect(err).To(BeNil())
			redirectURL, redirectCode = redirect.GetRedirect(url)
			Expect("https://www.example.com/new/target/to/promote?with=query").To(Equal(redirectURL))
			Expect(redirectCode).To(Equal(301))

			mapping := redirect.PathMaps
			mapping[0].To = "https://theotherserver.com/new/target/"
			url, err = url.Parse("https://example.com/old/path/to/promote?with=query")
			Expect(err).To(BeNil())
			redirectURL, redirectCode = redirect.GetRedirect(url)
			Expect("https://theotherserver.com/new/target/to/promote?with=query").To(Equal(redirectURL))
			Expect(redirectCode).To(Equal(301))

		})
		It("Domain struct redirect path mapping and not promotable", func() {
			redirect := &model.Redirect{
				RedirectFrom: "example.com",
				RedirectTo:   "https://www.example.com/",
				Code:         301,
				Promotable:   false,
				PathMaps: []model.PathMap{
					model.PathMap{From: "/old/path", To: "/new/target"},
				},
			}

			url, err := url.Parse("https://example.com/path/to/promote?with=query")
			Expect(err).To(BeNil())
			redirectURL, redirectCode := redirect.GetRedirect(url)
			Expect("https://www.example.com/").To(Equal(redirectURL))
			Expect(redirectCode).To(Equal(301))

			url, err = url.Parse("https://example.com/old/path/to/promote?with=query")
			Expect(err).To(BeNil())
			redirectURL, redirectCode = redirect.GetRedirect(url)
			Expect("https://www.example.com/new/target").To(Equal(redirectURL))
			Expect(redirectCode).To(Equal(301))

			mapping := redirect.PathMaps
			mapping[0].To = "https://theotherserver.com/new/target/"
			url, err = url.Parse("https://example.com/old/path/to/promote?with=query")
			Expect(err).To(BeNil())
			redirectURL, redirectCode = redirect.GetRedirect(url)
			Expect("https://theotherserver.com/new/target/to/promote").To(Equal(redirectURL))
			Expect(redirectCode).To(Equal(301))
		})
	})
})
