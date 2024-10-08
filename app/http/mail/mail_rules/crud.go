package mail_rules

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/amis"
	"github.com/go-home-admin/go-admin/app/entity/admin"
)

func (c *CrudContext) Common() {
	c.SetDb(admin.NewOrmMailRules())
}

func (c *CrudContext) Table(curd *amis.Crud) {
	curd.Column("自增", "id")
	curd.Column("文本", "text")
	curd.Column("图片", "image").Image()
	curd.Column("日期", "date").Date()
	curd.Column("进度", "progress").Progress()
	curd.Column("状态", "status").Status()
	curd.Column("开关", "switch").Switch()
	curd.Column("映射", "mapping").Mapping(map[string]string{})
	curd.Column("List", "list").List()
}

func (c *CrudContext) Form(form *amis.Form) {
	form.Input("text", "文本")
	form.Input("image", "图片")

}

func (c *Controller) GinHandleCurd(ctx *gin.Context) {
	var crud = &CrudContext{}
	crud.CurdController.Context = ctx
	crud.CurdController.Crud = crud

	page := crud.GetPage()
	page.AddBodyByJsonStr(`
{
  "type": "tpl",
  "tpl": "什么是代收功能？\n代收功能是：当前域名所有前缀子邮箱邮件，都会被代收邮箱代收。无需一个个创建，相当于有无数的子邮箱帐号，邮箱前缀自己定义",
  "inline": true,
  "wrapperComponent": "",
  "id": "u:4ec6bbdb8b8d"
}`)
	amis.GinHandleCurd(ctx, crud)
}

type CrudContext struct {
	amis.CurdController
}
