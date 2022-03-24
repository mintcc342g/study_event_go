package entity_test

import (
	"study-event-go/application/dto"
	"study-event-go/domain/entity"
	"study-event-go/types"

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
				Name:         "testname",
				Location:     "test location",
				LegionSystem: 2,
			}

			garden = &entity.Garden{
				Name:         "testname",
				Location:     "test location",
				LegionSystem: 2,
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

		Context("fail to create a new garden by invalid legion system", func() {
			BeforeEach(func() {
				gardenDTO.LegionSystem = types.NoneLegionSystem
				expectedError = errors.BadRequestf("invalid legion system")
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
			gardenDTO = &dto.Garden{ID: 1}
			garden = &entity.Garden{ID: 1, Name: "name"}
		})

		Context("pass to update a garden name", func() {
			BeforeEach(func() {
				gardenDTO.Name = "new name"
				comparable = nil
			})

			It("should not error", func() {
				err := garden.Update(gardenDTO, comparable)
				Expect(err).NotTo(HaveOccurred())
				Expect(garden.Name).To(Equal(gardenDTO.Name))
			})
		})

		Context("pass to update a garden location", func() {
			BeforeEach(func() {
				gardenDTO.Location = "new location"
				comparable = garden
			})

			It("should not error", func() {
				err := garden.Update(gardenDTO, comparable)
				Expect(err).NotTo(HaveOccurred())
				Expect(garden.Location).To(Equal(gardenDTO.Location))
			})
		})

		Context("pass to update a garden mentorship", func() {
			BeforeEach(func() {
				gardenDTO.MentorshipID = 2
				comparable = garden
			})

			It("should not error", func() {
				err := garden.Update(gardenDTO, comparable)
				Expect(err).NotTo(HaveOccurred())
				Expect(garden.MentorshipID).To(Equal(gardenDTO.MentorshipID))
			})
		})

		Context("pass to update a garden legion system", func() {
			BeforeEach(func() {
				gardenDTO.LegionSystem = 2
				comparable = garden
			})

			It("should not error", func() {
				err := garden.Update(gardenDTO, comparable)
				Expect(err).NotTo(HaveOccurred())
				Expect(garden.LegionSystem).To(Equal(gardenDTO.LegionSystem))
			})
		})

		Context("fail to update a garden by duplicated name", func() {
			BeforeEach(func() {
				comparable = &entity.Garden{
					ID:   2,
					Name: "name",
				}
				expectedError = errors.AlreadyExistsf("The name [%s]", comparable.Name)
			})

			It("should be error", func() {
				err := garden.Update(gardenDTO, comparable)
				Expect(err.Error()).To(Equal(expectedError.Error()))
			})
		})
	})
})
