package renderHandlers

import (
	"github.com/gin-gonic/gin"
)

type BaseHandler interface {
	List(ctx *gin.Context)
	Detail(ctx *gin.Context)
	Files(ctx *gin.Context)
	Blob(ctx *gin.Context)
	Commits(ctx *gin.Context)
	Commit(ctx *gin.Context)
}

type BaseHandlerImpl struct {
	resourceType string
	showTemplate string
}

func (b *BaseHandlerImpl) List(ctx *gin.Context) {
	renderTemplate(ctx, b.resourceType+"_index", nil)
}

func (b *BaseHandlerImpl) Detail(ctx *gin.Context) {
	b.renderShow(ctx, "show", "summary")
}

func (b *BaseHandlerImpl) Files(ctx *gin.Context) {
	b.renderShow(ctx, "files", "files")
}

func (b *BaseHandlerImpl) Blob(ctx *gin.Context) {
	b.renderShow(ctx, "blob", "files")
}

func (b *BaseHandlerImpl) Commits(ctx *gin.Context) {
	b.renderShow(ctx, "commits", "files")
}

func (b *BaseHandlerImpl) Commit(ctx *gin.Context) {
	commitId := ctx.Param("commit_id")
	b.renderShow(ctx, "commit", "files", map[string]interface{}{"commitId": commitId})
}

func (b *BaseHandlerImpl) renderShow(ctx *gin.Context, actionName, defaultTab string, extraData ...map[string]interface{}) {
	data := map[string]interface{}{
		"namespace":     ctx.Param("namespace"),
		"actionName":    actionName,
		"currentPath":   ctx.Param("path"),
		"currentBranch": ctx.Param("branch"),
		"defaultTab":    defaultTab,
	}

	if b.resourceType == "datasets" {
		data["datasetName"] = ctx.Param("dataset_name")
	} else if b.resourceType == "models" {
		data["modelName"] = ctx.Param("model_name")
	} else if b.resourceType == "spaces" {
		data["spaceName"] = ctx.Param("space_name")
	} else if b.resourceType == "codes" {
		data["codeName"] = ctx.Param("code_name")
	}

	for _, e := range extraData {
		for k, v := range e {
			data[k] = v
		}
	}
	renderTemplate(ctx, b.showTemplate, data)
}
