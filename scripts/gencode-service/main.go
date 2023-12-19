package main

import (
	"fmt"
	"github.com/gertd/go-pluralize"
	"lark/pkg/conf"
	"lark/scripts/gencode-service/template"
	"lark/scripts/gencode/config"
	"lark/scripts/gencode/utils"
	"lark/scripts/generate-gorm/gengorm"
	"strings"
)

func main() {
	var (
		cli         = pluralize.NewClient()
		packageName = "Phoenix" // 程序包名
		serviceName = "Auth"    // 服务实现
		apiName     = "SignUp"  // 服务下属 Api
		tableName   = "user_infos"
		modelName   = utils.ToCamelCase(cli.Singular(tableName)) // 表模型

		upperPackageName = utils.ToCamel(packageName)
		lowerPackageName = utils.FirstLower(upperPackageName)
		upperServiceName = utils.ToCamel(serviceName)
		lowerServiceName = utils.FirstLower(upperServiceName)
		upperApiName     = utils.ToCamel(apiName)
		lowerApiName     = utils.FirstLower(upperApiName)
		upperModelName   = utils.ToCamel(modelName)
		lowerModelName   = utils.FirstLower(upperModelName)

		cfg = &conf.Mysql{
			Address:      "127.0.0.1",
			Username:     "root",
			Password:     "",
			Db:           "canary",
			MaxOpenConns: 20,
			MaxIdleConns: 10,
			MaxLifetime:  120000,
			MaxIdleTime:  0,
			Charset:      "utf8mb4",
			Debug:        false,
		}
	)
	conf := config.GenConfig{
		Path:        "",
		Prefix:      "",
		PackageName: utils.CamelToSnake(upperPackageName),
		ServiceName: utils.CamelToSnake(upperServiceName),
		ApiName:     utils.CamelToSnake(upperApiName),
		ModelName:   utils.CamelToSnake(upperModelName),
		Dict: map[string]interface{}{
			"UpperPackageName":    upperPackageName,
			"LowerPackageName":    lowerPackageName,
			"UpperServiceName":    upperServiceName,
			"LowerServiceName":    lowerServiceName,
			"UpperApiName":        upperApiName,
			"LowerApiName":        lowerApiName,
			"UpperModelName":      upperModelName,
			"LowerModelName":      lowerModelName,
			"PackageName":         utils.CamelToSnake(upperPackageName),
			"ServiceName":         utils.CamelToSnake(upperServiceName),
			"ApiName":             utils.CamelToSnake(upperApiName),
			"ModelName":           utils.CamelToSnake(upperModelName),
			"AllUpperPackageName": strings.ToUpper(upperPackageName),
			"AllUpperServiceName": strings.ToUpper(upperServiceName),
			"AllUpperModelName":   strings.ToUpper(upperModelName),
		},
	}
	if conf.PackageName == "" {
		return
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
	//generateDomainPoCode(conf)
	generateDomainRepoCode(conf)
	gengorm.GenGorm(cfg, cli.Plural(tableName), "./domain/po/")

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

	if conf.ModelName != "" {
		conf.Path = fmt.Sprintf("./apps/apis/%s/dig", conf.PackageName)
		conf.ResetPrefSuf()
		conf.Filename = "dig_domain_" + conf.ModelName
		utils.GenCode(template.AppsDigModelTemplate, &conf)
	}

	if conf.ServiceName != "" {
		conf.Path = fmt.Sprintf("./apps/apis/%s/dig", conf.PackageName)
		conf.ResetPrefSuf()
		conf.Filename = "dig_service_" + conf.ServiceName
		utils.GenCode(template.DigServiceTemplate, &conf)
	}
}

// Domain
func generateDomainCacheCode(conf config.GenConfig) {
	if conf.ModelName == "" {
		return
	}
	conf.Path = "./domain/cache"
	conf.Filename = "cache_" + conf.ModelName
	utils.GenCode(template.DomainCacheTemplate, &conf)
}

func generateDomainCrConstCode(conf config.GenConfig) {
	if conf.ModelName == "" {
		return
	}
	conf.Path = fmt.Sprintf("./domain/cr/cr_%s", conf.ModelName)
	conf.Filename = "cr_" + conf.ModelName + "_const"
	utils.GenCode(template.DomainCrConstTemplate, &conf)
}

func generateDomainCrReadCode(conf config.GenConfig) {
	if conf.ModelName == "" {
		return
	}
	conf.Path = fmt.Sprintf("./domain/cr/cr_%s", conf.ModelName)
	conf.Filename = "cr_" + conf.ModelName
	utils.GenCode(template.DomainCrReadTemplate, &conf)
}

func generateDomainPoCode(conf config.GenConfig) {
	if conf.ModelName == "" {
		return
	}
	conf.Path = "./domain/po"
	conf.Filename = "po_" + conf.ModelName
	utils.GenCode(template.DomainPoTemplate, &conf)
}

func generateDomainRepoCode(conf config.GenConfig) {
	if conf.ModelName == "" {
		return
	}
	conf.Path = "./domain/repo"
	conf.Filename = "repo_" + conf.ModelName
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
	conf.Path = fmt.Sprintf("./apps/apis/%s/internal/ctrl/ctrl_%s", conf.PackageName, conf.ServiceName)
	conf.Filename = "ctrl_" + conf.ServiceName
	utils.GenCode(template.InternalCtrlTemplate, &conf)
	if conf.ApiName != "" {
		conf.Path = fmt.Sprintf("./apps/apis/%s/internal/ctrl/ctrl_%s", conf.PackageName, conf.ServiceName)
		conf.ResetPrefSuf()
		conf.Filename = "ctrl_" + utils.GetName(conf.ServiceName, conf.ApiName)
		utils.GenCode(template.InternalCtrlApiTemplate, &conf)
	}
}

func generateInternalDtoCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./apps/apis/%s/internal/dto/dto_%s", conf.PackageName, conf.ServiceName)
	conf.Filename = "dto_" + conf.ServiceName
	utils.GenCode(template.InternalDtoTemplate, &conf)

	if conf.ApiName != "" {
		conf.Path = fmt.Sprintf("./apps/apis/%s/internal/dto/dto_%s", conf.PackageName, conf.ServiceName)
		conf.ResetPrefSuf()
		conf.Filename = "dto_" + utils.GetName(conf.ServiceName, conf.ApiName)
		utils.GenCode(template.InterfacesDtoApiTemplate, &conf)
	}
}

func generateInternalRouterCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./apps/apis/%s/internal/router", conf.PackageName)
	conf.Filename = "router_api"
	utils.GenCode(template.InternalRouterTemplate, &conf)

	if conf.ApiName != "" {
		conf.Path = fmt.Sprintf("./apps/apis/%s/internal/router", conf.PackageName)
		conf.ResetPrefSuf()
		conf.Filename = "router_" + conf.ServiceName
		utils.GenCode(template.InternalRouterServiceTemplate, &conf)
	}
}

func generateInternalServerCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./apps/apis/%s/internal/server", conf.PackageName)
	conf.Filename = "server"
	utils.GenCode(template.InternalServerTemplate, &conf)
}

func generateInternalServiceCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./apps/apis/%s/internal/service/svc_%s", conf.PackageName, conf.ServiceName)
	conf.Filename = "svc_" + conf.ServiceName
	utils.GenCode(template.InternalServiceTemplate, &conf)
	if conf.ServiceName != "" {
		conf.Path = fmt.Sprintf("./apps/apis/%s/internal/service/svc_%s", conf.PackageName, conf.ServiceName)
		conf.ResetPrefSuf()
		conf.Filename = "svc_" + utils.GetName(conf.ServiceName, conf.ApiName)
		utils.GenCode(template.InternalServiceApiTemplate, &conf)
	}
}

func generateInternalServiceConstCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./apps/apis/%s/internal/service/svc_%s", conf.PackageName, conf.ServiceName)
	conf.Filename = "svc_" + conf.ServiceName + "_const"
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
