package entity

type Profile struct {
	StudentID string `valid:"required~Need StudentID, matches(^[BMS]\\d{7}$)"`
	Name      string `valid:"required~Need Name, stringlength(10|30)~Wrong length name 10-30!"`
	Email     string `valid:"email~Wrong Email Format!"`
	Password  string `valid:"required~Need Password, minstringlength(6)"`
	Phone     string `valid:"required~Need Phone Number, numeric~number 10 digits, stringlength(10|10)~number 10 digits"`
	Url       string `valid:"url~Wrong Format URL!"`
}
