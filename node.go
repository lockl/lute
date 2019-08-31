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

// BaseNode 描述了节点基础结构。
type BaseNode struct {
	typ             int       // 节点类型
	parent          *BaseNode // 父节点
	previous        *BaseNode // 前一个兄弟节点
	next            *BaseNode // 后一个兄弟节点
	firstChild      *BaseNode // 第一个子节点
	lastChild       *BaseNode // 最后一个子节点
	rawText         string    // 原始内容
	tokens          items     // 词法分析结果 tokens，语法分析阶段会继续操作这些 tokens
	close           bool      // 标识是否关闭
	lastLineBlank   bool      // 标识最后一行是否是空行
	lastLineChecked bool      // 标识最后一行是否检查过

	// 代码块

	isFenced    bool
	fenceChar   byte
	fenceLength int
	fenceOffset int
	info        items

	// HTML 块

	hType int // 规范中定义的 HTML 块类型（1-7）

	// 列表、列表项

	*listData

	// 链接、图片

	Destination items // 链接地址
	Title       items // 链接标题

	// 任务列表项 [ ]、[x] 或者 [X]

	checked bool // 是否勾选

	// 表

	Aligns []int // 从左到右每个表格节点的对齐方式，0：默认对齐，1：左对齐，2：居中对齐，3：右对齐
	Align  int   // 表的单元格对齐方式

	// 标题

	Level int // 1~6
}

// Type 返回节点类型。
func (n *BaseNode) Type() int {
	return n.typ
}

// IsOpen 返回节点是否是打开的。
func (n *BaseNode) IsOpen() bool {
	return !n.close
}

// IsClosed 返回节点是否是关闭的。
func (n *BaseNode) IsClosed() bool {
	return n.close
}

// Close 关闭节点。
func (n *BaseNode) Close() {
	n.close = true
}

// Finalize 节点最终化处理。比如围栏代码块提取 info 部分；HTML 代码块剔除结尾空格；段落需要解析链接引用定义等。
func (n *BaseNode) Finalize(context *Context) {
	switch n.typ {
	case NodeCodeBlock:
		n.CodeBlockFinalize(context)
	case NodeHTMLBlock:
		n.HTMLBlockFinalize(context)
	case NodeParagraph:
		n.ParagraphFinalize(context)
	case NodeList:
		n.ListFinalize(context)
	}
}

// Continue 判断节点是否可以继续处理，比如块引用需要 >，缩进代码块需要 4 空格，围栏代码块需要 ```。
// 如果可以继续处理返回 0，如果不能接续处理返回 1，如果返回 2（仅在围栏代码块闭合时）则说明可以继续下一行处理了。
func (n *BaseNode) Continue(context *Context) int {
	switch n.typ {
	case NodeCodeBlock:
		return n.CodeBlockContinue(context)
	case NodeHTMLBlock:
		return n.HTMLBlockContinue(context)
	case NodeParagraph:
		return n.ParagraphContinue(context)
	case NodeListItem:
		return n.ListItemContinue(context)
	case NodeBlockquote:
		return n.BlockquoteContinue(context)
	case NodeHeading, NodeThematicBreak:
		return 1
	}

	return 0
}

// AcceptLines 判断是否节点是否可以接受更多的文本行。比如 HTML 块、代码块和段落是可以接受更多的文本行的。
func (n *BaseNode) AcceptLines() bool {
	switch n.typ {
	case NodeParagraph, NodeCodeBlock, NodeHTMLBlock:
		return true
	}
	return false
}

// CanContain 判断是否能够包含 NodeType 指定类型的节点。 比如列表节点（一种块级容器）只能包含列表项节点，
// 块引用节点（另一种块级容器）可以包含任意节点；段落节点（一种叶子块节点）不能包含任何其他块级节点。
func (n *BaseNode) CanContain(nodeType int) bool {
	switch n.typ {
	case NodeCodeBlock, NodeHTMLBlock, NodeParagraph, NodeHeading, NodeThematicBreak:
		return false
	case NodeList:
		return NodeListItem == nodeType
	}

	return NodeListItem != nodeType
}

// LastLineBlank 判断节点最后一行是否是空行。
func (n *BaseNode) LastLineBlank() bool {
	return n.lastLineBlank
}

// SetLastLineBlank 设置节点最后一行是否是空行。
func (n *BaseNode) SetLastLineBlank(lastLineBlank bool) {
	n.lastLineBlank = lastLineBlank
}

// LastLineChecked 返回最后一行是否检查过。在判断列表是紧凑或松散模式时作为标识用。
func (n *BaseNode) LastLineChecked() bool {
	return n.lastLineChecked
}

// SetLastLineChecked 设置最后一行是否检查过。
func (n *BaseNode) SetLastLineChecked(lastLineChecked bool) {
	n.lastLineChecked = lastLineChecked
}

// Unlink 用于将节点从树上移除，后一个兄弟节点会接替该节点。
func (n *BaseNode) Unlink() {
	if nil != n.previous {
		n.previous.SetNext(n.next)
	} else if nil != n.parent {
		n.parent.SetFirstChild(n.next)
	}
	if nil != n.next {
		n.next.SetPrevious(n.previous)
	} else if nil != n.parent {
		n.parent.SetLastChild(n.previous)
	}
	n.parent = nil
	n.next = nil
	n.previous = nil
}

// Parent 返回父节点。
func (n *BaseNode) Parent() *BaseNode {
	return n.parent
}

// SetParent 设置父节点。
func (n *BaseNode) SetParent(parent *BaseNode) {
	n.parent = parent
}

// Next 返回后一个兄弟节点。
func (n *BaseNode) Next() *BaseNode {
	return n.next
}

// SetNext 设置后一个兄弟节点。
func (n *BaseNode) SetNext(next *BaseNode) {
	n.next = next
}

// Previous 返回前一个兄弟节点。
func (n *BaseNode) Previous() *BaseNode {
	return n.previous
}

// SetPrevious 设置前一个兄弟节点。
func (n *BaseNode) SetPrevious(previous *BaseNode) {
	n.previous = previous
}

// FirstChild 返回第一个子节点。
func (n *BaseNode) FirstChild() *BaseNode {
	return n.firstChild
}

// SetFirstChild 设置第一个子节点。
func (n *BaseNode) SetFirstChild(firstChild *BaseNode) {
	n.firstChild = firstChild
}

// 返回最后一个子节点。
func (n *BaseNode) LastChild() *BaseNode {
	return n.lastChild
}

// 设置最后一个子节点。
func (n *BaseNode) SetLastChild(lastChild *BaseNode) {
	n.lastChild = lastChild
}

// RawText 返回原始内容。
func (n *BaseNode) RawText() string {
	return n.rawText
}

// SetRawText 设置原始内容。
func (n *BaseNode) SetRawText(rawText string) {
	n.rawText = rawText
}

// AppendRawText 添加原始内容。
func (n *BaseNode) AppendRawText(rawText string) {
	n.rawText += rawText
}

// Tokens 返回所有 tokens。
func (n *BaseNode) Tokens() items {
	return n.tokens
}

// SetTokens 设置 tokens。
func (n *BaseNode) SetTokens(tokens items) {
	n.tokens = tokens
}

// AppendTokens 添加 tokens。
func (n *BaseNode) AppendTokens(tokens items) {
	n.tokens = append(n.tokens, tokens...)
}

// InsertAfter 在当前节点后插入一个兄弟节点。
func (n *BaseNode) InsertAfter(this *BaseNode, sibling *BaseNode) {
	sibling.Unlink()
	sibling.SetNext(n.next)
	if nil != sibling.Next() {
		sibling.Next().SetPrevious(sibling)
	}
	sibling.SetPrevious(this)
	n.next = sibling
	sibling.SetParent(n.parent)
	if nil == sibling.Next() {
		sibling.Parent().SetLastChild(sibling)
	}
}

// InsertBefore 在当前节点前插入一个兄弟节点。
func (n *BaseNode) InsertBefore(this *BaseNode, sibling *BaseNode) {
	sibling.Unlink()
	sibling.SetPrevious(n.previous)
	if nil != sibling.Previous() {
		sibling.Previous().SetNext(sibling)
	}
	sibling.SetNext(this)
	n.previous = sibling
	sibling.SetParent(n.parent)
	if nil == sibling.Previous() {
		sibling.Parent().SetFirstChild(sibling)
	}
}

// AppendChild 添加一个子节点。
func (n *BaseNode) AppendChild(this, child *BaseNode) {
	child.Unlink()
	child.SetParent(this)
	if nil != n.lastChild {
		n.lastChild.SetNext(child)
		child.SetPrevious(n.lastChild)
		n.lastChild = child
	} else {
		n.firstChild = child
		n.lastChild = child
	}
}

const (
	// CommonMark

	NodeDocument      = iota // 根节点类
	NodeParagraph            // 段落节点
	NodeHeading              // 标题节点
	NodeThematicBreak        // 分隔线节点
	NodeBlockquote           // 块引用节点
	NodeList                 // 列表节点
	NodeListItem             // 列表项节点
	NodeHTMLBlock            // HTML 块节点
	NodeInlineHTML           // 内联 HTML节点
	NodeCodeBlock            // 代码块节点
	NodeText                 // 文本节点
	NodeEmphasis             // 强调节点
	NodeStrong               // 加粗节点
	NodeCodeSpan             // 代码节点
	NodeHardBreak            // 硬换行节点
	NodeSoftBreak            // 软换行节点
	NodeLink                 // 链接节点
	NodeImage                // 图片节点

	// GFM

	NodeTaskListItemMarker // 任务列表项标记节点
	NodeStrikethrough      // 删除线节点
	NodeTable              // 表节点
	NodeTableHead          // 表头节点
	NodeTableRow           // 表行节点
	NodeTableCell          // 表格节点
)
