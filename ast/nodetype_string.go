// Code generated by "stringer -type=NodeType"; DO NOT EDIT.

package ast

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NodeDocument-0]
	_ = x[NodeParagraph-1]
	_ = x[NodeHeading-2]
	_ = x[NodeHeadingC8hMarker-3]
	_ = x[NodeThematicBreak-4]
	_ = x[NodeBlockquote-5]
	_ = x[NodeBlockquoteMarker-6]
	_ = x[NodeList-7]
	_ = x[NodeListItem-8]
	_ = x[NodeHTMLBlock-9]
	_ = x[NodeInlineHTML-10]
	_ = x[NodeCodeBlock-11]
	_ = x[NodeCodeBlockFenceOpenMarker-12]
	_ = x[NodeCodeBlockFenceCloseMarker-13]
	_ = x[NodeCodeBlockFenceInfoMarker-14]
	_ = x[NodeCodeBlockCode-15]
	_ = x[NodeText-16]
	_ = x[NodeEmphasis-17]
	_ = x[NodeEmA6kOpenMarker-18]
	_ = x[NodeEmA6kCloseMarker-19]
	_ = x[NodeEmU8eOpenMarker-20]
	_ = x[NodeEmU8eCloseMarker-21]
	_ = x[NodeStrong-22]
	_ = x[NodeStrongA6kOpenMarker-23]
	_ = x[NodeStrongA6kCloseMarker-24]
	_ = x[NodeStrongU8eOpenMarker-25]
	_ = x[NodeStrongU8eCloseMarker-26]
	_ = x[NodeCodeSpan-27]
	_ = x[NodeCodeSpanOpenMarker-28]
	_ = x[NodeCodeSpanContent-29]
	_ = x[NodeCodeSpanCloseMarker-30]
	_ = x[NodeHardBreak-31]
	_ = x[NodeSoftBreak-32]
	_ = x[NodeLink-33]
	_ = x[NodeImage-34]
	_ = x[NodeBang-35]
	_ = x[NodeOpenBracket-36]
	_ = x[NodeCloseBracket-37]
	_ = x[NodeOpenParen-38]
	_ = x[NodeCloseParen-39]
	_ = x[NodeLinkText-40]
	_ = x[NodeLinkDest-41]
	_ = x[NodeLinkTitle-42]
	_ = x[NodeLinkSpace-43]
	_ = x[NodeHTMLEntity-44]
	_ = x[NodeLinkRefDefBlock-45]
	_ = x[NodeLinkRefDef-46]
	_ = x[NodeTaskListItemMarker-100]
	_ = x[NodeStrikethrough-101]
	_ = x[NodeStrikethrough1OpenMarker-102]
	_ = x[NodeStrikethrough1CloseMarker-103]
	_ = x[NodeStrikethrough2OpenMarker-104]
	_ = x[NodeStrikethrough2CloseMarker-105]
	_ = x[NodeTable-106]
	_ = x[NodeTableHead-107]
	_ = x[NodeTableRow-108]
	_ = x[NodeTableCell-109]
	_ = x[NodeEmoji-200]
	_ = x[NodeEmojiUnicode-201]
	_ = x[NodeEmojiImg-202]
	_ = x[NodeEmojiAlias-203]
	_ = x[NodeMathBlock-300]
	_ = x[NodeMathBlockOpenMarker-301]
	_ = x[NodeMathBlockContent-302]
	_ = x[NodeMathBlockCloseMarker-303]
	_ = x[NodeInlineMath-304]
	_ = x[NodeInlineMathOpenMarker-305]
	_ = x[NodeInlineMathContent-306]
	_ = x[NodeInlineMathCloseMarker-307]
	_ = x[NodeBackslash-400]
	_ = x[NodeBackslashContent-401]
	_ = x[NodeVditorCaret-405]
	_ = x[NodeFootnotesDefBlock-410]
	_ = x[NodeFootnotesDef-411]
	_ = x[NodeFootnotesRef-412]
	_ = x[NodeToC-415]
	_ = x[NodeHeadingID-420]
	_ = x[NodeYamlFrontMatter-425]
	_ = x[NodeYamlFrontMatterOpenMarker-426]
	_ = x[NodeYamlFrontMatterContent-427]
	_ = x[NodeYamlFrontMatterCloseMarker-428]
	_ = x[NodeBlockRef-430]
	_ = x[NodeBlockRefID-431]
	_ = x[NodeBlockRefSpace-432]
	_ = x[NodeBlockRefText-433]
	_ = x[NodeBlockRefTextTplRenderResult-434]
	_ = x[NodeBlockEmbed-440]
	_ = x[NodeBlockEmbedID-441]
	_ = x[NodeBlockEmbedSpace-442]
	_ = x[NodeBlockEmbedText-443]
	_ = x[NodeBlockEmbedTextTplRenderResult-444]
	_ = x[NodeMark-450]
	_ = x[NodeMark1OpenMarker-451]
	_ = x[NodeMark1CloseMarker-452]
	_ = x[NodeMark2OpenMarker-453]
	_ = x[NodeMark2CloseMarker-454]
	_ = x[NodeKramdownBlockIAL-455]
	_ = x[NodeKramdownSpanIAL-456]
	_ = x[NodeTag-460]
	_ = x[NodeTagOpenMarker-461]
	_ = x[NodeTagCloseMarker-462]
	_ = x[NodeBlockQueryEmbed-465]
	_ = x[NodeOpenBrace-466]
	_ = x[NodeCloseBrace-467]
	_ = x[NodeBlockQueryEmbedScript-468]
	_ = x[NodeSuperBlock-475]
	_ = x[NodeSuperBlockOpenMarker-476]
	_ = x[NodeSuperBlockLayoutMarker-477]
	_ = x[NodeSuperBlockCloseMarker-478]
	_ = x[NodeSup-485]
	_ = x[NodeSupOpenMarker-486]
	_ = x[NodeSupCloseMarker-487]
	_ = x[NodeSub-490]
	_ = x[NodeSubOpenMarker-491]
	_ = x[NodeSubCloseMarker-492]
	_ = x[NodeGitConflict-495]
	_ = x[NodeGitConflictOpenMarker-496]
	_ = x[NodeGitConflictContent-497]
	_ = x[NodeGitConflictCloseMarker-498]
	_ = x[NodeIFrame-500]
	_ = x[NodeAudio-505]
	_ = x[NodeVideo-510]
	_ = x[NodeKbd-515]
	_ = x[NodeKbdOpenMarker-516]
	_ = x[NodeKbdCloseMarker-517]
	_ = x[NodeUnderline-520]
	_ = x[NodeUnderlineOpenMarker-521]
	_ = x[NodeUnderlineCloseMarker-522]
	_ = x[NodeTypeMaxVal-1024]
}

const _NodeType_name = "NodeDocumentNodeParagraphNodeHeadingNodeHeadingC8hMarkerNodeThematicBreakNodeBlockquoteNodeBlockquoteMarkerNodeListNodeListItemNodeHTMLBlockNodeInlineHTMLNodeCodeBlockNodeCodeBlockFenceOpenMarkerNodeCodeBlockFenceCloseMarkerNodeCodeBlockFenceInfoMarkerNodeCodeBlockCodeNodeTextNodeEmphasisNodeEmA6kOpenMarkerNodeEmA6kCloseMarkerNodeEmU8eOpenMarkerNodeEmU8eCloseMarkerNodeStrongNodeStrongA6kOpenMarkerNodeStrongA6kCloseMarkerNodeStrongU8eOpenMarkerNodeStrongU8eCloseMarkerNodeCodeSpanNodeCodeSpanOpenMarkerNodeCodeSpanContentNodeCodeSpanCloseMarkerNodeHardBreakNodeSoftBreakNodeLinkNodeImageNodeBangNodeOpenBracketNodeCloseBracketNodeOpenParenNodeCloseParenNodeLinkTextNodeLinkDestNodeLinkTitleNodeLinkSpaceNodeHTMLEntityNodeLinkRefDefBlockNodeLinkRefDefNodeTaskListItemMarkerNodeStrikethroughNodeStrikethrough1OpenMarkerNodeStrikethrough1CloseMarkerNodeStrikethrough2OpenMarkerNodeStrikethrough2CloseMarkerNodeTableNodeTableHeadNodeTableRowNodeTableCellNodeEmojiNodeEmojiUnicodeNodeEmojiImgNodeEmojiAliasNodeMathBlockNodeMathBlockOpenMarkerNodeMathBlockContentNodeMathBlockCloseMarkerNodeInlineMathNodeInlineMathOpenMarkerNodeInlineMathContentNodeInlineMathCloseMarkerNodeBackslashNodeBackslashContentNodeVditorCaretNodeFootnotesDefBlockNodeFootnotesDefNodeFootnotesRefNodeToCNodeHeadingIDNodeYamlFrontMatterNodeYamlFrontMatterOpenMarkerNodeYamlFrontMatterContentNodeYamlFrontMatterCloseMarkerNodeBlockRefNodeBlockRefIDNodeBlockRefSpaceNodeBlockRefTextNodeBlockRefTextTplRenderResultNodeBlockEmbedNodeBlockEmbedIDNodeBlockEmbedSpaceNodeBlockEmbedTextNodeBlockEmbedTextTplRenderResultNodeMarkNodeMark1OpenMarkerNodeMark1CloseMarkerNodeMark2OpenMarkerNodeMark2CloseMarkerNodeKramdownBlockIALNodeKramdownSpanIALNodeTagNodeTagOpenMarkerNodeTagCloseMarkerNodeBlockQueryEmbedNodeOpenBraceNodeCloseBraceNodeBlockQueryEmbedScriptNodeSuperBlockNodeSuperBlockOpenMarkerNodeSuperBlockLayoutMarkerNodeSuperBlockCloseMarkerNodeSupNodeSupOpenMarkerNodeSupCloseMarkerNodeSubNodeSubOpenMarkerNodeSubCloseMarkerNodeGitConflictNodeGitConflictOpenMarkerNodeGitConflictContentNodeGitConflictCloseMarkerNodeIFrameNodeAudioNodeVideoNodeKbdNodeKbdOpenMarkerNodeKbdCloseMarkerNodeUnderlineNodeUnderlineOpenMarkerNodeUnderlineCloseMarkerNodeTypeMaxVal"

var _NodeType_map = map[NodeType]string{
	0:    _NodeType_name[0:12],
	1:    _NodeType_name[12:25],
	2:    _NodeType_name[25:36],
	3:    _NodeType_name[36:56],
	4:    _NodeType_name[56:73],
	5:    _NodeType_name[73:87],
	6:    _NodeType_name[87:107],
	7:    _NodeType_name[107:115],
	8:    _NodeType_name[115:127],
	9:    _NodeType_name[127:140],
	10:   _NodeType_name[140:154],
	11:   _NodeType_name[154:167],
	12:   _NodeType_name[167:195],
	13:   _NodeType_name[195:224],
	14:   _NodeType_name[224:252],
	15:   _NodeType_name[252:269],
	16:   _NodeType_name[269:277],
	17:   _NodeType_name[277:289],
	18:   _NodeType_name[289:308],
	19:   _NodeType_name[308:328],
	20:   _NodeType_name[328:347],
	21:   _NodeType_name[347:367],
	22:   _NodeType_name[367:377],
	23:   _NodeType_name[377:400],
	24:   _NodeType_name[400:424],
	25:   _NodeType_name[424:447],
	26:   _NodeType_name[447:471],
	27:   _NodeType_name[471:483],
	28:   _NodeType_name[483:505],
	29:   _NodeType_name[505:524],
	30:   _NodeType_name[524:547],
	31:   _NodeType_name[547:560],
	32:   _NodeType_name[560:573],
	33:   _NodeType_name[573:581],
	34:   _NodeType_name[581:590],
	35:   _NodeType_name[590:598],
	36:   _NodeType_name[598:613],
	37:   _NodeType_name[613:629],
	38:   _NodeType_name[629:642],
	39:   _NodeType_name[642:656],
	40:   _NodeType_name[656:668],
	41:   _NodeType_name[668:680],
	42:   _NodeType_name[680:693],
	43:   _NodeType_name[693:706],
	44:   _NodeType_name[706:720],
	45:   _NodeType_name[720:739],
	46:   _NodeType_name[739:753],
	100:  _NodeType_name[753:775],
	101:  _NodeType_name[775:792],
	102:  _NodeType_name[792:820],
	103:  _NodeType_name[820:849],
	104:  _NodeType_name[849:877],
	105:  _NodeType_name[877:906],
	106:  _NodeType_name[906:915],
	107:  _NodeType_name[915:928],
	108:  _NodeType_name[928:940],
	109:  _NodeType_name[940:953],
	200:  _NodeType_name[953:962],
	201:  _NodeType_name[962:978],
	202:  _NodeType_name[978:990],
	203:  _NodeType_name[990:1004],
	300:  _NodeType_name[1004:1017],
	301:  _NodeType_name[1017:1040],
	302:  _NodeType_name[1040:1060],
	303:  _NodeType_name[1060:1084],
	304:  _NodeType_name[1084:1098],
	305:  _NodeType_name[1098:1122],
	306:  _NodeType_name[1122:1143],
	307:  _NodeType_name[1143:1168],
	400:  _NodeType_name[1168:1181],
	401:  _NodeType_name[1181:1201],
	405:  _NodeType_name[1201:1216],
	410:  _NodeType_name[1216:1237],
	411:  _NodeType_name[1237:1253],
	412:  _NodeType_name[1253:1269],
	415:  _NodeType_name[1269:1276],
	420:  _NodeType_name[1276:1289],
	425:  _NodeType_name[1289:1308],
	426:  _NodeType_name[1308:1337],
	427:  _NodeType_name[1337:1363],
	428:  _NodeType_name[1363:1393],
	430:  _NodeType_name[1393:1405],
	431:  _NodeType_name[1405:1419],
	432:  _NodeType_name[1419:1436],
	433:  _NodeType_name[1436:1452],
	434:  _NodeType_name[1452:1483],
	440:  _NodeType_name[1483:1497],
	441:  _NodeType_name[1497:1513],
	442:  _NodeType_name[1513:1532],
	443:  _NodeType_name[1532:1550],
	444:  _NodeType_name[1550:1583],
	450:  _NodeType_name[1583:1591],
	451:  _NodeType_name[1591:1610],
	452:  _NodeType_name[1610:1630],
	453:  _NodeType_name[1630:1649],
	454:  _NodeType_name[1649:1669],
	455:  _NodeType_name[1669:1689],
	456:  _NodeType_name[1689:1708],
	460:  _NodeType_name[1708:1715],
	461:  _NodeType_name[1715:1732],
	462:  _NodeType_name[1732:1750],
	465:  _NodeType_name[1750:1769],
	466:  _NodeType_name[1769:1782],
	467:  _NodeType_name[1782:1796],
	468:  _NodeType_name[1796:1821],
	475:  _NodeType_name[1821:1835],
	476:  _NodeType_name[1835:1859],
	477:  _NodeType_name[1859:1885],
	478:  _NodeType_name[1885:1910],
	485:  _NodeType_name[1910:1917],
	486:  _NodeType_name[1917:1934],
	487:  _NodeType_name[1934:1952],
	490:  _NodeType_name[1952:1959],
	491:  _NodeType_name[1959:1976],
	492:  _NodeType_name[1976:1994],
	495:  _NodeType_name[1994:2009],
	496:  _NodeType_name[2009:2034],
	497:  _NodeType_name[2034:2056],
	498:  _NodeType_name[2056:2082],
	500:  _NodeType_name[2082:2092],
	505:  _NodeType_name[2092:2101],
	510:  _NodeType_name[2101:2110],
	515:  _NodeType_name[2110:2117],
	516:  _NodeType_name[2117:2134],
	517:  _NodeType_name[2134:2152],
	520:  _NodeType_name[2152:2165],
	521:  _NodeType_name[2165:2188],
	522:  _NodeType_name[2188:2212],
	1024: _NodeType_name[2212:2226],
}

func (i NodeType) String() string {
	if str, ok := _NodeType_map[i]; ok {
		return str
	}
	return "NodeType(" + strconv.FormatInt(int64(i), 10) + ")"
}
