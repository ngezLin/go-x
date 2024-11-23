package mail

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"mime/multipart"
	"strings"
)

type email struct {
	from        string
	contentType string
	subject     string
	content     string
	to          []string
	cc          []string
	bcc         []string
	others      []string
	attachments []attachment
	signature   string
}
type attachment struct {
	name     string
	fileType string
	content  *bytes.Buffer
}

// Set subject of the email
func (m *email) SetSubject(ctx context.Context, subject string) {
	m.subject = subject
}

// Set sender name and email
func (m *email) SetFrom(ctx context.Context, fromName, fromEmail string) {
	headerName := "From:"
	fromEmail = fmt.Sprintf("<%s>", fromEmail)
	m.from = strings.Join([]string{headerName, fromName, fromEmail}, " ")
}

// Set recipient emails
func (m *email) SetTo(ctx context.Context, to []string) {
	m.to = to
}

// Set CC for the email
func (m *email) SetCC(ctx context.Context, cc []string) {
	m.cc = cc
}

// Set email as HTML based content
func (m *email) SetContentHTML(ctx context.Context) {
	headerName := "Content-Type:"
	m.contentType = strings.Join([]string{headerName, "text/html; charset=UTF-8"}, " ")
}

// Set content of the email
func (m *email) SetContent(ctx context.Context, content string) {
	m.content = content
}

// Set content of the email
func (m *email) SetOthers(ctx context.Context, content string) {
	m.others = append(m.others, content)
}

// Set attachments of the email
func (m *email) AddAttachment(ctx context.Context, name, fileType string, content *bytes.Buffer) {
	m.attachments = append(m.attachments, attachment{
		name:     name,
		fileType: fileType,
		content:  content,
	})
}

func (m *email) SetSignature(ctx context.Context, content string) {

}

func (m *email) SetDefaultSignature(ctx context.Context, value string) {
	if value == "" {
		value = DEFAULT_SIGNATURE
	}
}

// Build final content of the email.
func (m *email) build() string {
	var dataset []string

	if m.subject != "" {
		dataset = append(dataset, m.subject)
	}

	if m.from != "" {
		dataset = append(dataset, m.from)
	}

	if len(m.to) > 0 {
		headerName := "To:"
		headerValue := strings.Join(m.to, ",")
		toHeader := strings.Join([]string{headerName, headerValue}, " ")
		dataset = append(dataset, toHeader)
	}

	if len(m.cc) > 0 {
		headerName := "Cc:"
		headerValue := strings.Join(m.cc, ",")

		ccHeader := strings.Join([]string{headerName, headerValue}, " ")
		dataset = append(dataset, ccHeader)
	}

	if m.contentType != "" {
		dataset = append(dataset, m.contentType)
	}

	if m.content != "" {
		dataset = append(dataset, m.content)
	}

	//others cover custom need like custom content type
	dataset = append(dataset, m.others...)

	return strings.Join(dataset, "\r\n")
}

func (m *email) buildBytes() []byte {
	buf := bytes.NewBuffer(nil)
	withAttachments := len(m.attachments) > 0

	buf.WriteString("MIME-Version: 1.0\n")
	buf.WriteString(fmt.Sprintf("Subject: %s\n", m.subject))
	buf.WriteString(fmt.Sprintf("To: %s\n", strings.Join(m.to, ",")))
	if len(m.cc) > 0 {
		buf.WriteString(fmt.Sprintf("Cc: %s\n", strings.Join(m.cc, ",")))
	}
	if len(m.bcc) > 0 {
		buf.WriteString(fmt.Sprintf("Bcc: %s\n", strings.Join(m.bcc, ",")))
	}
	if m.signature != "" {
		buf.WriteString(fmt.Sprintf("%s\n", m.signature))
	}

	writer := multipart.NewWriter(buf)
	boundary := writer.Boundary()
	if len(m.attachments) > 0 {
		buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=\"%s\"\n", boundary))
		buf.WriteString(fmt.Sprintf("r\n--%s\r\n", boundary))
	} else {
		buf.WriteString("Content-Type: text/plain; charset=utf-8\n")
	}

	buf.WriteString("Content-Type: text/html; charset=utf-8\n")
	buf.WriteString("Content-Transfer-Encoding: 7bit\n")
	buf.WriteString(fmt.Sprintf("%s\n", m.content))

	if withAttachments {
		for _, v := range m.attachments {
			bites := v.content.Bytes()
			buf.WriteString(fmt.Sprintf("r\n--%s\r\n", boundary))
			buf.WriteString(fmt.Sprintf("Content-Type: %s\n; name=\"%s\"", v.fileType, v.name))
			buf.WriteString("Content-Transfer-Encoding: base64\n")
			buf.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=\"%s\"\n", v.name))

			b := make([]byte, base64.StdEncoding.EncodedLen(len(bites)))
			base64.StdEncoding.Encode(b, bites)
			buf.Write(b)
			buf.WriteString(fmt.Sprintf("\n--%s--", boundary))
		}
	}
	return buf.Bytes()
}
