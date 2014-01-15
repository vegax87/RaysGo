package helpers

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"github.com/astaxie/beego"
	"github.com/russross/blackfriday"
	"html/template"
	"strconv"
	"time"
)

func EncryptPassword(password string, salt []byte) string {
	if salt == nil {
		m := md5.New()
		m.Write([]byte(time.Now().String()))
		s := hex.EncodeToString(m.Sum(nil))
		salt = []byte(s[2:10])
	}
	mac := hmac.New(sha256.New, salt)
	mac.Write([]byte(password))
	//s := fmt.Sprintf("%x", (mac.Sum(salt)))
	s := hex.EncodeToString(mac.Sum(nil))

	hasher := sha1.New()
	hasher.Write([]byte(s))

	//result := fmt.Sprintf("%x", (hasher.Sum(nil)))
	result := hex.EncodeToString(hasher.Sum(nil))

	p := string(salt) + result

	return p
}

func ValidatePassword(hashed string, input_password string) bool {
	salt := hashed[0:8]
	if hashed == EncryptPassword(input_password, []byte(salt)) {
		return true
	} else {
		return false
	}
	return false
}

func Str2Int(s string) (int, error) {
	v, e := strconv.Atoi(s)
	return v, e
}

func Str2Int64(s string) (int64, error) {
	var (
		v int
		e error
	)
	if v, e = Str2Int(s); e != nil {
		return 0, e
	}
	return int64(v), e
}

func Markdown(raw []byte) []byte {
	return blackfriday.MarkdownCommon(raw)
}

func MarkdownHtml(s string) template.HTML {
	return beego.Str2html(string(Markdown([]byte(s))))
}

/* time */

// the begin time of today
func ThisDate() time.Time {
        t := time.Now()
        year, month, day := t.Date()

        return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

// the begin time of this week
func ThisWeek() time.Time {
        t := time.Now()
        year, month, day := t.AddDate(0, 0, -1*int(t.Weekday())).Date()

        return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

// the begin time of this month
func ThisMonth() time.Time {
        t := time.Now()
        year, month, _ := t.Date()

        return time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
}

// the begin time of this year
func ThisYear() time.Time {
        t := time.Now()
        year, _, _ := t.Date()

        return time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
}
