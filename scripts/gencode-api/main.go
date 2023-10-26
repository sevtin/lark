package main

import (
	"fmt"
	"lark/scripts/gencode-api/template"
	"lark/scripts/gencode/config"
	"lark/scripts/gencode/utils"
	"strings"
)

func main() {
	var (
		serviceName      = "Canary"
		upperServiceName = utils.ToCamel(serviceName)
		lowerServiceName = utils.FirstLower(upperServiceName)
		packageName      = utils.CamelToSnake(upperServiceName)
	)
	conf := config.GenConfig{
		Path:        "",
		Prefix:      "",
		PackageName: packageName,
		Dict: map[string]interface{}{
			"UpperServiceName": upperServiceName,
			"LowerServiceName": lowerServiceName,
			"PackageName":      packageName,
		},
	}

	// Cmd
	generateCmdTemplateCode(conf)

	// Configs
	generateConfigsYamlCode(conf)

	// Dig
	generateDigTemplateCode(conf)

	// Domain
	generateDomainCacheCode(conf)
	generateDomainCrConstCode(conf)
	generateDomainCrReadCode(conf)
	generateDomainPoCode(conf)
	generateDomainRepoCode(conf)

	// Internal
	generateInternalConfigCode(conf)
	generateInternalCtrlCode(conf)
	generateInternalDtoCode(conf)
	generateInternalRouterCode(conf)
	generateInternalServerCode(conf)
	generateInternalServiceCode(conf)
	generateInternalServiceConstCode(conf)

	// Pkg
	//generatePkgProtoCode(conf)
	//generatePkgProtoGoCode(conf)
	//generatePkgProtoRespCode(conf)
}

// Cmd
func generateCmdTemplateCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./apps/apis/%s/cmd", conf.PackageName)
	conf.Prefix = "main_"
	utils.GenCode(template.CmdTemplate, &conf)
}

// Configs
func generateConfigsYamlCode(conf config.GenConfig) {
	conf.Path = "./configs"
	conf.FileType = config.FILE_TYPE_YAML
	conf.Prefix = "api_"
	utils.GenCode(template.ConfigsYamlTemplate, &conf)
}

// Dig
func generateDigTemplateCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./apps/apis/%s/dig", conf.PackageName)
	conf.Filename = "dig"
	utils.GenCode(template.DigTemplate, &conf)
}

// Domain
func generateDomainCacheCode(conf config.GenConfig) {
	conf.Path = "./domain/cache"
	conf.Prefix = "cache_"
	utils.GenCode(template.DomainCacheTemplate, &conf)
}

func generateDomainCrConstCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./domain/cr/cr_%s", conf.PackageName)
	conf.Prefix = "cr_"
	conf.Suffix = "_const"
	utils.GenCode(template.DomainCrConstTemplate, &conf)
}

func generateDomainCrReadCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./domain/cr/cr_%s", conf.PackageName)
	conf.Prefix = "cr_"
	utils.GenCode(template.DomainCrReadTemplate, &conf)
}

func generateDomainPoCode(conf config.GenConfig) {
	conf.Path = "./domain/po"
	conf.Prefix = "po_"
	utils.GenCode(template.DomainPoTemplate, &conf)
}

func generateDomainRepoCode(conf config.GenConfig) {
	conf.Path = "./domain/repo"
	conf.Prefix = "repo_"
	utils.GenCode(template.DomainRepoTemplate, &conf)
}

// Internal
func generateInternalConfigCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./apps/apis/%s/internal/config", conf.PackageName)
	code := strings.ReplaceAll(template.InternalConfigCode, "\"yaml:", "`yaml:")
	code = strings.ReplaceAll(code, "\"\"", "\"`")
	conf.Filename = "config"
	utils.GenCode(template.ParseTemplate(code), &conf)
}

func generateInternalCtrlCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./apps/apis/%s/internal/ctrl", conf.PackageName)
	conf.Prefix = "ctrl_"
	utils.GenCode(template.InternalCtrlTemplate, &conf)
}

func generateInternalDtoCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./apps/apis/%s/internal/dto", conf.PackageName)
	conf.Prefix = "dto_"
	utils.GenCode(template.InternalDtoTemplate, &conf)
}

func generateInternalRouterCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./apps/apis/%s/internal/router", conf.PackageName)
	conf.Filename = "router_api"
	utils.GenCode(template.InternalRouterTemplate, &conf)
}

func generateInternalServerCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./apps/apis/%s/internal/server", conf.PackageName)
	conf.Filename = "server"
	utils.GenCode(template.InternalServerTemplate, &conf)
}

func generateInternalServiceCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./apps/apis/%s/internal/service", conf.PackageName)
	conf.Prefix = "svc_"
	utils.GenCode(template.InternalServiceTemplate, &conf)
}

func generateInternalServiceConstCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./apps/apis/%s/internal/service", conf.PackageName)
	conf.Prefix = "svc_"
	conf.Suffix = "_const"
	conf.Dict["AllUpperServiceName"] = strings.ToUpper(conf.PackageName)
	utils.GenCode(template.InternalServiceConstTemplate, &conf)
}

// Pkg
func generatePkgProtoCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./pkg/proto/pb_%s", conf.PackageName)
	conf.FileType = config.FILE_TYPE_PROTO
	utils.GenCode(template.PkgProtoTemplate, &conf)
}

func generatePkgProtoGoCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./pkg/proto/pb_%s", conf.PackageName)
	conf.Suffix = ".pb"
	utils.GenCode(template.PkgProtoGoTemplate, &conf)
}

func generatePkgProtoRespCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./pkg/proto/pb_%s", conf.PackageName)
	conf.PackageName = "resp"
	utils.GenCode(template.PkgProtoRespTemplate, &conf)
}
