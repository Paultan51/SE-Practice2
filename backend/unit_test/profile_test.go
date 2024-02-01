package unit_test

import (
	"fmt"
	"testing"

	"github.com/Paultan51/SE-Practice2/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/asaskevich/govalidator"
	GV "github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func Test_Profile(t *testing.T) {
	g := NewWithT(t)
	t.Run(`case success`, func(t *testing.T) {
		p := entity.Profile{
			StudentID: "B1234567",
			Name:      "abcdefghijklmnopqrstuvwxyz",
			Email:     "p@gmail.p",
			Password:  "123456",
			Phone:     "0123456789",
			Url:       "p.com",
		}
		ok, err := ValidateStruct(p)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}
func Test_Profile_StudentID_Fail(t *testing.T) {
	g := NewWithT(t)
	t.Run(`case studentID fail`, func(t *testing.T) {
		p := entity.Profile{
			StudentID: "",
			Name:      "abcdefghijklmnopqrstuvwxyz",
			Email:     "p@gmail.p",
			Password:  "123456",
			Phone:     "0123456789",
			Url:       "p.com",
		}
		ok, err := ValidateStruct(p)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Need StudentID"))
	})
	t.Run(`case studentID fail`, func(t *testing.T) {
		p := entity.Profile{
			StudentID: "B640",
			Name:      "abcdefghijklmnopqrstuvwxyz",
			Email:     "p@gmail.p",
			Password:  "123456",
			Phone:     "0123456789",
			Url:       "p.com",
		}
		ok, err := ValidateStruct(p)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal(fmt.Sprintf("StudentID: %s does not validate as matches(^[BMS]\\d{7}$)", p.StudentID)))
	})
}
func Test_Profile_Name_Fail(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run(`case name fail`, func(t *testing.T) {
		p := entity.Profile{
			StudentID: "B6405816",
			Name:      "",
			Email:     "p@gmail.p",
			Password:  "123456",
			Phone:     "0123456789",
			Url:       "p.com",
		}
		ok, err := govalidator.ValidateStruct(p)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Need Name"))
	})
	t.Run(`case name fail`, func(t *testing.T) {
		p := entity.Profile{
			StudentID: "B6405816",
			Name:      "abcde",
			Email:     "p@gmail.p",
			Password:  "123456",
			Phone:     "0123456789",
			Url:       "p.olo",
		}
		ok, err := govalidator.ValidateStruct(p)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Wrong length name 10-30!"))
	})
}
func Test_Profile_Email_Fail(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run(`case email fail`, func(t *testing.T) {
		p := entity.Profile{
			StudentID: "B6405816",
			Name:      "abcdefghijklmnopqrstuvwxyz",
			Email:     "aaa",
			Password:  "123456",
			Phone:     "0123456789",
			Url:       "p.com",
		}
		ok, err := GV.ValidateStruct(p)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Wrong Email Format!"))
	})
}
