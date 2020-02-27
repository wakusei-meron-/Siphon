package core

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	tm "github.com/buger/goterm"
	"github.com/eiannone/keyboard"
	"github.com/jhillyerd/enmime"
	"net/mail"
	"strings"
)

var (
	profile string
	bucket string
	prefix string
)

type (
	Peaberry interface {
		Run()
	}

	peaberryImpl struct {
		s3 *s3.S3
	}
)

func NewPeaberry() Peaberry {
	sess := session.Must(
		session.NewSessionWithOptions(
			session.Options{
				Profile: profile,
			}))
	sess.Config = sess.Config.WithRegion("ap-northeast-1")
	return &peaberryImpl{
		s3: s3.New(sess),
	}
}

func (p *peaberryImpl) setup() {
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
}

func (p *peaberryImpl) Run() {
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	mails, err := p.findMails()
	if err != nil {
		panic(err)
	}
	x, y := 1, 3
	for {
		tm.Clear() // Clear current screen
		for i, mail := range mails {
            tm.Println(fmt.Sprintf("%d | %s | %s", i, p.fmtAddress(mail.From), p.fmtText(mail.Text)))
		}
		tm.Flush() // Call it every time at the end of rendering
		char, key, err := keyboard.GetKey()
		if (err != nil) {
			panic(err)
		} else if key == keyboard.KeyEsc || key == keyboard.KeyCtrlC || char == 'q' {
			return
		}
		tm.MoveCursor(1, 2)
		tm.Printf("You pressed: %q, x:%d, y:%d", char, x, y)

		switch char {
		case 'h':
			if x > 1 {x -= 1}
		case 'j':
			if y < tm.Height() {y += 1}
		case 'k':
			if y > 1 {y -= 1}
		case 'l':
			if x < tm.Width() {x += 1}
		}
		tm.MoveCursor(x, y)
		tm.Print("*")
	}
}

func (p *peaberryImpl) fmtAddress(address *mail.Address) string {
	if address == nil {
		return ""
	}
	if address.Name != "" {
		return address.Name
	}
	return address.Address
}

func (p *peaberryImpl) fmtText(text string) string {
	s := convNewline(text, " ")
	return substr(s, int64(tm.Width() - 30))
}

func substr(s string, n int64) string {
	return string([]rune(s)[:n])
}

func convNewline(str, nlcode string) string {
	return strings.NewReplacer(
		"\r\n", nlcode,
		"\r", nlcode,
		"\n", nlcode,
	).Replace(str)
}

func (p *peaberryImpl) findMails() ([]*Mail, error) {
	input := &s3.ListObjectsV2Input{
		Bucket:  aws.String(bucket),
		Prefix:  aws.String(prefix),
		MaxKeys: aws.Int64(3),
	}
	resp, err := p.s3.ListObjectsV2(input)
	if err != nil {
		panic(err)
	}
	mails := make([]*Mail, len(resp.Contents))
	for i, content := range resp.Contents {
		obj, err := p.s3.GetObject(&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    content.Key,
		})
		if err != nil {
			panic(err)
		}
		envelop, err := enmime.ReadEnvelope(obj.Body)
		if err != nil {
			panic(err)
		}
		from, err := mail.ParseAddress(envelop.GetHeader("From"))
		if err != nil {
			return nil, err
		}
		toValues := envelop.GetHeaderValues("To")
		to := make([]*mail.Address, len(toValues))
		for i, v := range toValues {
			a, err := mail.ParseAddress(v)
			if err != nil {
				return nil, err
			}
			to[i] = a
		}
		ccValues := envelop.GetHeaderValues("Cc")
		cc := make([]*mail.Address, len(ccValues))
		for i, v := range ccValues {
			a, err := mail.ParseAddress(v)
			if err != nil {
				return nil, err
			}
			cc[i] = a
		}
		mails[i] = &Mail{
			From: from,
			To: to,
			CC: cc,
			Subject: envelop.GetHeader("Subject"),
			Text: envelop.Text,
		}
	}
	return mails, nil
}

