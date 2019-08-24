// Lute - A structured markdown engine.
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

package lute

// Image 描述了图片节点结构。
type Image struct {
	*BaseNode
	Destination string // 图片链接地址
	Title       string // 图片标题
}

// parseBang 解析 !，可能是图片标记开始 ![ 也可能是普通文本 !。
func (t *Tree) parseBang(tokens items) (ret Node) {
	var startPos = t.context.pos
	t.context.pos++
	if t.context.pos < len(tokens) && itemOpenBracket == tokens[t.context.pos] {
		t.context.pos++
		ret = &Text{tokens: toItems("![")}
		// 将图片开始标记入栈
		t.addBracket(ret, startPos+2, true)
		return
	}

	ret = &Text{tokens: toItems("!")}
	return
}
