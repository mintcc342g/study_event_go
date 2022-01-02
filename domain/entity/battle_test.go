package entity_test

import (
	"study-event-go/application/dto"
	"study-event-go/domain/entity"
	"study-event-go/domain/vo"

	"github.com/juju/errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Battle Domain", func() {

	var (
		alarmDTO *dto.Alarm
		alarm    *entity.Alarm

		expectedError error
	)

	Describe("Alarm", func() {
		BeforeEach(func() {
			alarmDTO = &dto.Alarm{
				GardenID:     1,
				CaveLocation: "shibuya",
				Huges: []*dto.Huge{&dto.Huge{
					Class: 1,
					Type:  0,
				}},
				TotalCount: 3,
				AlertLevel: 2,
			}

			alarm = &entity.Alarm{
				GardenID:     1,
				CaveLocation: "shibuya",
				Huges: []*vo.Huge{&vo.Huge{
					Class: 1,
					Type:  0,
				}},
				TotalCount: 3,
				AlertLevel: 2,
			}
		})

		Context("pass to create a new alarm", func() {
			It("should not error", func() {
				result, err := entity.NewAlarm(alarmDTO)
				Expect(result).To(Equal(alarm))
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("fail to create a new alarm by invalid garden id", func() {
			BeforeEach(func() {
				alarmDTO.GardenID = 0
				expectedError = errors.BadRequestf("invalid garden ID")
			})

			It("should be error", func() {
				result, err := entity.NewAlarm(alarmDTO)
				Expect(result).To(BeNil())
				Expect(err.Error()).To(Equal(expectedError.Error()))
			})
		})

		Context("fail to create a new alarm by invalid cave location", func() {
			BeforeEach(func() {
				alarmDTO.CaveLocation = ""
				expectedError = errors.BadRequestf("invalid cave location")
			})

			It("should be error", func() {
				result, err := entity.NewAlarm(alarmDTO)
				Expect(result).To(BeNil())
				Expect(err.Error()).To(Equal(expectedError.Error()))
			})
		})

		Context("fail to create a new alarm by invalid total count", func() {
			BeforeEach(func() {
				alarmDTO.TotalCount = 0
				expectedError = errors.BadRequestf("invalid total count")
			})

			It("should be error", func() {
				result, err := entity.NewAlarm(alarmDTO)
				Expect(result).To(BeNil())
				Expect(err.Error()).To(Equal(expectedError.Error()))
			})
		})

		Context("fail to create a new alarm by invalid huges", func() {
			BeforeEach(func() {
				alarmDTO.Huges = nil
				expectedError = errors.BadRequestf("invalid huges")
			})

			It("should be error", func() {
				result, err := entity.NewAlarm(alarmDTO)
				Expect(result).To(BeNil())
				Expect(err.Error()).To(Equal(expectedError.Error()))
			})
		})
	})
})
