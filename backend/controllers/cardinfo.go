package controllers

import (
	"github.com/gin-gonic/gin"
)

/**
 * Procura na lista de faturas do utilizador se existe uma fatura com o id enviado por parâmetro
 * nos parâmetros tem que contar o ID do user e o ID da fatura
**/
func GetIdentityCardInfo(c *gin.Context) {

	/*IdentitycardinfoId, err := strconv.ParseUint(c.Param("invoiceID"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid Invoice ID!"})
		return
	}

	uintID := uint(IdentitycardinfoId)

	// Vai buscar a lista de invoices e verifica se é nula ou não

	var listInvoices = GetAll(c)

	if listInvoices == nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Lista de faturas vazia!"})
		return
	}

	// Procura se o ID da fatura existe na lista de faturas do utiliador
	for _, invoice := range listInvoices {

		if invoice.ID == uintID {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": invoice})
			return
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Não existe na sua lista de faturas!"})
	return*/

}

/**
* Adiciona à tabela das faturas uma nova fatura com todos os parâmetros da mesma
**/
func AddIdentityCardInfo(c *gin.Context) {

	/*var addr = flag.String("addr", ":8080", "http server address")

	flag.Parse()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(w, r)
	})

	log.Fatal(http.ListenAndServe(*addr, nil))
	*/
	/*var invoice model.Invoice

	fmt.Println(c)

	if err := c.ShouldBindJSON(&invoice); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Check syntax!"})
		return
	}

	text, err := services.ProcessImage(invoice.Image)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Cannot process image!"})
		return
	}

	invoice.Info = text

	services.Db.Save(&invoice)

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Create successful!", "resourceId": invoice.ID})*/
}

/**
 * Procura na lista de faturas se existe uma fatura com o id enviado por parâmetro e elinina-a da tabela de faturas
**/
func DeleteIdentityCardInfo(c *gin.Context) {
	/*var invoice model.Invoice

	id := c.Param("id")

	services.Db.Where("id = ?", id).Find(&invoice, id)

	if invoice.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}

	services.Db.Delete(&invoice)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete succeeded!"})*/
}


func GetAllIdentityCardInfo(c *gin.Context){

}