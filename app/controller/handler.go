package controller


// AddAntrianHandler is a function to get all queue
func AddAntrianHandler(c *gin.Context){
	flag, err := model.AddAntrian()

	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "succes",
		})
	} else{
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error": err
		})
	}
}

//get all queue
func GetAntrianHandler(c *gin.Context){
	
	flag, data, err := model.GetAntrian()

	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "succes",
		})
	}else{
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error": err
		})
}


//function to update a queue
func UpdateAntrianHandler(c *gin.Context){
	idAntrian := c.Param("idAntrian")
	flag, data, err := model.UpdateAntrian()

	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "succes",
		})
	}else{
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error": err
		})
}

//delete a queue
func DeleteAntrianHandler(c *gin.Context){
	idAntrian := c.Param("idAntrian")
	flag, data, err := model.DeleteAntrian()

	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "succes",
		})
	}else{
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error": err
		})
}

func PageAntrianHandler(c *gin.Context){
	flag, result, err := model.GetAntrian()
	var currentAntrian map[string]interface{}

	for _, item := range result {
		if item != nil {
			currentAntrian = item
			break
		}
	}

	if flag && len(result)>0{
		c.HTML(http.StatusOK, "index.html", gin.H{
			"antrian": currentAntrian["id"]
		})
	}else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error": err,
		})
	}
}