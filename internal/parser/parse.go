package parser

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/gobuffalo/lush/ast"
)

var g = &grammar{
	rules: []*rule{
		{
			name: "DOC",
			pos:  position{line: 12, col: 1, offset: 115},
			expr: &actionExpr{
				pos: position{line: 12, col: 8, offset: 122},
				run: (*parser).callonDOC1,
				expr: &seqExpr{
					pos: position{line: 12, col: 8, offset: 122},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 12, col: 8, offset: 122},
							label: "b",
							expr: &zeroOrOneExpr{
								pos: position{line: 12, col: 10, offset: 124},
								expr: &ruleRefExpr{
									pos:  position{line: 12, col: 10, offset: 124},
									name: "SHEBANG",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 12, col: 19, offset: 133},
							label: "stmts",
							expr: &oneOrMoreExpr{
								pos: position{line: 12, col: 25, offset: 139},
								expr: &seqExpr{
									pos: position{line: 12, col: 26, offset: 140},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 12, col: 26, offset: 140},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 12, col: 28, offset: 142},
											name: "NODE",
										},
										&ruleRefExpr{
											pos:  position{line: 12, col: 33, offset: 147},
											name: "_",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 12, col: 37, offset: 151},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "NODE",
			pos:  position{line: 21, col: 1, offset: 313},
			expr: &actionExpr{
				pos: position{line: 21, col: 9, offset: 321},
				run: (*parser).callonNODE1,
				expr: &seqExpr{
					pos: position{line: 21, col: 9, offset: 321},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 21, col: 9, offset: 321},
							label: "s",
							expr: &choiceExpr{
								pos: position{line: 21, col: 13, offset: 325},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 21, col: 13, offset: 325},
										name: "CURRENT",
									},
									&ruleRefExpr{
										pos:  position{line: 21, col: 23, offset: 335},
										name: "IMPORT",
									},
									&ruleRefExpr{
										pos:  position{line: 21, col: 32, offset: 344},
										name: "GO",
									},
									&ruleRefExpr{
										pos:  position{line: 21, col: 37, offset: 349},
										name: "COMMENT",
									},
									&ruleRefExpr{
										pos:  position{line: 21, col: 47, offset: 359},
										name: "NULL",
									},
									&ruleRefExpr{
										pos:  position{line: 21, col: 54, offset: 366},
										name: "LET",
									},
									&ruleRefExpr{
										pos:  position{line: 21, col: 60, offset: 372},
										name: "VAR",
									},
									&ruleRefExpr{
										pos:  position{line: 21, col: 66, offset: 378},
										name: "ASSIGN",
									},
									&ruleRefExpr{
										pos:  position{line: 21, col: 75, offset: 387},
										name: "IF",
									},
									&ruleRefExpr{
										pos:  position{line: 21, col: 80, offset: 392},
										name: "IFOR",
									},
									&ruleRefExpr{
										pos:  position{line: 21, col: 87, offset: 399},
										name: "FOR",
									},
									&ruleRefExpr{
										pos:  position{line: 21, col: 93, offset: 405},
										name: "RANGE",
									},
									&ruleRefExpr{
										pos:  position{line: 21, col: 101, offset: 413},
										name: "GO",
									},
									&ruleRefExpr{
										pos:  position{line: 21, col: 106, offset: 418},
										name: "FCALL",
									},
									&ruleRefExpr{
										pos:  position{line: 21, col: 114, offset: 426},
										name: "RETURN",
									},
									&ruleRefExpr{
										pos:  position{line: 21, col: 123, offset: 435},
										name: "CALL",
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 21, col: 130, offset: 442},
							expr: &litMatcher{
								pos:        position{line: 21, col: 130, offset: 442},
								val:        ";",
								ignoreCase: false,
							},
						},
					},
				},
			},
		},
		{
			name: "IMPORT",
			pos:  position{line: 25, col: 1, offset: 468},
			expr: &actionExpr{
				pos: position{line: 25, col: 11, offset: 478},
				run: (*parser).callonIMPORT1,
				expr: &seqExpr{
					pos: position{line: 25, col: 11, offset: 478},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 25, col: 11, offset: 478},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 25, col: 13, offset: 480},
							val:        "import",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 25, col: 22, offset: 489},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 25, col: 24, offset: 491},
							label: "s",
							expr: &ruleRefExpr{
								pos:  position{line: 25, col: 26, offset: 493},
								name: "DQSTRING",
							},
						},
					},
				},
			},
		},
		{
			name: "CURRENT",
			pos:  position{line: 29, col: 1, offset: 532},
			expr: &actionExpr{
				pos: position{line: 29, col: 12, offset: 543},
				run: (*parser).callonCURRENT1,
				expr: &litMatcher{
					pos:        position{line: 29, col: 12, offset: 543},
					val:        "current",
					ignoreCase: false,
				},
			},
		},
		{
			name: "EXPRARG",
			pos:  position{line: 33, col: 1, offset: 599},
			expr: &choiceExpr{
				pos: position{line: 33, col: 13, offset: 611},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 33, col: 13, offset: 611},
						name: "CURRENT",
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 23, offset: 621},
						name: "POPEX",
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 31, offset: 629},
						name: "CALL",
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 38, offset: 636},
						name: "NUMBER",
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 47, offset: 645},
						name: "STRING",
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 56, offset: 654},
						name: "BOOL",
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 63, offset: 661},
						name: "ARRAY",
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 71, offset: 669},
						name: "MAP",
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 77, offset: 675},
						name: "NULL",
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 84, offset: 682},
						name: "IDENT",
					},
				},
			},
		},
		{
			name: "OP",
			pos:  position{line: 35, col: 1, offset: 690},
			expr: &actionExpr{
				pos: position{line: 35, col: 7, offset: 696},
				run: (*parser).callonOP1,
				expr: &choiceExpr{
					pos: position{line: 35, col: 9, offset: 698},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 35, col: 9, offset: 698},
							val:        "&&",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 35, col: 16, offset: 705},
							val:        "||",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 35, col: 23, offset: 712},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 35, col: 30, offset: 719},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 35, col: 37, offset: 726},
							val:        "==",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 35, col: 44, offset: 733},
							val:        "!=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 35, col: 51, offset: 740},
							val:        "~=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 35, col: 58, offset: 747},
							val:        "+",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 35, col: 64, offset: 753},
							val:        "-",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 35, col: 70, offset: 759},
							val:        "*",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 35, col: 76, offset: 765},
							val:        "/",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 35, col: 82, offset: 771},
							val:        "%",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 35, col: 88, offset: 777},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 35, col: 94, offset: 783},
							val:        "<",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OPEX",
			pos:  position{line: 39, col: 1, offset: 822},
			expr: &actionExpr{
				pos: position{line: 39, col: 9, offset: 830},
				run: (*parser).callonOPEX1,
				expr: &seqExpr{
					pos: position{line: 39, col: 9, offset: 830},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 39, col: 9, offset: 830},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 39, col: 11, offset: 832},
							label: "a",
							expr: &ruleRefExpr{
								pos:  position{line: 39, col: 14, offset: 835},
								name: "EXPRARG",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 39, col: 23, offset: 844},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 39, col: 25, offset: 846},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 39, col: 28, offset: 849},
								name: "OP",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 39, col: 31, offset: 852},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 39, col: 33, offset: 854},
							label: "b",
							expr: &choiceExpr{
								pos: position{line: 39, col: 36, offset: 857},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 39, col: 36, offset: 857},
										name: "OPEX",
									},
									&ruleRefExpr{
										pos:  position{line: 39, col: 43, offset: 864},
										name: "EXPRARG",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 39, col: 52, offset: 873},
							name: "_",
						},
						&zeroOrOneExpr{
							pos: position{line: 39, col: 54, offset: 875},
							expr: &litMatcher{
								pos:        position{line: 39, col: 54, offset: 875},
								val:        ";",
								ignoreCase: false,
							},
						},
					},
				},
			},
		},
		{
			name: "POPEX",
			pos:  position{line: 43, col: 1, offset: 923},
			expr: &actionExpr{
				pos: position{line: 43, col: 10, offset: 932},
				run: (*parser).callonPOPEX1,
				expr: &seqExpr{
					pos: position{line: 43, col: 10, offset: 932},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 43, col: 10, offset: 932},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 43, col: 12, offset: 934},
							name: "LP",
						},
						&ruleRefExpr{
							pos:  position{line: 43, col: 15, offset: 937},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 43, col: 17, offset: 939},
							label: "a",
							expr: &ruleRefExpr{
								pos:  position{line: 43, col: 20, offset: 942},
								name: "EXPRARG",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 43, col: 29, offset: 951},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 43, col: 31, offset: 953},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 43, col: 34, offset: 956},
								name: "OP",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 43, col: 37, offset: 959},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 43, col: 39, offset: 961},
							label: "b",
							expr: &ruleRefExpr{
								pos:  position{line: 43, col: 42, offset: 964},
								name: "EXPRARG",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 43, col: 51, offset: 973},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 43, col: 53, offset: 975},
							name: "RP",
						},
						&ruleRefExpr{
							pos:  position{line: 43, col: 56, offset: 978},
							name: "_",
						},
						&zeroOrOneExpr{
							pos: position{line: 43, col: 58, offset: 980},
							expr: &litMatcher{
								pos:        position{line: 43, col: 58, offset: 980},
								val:        ";",
								ignoreCase: false,
							},
						},
					},
				},
			},
		},
		{
			name: "IF",
			pos:  position{line: 47, col: 1, offset: 1029},
			expr: &actionExpr{
				pos: position{line: 47, col: 7, offset: 1035},
				run: (*parser).callonIF1,
				expr: &seqExpr{
					pos: position{line: 47, col: 7, offset: 1035},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 47, col: 7, offset: 1035},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 47, col: 9, offset: 1037},
							val:        "if",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 47, col: 14, offset: 1042},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 47, col: 16, offset: 1044},
							label: "p",
							expr: &zeroOrOneExpr{
								pos: position{line: 47, col: 18, offset: 1046},
								expr: &choiceExpr{
									pos: position{line: 47, col: 21, offset: 1049},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 47, col: 21, offset: 1049},
											name: "VAR",
										},
										&ruleRefExpr{
											pos:  position{line: 47, col: 27, offset: 1055},
											name: "LET",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 47, col: 35, offset: 1063},
							name: "_",
						},
						&zeroOrOneExpr{
							pos: position{line: 47, col: 37, offset: 1065},
							expr: &litMatcher{
								pos:        position{line: 47, col: 37, offset: 1065},
								val:        ";",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 47, col: 42, offset: 1070},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 47, col: 45, offset: 1073},
							label: "s",
							expr: &choiceExpr{
								pos: position{line: 47, col: 48, offset: 1076},
								alternatives: []interface{}{
									&seqExpr{
										pos: position{line: 47, col: 48, offset: 1076},
										exprs: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 47, col: 48, offset: 1076},
												name: "LP",
											},
											&ruleRefExpr{
												pos:  position{line: 47, col: 51, offset: 1079},
												name: "_",
											},
											&choiceExpr{
												pos: position{line: 47, col: 54, offset: 1082},
												alternatives: []interface{}{
													&ruleRefExpr{
														pos:  position{line: 47, col: 54, offset: 1082},
														name: "BOOL",
													},
													&ruleRefExpr{
														pos:  position{line: 47, col: 61, offset: 1089},
														name: "OPEX",
													},
													&ruleRefExpr{
														pos:  position{line: 47, col: 68, offset: 1096},
														name: "POPEX",
													},
												},
											},
											&ruleRefExpr{
												pos:  position{line: 47, col: 76, offset: 1104},
												name: "_",
											},
											&ruleRefExpr{
												pos:  position{line: 47, col: 78, offset: 1106},
												name: "RP",
											},
										},
									},
									&choiceExpr{
										pos: position{line: 47, col: 84, offset: 1112},
										alternatives: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 47, col: 84, offset: 1112},
												name: "BOOL",
											},
											&ruleRefExpr{
												pos:  position{line: 47, col: 91, offset: 1119},
												name: "OPEX",
											},
											&ruleRefExpr{
												pos:  position{line: 47, col: 98, offset: 1126},
												name: "POPEX",
											},
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 47, col: 107, offset: 1135},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 47, col: 109, offset: 1137},
							label: "b",
							expr: &ruleRefExpr{
								pos:  position{line: 47, col: 111, offset: 1139},
								name: "BLOCK",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 47, col: 117, offset: 1145},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 47, col: 119, offset: 1147},
							label: "e",
							expr: &zeroOrMoreExpr{
								pos: position{line: 47, col: 121, offset: 1149},
								expr: &choiceExpr{
									pos: position{line: 47, col: 122, offset: 1150},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 47, col: 122, offset: 1150},
											name: "ELSEIF",
										},
										&ruleRefExpr{
											pos:  position{line: 47, col: 131, offset: 1159},
											name: "ELSE",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "ELSE",
			pos:  position{line: 51, col: 1, offset: 1201},
			expr: &actionExpr{
				pos: position{line: 51, col: 9, offset: 1209},
				run: (*parser).callonELSE1,
				expr: &seqExpr{
					pos: position{line: 51, col: 9, offset: 1209},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 51, col: 9, offset: 1209},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 51, col: 11, offset: 1211},
							val:        "else",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 51, col: 18, offset: 1218},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 51, col: 20, offset: 1220},
							label: "b",
							expr: &ruleRefExpr{
								pos:  position{line: 51, col: 22, offset: 1222},
								name: "BLOCK",
							},
						},
					},
				},
			},
		},
		{
			name: "ELSEIF",
			pos:  position{line: 55, col: 1, offset: 1256},
			expr: &actionExpr{
				pos: position{line: 55, col: 11, offset: 1266},
				run: (*parser).callonELSEIF1,
				expr: &seqExpr{
					pos: position{line: 55, col: 11, offset: 1266},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 55, col: 11, offset: 1266},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 55, col: 13, offset: 1268},
							val:        "else",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 55, col: 20, offset: 1275},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 55, col: 22, offset: 1277},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 55, col: 24, offset: 1279},
								name: "IF",
							},
						},
					},
				},
			},
		},
		{
			name: "INTEGER",
			pos:  position{line: 59, col: 1, offset: 1312},
			expr: &actionExpr{
				pos: position{line: 59, col: 12, offset: 1323},
				run: (*parser).callonINTEGER1,
				expr: &seqExpr{
					pos: position{line: 59, col: 12, offset: 1323},
					exprs: []interface{}{
						&zeroOrOneExpr{
							pos: position{line: 59, col: 12, offset: 1323},
							expr: &litMatcher{
								pos:        position{line: 59, col: 12, offset: 1323},
								val:        "-",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 59, col: 17, offset: 1328},
							expr: &charClassMatcher{
								pos:        position{line: 59, col: 17, offset: 1328},
								val:        "[0-9]",
								ranges:     []rune{'0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
						},
					},
				},
			},
		},
		{
			name: "FLOAT",
			pos:  position{line: 63, col: 1, offset: 1371},
			expr: &actionExpr{
				pos: position{line: 63, col: 10, offset: 1380},
				run: (*parser).callonFLOAT1,
				expr: &seqExpr{
					pos: position{line: 63, col: 10, offset: 1380},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 63, col: 10, offset: 1380},
							name: "INTEGER",
						},
						&zeroOrOneExpr{
							pos: position{line: 63, col: 18, offset: 1388},
							expr: &litMatcher{
								pos:        position{line: 63, col: 18, offset: 1388},
								val:        ".",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 63, col: 23, offset: 1393},
							name: "INTEGER",
						},
					},
				},
			},
		},
		{
			name: "UN",
			pos:  position{line: 67, col: 1, offset: 1435},
			expr: &litMatcher{
				pos:        position{line: 67, col: 7, offset: 1441},
				val:        "_",
				ignoreCase: false,
			},
		},
		{
			name: "IDENT",
			pos:  position{line: 69, col: 1, offset: 1446},
			expr: &actionExpr{
				pos: position{line: 69, col: 10, offset: 1455},
				run: (*parser).callonIDENT1,
				expr: &seqExpr{
					pos: position{line: 69, col: 10, offset: 1455},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 69, col: 10, offset: 1455},
							expr: &charClassMatcher{
								pos:        position{line: 69, col: 10, offset: 1455},
								val:        "[a-zA-Z]",
								ranges:     []rune{'a', 'z', 'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 69, col: 20, offset: 1465},
							expr: &charClassMatcher{
								pos:        position{line: 69, col: 22, offset: 1467},
								val:        "[0-9a-zA-Z_]",
								chars:      []rune{'_'},
								ranges:     []rune{'0', '9', 'a', 'z', 'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
					},
				},
			},
		},
		{
			name: "NUMBER",
			pos:  position{line: 73, col: 1, offset: 1516},
			expr: &choiceExpr{
				pos: position{line: 73, col: 12, offset: 1527},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 73, col: 12, offset: 1527},
						name: "FLOAT",
					},
					&ruleRefExpr{
						pos:  position{line: 73, col: 20, offset: 1535},
						name: "INTEGER",
					},
				},
			},
		},
		{
			name: "LETV",
			pos:  position{line: 75, col: 1, offset: 1545},
			expr: &choiceExpr{
				pos: position{line: 75, col: 11, offset: 1555},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 75, col: 11, offset: 1555},
						name: "CURRENT",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 22, offset: 1566},
						name: "ACCESS",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 31, offset: 1575},
						name: "NULL",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 38, offset: 1582},
						name: "POPEX",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 46, offset: 1590},
						name: "OPEX",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 53, offset: 1597},
						name: "FCALL",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 61, offset: 1605},
						name: "FUNC",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 68, offset: 1612},
						name: "CALL",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 75, offset: 1619},
						name: "MAP",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 81, offset: 1625},
						name: "ARRAY",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 88, offset: 1632},
						name: "BOOL",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 95, offset: 1639},
						name: "STRING",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 104, offset: 1648},
						name: "NUMBER",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 113, offset: 1657},
						name: "NULL",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 120, offset: 1664},
						name: "IDENT",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 128, offset: 1672},
						name: "FOR",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 134, offset: 1678},
						name: "RANGE",
					},
				},
			},
		},
		{
			name: "LET",
			pos:  position{line: 77, col: 1, offset: 1688},
			expr: &actionExpr{
				pos: position{line: 77, col: 8, offset: 1695},
				run: (*parser).callonLET1,
				expr: &seqExpr{
					pos: position{line: 77, col: 8, offset: 1695},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 77, col: 8, offset: 1695},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 77, col: 10, offset: 1697},
							val:        "let",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 77, col: 16, offset: 1703},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 77, col: 18, offset: 1705},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 77, col: 23, offset: 1710},
								name: "IDENT",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 77, col: 29, offset: 1716},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 77, col: 31, offset: 1718},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 77, col: 35, offset: 1722},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 77, col: 37, offset: 1724},
							label: "value",
							expr: &ruleRefExpr{
								pos:  position{line: 77, col: 43, offset: 1730},
								name: "LETV",
							},
						},
					},
				},
			},
		},
		{
			name: "VAR",
			pos:  position{line: 81, col: 1, offset: 1772},
			expr: &actionExpr{
				pos: position{line: 81, col: 8, offset: 1779},
				run: (*parser).callonVAR1,
				expr: &seqExpr{
					pos: position{line: 81, col: 8, offset: 1779},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 81, col: 8, offset: 1779},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 81, col: 10, offset: 1781},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 81, col: 15, offset: 1786},
								name: "IDENT",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 81, col: 21, offset: 1792},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 81, col: 23, offset: 1794},
							val:        ":=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 81, col: 28, offset: 1799},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 81, col: 30, offset: 1801},
							label: "value",
							expr: &ruleRefExpr{
								pos:  position{line: 81, col: 36, offset: 1807},
								name: "LETV",
							},
						},
					},
				},
			},
		},
		{
			name: "ASSIGN",
			pos:  position{line: 85, col: 1, offset: 1849},
			expr: &actionExpr{
				pos: position{line: 85, col: 11, offset: 1859},
				run: (*parser).callonASSIGN1,
				expr: &seqExpr{
					pos: position{line: 85, col: 11, offset: 1859},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 85, col: 11, offset: 1859},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 85, col: 13, offset: 1861},
							label: "name",
							expr: &choiceExpr{
								pos: position{line: 85, col: 19, offset: 1867},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 85, col: 19, offset: 1867},
										name: "ACCESS",
									},
									&ruleRefExpr{
										pos:  position{line: 85, col: 26, offset: 1874},
										name: "IDENT",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 85, col: 33, offset: 1881},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 85, col: 35, offset: 1883},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 85, col: 39, offset: 1887},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 85, col: 41, offset: 1889},
							label: "value",
							expr: &ruleRefExpr{
								pos:  position{line: 85, col: 47, offset: 1895},
								name: "LETV",
							},
						},
					},
				},
			},
		},
		{
			name: "ARG",
			pos:  position{line: 89, col: 1, offset: 1940},
			expr: &actionExpr{
				pos: position{line: 89, col: 8, offset: 1947},
				run: (*parser).callonARG1,
				expr: &seqExpr{
					pos: position{line: 89, col: 8, offset: 1947},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 89, col: 8, offset: 1947},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 89, col: 10, offset: 1949},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 89, col: 12, offset: 1951},
								name: "LETV",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 89, col: 17, offset: 1956},
							name: "_",
						},
						&zeroOrOneExpr{
							pos: position{line: 89, col: 19, offset: 1958},
							expr: &litMatcher{
								pos:        position{line: 89, col: 19, offset: 1958},
								val:        ",",
								ignoreCase: false,
							},
						},
					},
				},
			},
		},
		{
			name: "BLOCK",
			pos:  position{line: 93, col: 1, offset: 1984},
			expr: &actionExpr{
				pos: position{line: 93, col: 10, offset: 1993},
				run: (*parser).callonBLOCK1,
				expr: &seqExpr{
					pos: position{line: 93, col: 10, offset: 1993},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 93, col: 10, offset: 1993},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 93, col: 12, offset: 1995},
							name: "LC",
						},
						&ruleRefExpr{
							pos:  position{line: 93, col: 15, offset: 1998},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 93, col: 17, offset: 2000},
							label: "s",
							expr: &zeroOrMoreExpr{
								pos: position{line: 93, col: 19, offset: 2002},
								expr: &choiceExpr{
									pos: position{line: 93, col: 21, offset: 2004},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 93, col: 21, offset: 2004},
											name: "CONTINUE",
										},
										&ruleRefExpr{
											pos:  position{line: 93, col: 32, offset: 2015},
											name: "BREAK",
										},
										&ruleRefExpr{
											pos:  position{line: 93, col: 40, offset: 2023},
											name: "NODE",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 93, col: 48, offset: 2031},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 93, col: 50, offset: 2033},
							name: "RC",
						},
					},
				},
			},
		},
		{
			name: "FOR",
			pos:  position{line: 97, col: 1, offset: 2065},
			expr: &actionExpr{
				pos: position{line: 97, col: 8, offset: 2072},
				run: (*parser).callonFOR1,
				expr: &seqExpr{
					pos: position{line: 97, col: 8, offset: 2072},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 97, col: 8, offset: 2072},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 97, col: 10, offset: 2074},
							val:        "for ",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 97, col: 17, offset: 2081},
							label: "ax",
							expr: &seqExpr{
								pos: position{line: 97, col: 22, offset: 2086},
								exprs: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 97, col: 22, offset: 2086},
										name: "LP",
									},
									&oneOrMoreExpr{
										pos: position{line: 97, col: 25, offset: 2089},
										expr: &ruleRefExpr{
											pos:  position{line: 97, col: 25, offset: 2089},
											name: "ARG",
										},
									},
									&ruleRefExpr{
										pos:  position{line: 97, col: 30, offset: 2094},
										name: "RP",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 97, col: 35, offset: 2099},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 97, col: 37, offset: 2101},
							val:        "in",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 97, col: 42, offset: 2106},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 97, col: 44, offset: 2108},
							label: "i",
							expr: &choiceExpr{
								pos: position{line: 97, col: 47, offset: 2111},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 97, col: 47, offset: 2111},
										name: "IDENT",
									},
									&ruleRefExpr{
										pos:  position{line: 97, col: 55, offset: 2119},
										name: "ARRAY",
									},
									&ruleRefExpr{
										pos:  position{line: 97, col: 63, offset: 2127},
										name: "MAP",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 97, col: 69, offset: 2133},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 97, col: 71, offset: 2135},
							label: "s",
							expr: &ruleRefExpr{
								pos:  position{line: 97, col: 73, offset: 2137},
								name: "BLOCK",
							},
						},
					},
				},
			},
		},
		{
			name: "RANGE",
			pos:  position{line: 101, col: 1, offset: 2177},
			expr: &actionExpr{
				pos: position{line: 101, col: 10, offset: 2186},
				run: (*parser).callonRANGE1,
				expr: &seqExpr{
					pos: position{line: 101, col: 10, offset: 2186},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 101, col: 10, offset: 2186},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 101, col: 12, offset: 2188},
							val:        "for ",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 101, col: 19, offset: 2195},
							label: "ax",
							expr: &oneOrMoreExpr{
								pos: position{line: 101, col: 24, offset: 2200},
								expr: &ruleRefExpr{
									pos:  position{line: 101, col: 24, offset: 2200},
									name: "ARG",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 101, col: 30, offset: 2206},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 101, col: 32, offset: 2208},
							val:        ":=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 101, col: 37, offset: 2213},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 101, col: 39, offset: 2215},
							val:        "range",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 101, col: 47, offset: 2223},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 101, col: 49, offset: 2225},
							label: "i",
							expr: &choiceExpr{
								pos: position{line: 101, col: 52, offset: 2228},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 101, col: 52, offset: 2228},
										name: "IDENT",
									},
									&ruleRefExpr{
										pos:  position{line: 101, col: 60, offset: 2236},
										name: "ARRAY",
									},
									&ruleRefExpr{
										pos:  position{line: 101, col: 68, offset: 2244},
										name: "MAP",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 101, col: 74, offset: 2250},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 101, col: 76, offset: 2252},
							label: "s",
							expr: &ruleRefExpr{
								pos:  position{line: 101, col: 78, offset: 2254},
								name: "BLOCK",
							},
						},
					},
				},
			},
		},
		{
			name: "IFOR",
			pos:  position{line: 105, col: 1, offset: 2296},
			expr: &actionExpr{
				pos: position{line: 105, col: 9, offset: 2304},
				run: (*parser).callonIFOR1,
				expr: &seqExpr{
					pos: position{line: 105, col: 9, offset: 2304},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 105, col: 9, offset: 2304},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 105, col: 11, offset: 2306},
							val:        "for",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 105, col: 17, offset: 2312},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 105, col: 19, offset: 2314},
							label: "s",
							expr: &ruleRefExpr{
								pos:  position{line: 105, col: 21, offset: 2316},
								name: "BLOCK",
							},
						},
					},
				},
			},
		},
		{
			name: "ARRAY",
			pos:  position{line: 109, col: 1, offset: 2359},
			expr: &actionExpr{
				pos: position{line: 109, col: 10, offset: 2368},
				run: (*parser).callonARRAY1,
				expr: &seqExpr{
					pos: position{line: 109, col: 10, offset: 2368},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 109, col: 10, offset: 2368},
							name: "LB",
						},
						&ruleRefExpr{
							pos:  position{line: 109, col: 13, offset: 2371},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 109, col: 15, offset: 2373},
							label: "ax",
							expr: &zeroOrMoreExpr{
								pos: position{line: 109, col: 18, offset: 2376},
								expr: &ruleRefExpr{
									pos:  position{line: 109, col: 18, offset: 2376},
									name: "ARG",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 109, col: 23, offset: 2381},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 109, col: 25, offset: 2383},
							name: "RB",
						},
					},
				},
			},
		},
		{
			name: "ACCESS",
			pos:  position{line: 113, col: 1, offset: 2416},
			expr: &actionExpr{
				pos: position{line: 113, col: 11, offset: 2426},
				run: (*parser).callonACCESS1,
				expr: &seqExpr{
					pos: position{line: 113, col: 11, offset: 2426},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 113, col: 11, offset: 2426},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 113, col: 14, offset: 2429},
								name: "IDENT",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 113, col: 21, offset: 2436},
							name: "LB",
						},
						&labeledExpr{
							pos:   position{line: 113, col: 24, offset: 2439},
							label: "key",
							expr: &choiceExpr{
								pos: position{line: 113, col: 29, offset: 2444},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 113, col: 29, offset: 2444},
										name: "NUMBER",
									},
									&ruleRefExpr{
										pos:  position{line: 113, col: 36, offset: 2451},
										name: "IDENT",
									},
									&ruleRefExpr{
										pos:  position{line: 113, col: 42, offset: 2457},
										name: "STRING",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 113, col: 50, offset: 2465},
							name: "RB",
						},
					},
				},
			},
		},
		{
			name: "RETURN",
			pos:  position{line: 117, col: 1, offset: 2503},
			expr: &actionExpr{
				pos: position{line: 117, col: 11, offset: 2513},
				run: (*parser).callonRETURN1,
				expr: &seqExpr{
					pos: position{line: 117, col: 11, offset: 2513},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 117, col: 11, offset: 2513},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 117, col: 13, offset: 2515},
							val:        "return",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 117, col: 22, offset: 2524},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 117, col: 24, offset: 2526},
							label: "s",
							expr: &zeroOrMoreExpr{
								pos: position{line: 117, col: 26, offset: 2528},
								expr: &ruleRefExpr{
									pos:  position{line: 117, col: 26, offset: 2528},
									name: "ARG",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "FARG",
			pos:  position{line: 121, col: 1, offset: 2563},
			expr: &actionExpr{
				pos: position{line: 121, col: 9, offset: 2571},
				run: (*parser).callonFARG1,
				expr: &seqExpr{
					pos: position{line: 121, col: 9, offset: 2571},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 121, col: 9, offset: 2571},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 121, col: 11, offset: 2573},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 121, col: 14, offset: 2576},
								name: "IDENT",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 121, col: 21, offset: 2583},
							name: "_",
						},
						&zeroOrOneExpr{
							pos: position{line: 121, col: 23, offset: 2585},
							expr: &litMatcher{
								pos:        position{line: 121, col: 23, offset: 2585},
								val:        ",",
								ignoreCase: false,
							},
						},
					},
				},
			},
		},
		{
			name: "FUNC",
			pos:  position{line: 125, col: 1, offset: 2611},
			expr: &actionExpr{
				pos: position{line: 125, col: 9, offset: 2619},
				run: (*parser).callonFUNC1,
				expr: &seqExpr{
					pos: position{line: 125, col: 9, offset: 2619},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 125, col: 9, offset: 2619},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 125, col: 11, offset: 2621},
							val:        "func",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 125, col: 18, offset: 2628},
							name: "LP",
						},
						&labeledExpr{
							pos:   position{line: 125, col: 21, offset: 2631},
							label: "ax",
							expr: &zeroOrMoreExpr{
								pos: position{line: 125, col: 24, offset: 2634},
								expr: &ruleRefExpr{
									pos:  position{line: 125, col: 24, offset: 2634},
									name: "FARG",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 125, col: 30, offset: 2640},
							name: "RP",
						},
						&ruleRefExpr{
							pos:  position{line: 125, col: 33, offset: 2643},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 125, col: 35, offset: 2645},
							label: "s",
							expr: &ruleRefExpr{
								pos:  position{line: 125, col: 37, offset: 2647},
								name: "BLOCK",
							},
						},
					},
				},
			},
		},
		{
			name: "FCALL",
			pos:  position{line: 129, col: 1, offset: 2685},
			expr: &actionExpr{
				pos: position{line: 129, col: 10, offset: 2694},
				run: (*parser).callonFCALL1,
				expr: &seqExpr{
					pos: position{line: 129, col: 10, offset: 2694},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 129, col: 10, offset: 2694},
							label: "f",
							expr: &ruleRefExpr{
								pos:  position{line: 129, col: 12, offset: 2696},
								name: "FUNC",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 129, col: 17, offset: 2701},
							name: "LP",
						},
						&labeledExpr{
							pos:   position{line: 129, col: 20, offset: 2704},
							label: "ax",
							expr: &zeroOrMoreExpr{
								pos: position{line: 129, col: 23, offset: 2707},
								expr: &ruleRefExpr{
									pos:  position{line: 129, col: 23, offset: 2707},
									name: "ARG",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 129, col: 28, offset: 2712},
							name: "RP",
						},
						&ruleRefExpr{
							pos:  position{line: 129, col: 31, offset: 2715},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 129, col: 33, offset: 2717},
							label: "b",
							expr: &zeroOrOneExpr{
								pos: position{line: 129, col: 35, offset: 2719},
								expr: &ruleRefExpr{
									pos:  position{line: 129, col: 36, offset: 2720},
									name: "BLOCK",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "GO",
			pos:  position{line: 133, col: 1, offset: 2768},
			expr: &actionExpr{
				pos: position{line: 133, col: 7, offset: 2774},
				run: (*parser).callonGO1,
				expr: &seqExpr{
					pos: position{line: 133, col: 7, offset: 2774},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 133, col: 7, offset: 2774},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 133, col: 9, offset: 2776},
							val:        "go",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 133, col: 14, offset: 2781},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 133, col: 16, offset: 2783},
							label: "s",
							expr: &choiceExpr{
								pos: position{line: 133, col: 20, offset: 2787},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 133, col: 20, offset: 2787},
										name: "FCALL",
									},
									&ruleRefExpr{
										pos:  position{line: 133, col: 28, offset: 2795},
										name: "CALL",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 133, col: 34, offset: 2801},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "CALL",
			pos:  position{line: 141, col: 1, offset: 2936},
			expr: &actionExpr{
				pos: position{line: 141, col: 9, offset: 2944},
				run: (*parser).callonCALL1,
				expr: &seqExpr{
					pos: position{line: 141, col: 9, offset: 2944},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 141, col: 9, offset: 2944},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 141, col: 11, offset: 2946},
							label: "i",
							expr: &choiceExpr{
								pos: position{line: 141, col: 15, offset: 2950},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 141, col: 15, offset: 2950},
										name: "IDENT",
									},
									&ruleRefExpr{
										pos:  position{line: 141, col: 21, offset: 2956},
										name: "CURRENT",
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 141, col: 32, offset: 2967},
							expr: &litMatcher{
								pos:        position{line: 141, col: 32, offset: 2967},
								val:        ".",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 141, col: 38, offset: 2973},
							label: "y",
							expr: &zeroOrOneExpr{
								pos: position{line: 141, col: 40, offset: 2975},
								expr: &ruleRefExpr{
									pos:  position{line: 141, col: 41, offset: 2976},
									name: "IDENT",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 141, col: 49, offset: 2984},
							name: "LP",
						},
						&labeledExpr{
							pos:   position{line: 141, col: 52, offset: 2987},
							label: "ax",
							expr: &zeroOrMoreExpr{
								pos: position{line: 141, col: 55, offset: 2990},
								expr: &ruleRefExpr{
									pos:  position{line: 141, col: 55, offset: 2990},
									name: "ARG",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 141, col: 60, offset: 2995},
							name: "RP",
						},
						&ruleRefExpr{
							pos:  position{line: 141, col: 63, offset: 2998},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 141, col: 65, offset: 3000},
							label: "b",
							expr: &zeroOrOneExpr{
								pos: position{line: 141, col: 67, offset: 3002},
								expr: &ruleRefExpr{
									pos:  position{line: 141, col: 68, offset: 3003},
									name: "BLOCK",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "BOOL",
			pos:  position{line: 145, col: 1, offset: 3049},
			expr: &actionExpr{
				pos: position{line: 145, col: 9, offset: 3057},
				run: (*parser).callonBOOL1,
				expr: &choiceExpr{
					pos: position{line: 145, col: 10, offset: 3058},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 145, col: 10, offset: 3058},
							val:        "false",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 145, col: 20, offset: 3068},
							val:        "true",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "STRING",
			pos:  position{line: 149, col: 1, offset: 3109},
			expr: &actionExpr{
				pos: position{line: 149, col: 11, offset: 3119},
				run: (*parser).callonSTRING1,
				expr: &labeledExpr{
					pos:   position{line: 149, col: 11, offset: 3119},
					label: "s",
					expr: &choiceExpr{
						pos: position{line: 149, col: 14, offset: 3122},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 149, col: 14, offset: 3122},
								name: "MLSTRING",
							},
							&ruleRefExpr{
								pos:  position{line: 149, col: 25, offset: 3133},
								name: "DQSTRING",
							},
						},
					},
				},
			},
		},
		{
			name: "MLSTRING",
			pos:  position{line: 153, col: 1, offset: 3165},
			expr: &actionExpr{
				pos: position{line: 153, col: 13, offset: 3177},
				run: (*parser).callonMLSTRING1,
				expr: &seqExpr{
					pos: position{line: 153, col: 13, offset: 3177},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 153, col: 13, offset: 3177},
							val:        "`",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 153, col: 17, offset: 3181},
							expr: &charClassMatcher{
								pos:        position{line: 153, col: 17, offset: 3181},
								val:        "[^`]",
								chars:      []rune{'`'},
								ignoreCase: false,
								inverted:   true,
							},
						},
						&litMatcher{
							pos:        position{line: 153, col: 23, offset: 3187},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "DQSTRING",
			pos:  position{line: 157, col: 1, offset: 3218},
			expr: &actionExpr{
				pos: position{line: 157, col: 13, offset: 3230},
				run: (*parser).callonDQSTRING1,
				expr: &seqExpr{
					pos: position{line: 157, col: 13, offset: 3230},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 157, col: 13, offset: 3230},
							val:        "\"",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 157, col: 17, offset: 3234},
							expr: &choiceExpr{
								pos: position{line: 157, col: 19, offset: 3236},
								alternatives: []interface{}{
									&seqExpr{
										pos: position{line: 157, col: 19, offset: 3236},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 157, col: 19, offset: 3236},
												expr: &ruleRefExpr{
													pos:  position{line: 157, col: 20, offset: 3237},
													name: "DQEscapeChar",
												},
											},
											&anyMatcher{
												line: 157, col: 33, offset: 3250,
											},
										},
									},
									&seqExpr{
										pos: position{line: 157, col: 37, offset: 3254},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 157, col: 37, offset: 3254},
												val:        "\\",
												ignoreCase: false,
											},
											&ruleRefExpr{
												pos:  position{line: 157, col: 42, offset: 3259},
												name: "DQEscapeSequence",
											},
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 157, col: 62, offset: 3279},
							val:        "\"",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "BREAK",
			pos:  position{line: 161, col: 1, offset: 3310},
			expr: &actionExpr{
				pos: position{line: 161, col: 10, offset: 3319},
				run: (*parser).callonBREAK1,
				expr: &seqExpr{
					pos: position{line: 161, col: 10, offset: 3319},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 161, col: 10, offset: 3319},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 161, col: 12, offset: 3321},
							val:        "break",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 161, col: 20, offset: 3329},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "CONTINUE",
			pos:  position{line: 165, col: 1, offset: 3375},
			expr: &actionExpr{
				pos: position{line: 165, col: 13, offset: 3387},
				run: (*parser).callonCONTINUE1,
				expr: &seqExpr{
					pos: position{line: 165, col: 13, offset: 3387},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 165, col: 13, offset: 3387},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 165, col: 15, offset: 3389},
							val:        "continue",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 165, col: 26, offset: 3400},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "NULL",
			pos:  position{line: 169, col: 1, offset: 3449},
			expr: &actionExpr{
				pos: position{line: 169, col: 8, offset: 3458},
				run: (*parser).callonNULL1,
				expr: &litMatcher{
					pos:        position{line: 169, col: 10, offset: 3460},
					val:        "nil",
					ignoreCase: false,
				},
			},
		},
		{
			name: "MAP",
			pos:  position{line: 173, col: 1, offset: 3510},
			expr: &actionExpr{
				pos: position{line: 173, col: 8, offset: 3517},
				run: (*parser).callonMAP1,
				expr: &seqExpr{
					pos: position{line: 173, col: 8, offset: 3517},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 173, col: 8, offset: 3517},
							name: "LC",
						},
						&ruleRefExpr{
							pos:  position{line: 173, col: 11, offset: 3520},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 173, col: 13, offset: 3522},
							label: "vals",
							expr: &zeroOrMoreExpr{
								pos: position{line: 173, col: 18, offset: 3527},
								expr: &seqExpr{
									pos: position{line: 173, col: 20, offset: 3529},
									exprs: []interface{}{
										&choiceExpr{
											pos: position{line: 173, col: 21, offset: 3530},
											alternatives: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 173, col: 21, offset: 3530},
													name: "IDENT",
												},
												&ruleRefExpr{
													pos:  position{line: 173, col: 27, offset: 3536},
													name: "STRING",
												},
											},
										},
										&ruleRefExpr{
											pos:  position{line: 173, col: 35, offset: 3544},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 173, col: 37, offset: 3546},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 173, col: 41, offset: 3550},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 173, col: 44, offset: 3553},
											name: "ARG",
										},
										&ruleRefExpr{
											pos:  position{line: 173, col: 49, offset: 3558},
											name: "_",
										},
										&zeroOrMoreExpr{
											pos: position{line: 173, col: 51, offset: 3560},
											expr: &litMatcher{
												pos:        position{line: 173, col: 52, offset: 3561},
												val:        ",",
												ignoreCase: false,
											},
										},
										&ruleRefExpr{
											pos:  position{line: 173, col: 58, offset: 3567},
											name: "_",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 173, col: 64, offset: 3573},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 173, col: 66, offset: 3575},
							name: "RC",
						},
					},
				},
			},
		},
		{
			name: "COMMENT",
			pos:  position{line: 177, col: 1, offset: 3608},
			expr: &actionExpr{
				pos: position{line: 177, col: 12, offset: 3619},
				run: (*parser).callonCOMMENT1,
				expr: &seqExpr{
					pos: position{line: 177, col: 12, offset: 3619},
					exprs: []interface{}{
						&choiceExpr{
							pos: position{line: 177, col: 13, offset: 3620},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 177, col: 13, offset: 3620},
									val:        "#",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 177, col: 19, offset: 3626},
									val:        "//",
									ignoreCase: false,
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 177, col: 26, offset: 3633},
							label: "a",
							expr: &zeroOrMoreExpr{
								pos: position{line: 177, col: 28, offset: 3635},
								expr: &charClassMatcher{
									pos:        position{line: 177, col: 28, offset: 3635},
									val:        "[^\\n\\r]",
									chars:      []rune{'\n', '\r'},
									ignoreCase: false,
									inverted:   true,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "SHEBANG",
			pos:  position{line: 181, col: 1, offset: 3680},
			expr: &actionExpr{
				pos: position{line: 181, col: 12, offset: 3691},
				run: (*parser).callonSHEBANG1,
				expr: &seqExpr{
					pos: position{line: 181, col: 12, offset: 3691},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 181, col: 12, offset: 3691},
							val:        "#!",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 181, col: 17, offset: 3696},
							label: "a",
							expr: &zeroOrMoreExpr{
								pos: position{line: 181, col: 19, offset: 3698},
								expr: &charClassMatcher{
									pos:        position{line: 181, col: 19, offset: 3698},
									val:        "[^\\n\\r]",
									chars:      []rune{'\n', '\r'},
									ignoreCase: false,
									inverted:   true,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "LP",
			pos:  position{line: 185, col: 1, offset: 3744},
			expr: &actionExpr{
				pos: position{line: 185, col: 7, offset: 3750},
				run: (*parser).callonLP1,
				expr: &litMatcher{
					pos:        position{line: 185, col: 7, offset: 3750},
					val:        "(",
					ignoreCase: false,
				},
			},
		},
		{
			name: "RP",
			pos:  position{line: 186, col: 1, offset: 3775},
			expr: &actionExpr{
				pos: position{line: 186, col: 7, offset: 3781},
				run: (*parser).callonRP1,
				expr: &litMatcher{
					pos:        position{line: 186, col: 7, offset: 3781},
					val:        ")",
					ignoreCase: false,
				},
			},
		},
		{
			name: "LB",
			pos:  position{line: 187, col: 1, offset: 3806},
			expr: &actionExpr{
				pos: position{line: 187, col: 7, offset: 3812},
				run: (*parser).callonLB1,
				expr: &litMatcher{
					pos:        position{line: 187, col: 7, offset: 3812},
					val:        "[",
					ignoreCase: false,
				},
			},
		},
		{
			name: "RB",
			pos:  position{line: 188, col: 1, offset: 3837},
			expr: &actionExpr{
				pos: position{line: 188, col: 7, offset: 3843},
				run: (*parser).callonRB1,
				expr: &litMatcher{
					pos:        position{line: 188, col: 7, offset: 3843},
					val:        "]",
					ignoreCase: false,
				},
			},
		},
		{
			name: "LC",
			pos:  position{line: 189, col: 1, offset: 3868},
			expr: &actionExpr{
				pos: position{line: 189, col: 7, offset: 3874},
				run: (*parser).callonLC1,
				expr: &litMatcher{
					pos:        position{line: 189, col: 7, offset: 3874},
					val:        "{",
					ignoreCase: false,
				},
			},
		},
		{
			name: "RC",
			pos:  position{line: 190, col: 1, offset: 3899},
			expr: &actionExpr{
				pos: position{line: 190, col: 7, offset: 3905},
				run: (*parser).callonRC1,
				expr: &litMatcher{
					pos:        position{line: 190, col: 7, offset: 3905},
					val:        "}",
					ignoreCase: false,
				},
			},
		},
		{
			name: "MLEscapeChar",
			pos:  position{line: 192, col: 1, offset: 3931},
			expr: &charClassMatcher{
				pos:        position{line: 192, col: 17, offset: 3947},
				val:        "[\\x00-\\x1f\\\\`]",
				chars:      []rune{'\\', '`'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "MLEscapeSequence",
			pos:  position{line: 193, col: 1, offset: 3962},
			expr: &choiceExpr{
				pos: position{line: 193, col: 21, offset: 3982},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 193, col: 21, offset: 3982},
						name: "MLSingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 193, col: 42, offset: 4003},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "MLSingleCharEscape",
			pos:  position{line: 194, col: 1, offset: 4017},
			expr: &charClassMatcher{
				pos:        position{line: 194, col: 23, offset: 4039},
				val:        "['\\\\/bfnrt]",
				chars:      []rune{'\'', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DQEscapeChar",
			pos:  position{line: 196, col: 1, offset: 4052},
			expr: &charClassMatcher{
				pos:        position{line: 196, col: 17, offset: 4068},
				val:        "[\\x00-\\x1f\\\\\"]",
				chars:      []rune{'\\', '"'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DQEscapeSequence",
			pos:  position{line: 197, col: 1, offset: 4083},
			expr: &choiceExpr{
				pos: position{line: 197, col: 21, offset: 4103},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 197, col: 21, offset: 4103},
						name: "DQSingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 197, col: 42, offset: 4124},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "DQSingleCharEscape",
			pos:  position{line: 198, col: 1, offset: 4138},
			expr: &charClassMatcher{
				pos:        position{line: 198, col: 23, offset: 4160},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "UnicodeEscape",
			pos:  position{line: 200, col: 1, offset: 4173},
			expr: &seqExpr{
				pos: position{line: 200, col: 17, offset: 4191},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 200, col: 17, offset: 4191},
						val:        "u",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 200, col: 21, offset: 4195},
						name: "HexDigit",
					},
					&ruleRefExpr{
						pos:  position{line: 200, col: 30, offset: 4204},
						name: "HexDigit",
					},
					&ruleRefExpr{
						pos:  position{line: 200, col: 39, offset: 4213},
						name: "HexDigit",
					},
					&ruleRefExpr{
						pos:  position{line: 200, col: 48, offset: 4222},
						name: "HexDigit",
					},
				},
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 202, col: 1, offset: 4232},
			expr: &charClassMatcher{
				pos:        position{line: 202, col: 16, offset: 4249},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 204, col: 1, offset: 4256},
			expr: &charClassMatcher{
				pos:        position{line: 204, col: 23, offset: 4280},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 206, col: 1, offset: 4287},
			expr: &charClassMatcher{
				pos:        position{line: 206, col: 12, offset: 4300},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "_",
			displayName: "\"whitespace\"",
			pos:         position{line: 208, col: 1, offset: 4311},
			expr: &actionExpr{
				pos: position{line: 208, col: 18, offset: 4330},
				run: (*parser).callon_1,
				expr: &zeroOrMoreExpr{
					pos: position{line: 208, col: 18, offset: 4330},
					expr: &charClassMatcher{
						pos:        position{line: 208, col: 18, offset: 4330},
						val:        "[ \\t\\r\\n]",
						chars:      []rune{' ', '\t', '\r', '\n'},
						ignoreCase: false,
						inverted:   false,
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 212, col: 1, offset: 4366},
			expr: &actionExpr{
				pos: position{line: 212, col: 8, offset: 4373},
				run: (*parser).callonEOF1,
				expr: &notExpr{
					pos: position{line: 212, col: 8, offset: 4373},
					expr: &anyMatcher{
						line: 212, col: 9, offset: 4374,
					},
				},
			},
		},
	},
}

func (c *current) onDOC1(b, stmts interface{}) (interface{}, error) {
	s, err := ast.NewNodes(stmts)
	if b != nil {
		bang := b.(ast.Shebang)
		s = append(ast.Nodes{bang}, s...)
	}
	return ast.Script{Nodes: s}, err
}

func (p *parser) callonDOC1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDOC1(stack["b"], stack["stmts"])
}

func (c *current) onNODE1(s interface{}) (interface{}, error) {
	return s, nil
}

func (p *parser) callonNODE1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onNODE1(stack["s"])
}

func (c *current) onIMPORT1(s interface{}) (interface{}, error) {
	return newImport(c, s)
}

func (p *parser) callonIMPORT1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIMPORT1(stack["s"])
}

func (c *current) onCURRENT1() (interface{}, error) {
	return ast.Current{Meta: meta(c)}, nil
}

func (p *parser) callonCURRENT1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCURRENT1()
}

func (c *current) onOP1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonOP1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOP1()
}

func (c *current) onOPEX1(a, op, b interface{}) (interface{}, error) {
	return newOpExpression(c, a, op, b)
}

func (p *parser) callonOPEX1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOPEX1(stack["a"], stack["op"], stack["b"])
}

func (c *current) onPOPEX1(a, op, b interface{}) (interface{}, error) {
	return newPopExpression(c, a, op, b)
}

func (p *parser) callonPOPEX1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPOPEX1(stack["a"], stack["op"], stack["b"])
}

func (c *current) onIF1(p, s, b, e interface{}) (interface{}, error) {
	return newIf(c, p, s, b, e)
}

func (p *parser) callonIF1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIF1(stack["p"], stack["s"], stack["b"], stack["e"])
}

func (c *current) onELSE1(b interface{}) (interface{}, error) {
	return newElse(c, b)
}

func (p *parser) callonELSE1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onELSE1(stack["b"])
}

func (c *current) onELSEIF1(i interface{}) (interface{}, error) {
	return newElseIf(c, i)
}

func (p *parser) callonELSEIF1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onELSEIF1(stack["i"])
}

func (c *current) onINTEGER1() (interface{}, error) {
	return newInteger(c, c.text)
}

func (p *parser) callonINTEGER1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onINTEGER1()
}

func (c *current) onFLOAT1() (interface{}, error) {
	return newFloat(c, c.text)
}

func (p *parser) callonFLOAT1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFLOAT1()
}

func (c *current) onIDENT1() (interface{}, error) {
	return newIdent(c, c.text)
}

func (p *parser) callonIDENT1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIDENT1()
}

func (c *current) onLET1(name, value interface{}) (interface{}, error) {
	return newLet(c, name, value)
}

func (p *parser) callonLET1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLET1(stack["name"], stack["value"])
}

func (c *current) onVAR1(name, value interface{}) (interface{}, error) {
	return newVar(c, name, value)
}

func (p *parser) callonVAR1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onVAR1(stack["name"], stack["value"])
}

func (c *current) onASSIGN1(name, value interface{}) (interface{}, error) {
	return newAssign(c, name, value)
}

func (p *parser) callonASSIGN1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onASSIGN1(stack["name"], stack["value"])
}

func (c *current) onARG1(i interface{}) (interface{}, error) {
	return i, nil
}

func (p *parser) callonARG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onARG1(stack["i"])
}

func (c *current) onBLOCK1(s interface{}) (interface{}, error) {
	return newBlock(c, s)
}

func (p *parser) callonBLOCK1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBLOCK1(stack["s"])
}

func (c *current) onFOR1(ax, i, s interface{}) (interface{}, error) {
	return newFor(c, i, ax, s)
}

func (p *parser) callonFOR1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFOR1(stack["ax"], stack["i"], stack["s"])
}

func (c *current) onRANGE1(ax, i, s interface{}) (interface{}, error) {
	return newRange(c, i, ax, s)
}

func (p *parser) callonRANGE1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onRANGE1(stack["ax"], stack["i"], stack["s"])
}

func (c *current) onIFOR1(s interface{}) (interface{}, error) {
	return newFor(c, nil, nil, s)
}

func (p *parser) callonIFOR1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIFOR1(stack["s"])
}

func (c *current) onARRAY1(ax interface{}) (interface{}, error) {
	return newArray(c, ax)
}

func (p *parser) callonARRAY1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onARRAY1(stack["ax"])
}

func (c *current) onACCESS1(i, key interface{}) (interface{}, error) {
	return newAccess(c, i, key)
}

func (p *parser) callonACCESS1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onACCESS1(stack["i"], stack["key"])
}

func (c *current) onRETURN1(s interface{}) (interface{}, error) {
	return newReturn(c, s)
}

func (p *parser) callonRETURN1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onRETURN1(stack["s"])
}

func (c *current) onFARG1(i interface{}) (interface{}, error) {
	return i, nil
}

func (p *parser) callonFARG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFARG1(stack["i"])
}

func (c *current) onFUNC1(ax, s interface{}) (interface{}, error) {
	return newFunc(c, ax, s)
}

func (p *parser) callonFUNC1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFUNC1(stack["ax"], stack["s"])
}

func (c *current) onFCALL1(f, ax, b interface{}) (interface{}, error) {
	return newCall(c, f, nil, ax, b)
}

func (p *parser) callonFCALL1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFCALL1(stack["f"], stack["ax"], stack["b"])
}

func (c *current) onGO1(s interface{}) (interface{}, error) {
	ca, ok := s.(ast.Call)
	if !ok {
		return nil, fmt.Errorf("expected ast.Call got %T", s)
	}
	return ast.NewGoroutine(ca)
}

func (p *parser) callonGO1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onGO1(stack["s"])
}

func (c *current) onCALL1(i, y, ax, b interface{}) (interface{}, error) {
	return newCall(c, i, y, ax, b)
}

func (p *parser) callonCALL1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCALL1(stack["i"], stack["y"], stack["ax"], stack["b"])
}

func (c *current) onBOOL1() (interface{}, error) {
	return newBool(c, c.text)
}

func (p *parser) callonBOOL1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBOOL1()
}

func (c *current) onSTRING1(s interface{}) (interface{}, error) {
	return s, nil
}

func (p *parser) callonSTRING1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSTRING1(stack["s"])
}

func (c *current) onMLSTRING1() (interface{}, error) {
	return newString(c)
}

func (p *parser) callonMLSTRING1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMLSTRING1()
}

func (c *current) onDQSTRING1() (interface{}, error) {
	return newString(c)
}

func (p *parser) callonDQSTRING1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDQSTRING1()
}

func (c *current) onBREAK1() (interface{}, error) {
	return ast.Break{Meta: meta(c)}, nil
}

func (p *parser) callonBREAK1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBREAK1()
}

func (c *current) onCONTINUE1() (interface{}, error) {
	return ast.Continue{Meta: meta(c)}, nil
}

func (p *parser) callonCONTINUE1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCONTINUE1()
}

func (c *current) onNULL1() (interface{}, error) {
	return ast.Nil{Meta: meta(c)}, nil
}

func (p *parser) callonNULL1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onNULL1()
}

func (c *current) onMAP1(vals interface{}) (interface{}, error) {
	return newMap(c, vals)
}

func (p *parser) callonMAP1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMAP1(stack["vals"])
}

func (c *current) onCOMMENT1(a interface{}) (interface{}, error) {
	return newComment(c, c.text)
}

func (p *parser) callonCOMMENT1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCOMMENT1(stack["a"])
}

func (c *current) onSHEBANG1(a interface{}) (interface{}, error) {
	return ast.NewShebang(c.text)
}

func (p *parser) callonSHEBANG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSHEBANG1(stack["a"])
}

func (c *current) onLP1() (interface{}, error) {
	return newNoop(c)
}

func (p *parser) callonLP1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLP1()
}

func (c *current) onRP1() (interface{}, error) {
	return newNoop(c)
}

func (p *parser) callonRP1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onRP1()
}

func (c *current) onLB1() (interface{}, error) {
	return newNoop(c)
}

func (p *parser) callonLB1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLB1()
}

func (c *current) onRB1() (interface{}, error) {
	return newNoop(c)
}

func (p *parser) callonRB1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onRB1()
}

func (c *current) onLC1() (interface{}, error) {
	return newNoop(c)
}

func (p *parser) callonLC1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLC1()
}

func (c *current) onRC1() (interface{}, error) {
	return newNoop(c)
}

func (p *parser) callonRC1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onRC1()
}

func (c *current) on_1() (interface{}, error) {
	return newNoop(c)
}

func (p *parser) callon_1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.on_1()
}

func (c *current) onEOF1() (interface{}, error) {
	return ast.Noop{}, nil
}

func (p *parser) callonEOF1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEOF1()
}

var (
	// errNoRule is returned when the grammar to parse has no rule.
	errNoRule = errors.New("grammar has no rule")

	// errInvalidEncoding is returned when the source is not properly
	// utf8-encoded.
	errInvalidEncoding = errors.New("invalid encoding")

	// errNoMatch is returned if no match could be found.
	errNoMatch = errors.New("no match found")
)

// Option is a function that can set an option on the parser. It returns
// the previous setting as an Option.
type Option func(*parser) Option

// Debug creates an Option to set the debug flag to b. When set to true,
// debugging information is printed to stdout while parsing.
//
// The default is false.
func Debug(b bool) Option {
	return func(p *parser) Option {
		old := p.debug
		p.debug = b
		return Debug(old)
	}
}

// Memoize creates an Option to set the memoize flag to b. When set to true,
// the parser will cache all results so each expression is evaluated only
// once. This guarantees linear parsing time even for pathological cases,
// at the expense of more memory and slower times for typical cases.
//
// The default is false.
func Memoize(b bool) Option {
	return func(p *parser) Option {
		old := p.memoize
		p.memoize = b
		return Memoize(old)
	}
}

// Recover creates an Option to set the recover flag to b. When set to
// true, this causes the parser to recover from panics and convert it
// to an error. Setting it to false can be useful while debugging to
// access the full stack trace.
//
// The default is true.
func Recover(b bool) Option {
	return func(p *parser) Option {
		old := p.recover
		p.recover = b
		return Recover(old)
	}
}

// ParseFile parses the file identified by filename.
func ParseFile(filename string, opts ...Option) (interface{}, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ParseReader(filename, f, opts...)
}

// ParseReader parses the data from r using filename as information in the
// error messages.
func ParseReader(filename string, r io.Reader, opts ...Option) (interface{}, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return Parse(filename, b, opts...)
}

// Parse parses the data from b using filename as information in the
// error messages.
func Parse(filename string, b []byte, opts ...Option) (interface{}, error) {
	return newParser(filename, b, opts...).parse(g)
}

// position records a position in the text.
type position struct {
	line, col, offset int
}

func (p position) String() string {
	return fmt.Sprintf("%d:%d [%d]", p.line, p.col, p.offset)
}

// savepoint stores all state required to go back to this point in the
// parser.
type savepoint struct {
	position
	rn rune
	w  int
}

type current struct {
	pos  position // start position of the match
	text []byte   // raw text of the match
}

// the AST types...

type grammar struct {
	pos   position
	rules []*rule
}

type rule struct {
	pos         position
	name        string
	displayName string
	expr        interface{}
}

type choiceExpr struct {
	pos          position
	alternatives []interface{}
}

type actionExpr struct {
	pos  position
	expr interface{}
	run  func(*parser) (interface{}, error)
}

type seqExpr struct {
	pos   position
	exprs []interface{}
}

type labeledExpr struct {
	pos   position
	label string
	expr  interface{}
}

type expr struct {
	pos  position
	expr interface{}
}

type andExpr expr
type notExpr expr
type zeroOrOneExpr expr
type zeroOrMoreExpr expr
type oneOrMoreExpr expr

type ruleRefExpr struct {
	pos  position
	name string
}

type andCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type notCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type litMatcher struct {
	pos        position
	val        string
	ignoreCase bool
}

type charClassMatcher struct {
	pos        position
	val        string
	chars      []rune
	ranges     []rune
	classes    []*unicode.RangeTable
	ignoreCase bool
	inverted   bool
}

type anyMatcher position

// errList cumulates the errors found by the parser.
type errList []error

func (e *errList) add(err error) {
	*e = append(*e, err)
}

func (e errList) err() error {
	if len(e) == 0 {
		return nil
	}
	e.dedupe()
	return e
}

func (e *errList) dedupe() {
	var cleaned []error
	set := make(map[string]bool)
	for _, err := range *e {
		if msg := err.Error(); !set[msg] {
			set[msg] = true
			cleaned = append(cleaned, err)
		}
	}
	*e = cleaned
}

func (e errList) Error() string {
	switch len(e) {
	case 0:
		return ""
	case 1:
		return e[0].Error()
	default:
		var buf bytes.Buffer

		for i, err := range e {
			if i > 0 {
				buf.WriteRune('\n')
			}
			buf.WriteString(err.Error())
		}
		return buf.String()
	}
}

// parserError wraps an error with a prefix indicating the rule in which
// the error occurred. The original error is stored in the Inner field.
type parserError struct {
	Inner  error
	pos    position
	prefix string
}

// Error returns the error message.
func (p *parserError) Error() string {
	return p.prefix + ": " + p.Inner.Error()
}

// newParser creates a parser with the specified input source and options.
func newParser(filename string, b []byte, opts ...Option) *parser {
	p := &parser{
		filename: filename,
		errs:     new(errList),
		data:     b,
		pt:       savepoint{position: position{line: 1}},
		recover:  true,
	}
	p.setOptions(opts)
	return p
}

// setOptions applies the options to the parser.
func (p *parser) setOptions(opts []Option) {
	for _, opt := range opts {
		opt(p)
	}
}

type resultTuple struct {
	v   interface{}
	b   bool
	end savepoint
}

type parser struct {
	filename string
	pt       savepoint
	cur      current

	data []byte
	errs *errList

	recover bool
	debug   bool
	depth   int

	memoize bool
	// memoization table for the packrat algorithm:
	// map[offset in source] map[expression or rule] {value, match}
	memo map[int]map[interface{}]resultTuple

	// rules table, maps the rule identifier to the rule node
	rules map[string]*rule
	// variables stack, map of label to value
	vstack []map[string]interface{}
	// rule stack, allows identification of the current rule in errors
	rstack []*rule

	// stats
	exprCnt int
}

// push a variable set on the vstack.
func (p *parser) pushV() {
	if cap(p.vstack) == len(p.vstack) {
		// create new empty slot in the stack
		p.vstack = append(p.vstack, nil)
	} else {
		// slice to 1 more
		p.vstack = p.vstack[:len(p.vstack)+1]
	}

	// get the last args set
	m := p.vstack[len(p.vstack)-1]
	if m != nil && len(m) == 0 {
		// empty map, all good
		return
	}

	m = make(map[string]interface{})
	p.vstack[len(p.vstack)-1] = m
}

// pop a variable set from the vstack.
func (p *parser) popV() {
	// if the map is not empty, clear it
	m := p.vstack[len(p.vstack)-1]
	if len(m) > 0 {
		// GC that map
		p.vstack[len(p.vstack)-1] = nil
	}
	p.vstack = p.vstack[:len(p.vstack)-1]
}

func (p *parser) print(prefix, s string) string {
	if !p.debug {
		return s
	}

	fmt.Printf("%s %d:%d:%d: %s [%#U]\n",
		prefix, p.pt.line, p.pt.col, p.pt.offset, s, p.pt.rn)
	return s
}

func (p *parser) in(s string) string {
	p.depth++
	return p.print(strings.Repeat(" ", p.depth)+">", s)
}

func (p *parser) out(s string) string {
	p.depth--
	return p.print(strings.Repeat(" ", p.depth)+"<", s)
}

func (p *parser) addErr(err error) {
	p.addErrAt(err, p.pt.position)
}

func (p *parser) addErrAt(err error, pos position) {
	var buf bytes.Buffer
	if p.filename != "" {
		buf.WriteString(p.filename)
	}
	if buf.Len() > 0 {
		buf.WriteString(":")
	}
	buf.WriteString(fmt.Sprintf("%d:%d (%d)", pos.line, pos.col, pos.offset))
	if len(p.rstack) > 0 {
		if buf.Len() > 0 {
			buf.WriteString(": ")
		}
		rule := p.rstack[len(p.rstack)-1]
		if rule.displayName != "" {
			buf.WriteString("rule " + rule.displayName)
		} else {
			buf.WriteString("rule " + rule.name)
		}
	}
	pe := &parserError{Inner: err, pos: pos, prefix: buf.String()}
	p.errs.add(pe)
}

// read advances the parser to the next rune.
func (p *parser) read() {
	p.pt.offset += p.pt.w
	rn, n := utf8.DecodeRune(p.data[p.pt.offset:])
	p.pt.rn = rn
	p.pt.w = n
	p.pt.col++
	if rn == '\n' {
		p.pt.line++
		p.pt.col = 0
	}

	if rn == utf8.RuneError {
		if n == 1 {
			p.addErr(errInvalidEncoding)
		}
	}
}

// restore parser position to the savepoint pt.
func (p *parser) restore(pt savepoint) {
	if p.debug {
		defer p.out(p.in("restore"))
	}
	if pt.offset == p.pt.offset {
		return
	}
	p.pt = pt
}

// get the slice of bytes from the savepoint start to the current position.
func (p *parser) sliceFrom(start savepoint) []byte {
	return p.data[start.position.offset:p.pt.position.offset]
}

func (p *parser) getMemoized(node interface{}) (resultTuple, bool) {
	if len(p.memo) == 0 {
		return resultTuple{}, false
	}
	m := p.memo[p.pt.offset]
	if len(m) == 0 {
		return resultTuple{}, false
	}
	res, ok := m[node]
	return res, ok
}

func (p *parser) setMemoized(pt savepoint, node interface{}, tuple resultTuple) {
	if p.memo == nil {
		p.memo = make(map[int]map[interface{}]resultTuple)
	}
	m := p.memo[pt.offset]
	if m == nil {
		m = make(map[interface{}]resultTuple)
		p.memo[pt.offset] = m
	}
	m[node] = tuple
}

func (p *parser) buildRulesTable(g *grammar) {
	p.rules = make(map[string]*rule, len(g.rules))
	for _, r := range g.rules {
		p.rules[r.name] = r
	}
}

func (p *parser) parse(g *grammar) (val interface{}, err error) {
	if len(g.rules) == 0 {
		p.addErr(errNoRule)
		return nil, p.errs.err()
	}

	// TODO : not super critical but this could be generated
	p.buildRulesTable(g)

	if p.recover {
		// panic can be used in action code to stop parsing immediately
		// and return the panic as an error.
		defer func() {
			if e := recover(); e != nil {
				if p.debug {
					defer p.out(p.in("panic handler"))
				}
				val = nil
				switch e := e.(type) {
				case error:
					p.addErr(e)
				default:
					p.addErr(fmt.Errorf("%v", e))
				}
				err = p.errs.err()
			}
		}()
	}

	// start rule is rule [0]
	p.read() // advance to first rune
	val, ok := p.parseRule(g.rules[0])
	if !ok {
		if len(*p.errs) == 0 {
			// make sure this doesn't go out silently
			p.addErr(errNoMatch)
		}
		return nil, p.errs.err()
	}
	return val, p.errs.err()
}

func (p *parser) parseRule(rule *rule) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRule " + rule.name))
	}

	if p.memoize {
		res, ok := p.getMemoized(rule)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
	}

	start := p.pt
	p.rstack = append(p.rstack, rule)
	p.pushV()
	val, ok := p.parseExpr(rule.expr)
	p.popV()
	p.rstack = p.rstack[:len(p.rstack)-1]
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}

	if p.memoize {
		p.setMemoized(start, rule, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseExpr(expr interface{}) (interface{}, bool) {
	var pt savepoint
	var ok bool

	if p.memoize {
		res, ok := p.getMemoized(expr)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
		pt = p.pt
	}

	p.exprCnt++
	var val interface{}
	switch expr := expr.(type) {
	case *actionExpr:
		val, ok = p.parseActionExpr(expr)
	case *andCodeExpr:
		val, ok = p.parseAndCodeExpr(expr)
	case *andExpr:
		val, ok = p.parseAndExpr(expr)
	case *anyMatcher:
		val, ok = p.parseAnyMatcher(expr)
	case *charClassMatcher:
		val, ok = p.parseCharClassMatcher(expr)
	case *choiceExpr:
		val, ok = p.parseChoiceExpr(expr)
	case *labeledExpr:
		val, ok = p.parseLabeledExpr(expr)
	case *litMatcher:
		val, ok = p.parseLitMatcher(expr)
	case *notCodeExpr:
		val, ok = p.parseNotCodeExpr(expr)
	case *notExpr:
		val, ok = p.parseNotExpr(expr)
	case *oneOrMoreExpr:
		val, ok = p.parseOneOrMoreExpr(expr)
	case *ruleRefExpr:
		val, ok = p.parseRuleRefExpr(expr)
	case *seqExpr:
		val, ok = p.parseSeqExpr(expr)
	case *zeroOrMoreExpr:
		val, ok = p.parseZeroOrMoreExpr(expr)
	case *zeroOrOneExpr:
		val, ok = p.parseZeroOrOneExpr(expr)
	default:
		panic(fmt.Sprintf("unknown expression type %T", expr))
	}
	if p.memoize {
		p.setMemoized(pt, expr, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseActionExpr(act *actionExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseActionExpr"))
	}

	start := p.pt
	val, ok := p.parseExpr(act.expr)
	if ok {
		p.cur.pos = start.position
		p.cur.text = p.sliceFrom(start)
		actVal, err := act.run(p)
		if err != nil {
			p.addErrAt(err, start.position)
		}
		val = actVal
	}
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}
	return val, ok
}

func (p *parser) parseAndCodeExpr(and *andCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndCodeExpr"))
	}

	ok, err := and.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, ok
}

func (p *parser) parseAndExpr(and *andExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndExpr"))
	}

	pt := p.pt
	p.pushV()
	_, ok := p.parseExpr(and.expr)
	p.popV()
	p.restore(pt)
	return nil, ok
}

func (p *parser) parseAnyMatcher(any *anyMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAnyMatcher"))
	}

	if p.pt.rn != utf8.RuneError {
		start := p.pt
		p.read()
		return p.sliceFrom(start), true
	}
	return nil, false
}

func (p *parser) parseCharClassMatcher(chr *charClassMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseCharClassMatcher"))
	}

	cur := p.pt.rn
	// can't match EOF
	if cur == utf8.RuneError {
		return nil, false
	}
	start := p.pt
	if chr.ignoreCase {
		cur = unicode.ToLower(cur)
	}

	// try to match in the list of available chars
	for _, rn := range chr.chars {
		if rn == cur {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of ranges
	for i := 0; i < len(chr.ranges); i += 2 {
		if cur >= chr.ranges[i] && cur <= chr.ranges[i+1] {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of Unicode classes
	for _, cl := range chr.classes {
		if unicode.Is(cl, cur) {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	if chr.inverted {
		p.read()
		return p.sliceFrom(start), true
	}
	return nil, false
}

func (p *parser) parseChoiceExpr(ch *choiceExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseChoiceExpr"))
	}

	for _, alt := range ch.alternatives {
		p.pushV()
		val, ok := p.parseExpr(alt)
		p.popV()
		if ok {
			return val, ok
		}
	}
	return nil, false
}

func (p *parser) parseLabeledExpr(lab *labeledExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLabeledExpr"))
	}

	p.pushV()
	val, ok := p.parseExpr(lab.expr)
	p.popV()
	if ok && lab.label != "" {
		m := p.vstack[len(p.vstack)-1]
		m[lab.label] = val
	}
	return val, ok
}

func (p *parser) parseLitMatcher(lit *litMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLitMatcher"))
	}

	start := p.pt
	for _, want := range lit.val {
		cur := p.pt.rn
		if lit.ignoreCase {
			cur = unicode.ToLower(cur)
		}
		if cur != want {
			p.restore(start)
			return nil, false
		}
		p.read()
	}
	return p.sliceFrom(start), true
}

func (p *parser) parseNotCodeExpr(not *notCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotCodeExpr"))
	}

	ok, err := not.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, !ok
}

func (p *parser) parseNotExpr(not *notExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotExpr"))
	}

	pt := p.pt
	p.pushV()
	_, ok := p.parseExpr(not.expr)
	p.popV()
	p.restore(pt)
	return nil, !ok
}

func (p *parser) parseOneOrMoreExpr(expr *oneOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseOneOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			if len(vals) == 0 {
				// did not match once, no match
				return nil, false
			}
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseRuleRefExpr(ref *ruleRefExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRuleRefExpr " + ref.name))
	}

	if ref.name == "" {
		panic(fmt.Sprintf("%s: invalid rule: missing name", ref.pos))
	}

	rule := p.rules[ref.name]
	if rule == nil {
		p.addErr(fmt.Errorf("undefined rule: %s", ref.name))
		return nil, false
	}
	return p.parseRule(rule)
}

func (p *parser) parseSeqExpr(seq *seqExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseSeqExpr"))
	}

	var vals []interface{}

	pt := p.pt
	for _, expr := range seq.exprs {
		val, ok := p.parseExpr(expr)
		if !ok {
			p.restore(pt)
			return nil, false
		}
		vals = append(vals, val)
	}
	return vals, true
}

func (p *parser) parseZeroOrMoreExpr(expr *zeroOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseZeroOrOneExpr(expr *zeroOrOneExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrOneExpr"))
	}

	p.pushV()
	val, _ := p.parseExpr(expr.expr)
	p.popV()
	// whether it matched or not, consider it a match
	return val, true
}

func rangeTable(class string) *unicode.RangeTable {
	if rt, ok := unicode.Categories[class]; ok {
		return rt
	}
	if rt, ok := unicode.Properties[class]; ok {
		return rt
	}
	if rt, ok := unicode.Scripts[class]; ok {
		return rt
	}

	// cannot happen
	panic(fmt.Sprintf("invalid Unicode class: %s", class))
}
