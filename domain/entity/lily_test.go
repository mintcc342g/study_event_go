package entity_test

import (
	"study-event-go/application/dto"
	"study-event-go/domain/entity"
	"study-event-go/domain/vo"
	"study-event-go/types"

	"github.com/juju/errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Lily Domain", func() {

	var (
		lilyDTO *dto.Lily
		lily    *entity.Lily

		charmModelIDs []types.CharmModelID
		charmModels   []*entity.CharmModel

		expectedError error
	)

	Describe("Lily Create", func() {
		BeforeEach(func() {
			lilyDTO = &dto.Lily{
				FirstName:  "mirai",
				MiddleName: "maria",
				LastName:   "kishimoto",
				GardenID:   1,
			}

			lily = &entity.Lily{
				Name:     vo.NewName("mirai", "maria", "kishimoto"),
				GardenID: 1,
			}
		})

		Context("pass to create a new lily without birth", func() {
			It("should not error", func() {
				result, err := entity.NewLily(lilyDTO)
				Expect(result).To(Equal(lily))
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("pass to create a new lily without middle name", func() {
			BeforeEach(func() {
				lilyDTO.MiddleName = ""
				lily.Name.Middle = ""
			})

			It("should not error", func() {
				result, err := entity.NewLily(lilyDTO)
				Expect(result).To(Equal(lily))
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("fail to create a new lily by invalid first name", func() {
			BeforeEach(func() {
				lilyDTO.FirstName = ""
				expectedError = errors.BadRequestf("invalid name")
			})

			It("should be error", func() {
				result, err := entity.NewLily(lilyDTO)
				Expect(result).To(BeNil())
				Expect(err.Error()).To(Equal(expectedError.Error()))
			})
		})

		Context("fail to create a new lily by invalid last name", func() {
			BeforeEach(func() {
				lilyDTO.LastName = ""
				expectedError = errors.BadRequestf("invalid name")
			})

			It("should be error", func() {
				result, err := entity.NewLily(lilyDTO)
				Expect(result).To(BeNil())
				Expect(err.Error()).To(Equal(expectedError.Error()))
			})
		})

		Context("fail to create a new lily by invalid garden id", func() {
			BeforeEach(func() {
				lilyDTO.GardenID = 0
				expectedError = errors.BadRequestf("invalid garden id")
			})

			It("should be error", func() {
				result, err := entity.NewLily(lilyDTO)
				Expect(result).To(BeNil())
				Expect(err.Error()).To(Equal(expectedError.Error()))
			})
		})
	})

	Describe("Lily Contract", func() {
		BeforeEach(func() {
			charmModelIDs = append(charmModelIDs,
				types.CharmModelID(1),
				types.CharmModelID(2),
				types.CharmModelID(3),
			)

			charmModels = append(charmModels,
				&entity.CharmModel{ID: 1},
				&entity.CharmModel{ID: 2},
				&entity.CharmModel{ID: 3},
			)

			lily = &entity.Lily{
				ID:     30,
				Charms: []*entity.Charm{},
			}

			lily.Charms = append(lily.Charms,
				&entity.Charm{ModelID: 1, OwnerID: lily.ID},
				&entity.Charm{ModelID: 2, OwnerID: lily.ID},
				&entity.Charm{ModelID: 3, OwnerID: lily.ID},
			)
		})

		Context("a lily success to contract with charms", func() {
			It("should not error", func() {
				err := lily.ContractWith(charmModelIDs, charmModels...)
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("a lily fail to contract with charms by invalid model ids", func() {
			BeforeEach(func() {
				charmModels = charmModels[1:]

				expectedError = errors.BadRequestf("invalid charm model ids")
			})

			It("should be error", func() {
				err := lily.ContractWith(charmModelIDs, charmModels...)
				Expect(err.Error()).To(Equal(expectedError.Error()))
			})
		})
	})
})
