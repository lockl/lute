// Lute - 一款对中文语境优化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package test

import (
	"strings"
	"testing"

	"github.com/88250/lute"
)

var spinVditorIRBlockDOMTests = []*parseTest{

	{"27", "<div data-block=\"0\" data-node-id=\"20200825235504-ect683e\" data-type=\"code-block\" class=\"vditor-ir__node\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"\"><use xlink:href=\"#vditor-icon-code\"></use></svg></span><span data-type=\"code-block-open-marker\">```</span><span class=\"vditor-ir__marker vditor-ir__marker--info\" data-type=\"code-block-info\">​</span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code>f<wbr></code></pre><pre class=\"vditor-ir__preview\" data-render=\"1\"><div class=\"vditor-copy\"><textarea></textarea><span aria-label=\"复制\" onmouseover=\"this.setAttribute('aria-label', '复制')\" class=\"vditor-tooltipped vditor-tooltipped__w\" onclick=\"this.previousElementSibling.select();document.execCommand('copy');this.setAttribute('aria-label', '已复制')\"><svg><use xlink:href=\"#vditor-icon-copy\"></use></svg></span></div><code class=\"hljs\"></code></pre><span data-type=\"code-block-close-marker\">```</span></div>", "<div data-block=\"0\" data-node-id=\"20200825235504-ect683e\" data-type=\"code-block\" class=\"vditor-ir__node vditor-ir__node--expand\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#vditor-icon-code\"></use></svg></span>\u200b<span data-type=\"code-block-open-marker\">```</span><span class=\"vditor-ir__marker vditor-ir__marker--info\" data-type=\"code-block-info\">\u200b</span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code>f<wbr>\n</code></pre><pre class=\"vditor-ir__preview\" data-render=\"2\"><code>f\n</code></pre><span data-type=\"code-block-close-marker\">```</span></div>"},
	{"26", "<div data-block=\"0\" data-node-id=\"20200825234341-mq1uocf\" data-type=\"math-block\" class=\"vditor-ir__node\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"\"><use xlink:href=\"#iconMath\n\"></use></svg></span><span data-type=\"math-block-open-marker\">$$</span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code data-type=\"math-block\" class=\"language-math\">f<wbr></code></pre><pre class=\"vditor-ir__preview\" data-render=\"1\"><code class=\"language-math\"><div class=\"vditor-math\" data-math=\"\"><span class=\"katex-display\"><span class=\"katex\"><span class=\"katex-html\" aria-hidden=\"true\"></span></span></span></div></code></pre><span data-type=\"math-block-close-marker\">$$</span></div>", "<div data-block=\"0\" data-node-id=\"20200825234341-mq1uocf\" data-type=\"math-block\" class=\"vditor-ir__node vditor-ir__node--expand\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#iconMath\n\"></use></svg></span>\u200b<span data-type=\"math-block-open-marker\">$$</span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code data-type=\"math-block\" class=\"language-math\">f<wbr></code></pre><pre class=\"vditor-ir__preview\" data-render=\"2\"><code data-type=\"math-block\" class=\"language-math\">f</code></pre><span data-type=\"math-block-close-marker\">$$</span></div>"},
	{"25", "<p data-block=\"0\" data-node-id=\"20200815231623-j492vp3\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#iconParagraph\"></use></svg></span><span data-type=\"inline-math\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker\">$</span><code data-newline=\"1\" class=\"vditor-ir__marker vditor-ir__marker--pre\" data-type=\"math-inline\">foo</code><span class=\"vditor-ir__preview\" data-render=\"2\"><code class=\"language-math\">foo</code></span><span class=\"vditor-ir__marker\">$</span></span></p>", "<p data-block=\"0\" data-node-id=\"20200815231623-j492vp3\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#iconParagraph\"></use></svg></span>\u200b<span data-type=\"inline-math\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker\">$</span><code data-newline=\"1\" class=\"vditor-ir__marker vditor-ir__marker--pre\" data-type=\"math-inline\">foo</code><span class=\"vditor-ir__preview\" data-render=\"2\"><code class=\"language-math\">foo</code></span><span class=\"vditor-ir__marker\">$</span></span></p>"},
	{"24", "<p data-block=\"0\" data-node-id=\"20200825174808-w1hxf9h\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#iconParagraph\"></use></svg></span>​</p>| col1<wbr> | col2 | col3 |\n| --- | --- | --- |\n|  |  |  |\n|  |  |  |", "<table data-block=\"0\" data-type=\"table\" data-node-id=\"20200825174808-w1hxf9h\"><thead><tr><th>col1<wbr></th><th>col2</th><th>col3</th></tr></thead><tbody><tr><td> </td><td> </td><td> </td></tr><tr><td> </td><td> </td><td> </td></tr></tbody></table>"},
	{"23", "<h1 data-block=\"0\" class=\"vditor-ir__node\" data-node-id=\"20200824174550-e8couax\" id=\"ir-f\" data-marker=\"#\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#vditor-icon-headings\"></use></svg></span><span class=\"vditor-ir__marker vditor-ir__marker--heading\" data-type=\"heading-marker\"># </span>foo<wbr></h1>", "<h1 data-block=\"0\" class=\"vditor-ir__node vditor-ir__node--expand\" data-node-id=\"20200824174550-e8couax\" id=\"ir-foo\" data-marker=\"#\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#iconH1\"></use></svg></span><span class=\"vditor-ir__marker vditor-ir__marker--heading\" data-type=\"heading-marker\"># </span>foo<wbr></h1>"},
	{"22", "<p data-block=\"0\" data-node-id=\"20200821122908-xte3idk\">##<wbr><span>foo</span></p>", "<p data-block=\"0\" data-node-id=\"20200821122908-xte3idk\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#iconParagraph\"></use></svg></span>##<wbr>foo</p>"},
	{"21", "<p data-block=\"0\" data-node-id=\"20200821090811-hrewhtz\">​<span data-type=\"emoji\" class=\"vditor-ir__node\"><span data-render=\"2\">😄</span><span class=\"vditor-ir__marker\">:smile:foo<wbr></span></span></p>", "<p data-block=\"0\" data-node-id=\"20200821090811-hrewhtz\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#iconParagraph\"></use></svg></span>\u200b<span data-type=\"emoji\" class=\"vditor-ir__node\"><span data-render=\"2\">😄</span><span class=\"vditor-ir__marker\">:smile:</span></span>foo<wbr></p>"},
	{"20", "<p data-block=\"0\" data-node-id=\"20200817164045-vtlwy31\"><span data-type=\"html-entity\" class=\"vditor-ir__node\"><code data-newline=\"1\" class=\"vditor-ir__marker vditor-ir__marker--pre\" data-type=\"html-entity\">&amp;emsp;</code><span class=\"vditor-ir__preview\" data-render=\"2\"></span></span></p>", "<p data-block=\"0\" data-node-id=\"20200817164045-vtlwy31\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#iconParagraph\"></use></svg></span>\u200b<span data-type=\"html-entity\" class=\"vditor-ir__node\"><code data-newline=\"1\" class=\"vditor-ir__marker vditor-ir__marker--pre\" data-type=\"html-entity\">&amp;emsp;</code><span class=\"vditor-ir__preview\" data-render=\"2\"><code>\u2003</code></span></span></p>"},
	{"19", "<p data-block=\"0\" data-node-id=\"20200817131123-tyrwvzd\"><span data-type=\"inline-math\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker\">$</span><code data-newline=\"1\" class=\"vditor-ir__marker vditor-ir__marker--pre\" data-type=\"math-inline\">foo</code><span class=\"vditor-ir__preview\" data-render=\"2\"></span><span class=\"vditor-ir__marker\">$b<wbr></span></span></p>", "<p data-block=\"0\" data-node-id=\"20200817131123-tyrwvzd\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#iconParagraph\"></use></svg></span><span data-type=\"inline-math\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker\">$</span><code data-newline=\"1\" class=\"vditor-ir__marker vditor-ir__marker--pre\" data-type=\"math-inline\">foo</code><span class=\"vditor-ir__preview\" data-render=\"2\"><code class=\"language-math\">foo</code></span><span class=\"vditor-ir__marker\">$</span></span>b<wbr></p>"},
	// TODO 重写嵌入实现 {"18", "<p data-block=\"0\" data-node-id=\"20200815230922-cap9upn\">foo<span data-type=\"block-ref-embed\" class=\"vditor-ir__node\"><span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">20200813085553-gjpeb40</span> <span class=\"vditor-ir__blockref\">\"*\"</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span></span></p><p data-block=\"0\" data-node-id=\"20200815230922-cap9upn\"><font color=\"#008000\"><wbr><br></font><span data-type=\"block-ref-embed\" class=\"vditor-ir__node\"><span data-block-def-id=\"20200813085553-gjpeb40\" data-render=\"1\" data-type=\"block-render\" style=\"height: 0px;\"></span></span></p>", "<p data-block=\"0\" data-node-id=\"20200815230922-cap9upn\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#iconParagraph\"></use></svg></span>foo<span data-type=\"block-ref-embed\" class=\"vditor-ir__node\"><span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">20200813085553-gjpeb40</span> <span class=\"vditor-ir__blockref\">&#34;*&#34;</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span><span data-block-def-id=\"20200813085553-gjpeb40\" data-render=\"2\" data-type=\"block-render\"></span></span></p><p data-block=\"0\" <span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#iconParagraph\"></use></svg></span><wbr></p>"},
	{"17", "<p data-block=\"0\" data-node-id=\"20200815151746-idfp3ma\">foo((20200815093609-63sg22j \"*<wbr>\"))bar</p>", "<p data-block=\"0\" data-node-id=\"20200815151746-idfp3ma\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#iconParagraph\"></use></svg></span>foo<span data-type=\"block-ref-embed\" class=\"vditor-ir__node vditor-ir__node--expand\"><span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">20200815093609-63sg22j</span> <span class=\"vditor-ir__blockref\">&#34;*<wbr>&#34;</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span><span data-block-def-id=\"20200815093609-63sg22j\" data-render=\"2\" data-type=\"block-render\"></span></span>bar</p>"},
	{"16", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\" data-node-id=\"20200814215643-77a63qr\"><li data-marker=\"*\" class=\"vditor-task\" data-node-id=\"\"><input checked=\"\" type=\"checkbox\"> foo<wbr></li></ul>", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\" data-node-id=\"20200814215643-77a63qr\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#vditor-icon-check\"></use></svg></span><li data-marker=\"*\" class=\"vditor-task\" data-node-id=\"\"><input checked=\"\" type=\"checkbox\"/> foo<wbr></li></ul>"},
	{"15", "<p data-block=\"0\" data-node-id=\"20200813185628-tdpjvvr\">foo</p><p data-block=\"0\" data-node-id=\"20200813185636-sp1i1bp\"><span data-type=\"block-ref\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">20200813185628-tdpjvvr</span> <span class=\"vditor-ir__blockref\">\"bar\"</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span> baz\n```\n[text](addr)\n```<wbr></p>", "<p data-block=\"0\" data-node-id=\"20200813185628-tdpjvvr\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#iconParagraph\"></use></svg></span>foo</p><p data-block=\"0\" data-node-id=\"20200813185636-sp1i1bp\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#iconParagraph\"></use></svg></span><span data-type=\"block-ref\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">20200813185628-tdpjvvr</span> <span class=\"vditor-ir__blockref\">&#34;bar&#34;</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span> baz</p><div data-block=\"0\" data-node-id=\"20200826173100-gfbf08r\" data-type=\"code-block\" class=\"vditor-ir__node vditor-ir__node--expand\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#vditor-icon-code\"></use></svg></span>​<span data-type=\"code-block-open-marker\">```</span><span class=\"vditor-ir__marker vditor-ir__marker--info\" data-type=\"code-block-info\">​</span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code>[text](addr)<wbr>\n</code></pre><pre class=\"vditor-ir__preview\" data-render=\"2\"><code>[text](addr)</code></pre><span data-type=\"code-block-close-marker\">```</span></div>"},
	{"14", "<p data-block=\"0\" data-node-id=\"20200813113846-42h0ba1\">![foo](bar)<wbr></p>", "<p data-block=\"0\" data-node-id=\"20200813113846-42h0ba1\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#iconParagraph\"></use></svg></span><span class=\"vditor-ir__node vditor-ir__node--expand\" data-type=\"img\"><span class=\"vditor-ir__marker\">!</span><span class=\"vditor-ir__marker vditor-ir__marker--bracket\">[</span><span class=\"vditor-ir__marker vditor-ir__marker--bracket\">foo</span><span class=\"vditor-ir__marker vditor-ir__marker--bracket\">]</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">bar</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><wbr><img src=\"bar\" alt=\"foo\"/></span></p>"},
	{"13", "<p data-block=\"0\" data-node-id=\"20200811153824-grm842y\">foo</p><blockquote data-block=\"0\" data-node-id=\"20200811153825-amjdbjz\"><p data-block=\"0\" data-node-id=\"\"><wbr><br></p></blockquote>", "<p data-block=\"0\" data-node-id=\"20200811153824-grm842y\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#iconParagraph\"></use></svg></span>foo</p><p data-block=\"0\" data-node-id=\"20200811153825-amjdbjz\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#iconParagraph\"></use></svg></span>&gt; <wbr></p>"},
	{"12", "<p data-block=\"0\" data-node-id=\"20200811153040-mrqu125\"><span data-type=\"a\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--bracket\">[</span><span class=\"vditor-ir__link\">foo</span><span class=\"vditor-ir__marker vditor-ir__marker--bracket\">]</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">bar</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)b<wbr></span></span></p>", "<p data-block=\"0\" data-node-id=\"20200811153040-mrqu125\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#iconParagraph\"></use></svg></span><span data-type=\"a\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--bracket\">[</span><span class=\"vditor-ir__link\">foo</span><span class=\"vditor-ir__marker vditor-ir__marker--bracket\">]</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">bar</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span>b<wbr></p>"},
	{"11", "<h1 data-block=\"0\" class=\"vditor-ir__node\" data-node-id=\"20200825151101-01fo919\" id=\"ir-ba<wbr>\" data-marker=\"#\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#vditor-icon-headings\"></use></svg></span><span class=\"vditor-ir__marker vditor-ir__marker--heading\" data-type=\"heading-marker\"># </span>foo<span data-type=\"heading-id\" class=\"vditor-ir__marker\"> {bar<wbr>}</span></h1>", "<h1 data-block=\"0\" class=\"vditor-ir__node vditor-ir__node--expand\" data-node-id=\"20200825151101-01fo919\" id=\"ir-bar&lt;wbr&gt;\" data-marker=\"#\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#iconH1\"></use></svg></span><span class=\"vditor-ir__marker vditor-ir__marker--heading\" data-type=\"heading-marker\"># </span>foo<span data-type=\"heading-id\" class=\"vditor-ir__marker\"> {bar<wbr>}</span></h1>"},
	{"10", "<p data-block=\"0\" data-node-id=\"20200811140724-dxmm3jo\"><span data-type=\"em\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--bi\">*</span><em data-newline=\"1\">foo</em><span class=\"vditor-ir__marker vditor-ir__marker--bi\">*b<wbr></span></span></p>", "<p data-block=\"0\" data-node-id=\"20200811140724-dxmm3jo\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#iconParagraph\"></use></svg></span><span data-type=\"em\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--bi\">*</span><em data-newline=\"1\">foo</em><span class=\"vditor-ir__marker vditor-ir__marker--bi\">*</span></span>b<wbr></p>"},
	{"9", "<p data-block=\"0\" data-node-id=\"20200810211034-9d34ae\"><span data-type=\"block-ref\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">20200810191413-6f5616</span> <span class=\"vditor-ir__blockref\">\"foo\"</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">) <wbr></span></span></p>", "<p data-block=\"0\" data-node-id=\"20200810211034-9d34ae\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#iconParagraph\"></use></svg></span><span data-type=\"block-ref\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">20200810191413-6f5616</span> <span class=\"vditor-ir__blockref\">&#34;foo&#34;</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span> <wbr></p>"},
	{"8", "<p data-block=\"0\" data-node-id=\"20200809211933-b81ed7\">* foo<wbr></p>", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\" data-node-id=\"20200809211933-b81ed7\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#vditor-icon-list\"></use></svg></span><li data-marker=\"*\" data-node-id=\"\">foo<wbr></li></ul>"},
	{"7", "<p data-block=\"0\" data-node-id=\"20200809184752-a537de\">&gt; foo<wbr></p>", "<blockquote data-block=\"0\" data-node-id=\"20200809184752-a537de\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#vditor-icon-quote\"></use></svg></span><p data-block=\"0\" data-node-id=\"\">foo<wbr></p></blockquote>"},
	{"6", "<p data-block=\"0\" data-node-id=\"20200809093825-b06abb\"><span data-type=\"block-ref\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">foo</span> <span class=\"vditor-ir__blockref\">\"text<wbr>\"</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span></p>", "<p data-block=\"0\" data-node-id=\"20200809093825-b06abb\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#iconParagraph\"></use></svg></span><span data-type=\"block-ref\" class=\"vditor-ir__node vditor-ir__node--expand\"><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">foo</span> <span class=\"vditor-ir__blockref\">&#34;text<wbr>&#34;</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span></p>"},
	{"5", "<p data-block=\"0\" data-node-id=\"1596459249782\">foo ((bar)) <wbr></p>", "<p data-block=\"0\" data-node-id=\"1596459249782\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#iconParagraph\"></use></svg></span>foo <span data-type=\"block-ref\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">bar</span> <span class=\"vditor-ir__blockref\">&#34;bar&#34;</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span> <wbr></p>"},
	{"4", "<p data-block=\"0\" data-node-id=\"1596459249782\">((foo \"text\")<wbr></p>\n", "<p data-block=\"0\" data-node-id=\"1596459249782\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#iconParagraph\"></use></svg></span>((foo &#34;text&#34;)<wbr></p>"},
	{"3", "<p data-block=\"0\" data-node-id=\"1596459249782\">((foo \"text\"))<wbr></p>\n", "<p data-block=\"0\" data-node-id=\"1596459249782\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#iconParagraph\"></use></svg></span><span data-type=\"block-ref\" class=\"vditor-ir__node vditor-ir__node--expand\"><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">foo</span> <span class=\"vditor-ir__blockref\">&#34;text&#34;</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span><wbr></p>"},
	{"2", "<p data-block=\"0\" data-node-id=\"1596459249782\">((foo))<wbr></p>\n", "<p data-block=\"0\" data-node-id=\"1596459249782\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#iconParagraph\"></use></svg></span><span data-type=\"block-ref\" class=\"vditor-ir__node vditor-ir__node--expand\"><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">foo</span> <span class=\"vditor-ir__blockref\">&#34;foo&#34;</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span><wbr></p>"},
	{"1", "<p data-block=\"0\" data-node-id=\"1\">foo</p><p data-block=\"0\" data-node-id=\"2\"><wbr><br></p>", "<p data-block=\"0\" data-node-id=\"1\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#iconParagraph\"></use></svg></span>foo</p><p data-block=\"0\" data-node-id=\"2\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#iconParagraph\"></use></svg></span>\u200b<wbr></p>"},
	{"0", "<p data-block=\"0\" data-node-id=\"1\">foo</p><p data-block=\"0\" data-node-id=\"20200811112006-322210\"><wbr><br></p>", "<p data-block=\"0\" data-node-id=\"1\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#iconParagraph\"></use></svg></span>foo</p><p data-block=\"0\" data-node-id=\"20200811112006-322210\"><span class=\"vditor-ir__menu\" data-menu=\"0\"><svg class=\"fn__none\"><use xlink:href=\"#iconParagraph\"></use></svg></span>\u200b<wbr></p>"},
}

func TestSpinVditorIRBlockDOM(t *testing.T) {
	luteEngine := lute.New()
	luteEngine.BlockRef = true

	for _, test := range spinVditorIRBlockDOMTests {
		html := luteEngine.SpinVditorIRBlockDOM(test.from)

		if "15" == test.name || "18" == test.name {
			// 去掉最后一个新生成的节点 ID，因为这个节点 ID 是随机生成，每次测试用例运行都不一样，比较没有意义，长度一致即可
			lastNodeIDIdx := strings.LastIndex(html, "data-node-id=")
			length := len("data-node-id=\"20200813190226-1234567\" ")
			html = html[:lastNodeIDIdx] + html[lastNodeIDIdx+length:]
			lastNodeIDIdx = strings.LastIndex(test.to, "data-node-id=")
			test.to = test.to[:lastNodeIDIdx] + test.to[lastNodeIDIdx+length:]
		}

		if test.to != html {
			t.Fatalf("test case [%s] failed\nexpected\n\t%q\ngot\n\t%q\noriginal html\n\t%q", test.name, test.to, html, test.from)
		}
	}
}
