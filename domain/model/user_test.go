package model_test

import (
	"github.com/EusRique/money-transfer/domain/model"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestNewUser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NewUser Suite")
}

var _ = Describe("User.isValid()", func() {
	Context("When validation user error does not return error", func() {
		It("Returns nil", func() {
			userValidation := model.User{
				Name:     "any_name",
				LastName: "any_last_name",
				Document: "any_document",
				Email:    "any_email@mail.com",
				Password: "any_password",
				IsSeller: false,
			}

			response := userValidation.IsValid()
			Expect(response).To(BeNil())
		})
	})

	Context("When validation user error return an error", func() {
		It("Returns error", func() {
			userValidation := model.User{
				LastName: "any_last_name",
				Document: "any_document",
				Email:    "any_email@mail.com",
				Password: "any_password",
				IsSeller: false,
			}

			response := userValidation.IsValid()
			Expect(response).To(MatchError(response.Error()))
		})
	})

})
