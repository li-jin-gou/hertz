/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package generator

import "path/filepath"

//-----------------------------------Default Layout-----------------------------------------

const (
	sp = string(filepath.Separator)

	defaultBizDir     = "biz"
	defaultModelDir   = "biz" + sp + "model"
	defaultHandlerDir = "biz" + sp + "handler"
	defaultServiceDir = "biz" + sp + "service"
	defaultDalDir     = "biz" + sp + "dal"
	defaultScriptDir  = "script"
	defaultConfDir    = "conf"
	defaultRouterDir  = "biz" + sp + "router"
	defaultClientDir  = "biz" + sp + "client"
)

const (
	routerGenIndex = 9
	routerIndex    = 10

	RegisterFile = "router_gen.go"
)

var defaultLayoutConfig = TemplateConfig{
	Layouts: []Template{
		{
			Path: defaultDalDir + sp,
		},
		{
			Path: defaultHandlerDir + sp,
		},
		{
			Path: defaultModelDir + sp,
		},
		{
			Path: defaultServiceDir + sp,
		},
		{
			Path: "main.go",
			Body: `// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default()

	register(h)
	h.Spin()
}
			`,
		},
		{
			Path:   "go.mod",
			Delims: [2]string{"{{", "}}"},
			Body: `module {{.GoModule}}
{{- if .UseApacheThrift}}
replace github.com/apache/thrift => github.com/apache/thrift v0.13.0
{{- end}}
			`,
		},
		{
			Path: ".gitignore",
			Body: `*.o
*.a
*.so
_obj
_test
*.[568vq]
[568vq].out
*.cgo1.go
*.cgo2.c
_cgo_defun.c
_cgo_gotypes.go
_cgo_export.*
_testmain.go
*.exe
*.exe~
*.test
*.prof
*.rar
*.zip
*.gz
*.psd
*.bmd
*.cfg
*.pptx
*.log
*nohup.out
*settings.pyc
*.sublime-project
*.sublime-workspace
!.gitkeep
.DS_Store
/.idea
/.vscode
/output
*.local.yml
dumped_hertz_remote_config.json
		  `,
		},
		{
			Path:   ".hz",
			Delims: [2]string{"{{", "}}"},
			Body: `// Code generated by hz. DO NOT EDIT.

hz version: {{.hzVersion}}`,
		},
		{
			Path: defaultHandlerDir + sp + "ping.go",
			Body: `// Code generated by hertz generator.

package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Ping .
func Ping(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, utils.H{
		"message": "pong",
	})
}
`,
		},
		{
			Path: RegisterFile,
			Body: `// Code generated by hertz generator. DO NOT EDIT.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	router "{{.RouterPkgPath}}"
)

// register registers all routers.
func register(r *server.Hertz) {

	router.GeneratedRegister(r)

	customizedRegister(r)
}
`,
		},
		{
			Path: "router.go",
			Body: `// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	handler "{{.HandlerPkgPath}}"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz){
	r.GET("/ping", handler.Ping)

	// your code ...
}
`,
		},
		{
			Path: defaultRouterDir + sp + registerTplName,
			Body: `// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz){
	` + insertPointNew + `
}
`,
		},
		{
			Path: "build.sh",
			Body: `
#!/bin/bash
RUN_NAME={{.ServiceName}}
mkdir -p output/bin
cp script/* output 2>/dev/null
chmod +x output/bootstrap.sh
go build -o output/bin/${RUN_NAME}
`,
		},
		{
			Path: defaultScriptDir + sp + "bootstrap.sh",
			Body: `
#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName={{.ServiceName}}
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}
`,
		},
	},
}
