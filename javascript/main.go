// Lute - 一款对中文语境优化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under the Mulan PSL v1.
// You can use this software according to the terms and conditions of the Mulan PSL v1.
// You may obtain a copy of Mulan PSL v1 at:
//     http://license.coscl.org.cn/MulanPSL
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR
// PURPOSE.
// See the Mulan PSL v1 for more details.

package main

import (
	"github.com/88250/lute"
	"github.com/gopherjs/gopherjs/js"
)

func New(options map[string]map[string]*js.Object) *js.Object {
	engine := lute.New()
	registerRenderers(engine, options)
	return js.MakeWrapper(engine)
}

func registerRenderers(engine *lute.Lute, options map[string]map[string]*js.Object) {
	for rendererType, extRenderer := range options["renderers"] {
		switch extRenderer.Interface().(type) {
		case map[string]interface{}:
			break
		default:
			continue
		}

		var rendererFuncs map[lute.NodeType]lute.RendererFunc
		if "format" == rendererType {
			rendererFuncs = engine.FormatRendererFuncs
		} else if "vditor" == rendererType {
			rendererFuncs = engine.VditorRendererFuncs
		} else if "html" == rendererType {
			rendererFuncs = engine.HTMLRendererFuncs
		} else {
			continue
		}

		renderFuncs := extRenderer.Interface().(map[string]interface{})
		for funcName, _ := range renderFuncs {
			nodeType := "Node" + funcName[len("render"):]
			rendererFuncs[lute.Str2NodeType(nodeType)] = func(n *lute.Node, entering bool) (status lute.WalkStatus, err error) {
				nodeType := n.Typ.String()
				funcName = "render" + nodeType[len("Node"):]
				walkStatus := extRenderer.Call(funcName, js.MakeWrapper(n), entering).Int()
				return lute.WalkStatus(walkStatus), nil
			}
		}
	}
}

func main() {
	js.Global.Set("Lute", map[string]interface{}{
		"Version":          lute.Version,
		"New":              New,
		"WalkStop":         lute.WalkStop,
		"WalkSkipChildren": lute.WalkSkipChildren,
		"WalkContinue":     lute.WalkContinue,
	})
}
