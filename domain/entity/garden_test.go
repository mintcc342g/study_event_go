package entity_test

import (
	"study-event-go/application/dto"
	"study-event-go/domain/entity"

	"github.com/juju/errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Garden Domain", func() {

	var (
		gardenDTO  *dto.Garden
		garden     *entity.Garden
		comparable *entity.Garden

		expectedError error
	)

	Describe("Garden Create", func() {
		BeforeEach(func() {
			gardenDTO = &dto.Garden{
				Name:     "testname",
				Location: "test location",
			}

			garden = &entity.Garden{
				Name:     "testname",
				Location: "test location",
			}
		})

		Context("pass to create a new garden without mentorship id", func() {
			It("should not error", func() {
				result, err := entity.NewGarden(gardenDTO)
				Expect(result).To(Equal(garden))
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("pass to create a new garden with mentorship id", func() {
			BeforeEach(func() {
				gardenDTO.MentorshipID = 1
				garden.MentorshipID = 1
			})

			It("should not error", func() {
				result, err := entity.NewGarden(gardenDTO)
				Expect(result).To(Equal(garden))
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("fail to create a new garden by invalid name", func() {
			BeforeEach(func() {
				gardenDTO.Name = ""
				expectedError = errors.BadRequestf("invalid name")
			})

			It("should be error", func() {
				result, err := entity.NewGarden(gardenDTO)
				Expect(result).To(BeNil())
				Expect(err.Error()).To(Equal(expectedError.Error()))
			})
		})

		Context("fail to create a new garden by invalid location", func() {
			BeforeEach(func() {
				gardenDTO.Location = ""
				expectedError = errors.BadRequestf("invalid location")
			})

			It("should be error", func() {
				result, err := entity.NewGarden(gardenDTO)
				Expect(result).To(BeNil())
				Expect(err.Error()).To(Equal(expectedError.Error()))
			})
		})
	})

	Describe("Garden Update", func() {
		BeforeEach(func() {
			gardenDTO = &dto.Garden{
				ID:       1,
				Name:     "updatedname",
				Location: "somewhere",
			}

			garden = &entity.Garden{
				ID:       1,
				Name:     "updatedname",
				Location: "somewhere",
			}

			comparable = &entity.Garden{
				ID:       1,
				Name:     "updatedname",
				Location: "somewhere",
			}
		})

		Context("pass to update a garden", func() {
			It("should not error", func() {
				err := garden.Update(gardenDTO, comparable)
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("fail to update a garden by duplicated name", func() {
			BeforeEach(func() {
				comparable.ID = 2
				expectedError = errors.AlreadyExistsf("The name [%s]", comparable.Name)
			})

			It("should be error", func() {
				err := garden.Update(gardenDTO, comparable)
				Expect(err.Error()).To(Equal(expectedError.Error()))
			})
		})
	})
})
