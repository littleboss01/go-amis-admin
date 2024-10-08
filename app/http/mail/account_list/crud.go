package account_list

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/amis"
	"github.com/go-home-admin/go-admin/app/entity/admin"
	"os"
)

func (c *CrudContext) Common() {
	c.SetDb(admin.NewOrmHmAccounts())
	c.SetPrimary("accountid")
}

func (c *CrudContext) Table(curd *amis.Crud) {
	//form := &amis.Form{}

	btnImport := curd.AddCustomButton("导入")
	//form1 := amis.NewForm(c.Context)
	//c.Crud.Form(form1)
	//f := form1.SetApi(amis.GetUrl(c.Context, ""), "post")
	//btnImport.SetDialogForm(f)
	page := amis.NewPage("a")
	bytes, _ := os.ReadFile("./page/accountImport.json")
	page.AddBodyByJsonStr(string(bytes))
	btnImport.SetDialog(page)
	//btnImport.SetAjax("导入", amis.GetUrl(c.Context, "/import"))
	btnImport.Level = "success"
	//btnImport.ActionType = "dialog"

	btnExport := curd.AddCustomButton("导出")
	btnExport.Level = "success"
	page1 := amis.NewPage("导出")
	bytes, _ = os.ReadFile("page/accountExport.json")
	page1.AddBodyByJsonStr(string(bytes))
	btnExport.SetDialog(page1)

	curd.AutoGenerateFilter()
	curd.Column("id", "accountid")
	//curd.Column("accountdomainid", "accountdomainid")
	curd.Column("账号", "accountaddress").SearchableInput()
	curd.Column("密码", "accountpassword").SearchableInput()
	curd.Column("是否启用", "accountactive")
	//curd.Column("语言", "language")
	btnInmail := curd.Operation().AddButton("进入邮箱")
	btnInmail.Level = "success"
}

func (c *CrudContext) Form(form *amis.Form) {
	//form.Input("accountid", "id").SaveInt()
	//form.SetData(map[string]interface{}{
	//	"accountid": 2,
	//})
	//form.CreateBefore(func(form *amis.Form) {
	//	form.AddData("accountactive", 2)
	//})
	//form.AddData("accountid", int(2))

	id := form.Input("accountid", "id")
	id.Type = "hidden"
	id.SaveInt()

	accountdomainid := form.Input("accountdomainid", "accountdomainid")
	accountdomainid.Type = "hidden"
	accountdomainid.SaveInt()

	accountactive := form.Input("accountactive", "是否启用")
	accountactive.Type = "hidden"
	accountactive.SaveInt()

	form.Input("accountaddress", "账号")
	accountpassword := form.Input("accountpassword", "密码")
	accountpassword.Id = "accountpassword" //设置id是为了方便js脚本SetValue

	bytes, _ := os.ReadFile("./page/GenPassWord.json")
	form.AddBodyByJsonStr(string(bytes))

	form.AddCreatedAndUpdatedAt()

}

func (c *Controller) GinHandleCurd(ctx *gin.Context) {
	var crud = &CrudContext{}
	crud.CurdController.Context = ctx
	crud.CurdController.Crud = crud

	//可以在这里修改page的body增加其他amis组件
	page := crud.GetPage()
	page.Title = "aaaa"
	//page.Body =""
	//keyValueMap := map[string]interface{}{
	//	"title":   "页面标题",
	//	"content": "这是页面内容",
	//	"number":  42,
	//}
	//page.AddBody(keyValueMap)

	page.AddBodyByJsonStr(`[{
      "type": "tpl",
      "tpl": "初始页面",
      "wrapperComponent": "",
      "inline": false,
      "id": "u:42a3c9071c61"
    },
    {
      "type": "input-text",
      "label": "文本",
      "name": "text",
      "id": "u:124601609c22"
    },
    {
      "type": "input-text",
      "label": "文本",
      "name": "text",
      "id": "u:8713f0c158e5"
    },
    {
      "type": "select",
      "label": "选项",
      "name": "select",
      "options": [
        {
          "label": "选项A",
          "value": "A"
        },
        {
          "label": "选项B",
          "value": "B"
        }
      ],
      "id": "u:5dfaa4bffafc"
    }]`)
	//crud.GetPage().
	amis.GinHandleCurd(ctx, crud)
}

type CrudContext struct {
	amis.CurdController
}
