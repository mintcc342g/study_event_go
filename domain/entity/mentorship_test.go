package entity_test

import (
	"study-event-go/application/dto"
	"study-event-go/domain/entity"

	"github.com/juju/errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Mentorship Domain", func() {

	var (
		mentorshipDTO *dto.Mentorship
		mentorship    *entity.Mentorship

		expectedError error
	)

	Describe("Mentorship Create", func() {
		BeforeEach(func() {
			mentorshipDTO = &dto.Mentorship{
				Name: " testName  ",
			}

			mentorship = &entity.Mentorship{
				Name: "testname",
			}
		})

		Context("pass to create a new mentorship", func() {
			It("should not error", func() {
				result, err := entity.NewMentorship(mentorshipDTO)
				Expect(result).To(Equal(mentorship))
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("fail to create a new mentorship by invalid name", func() {
			BeforeEach(func() {
				mentorshipDTO.Name = "   "
				expectedError = errors.BadRequestf("invalid name")
			})

			It("should be error", func() {
				result, err := entity.NewMentorship(mentorshipDTO)
				Expect(result).To(BeNil())
				Expect(err.Error()).To(Equal(expectedError.Error()))
			})
		})
	})

	Describe("Mentorship Update", func() {
		BeforeEach(func() {
			mentorshipDTO = &dto.Mentorship{
				Name: " updatedName  ",
			}

			mentorship = &entity.Mentorship{
				Name: "updatedname",
			}
		})

		Context("pass to update a mentorship", func() {
			It("should not error", func() {
				err := mentorship.Update(mentorshipDTO)
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("fail to update a mentorship by invalid name", func() {
			BeforeEach(func() {
				mentorshipDTO.Name = "   "
				expectedError = errors.BadRequestf("invalid name")
			})

			It("should be error", func() {
				err := mentorship.Update(mentorshipDTO)
				Expect(err.Error()).To(Equal(expectedError.Error()))
			})
		})
	})
})
