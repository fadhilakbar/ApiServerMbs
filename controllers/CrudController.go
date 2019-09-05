package controllers

import (
	"ApiServerMbs/database"
	"ApiServerMbs/functions"
	"ApiServerMbs/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//noinspection SqlDialectInspection,SqlNoDataSourceInspection
func InsertOutbox(context *gin.Context) {
	var params models.Outbox
	err := context.BindJSON(&params)
	if err != nil {
		context.JSON(http.StatusOK, functions.GetResponse("01", "BINDING JSON ERROR"))
		return
	}
	sql := `INSERT INTO outbox(receiver_number,message,status)VALUES(?,?,0)`
	stmt, err := database.ConnectDB().Prepare(sql)
	if err != nil {
		context.JSON(http.StatusOK, functions.GetResponse("01", "INSERT OUTBOX GAGAL"))
	} else {
		_, result := stmt.Exec(params.Rphone, params.Rmessage)
		if result != nil {
			context.JSON(http.StatusOK, functions.GetResponse("01", "INSERT OUTBOX GAGAL"))
		} else {
			context.JSON(http.StatusOK, functions.GetResponse("00", "INSERT OUTBOX SUKSES"))
		}
	}
}

func InsertOutboxMail(context *gin.Context) {
	var params models.OutboxMail
	err := context.BindJSON(&params)
	if err != nil {
		context.JSON(http.StatusOK, functions.GetResponse("01", "BINDING JSON ERROR"))
		return
	}
	sqlstatement := `INSERT INTO outbox_email(email,content,subject)VALUES(?,?,?)`
	stmt, err := database.ConnectDB().Prepare(sqlstatement)
	if err != nil {
		context.JSON(http.StatusOK, functions.GetResponse("01", "INSERT OUTBOX MAIL GAGAL"))
	} else {
		_, result := stmt.Exec(params.Remail, params.Rcontent, params.Rsubject)
		if result != nil {
			context.JSON(http.StatusOK, functions.GetResponse("01", "INSERT OUTBOX MAIL GAGAL"))
		} else {
			context.JSON(http.StatusOK, functions.GetResponse("00", "INSERT OUTBOX MAIL SUKSES"))
		}
	}
}
