#API接口说明
    · 1. 接口成功返回code为1
    · 2. 接口错误的返回code为0
    · 3. 接口函数统一命名为func XXXXXXHandler(c *gin.Context, commArgs *models.WebReqCommArgs, values url.Values)格式
    · 4. 接口返回json统一格式： c:一级错误码，m:消息， d：数据
            {
                "c":1,
                "d":{
                    "code":0,
                    "money":600,
                    "msg":"成功"
                },
                "m":"queryDownScoreHandler"
            }