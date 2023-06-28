package repositories

import (
	"github.com/go-gomail/gomail"
	"github.com/sirupsen/logrus"
)

var submissionsRepo SubmissionsRepository = SubmissionsRepository{}

func SendEmail(reciept string, subject string, body string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "support@sistadm.us")
	m.SetHeader("To", reciept)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.elasticemail.com", 2525, "support@sistadm.us", "A48B8D1F28A8F882EE876A5D6037712AA9A2")
	if err := d.DialAndSend(m); err != nil {
		logrus.Error("Failed to send email to ", reciept, " with subject ", subject, " and body ", body, " with error ", err)
	}
}

func SendMeritEmail(submissionId int) {
	s, err := submissionsRepo.GetSubmissionById(submissionId)
	logrus.Info("Merit Mail Sent ", s.Application.BasicDetails.Email)
	if err != nil {
		logrus.Error("Failed to get submission with id ", submissionId, " with error ", err)
		return
	}

	SendEmail(s.Application.BasicDetails.Email, "Admission Approved", `
		<p>Dear `+s.Application.BasicDetails.Name+`,</p>
		<p>Congratulations! Your admission has been approved. Please check your admission status on <a href="https://sistadm.us">sistadm.us</a>
			and follow the instructions to complete your admission process.</p>
		</p>
		<p>Thank you for your participation.</p>
		<p>Best regards,</p>
		<p>SIST Admission Team</p>
	`)

}

func SendSuccessfulAdmissionEmail(submissionId int) {
	s, err := submissionsRepo.GetSubmissionById(submissionId)
	logrus.Info("Successful Mail Sent ", s.Application.BasicDetails.Email)
	if err != nil {
		logrus.Error("Failed to get submission with id ", submissionId, " with error ", err)
		return
	}

	SendEmail(s.Application.BasicDetails.Email, "Admission Approved", `
		<p>Dear `+s.Application.BasicDetails.Name+`,</p>
		<p>Congratulations! You have successfully 
			completed your admission process. Please check your admission status on <a href="https://sistadm.us">sistadm.us</a>
			You can download your admission letter from the website.</p>
		</p>
		<p>
			Pleas reach out to the college with physical copies of the documents you have uploaded during the admission process.
			Post verification, you will be able to start your classes.
		</p>
		</p>
		<p>Thank you for your participation.</p>
		<p>Best regards,</p>
		<p>SIST Admission Team</p>
	`)
}
