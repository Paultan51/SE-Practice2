package unit_test

import (
	"testing"

	"github.com/Paultan51/SE-Practice2/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func Test_Profile(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run(`case success`, func(t *testing.T) {
		p := entity.Profile{
			StudentID: "B1234567",
			Name:      "abcdefghijklmnopqrstuvwxyz",
			Email:     "p@gmail.p",
			Password:  "123456",
			Phone:     "0123456789",
			Url:       "p.com",
		}
		ok, err := govalidator.ValidateStruct(p)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}
func Test_Profile_StudentID_Fail(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run(`case studentID fail`, func(t *testing.T) {
		p := entity.Profile{
			StudentID: "",
			Name:      "abcdefghijklmnopqrstuvwxyz",
			Email:     "p@gmail.p",
			Password:  "123456",
			Phone:     "0123456789",
			Url:       "p.com",
		}
		ok, err := govalidator.ValidateStruct(p)
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
		ok, err := govalidator.ValidateStruct(p)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Wrong Format StudentID!"))
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
			Url:       "p.com",
		}
		ok, err := govalidator.ValidateStruct(p)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Wrong length name 10-30!"))
	})
}
