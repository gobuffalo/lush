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
														name: "EXPR",
													},
												},
											},
											&ruleRefExpr{
												pos:  position{line: 47, col: 68, offset: 1096},
												name: "_",
											},
											&ruleRefExpr{
												pos:  position{line: 47, col: 70, offset: 1098},
												name: "RP",
											},
										},
									},
									&choiceExpr{
										pos: position{line: 47, col: 76, offset: 1104},
										alternatives: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 47, col: 76, offset: 1104},
												name: "BOOL",
											},
											&ruleRefExpr{
												pos:  position{line: 47, col: 83, offset: 1111},
												name: "EXPR",
											},
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 47, col: 92, offset: 1120},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 47, col: 94, offset: 1122},
							label: "b",
							expr: &ruleRefExpr{
								pos:  position{line: 47, col: 96, offset: 1124},
								name: "BLOCK",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 47, col: 102, offset: 1130},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 47, col: 104, offset: 1132},
							label: "e",
							expr: &zeroOrMoreExpr{
								pos: position{line: 47, col: 106, offset: 1134},
								expr: &choiceExpr{
									pos: position{line: 47, col: 107, offset: 1135},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 47, col: 107, offset: 1135},
											name: "ELSEIF",
										},
										&ruleRefExpr{
											pos:  position{line: 47, col: 116, offset: 1144},
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
			pos:  position{line: 51, col: 1, offset: 1186},
			expr: &actionExpr{
				pos: position{line: 51, col: 9, offset: 1194},
				run: (*parser).callonELSE1,
				expr: &seqExpr{
					pos: position{line: 51, col: 9, offset: 1194},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 51, col: 9, offset: 1194},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 51, col: 11, offset: 1196},
							val:        "else",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 51, col: 18, offset: 1203},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 51, col: 20, offset: 1205},
							label: "b",
							expr: &ruleRefExpr{
								pos:  position{line: 51, col: 22, offset: 1207},
								name: "BLOCK",
							},
						},
					},
				},
			},
		},
		{
			name: "ELSEIF",
			pos:  position{line: 55, col: 1, offset: 1241},
			expr: &actionExpr{
				pos: position{line: 55, col: 11, offset: 1251},
				run: (*parser).callonELSEIF1,
				expr: &seqExpr{
					pos: position{line: 55, col: 11, offset: 1251},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 55, col: 11, offset: 1251},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 55, col: 13, offset: 1253},
							val:        "else",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 55, col: 20, offset: 1260},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 55, col: 22, offset: 1262},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 55, col: 24, offset: 1264},
								name: "IF",
							},
						},
					},
				},
			},
		},
		{
			name: "INTEGER",
			pos:  position{line: 59, col: 1, offset: 1297},
			expr: &actionExpr{
				pos: position{line: 59, col: 12, offset: 1308},
				run: (*parser).callonINTEGER1,
				expr: &seqExpr{
					pos: position{line: 59, col: 12, offset: 1308},
					exprs: []interface{}{
						&zeroOrOneExpr{
							pos: position{line: 59, col: 12, offset: 1308},
							expr: &litMatcher{
								pos:        position{line: 59, col: 12, offset: 1308},
								val:        "-",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 59, col: 17, offset: 1313},
							expr: &charClassMatcher{
								pos:        position{line: 59, col: 17, offset: 1313},
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
			pos:  position{line: 63, col: 1, offset: 1356},
			expr: &actionExpr{
				pos: position{line: 63, col: 10, offset: 1365},
				run: (*parser).callonFLOAT1,
				expr: &seqExpr{
					pos: position{line: 63, col: 10, offset: 1365},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 63, col: 10, offset: 1365},
							name: "INTEGER",
						},
						&zeroOrOneExpr{
							pos: position{line: 63, col: 18, offset: 1373},
							expr: &litMatcher{
								pos:        position{line: 63, col: 18, offset: 1373},
								val:        ".",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 63, col: 23, offset: 1378},
							name: "INTEGER",
						},
					},
				},
			},
		},
		{
			name: "UN",
			pos:  position{line: 67, col: 1, offset: 1420},
			expr: &litMatcher{
				pos:        position{line: 67, col: 7, offset: 1426},
				val:        "_",
				ignoreCase: false,
			},
		},
		{
			name: "IDENT",
			pos:  position{line: 69, col: 1, offset: 1431},
			expr: &actionExpr{
				pos: position{line: 69, col: 10, offset: 1440},
				run: (*parser).callonIDENT1,
				expr: &seqExpr{
					pos: position{line: 69, col: 10, offset: 1440},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 69, col: 10, offset: 1440},
							expr: &charClassMatcher{
								pos:        position{line: 69, col: 10, offset: 1440},
								val:        "[a-zA-Z]",
								ranges:     []rune{'a', 'z', 'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 69, col: 20, offset: 1450},
							expr: &charClassMatcher{
								pos:        position{line: 69, col: 22, offset: 1452},
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
			pos:  position{line: 73, col: 1, offset: 1501},
			expr: &choiceExpr{
				pos: position{line: 73, col: 12, offset: 1512},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 73, col: 12, offset: 1512},
						name: "FLOAT",
					},
					&ruleRefExpr{
						pos:  position{line: 73, col: 20, offset: 1520},
						name: "INTEGER",
					},
				},
			},
		},
		{
			name: "LETV",
			pos:  position{line: 75, col: 1, offset: 1530},
			expr: &choiceExpr{
				pos: position{line: 75, col: 11, offset: 1540},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 75, col: 11, offset: 1540},
						name: "CURRENT",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 22, offset: 1551},
						name: "ACCESS",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 31, offset: 1560},
						name: "NULL",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 38, offset: 1567},
						name: "POPEX",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 46, offset: 1575},
						name: "OPEX",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 53, offset: 1582},
						name: "FCALL",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 61, offset: 1590},
						name: "FUNC",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 68, offset: 1597},
						name: "CALL",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 75, offset: 1604},
						name: "MAP",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 81, offset: 1610},
						name: "ARRAY",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 88, offset: 1617},
						name: "BOOL",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 95, offset: 1624},
						name: "STRING",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 104, offset: 1633},
						name: "NUMBER",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 113, offset: 1642},
						name: "NULL",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 120, offset: 1649},
						name: "IDENT",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 128, offset: 1657},
						name: "FOR",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 134, offset: 1663},
						name: "RANGE",
					},
				},
			},
		},
		{
			name: "COMPARATOR",
			pos:  position{line: 79, col: 1, offset: 1687},
			expr: &choiceExpr{
				pos: position{line: 79, col: 15, offset: 1701},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 79, col: 15, offset: 1701},
						val:        "&&",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 79, col: 22, offset: 1708},
						val:        "||",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 79, col: 29, offset: 1715},
						val:        "<=",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 79, col: 36, offset: 1722},
						val:        ">=",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 79, col: 43, offset: 1729},
						val:        "==",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 79, col: 50, offset: 1736},
						val:        "!=",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 79, col: 57, offset: 1743},
						run: (*parser).callonCOMPARATOR8,
						expr: &litMatcher{
							pos:        position{line: 79, col: 57, offset: 1743},
							val:        "~=",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OPADD",
			pos:  position{line: 83, col: 1, offset: 1782},
			expr: &actionExpr{
				pos: position{line: 83, col: 10, offset: 1791},
				run: (*parser).callonOPADD1,
				expr: &choiceExpr{
					pos: position{line: 83, col: 12, offset: 1793},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 83, col: 12, offset: 1793},
							val:        "+",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 83, col: 18, offset: 1799},
							val:        "-",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OPMUL",
			pos:  position{line: 86, col: 1, offset: 1838},
			expr: &actionExpr{
				pos: position{line: 86, col: 10, offset: 1847},
				run: (*parser).callonOPMUL1,
				expr: &choiceExpr{
					pos: position{line: 86, col: 12, offset: 1849},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 86, col: 12, offset: 1849},
							val:        "*",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 86, col: 18, offset: 1855},
							val:        "/",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EXPR",
			pos:  position{line: 90, col: 1, offset: 1895},
			expr: &actionExpr{
				pos: position{line: 90, col: 9, offset: 1903},
				run: (*parser).callonEXPR1,
				expr: &seqExpr{
					pos: position{line: 90, col: 9, offset: 1903},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 90, col: 9, offset: 1903},
							label: "head",
							expr: &ruleRefExpr{
								pos:  position{line: 90, col: 14, offset: 1908},
								name: "ADDITIVE",
							},
						},
						&labeledExpr{
							pos:   position{line: 90, col: 23, offset: 1917},
							label: "tail",
							expr: &zeroOrMoreExpr{
								pos: position{line: 90, col: 28, offset: 1922},
								expr: &seqExpr{
									pos: position{line: 90, col: 30, offset: 1924},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 90, col: 30, offset: 1924},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 90, col: 32, offset: 1926},
											name: "COMPARATOR",
										},
										&ruleRefExpr{
											pos:  position{line: 90, col: 43, offset: 1937},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 90, col: 45, offset: 1939},
											name: "ADDITIVE",
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
			name: "ADDITIVE",
			pos:  position{line: 94, col: 1, offset: 1994},
			expr: &actionExpr{
				pos: position{line: 94, col: 13, offset: 2006},
				run: (*parser).callonADDITIVE1,
				expr: &seqExpr{
					pos: position{line: 94, col: 13, offset: 2006},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 94, col: 13, offset: 2006},
							label: "head",
							expr: &ruleRefExpr{
								pos:  position{line: 94, col: 18, offset: 2011},
								name: "MULTITIVE",
							},
						},
						&labeledExpr{
							pos:   position{line: 94, col: 28, offset: 2021},
							label: "tail",
							expr: &zeroOrMoreExpr{
								pos: position{line: 94, col: 33, offset: 2026},
								expr: &seqExpr{
									pos: position{line: 94, col: 35, offset: 2028},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 94, col: 35, offset: 2028},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 94, col: 37, offset: 2030},
											name: "OPADD",
										},
										&ruleRefExpr{
											pos:  position{line: 94, col: 43, offset: 2036},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 94, col: 45, offset: 2038},
											name: "MULTITIVE",
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
			name: "MULTITIVE",
			pos:  position{line: 98, col: 1, offset: 2094},
			expr: &actionExpr{
				pos: position{line: 98, col: 14, offset: 2107},
				run: (*parser).callonMULTITIVE1,
				expr: &seqExpr{
					pos: position{line: 98, col: 14, offset: 2107},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 98, col: 14, offset: 2107},
							label: "head",
							expr: &ruleRefExpr{
								pos:  position{line: 98, col: 19, offset: 2112},
								name: "UNARY",
							},
						},
						&labeledExpr{
							pos:   position{line: 98, col: 25, offset: 2118},
							label: "tail",
							expr: &zeroOrMoreExpr{
								pos: position{line: 98, col: 30, offset: 2123},
								expr: &seqExpr{
									pos: position{line: 98, col: 32, offset: 2125},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 98, col: 32, offset: 2125},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 98, col: 34, offset: 2127},
											name: "OPMUL",
										},
										&ruleRefExpr{
											pos:  position{line: 98, col: 40, offset: 2133},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 98, col: 42, offset: 2135},
											name: "UNARY",
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
			name: "UNARY",
			pos:  position{line: 102, col: 1, offset: 2187},
			expr: &choiceExpr{
				pos: position{line: 102, col: 10, offset: 2196},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 102, col: 10, offset: 2196},
						run: (*parser).callonUNARY2,
						expr: &labeledExpr{
							pos:   position{line: 102, col: 10, offset: 2196},
							label: "num",
							expr: &ruleRefExpr{
								pos:  position{line: 102, col: 14, offset: 2200},
								name: "NUMBER",
							},
						},
					},
					&actionExpr{
						pos: position{line: 104, col: 5, offset: 2231},
						run: (*parser).callonUNARY5,
						expr: &seqExpr{
							pos: position{line: 104, col: 5, offset: 2231},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 104, col: 5, offset: 2231},
									val:        "(",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 104, col: 9, offset: 2235},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 104, col: 14, offset: 2240},
										name: "EXPR",
									},
								},
								&litMatcher{
									pos:        position{line: 104, col: 19, offset: 2245},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 106, col: 5, offset: 2274},
						run: (*parser).callonUNARY11,
						expr: &labeledExpr{
							pos:   position{line: 106, col: 5, offset: 2274},
							label: "call",
							expr: &ruleRefExpr{
								pos:  position{line: 106, col: 10, offset: 2279},
								name: "CALLEXPR",
							},
						},
					},
				},
			},
		},
		{
			name: "CALLEXPR",
			pos:  position{line: 110, col: 1, offset: 2312},
			expr: &actionExpr{
				pos: position{line: 110, col: 13, offset: 2324},
				run: (*parser).callonCALLEXPR1,
				expr: &seqExpr{
					pos: position{line: 110, col: 13, offset: 2324},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 110, col: 13, offset: 2324},
							label: "head",
							expr: &choiceExpr{
								pos: position{line: 110, col: 19, offset: 2330},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 110, col: 19, offset: 2330},
										name: "REF",
									},
									&ruleRefExpr{
										pos:  position{line: 110, col: 25, offset: 2336},
										name: "FCALL",
									},
									&ruleRefExpr{
										pos:  position{line: 110, col: 33, offset: 2344},
										name: "IDENT",
									},
									&ruleRefExpr{
										pos:  position{line: 110, col: 41, offset: 2352},
										name: "VAR",
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 110, col: 46, offset: 2357},
							label: "tail",
							expr: &zeroOrMoreExpr{
								pos: position{line: 110, col: 51, offset: 2362},
								expr: &seqExpr{
									pos: position{line: 110, col: 52, offset: 2363},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 110, col: 52, offset: 2363},
											val:        ".",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 110, col: 56, offset: 2367},
											name: "IDENT",
										},
										&zeroOrOneExpr{
											pos: position{line: 110, col: 62, offset: 2373},
											expr: &ruleRefExpr{
												pos:  position{line: 110, col: 62, offset: 2373},
												name: "ARGS",
											},
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
			name: "ARGS",
			pos:  position{line: 114, col: 1, offset: 2422},
			expr: &actionExpr{
				pos: position{line: 114, col: 9, offset: 2430},
				run: (*parser).callonARGS1,
				expr: &seqExpr{
					pos: position{line: 114, col: 9, offset: 2430},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 114, col: 9, offset: 2430},
							val:        "(",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 114, col: 13, offset: 2434},
							label: "head",
							expr: &zeroOrOneExpr{
								pos: position{line: 114, col: 18, offset: 2439},
								expr: &choiceExpr{
									pos: position{line: 114, col: 19, offset: 2440},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 114, col: 19, offset: 2440},
											name: "NUMBER",
										},
										&ruleRefExpr{
											pos:  position{line: 114, col: 26, offset: 2447},
											name: "STRING",
										},
										&ruleRefExpr{
											pos:  position{line: 114, col: 33, offset: 2454},
											name: "EXPR",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 114, col: 40, offset: 2461},
							label: "tail",
							expr: &zeroOrMoreExpr{
								pos: position{line: 114, col: 45, offset: 2466},
								expr: &seqExpr{
									pos: position{line: 114, col: 46, offset: 2467},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 114, col: 46, offset: 2467},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 114, col: 50, offset: 2471},
											name: "_",
										},
										&choiceExpr{
											pos: position{line: 114, col: 53, offset: 2474},
											alternatives: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 114, col: 53, offset: 2474},
													name: "STRING",
												},
												&ruleRefExpr{
													pos:  position{line: 114, col: 60, offset: 2481},
													name: "NUMBER",
												},
												&ruleRefExpr{
													pos:  position{line: 114, col: 67, offset: 2488},
													name: "EXPR",
												},
											},
										},
										&ruleRefExpr{
											pos:  position{line: 114, col: 73, offset: 2494},
											name: "_",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 114, col: 78, offset: 2499},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 114, col: 80, offset: 2501},
							val:        ")",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 114, col: 84, offset: 2505},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "REF",
			pos:  position{line: 118, col: 1, offset: 2547},
			expr: &actionExpr{
				pos: position{line: 118, col: 8, offset: 2554},
				run: (*parser).callonREF1,
				expr: &labeledExpr{
					pos:   position{line: 118, col: 8, offset: 2554},
					label: "name",
					expr: &ruleRefExpr{
						pos:  position{line: 118, col: 13, offset: 2559},
						name: "IDENT",
					},
				},
			},
		},
		{
			name: "LET",
			pos:  position{line: 122, col: 1, offset: 2598},
			expr: &actionExpr{
				pos: position{line: 122, col: 8, offset: 2605},
				run: (*parser).callonLET1,
				expr: &seqExpr{
					pos: position{line: 122, col: 8, offset: 2605},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 122, col: 8, offset: 2605},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 122, col: 10, offset: 2607},
							val:        "let",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 122, col: 16, offset: 2613},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 122, col: 18, offset: 2615},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 122, col: 23, offset: 2620},
								name: "IDENT",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 122, col: 29, offset: 2626},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 122, col: 31, offset: 2628},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 122, col: 35, offset: 2632},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 122, col: 37, offset: 2634},
							label: "value",
							expr: &ruleRefExpr{
								pos:  position{line: 122, col: 43, offset: 2640},
								name: "EXPR",
							},
						},
					},
				},
			},
		},
		{
			name: "VAR",
			pos:  position{line: 126, col: 1, offset: 2682},
			expr: &actionExpr{
				pos: position{line: 126, col: 8, offset: 2689},
				run: (*parser).callonVAR1,
				expr: &seqExpr{
					pos: position{line: 126, col: 8, offset: 2689},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 126, col: 8, offset: 2689},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 126, col: 10, offset: 2691},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 126, col: 15, offset: 2696},
								name: "REF",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 126, col: 19, offset: 2700},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 126, col: 21, offset: 2702},
							val:        ":=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 126, col: 26, offset: 2707},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 126, col: 28, offset: 2709},
							label: "value",
							expr: &ruleRefExpr{
								pos:  position{line: 126, col: 34, offset: 2715},
								name: "LETV",
							},
						},
					},
				},
			},
		},
		{
			name: "ASSIGN",
			pos:  position{line: 130, col: 1, offset: 2757},
			expr: &actionExpr{
				pos: position{line: 130, col: 11, offset: 2767},
				run: (*parser).callonASSIGN1,
				expr: &seqExpr{
					pos: position{line: 130, col: 11, offset: 2767},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 130, col: 11, offset: 2767},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 130, col: 13, offset: 2769},
							label: "name",
							expr: &choiceExpr{
								pos: position{line: 130, col: 19, offset: 2775},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 130, col: 19, offset: 2775},
										name: "ACCESS",
									},
									&ruleRefExpr{
										pos:  position{line: 130, col: 26, offset: 2782},
										name: "IDENT",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 130, col: 33, offset: 2789},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 130, col: 35, offset: 2791},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 130, col: 39, offset: 2795},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 130, col: 41, offset: 2797},
							label: "value",
							expr: &ruleRefExpr{
								pos:  position{line: 130, col: 47, offset: 2803},
								name: "LETV",
							},
						},
					},
				},
			},
		},
		{
			name: "ARG",
			pos:  position{line: 134, col: 1, offset: 2848},
			expr: &actionExpr{
				pos: position{line: 134, col: 8, offset: 2855},
				run: (*parser).callonARG1,
				expr: &seqExpr{
					pos: position{line: 134, col: 8, offset: 2855},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 134, col: 8, offset: 2855},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 134, col: 10, offset: 2857},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 134, col: 12, offset: 2859},
								name: "EXPR",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 134, col: 17, offset: 2864},
							name: "_",
						},
						&zeroOrOneExpr{
							pos: position{line: 134, col: 19, offset: 2866},
							expr: &litMatcher{
								pos:        position{line: 134, col: 19, offset: 2866},
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
			pos:  position{line: 138, col: 1, offset: 2892},
			expr: &actionExpr{
				pos: position{line: 138, col: 10, offset: 2901},
				run: (*parser).callonBLOCK1,
				expr: &seqExpr{
					pos: position{line: 138, col: 10, offset: 2901},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 138, col: 10, offset: 2901},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 138, col: 12, offset: 2903},
							name: "LC",
						},
						&ruleRefExpr{
							pos:  position{line: 138, col: 15, offset: 2906},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 138, col: 17, offset: 2908},
							label: "s",
							expr: &zeroOrMoreExpr{
								pos: position{line: 138, col: 19, offset: 2910},
								expr: &choiceExpr{
									pos: position{line: 138, col: 21, offset: 2912},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 138, col: 21, offset: 2912},
											name: "CONTINUE",
										},
										&ruleRefExpr{
											pos:  position{line: 138, col: 32, offset: 2923},
											name: "BREAK",
										},
										&ruleRefExpr{
											pos:  position{line: 138, col: 40, offset: 2931},
											name: "NODE",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 138, col: 48, offset: 2939},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 138, col: 50, offset: 2941},
							name: "RC",
						},
					},
				},
			},
		},
		{
			name: "FOR",
			pos:  position{line: 142, col: 1, offset: 2973},
			expr: &actionExpr{
				pos: position{line: 142, col: 8, offset: 2980},
				run: (*parser).callonFOR1,
				expr: &seqExpr{
					pos: position{line: 142, col: 8, offset: 2980},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 142, col: 8, offset: 2980},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 142, col: 10, offset: 2982},
							val:        "for ",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 142, col: 17, offset: 2989},
							label: "ax",
							expr: &seqExpr{
								pos: position{line: 142, col: 22, offset: 2994},
								exprs: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 142, col: 22, offset: 2994},
										name: "LP",
									},
									&oneOrMoreExpr{
										pos: position{line: 142, col: 25, offset: 2997},
										expr: &ruleRefExpr{
											pos:  position{line: 142, col: 25, offset: 2997},
											name: "ARG",
										},
									},
									&ruleRefExpr{
										pos:  position{line: 142, col: 30, offset: 3002},
										name: "RP",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 142, col: 35, offset: 3007},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 142, col: 37, offset: 3009},
							val:        "in",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 142, col: 42, offset: 3014},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 142, col: 44, offset: 3016},
							label: "i",
							expr: &choiceExpr{
								pos: position{line: 142, col: 47, offset: 3019},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 142, col: 47, offset: 3019},
										name: "IDENT",
									},
									&ruleRefExpr{
										pos:  position{line: 142, col: 55, offset: 3027},
										name: "ARRAY",
									},
									&ruleRefExpr{
										pos:  position{line: 142, col: 63, offset: 3035},
										name: "MAP",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 142, col: 69, offset: 3041},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 142, col: 71, offset: 3043},
							label: "s",
							expr: &ruleRefExpr{
								pos:  position{line: 142, col: 73, offset: 3045},
								name: "BLOCK",
							},
						},
					},
				},
			},
		},
		{
			name: "RANGE",
			pos:  position{line: 146, col: 1, offset: 3085},
			expr: &actionExpr{
				pos: position{line: 146, col: 10, offset: 3094},
				run: (*parser).callonRANGE1,
				expr: &seqExpr{
					pos: position{line: 146, col: 10, offset: 3094},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 146, col: 10, offset: 3094},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 146, col: 12, offset: 3096},
							val:        "for ",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 146, col: 19, offset: 3103},
							label: "ax",
							expr: &oneOrMoreExpr{
								pos: position{line: 146, col: 24, offset: 3108},
								expr: &ruleRefExpr{
									pos:  position{line: 146, col: 24, offset: 3108},
									name: "ARG",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 146, col: 30, offset: 3114},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 146, col: 32, offset: 3116},
							val:        ":=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 146, col: 37, offset: 3121},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 146, col: 39, offset: 3123},
							val:        "range",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 146, col: 47, offset: 3131},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 146, col: 49, offset: 3133},
							label: "i",
							expr: &choiceExpr{
								pos: position{line: 146, col: 52, offset: 3136},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 146, col: 52, offset: 3136},
										name: "IDENT",
									},
									&ruleRefExpr{
										pos:  position{line: 146, col: 60, offset: 3144},
										name: "ARRAY",
									},
									&ruleRefExpr{
										pos:  position{line: 146, col: 68, offset: 3152},
										name: "MAP",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 146, col: 74, offset: 3158},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 146, col: 76, offset: 3160},
							label: "s",
							expr: &ruleRefExpr{
								pos:  position{line: 146, col: 78, offset: 3162},
								name: "BLOCK",
							},
						},
					},
				},
			},
		},
		{
			name: "IFOR",
			pos:  position{line: 150, col: 1, offset: 3204},
			expr: &actionExpr{
				pos: position{line: 150, col: 9, offset: 3212},
				run: (*parser).callonIFOR1,
				expr: &seqExpr{
					pos: position{line: 150, col: 9, offset: 3212},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 150, col: 9, offset: 3212},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 150, col: 11, offset: 3214},
							val:        "for",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 150, col: 17, offset: 3220},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 150, col: 19, offset: 3222},
							label: "s",
							expr: &ruleRefExpr{
								pos:  position{line: 150, col: 21, offset: 3224},
								name: "BLOCK",
							},
						},
					},
				},
			},
		},
		{
			name: "ARRAY",
			pos:  position{line: 154, col: 1, offset: 3267},
			expr: &actionExpr{
				pos: position{line: 154, col: 10, offset: 3276},
				run: (*parser).callonARRAY1,
				expr: &seqExpr{
					pos: position{line: 154, col: 10, offset: 3276},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 154, col: 10, offset: 3276},
							name: "LB",
						},
						&ruleRefExpr{
							pos:  position{line: 154, col: 13, offset: 3279},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 154, col: 15, offset: 3281},
							label: "ax",
							expr: &zeroOrMoreExpr{
								pos: position{line: 154, col: 18, offset: 3284},
								expr: &ruleRefExpr{
									pos:  position{line: 154, col: 18, offset: 3284},
									name: "ARG",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 154, col: 23, offset: 3289},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 154, col: 25, offset: 3291},
							name: "RB",
						},
					},
				},
			},
		},
		{
			name: "ACCESS",
			pos:  position{line: 158, col: 1, offset: 3324},
			expr: &actionExpr{
				pos: position{line: 158, col: 11, offset: 3334},
				run: (*parser).callonACCESS1,
				expr: &seqExpr{
					pos: position{line: 158, col: 11, offset: 3334},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 158, col: 11, offset: 3334},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 158, col: 14, offset: 3337},
								name: "IDENT",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 158, col: 21, offset: 3344},
							name: "LB",
						},
						&labeledExpr{
							pos:   position{line: 158, col: 24, offset: 3347},
							label: "key",
							expr: &choiceExpr{
								pos: position{line: 158, col: 29, offset: 3352},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 158, col: 29, offset: 3352},
										name: "NUMBER",
									},
									&ruleRefExpr{
										pos:  position{line: 158, col: 36, offset: 3359},
										name: "IDENT",
									},
									&ruleRefExpr{
										pos:  position{line: 158, col: 42, offset: 3365},
										name: "STRING",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 158, col: 50, offset: 3373},
							name: "RB",
						},
					},
				},
			},
		},
		{
			name: "RETURN",
			pos:  position{line: 162, col: 1, offset: 3411},
			expr: &actionExpr{
				pos: position{line: 162, col: 11, offset: 3421},
				run: (*parser).callonRETURN1,
				expr: &seqExpr{
					pos: position{line: 162, col: 11, offset: 3421},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 162, col: 11, offset: 3421},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 162, col: 13, offset: 3423},
							val:        "return",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 162, col: 22, offset: 3432},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 162, col: 24, offset: 3434},
							label: "s",
							expr: &choiceExpr{
								pos: position{line: 162, col: 27, offset: 3437},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 162, col: 27, offset: 3437},
										name: "CALL",
									},
									&ruleRefExpr{
										pos:  position{line: 162, col: 34, offset: 3444},
										name: "CALLEXPR",
									},
									&ruleRefExpr{
										pos:  position{line: 162, col: 45, offset: 3455},
										name: "EXPR",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 162, col: 51, offset: 3461},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "FARG",
			pos:  position{line: 166, col: 1, offset: 3493},
			expr: &actionExpr{
				pos: position{line: 166, col: 9, offset: 3501},
				run: (*parser).callonFARG1,
				expr: &seqExpr{
					pos: position{line: 166, col: 9, offset: 3501},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 166, col: 9, offset: 3501},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 166, col: 11, offset: 3503},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 166, col: 14, offset: 3506},
								name: "IDENT",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 166, col: 21, offset: 3513},
							name: "_",
						},
						&zeroOrOneExpr{
							pos: position{line: 166, col: 23, offset: 3515},
							expr: &litMatcher{
								pos:        position{line: 166, col: 23, offset: 3515},
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
			pos:  position{line: 170, col: 1, offset: 3541},
			expr: &actionExpr{
				pos: position{line: 170, col: 9, offset: 3549},
				run: (*parser).callonFUNC1,
				expr: &seqExpr{
					pos: position{line: 170, col: 9, offset: 3549},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 170, col: 9, offset: 3549},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 170, col: 11, offset: 3551},
							val:        "func",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 170, col: 18, offset: 3558},
							name: "LP",
						},
						&labeledExpr{
							pos:   position{line: 170, col: 21, offset: 3561},
							label: "ax",
							expr: &zeroOrMoreExpr{
								pos: position{line: 170, col: 24, offset: 3564},
								expr: &ruleRefExpr{
									pos:  position{line: 170, col: 24, offset: 3564},
									name: "FARG",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 170, col: 30, offset: 3570},
							name: "RP",
						},
						&ruleRefExpr{
							pos:  position{line: 170, col: 33, offset: 3573},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 170, col: 35, offset: 3575},
							label: "s",
							expr: &ruleRefExpr{
								pos:  position{line: 170, col: 37, offset: 3577},
								name: "BLOCK",
							},
						},
					},
				},
			},
		},
		{
			name: "FCALL",
			pos:  position{line: 174, col: 1, offset: 3615},
			expr: &actionExpr{
				pos: position{line: 174, col: 10, offset: 3624},
				run: (*parser).callonFCALL1,
				expr: &seqExpr{
					pos: position{line: 174, col: 10, offset: 3624},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 174, col: 10, offset: 3624},
							label: "f",
							expr: &ruleRefExpr{
								pos:  position{line: 174, col: 12, offset: 3626},
								name: "FUNC",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 174, col: 17, offset: 3631},
							name: "LP",
						},
						&labeledExpr{
							pos:   position{line: 174, col: 20, offset: 3634},
							label: "ax",
							expr: &zeroOrMoreExpr{
								pos: position{line: 174, col: 23, offset: 3637},
								expr: &ruleRefExpr{
									pos:  position{line: 174, col: 23, offset: 3637},
									name: "ARG",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 174, col: 28, offset: 3642},
							name: "RP",
						},
						&ruleRefExpr{
							pos:  position{line: 174, col: 31, offset: 3645},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 174, col: 33, offset: 3647},
							label: "b",
							expr: &zeroOrOneExpr{
								pos: position{line: 174, col: 35, offset: 3649},
								expr: &ruleRefExpr{
									pos:  position{line: 174, col: 36, offset: 3650},
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
			pos:  position{line: 178, col: 1, offset: 3698},
			expr: &actionExpr{
				pos: position{line: 178, col: 7, offset: 3704},
				run: (*parser).callonGO1,
				expr: &seqExpr{
					pos: position{line: 178, col: 7, offset: 3704},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 178, col: 7, offset: 3704},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 178, col: 9, offset: 3706},
							val:        "go",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 178, col: 14, offset: 3711},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 178, col: 16, offset: 3713},
							label: "s",
							expr: &choiceExpr{
								pos: position{line: 178, col: 20, offset: 3717},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 178, col: 20, offset: 3717},
										name: "FCALL",
									},
									&ruleRefExpr{
										pos:  position{line: 178, col: 28, offset: 3725},
										name: "CALL",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 178, col: 34, offset: 3731},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "CALL",
			pos:  position{line: 186, col: 1, offset: 3866},
			expr: &actionExpr{
				pos: position{line: 186, col: 9, offset: 3874},
				run: (*parser).callonCALL1,
				expr: &seqExpr{
					pos: position{line: 186, col: 9, offset: 3874},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 186, col: 9, offset: 3874},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 186, col: 11, offset: 3876},
							label: "i",
							expr: &choiceExpr{
								pos: position{line: 186, col: 15, offset: 3880},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 186, col: 15, offset: 3880},
										name: "IDENT",
									},
									&ruleRefExpr{
										pos:  position{line: 186, col: 21, offset: 3886},
										name: "CURRENT",
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 186, col: 32, offset: 3897},
							expr: &litMatcher{
								pos:        position{line: 186, col: 32, offset: 3897},
								val:        ".",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 186, col: 38, offset: 3903},
							label: "y",
							expr: &zeroOrOneExpr{
								pos: position{line: 186, col: 40, offset: 3905},
								expr: &ruleRefExpr{
									pos:  position{line: 186, col: 41, offset: 3906},
									name: "IDENT",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 186, col: 49, offset: 3914},
							name: "LP",
						},
						&labeledExpr{
							pos:   position{line: 186, col: 52, offset: 3917},
							label: "ax",
							expr: &zeroOrMoreExpr{
								pos: position{line: 186, col: 55, offset: 3920},
								expr: &ruleRefExpr{
									pos:  position{line: 186, col: 55, offset: 3920},
									name: "ARG",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 186, col: 60, offset: 3925},
							name: "RP",
						},
						&ruleRefExpr{
							pos:  position{line: 186, col: 63, offset: 3928},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 186, col: 65, offset: 3930},
							label: "b",
							expr: &zeroOrOneExpr{
								pos: position{line: 186, col: 67, offset: 3932},
								expr: &ruleRefExpr{
									pos:  position{line: 186, col: 68, offset: 3933},
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
			pos:  position{line: 190, col: 1, offset: 3979},
			expr: &actionExpr{
				pos: position{line: 190, col: 9, offset: 3987},
				run: (*parser).callonBOOL1,
				expr: &choiceExpr{
					pos: position{line: 190, col: 10, offset: 3988},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 190, col: 10, offset: 3988},
							val:        "false",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 190, col: 20, offset: 3998},
							val:        "true",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "STRING",
			pos:  position{line: 194, col: 1, offset: 4039},
			expr: &actionExpr{
				pos: position{line: 194, col: 11, offset: 4049},
				run: (*parser).callonSTRING1,
				expr: &labeledExpr{
					pos:   position{line: 194, col: 11, offset: 4049},
					label: "s",
					expr: &choiceExpr{
						pos: position{line: 194, col: 14, offset: 4052},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 194, col: 14, offset: 4052},
								name: "MLSTRING",
							},
							&ruleRefExpr{
								pos:  position{line: 194, col: 25, offset: 4063},
								name: "DQSTRING",
							},
						},
					},
				},
			},
		},
		{
			name: "MLSTRING",
			pos:  position{line: 198, col: 1, offset: 4095},
			expr: &actionExpr{
				pos: position{line: 198, col: 13, offset: 4107},
				run: (*parser).callonMLSTRING1,
				expr: &seqExpr{
					pos: position{line: 198, col: 13, offset: 4107},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 198, col: 13, offset: 4107},
							val:        "`",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 198, col: 17, offset: 4111},
							expr: &charClassMatcher{
								pos:        position{line: 198, col: 17, offset: 4111},
								val:        "[^`]",
								chars:      []rune{'`'},
								ignoreCase: false,
								inverted:   true,
							},
						},
						&litMatcher{
							pos:        position{line: 198, col: 23, offset: 4117},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "DQSTRING",
			pos:  position{line: 202, col: 1, offset: 4148},
			expr: &actionExpr{
				pos: position{line: 202, col: 13, offset: 4160},
				run: (*parser).callonDQSTRING1,
				expr: &seqExpr{
					pos: position{line: 202, col: 13, offset: 4160},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 202, col: 13, offset: 4160},
							val:        "\"",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 202, col: 17, offset: 4164},
							expr: &choiceExpr{
								pos: position{line: 202, col: 19, offset: 4166},
								alternatives: []interface{}{
									&seqExpr{
										pos: position{line: 202, col: 19, offset: 4166},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 202, col: 19, offset: 4166},
												expr: &ruleRefExpr{
													pos:  position{line: 202, col: 20, offset: 4167},
													name: "DQEscapeChar",
												},
											},
											&anyMatcher{
												line: 202, col: 33, offset: 4180,
											},
										},
									},
									&seqExpr{
										pos: position{line: 202, col: 37, offset: 4184},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 202, col: 37, offset: 4184},
												val:        "\\",
												ignoreCase: false,
											},
											&ruleRefExpr{
												pos:  position{line: 202, col: 42, offset: 4189},
												name: "DQEscapeSequence",
											},
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 202, col: 62, offset: 4209},
							val:        "\"",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "BREAK",
			pos:  position{line: 206, col: 1, offset: 4240},
			expr: &actionExpr{
				pos: position{line: 206, col: 10, offset: 4249},
				run: (*parser).callonBREAK1,
				expr: &seqExpr{
					pos: position{line: 206, col: 10, offset: 4249},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 206, col: 10, offset: 4249},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 206, col: 12, offset: 4251},
							val:        "break",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 206, col: 20, offset: 4259},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "CONTINUE",
			pos:  position{line: 210, col: 1, offset: 4305},
			expr: &actionExpr{
				pos: position{line: 210, col: 13, offset: 4317},
				run: (*parser).callonCONTINUE1,
				expr: &seqExpr{
					pos: position{line: 210, col: 13, offset: 4317},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 210, col: 13, offset: 4317},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 210, col: 15, offset: 4319},
							val:        "continue",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 210, col: 26, offset: 4330},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "NULL",
			pos:  position{line: 214, col: 1, offset: 4379},
			expr: &actionExpr{
				pos: position{line: 214, col: 8, offset: 4388},
				run: (*parser).callonNULL1,
				expr: &litMatcher{
					pos:        position{line: 214, col: 10, offset: 4390},
					val:        "nil",
					ignoreCase: false,
				},
			},
		},
		{
			name: "MAP",
			pos:  position{line: 218, col: 1, offset: 4440},
			expr: &actionExpr{
				pos: position{line: 218, col: 8, offset: 4447},
				run: (*parser).callonMAP1,
				expr: &seqExpr{
					pos: position{line: 218, col: 8, offset: 4447},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 218, col: 8, offset: 4447},
							name: "LC",
						},
						&ruleRefExpr{
							pos:  position{line: 218, col: 11, offset: 4450},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 218, col: 13, offset: 4452},
							label: "vals",
							expr: &zeroOrMoreExpr{
								pos: position{line: 218, col: 18, offset: 4457},
								expr: &seqExpr{
									pos: position{line: 218, col: 20, offset: 4459},
									exprs: []interface{}{
										&choiceExpr{
											pos: position{line: 218, col: 21, offset: 4460},
											alternatives: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 218, col: 21, offset: 4460},
													name: "IDENT",
												},
												&ruleRefExpr{
													pos:  position{line: 218, col: 27, offset: 4466},
													name: "STRING",
												},
											},
										},
										&ruleRefExpr{
											pos:  position{line: 218, col: 35, offset: 4474},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 218, col: 37, offset: 4476},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 218, col: 41, offset: 4480},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 218, col: 44, offset: 4483},
											name: "ARG",
										},
										&ruleRefExpr{
											pos:  position{line: 218, col: 49, offset: 4488},
											name: "_",
										},
										&zeroOrMoreExpr{
											pos: position{line: 218, col: 51, offset: 4490},
											expr: &litMatcher{
												pos:        position{line: 218, col: 52, offset: 4491},
												val:        ",",
												ignoreCase: false,
											},
										},
										&ruleRefExpr{
											pos:  position{line: 218, col: 58, offset: 4497},
											name: "_",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 218, col: 64, offset: 4503},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 218, col: 66, offset: 4505},
							name: "RC",
						},
					},
				},
			},
		},
		{
			name: "COMMENT",
			pos:  position{line: 222, col: 1, offset: 4538},
			expr: &actionExpr{
				pos: position{line: 222, col: 12, offset: 4549},
				run: (*parser).callonCOMMENT1,
				expr: &seqExpr{
					pos: position{line: 222, col: 12, offset: 4549},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 222, col: 12, offset: 4549},
							name: "_",
						},
						&choiceExpr{
							pos: position{line: 222, col: 15, offset: 4552},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 222, col: 15, offset: 4552},
									val:        "#",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 222, col: 21, offset: 4558},
									val:        "//",
									ignoreCase: false,
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 222, col: 28, offset: 4565},
							label: "a",
							expr: &zeroOrMoreExpr{
								pos: position{line: 222, col: 30, offset: 4567},
								expr: &charClassMatcher{
									pos:        position{line: 222, col: 30, offset: 4567},
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
			pos:  position{line: 226, col: 1, offset: 4612},
			expr: &actionExpr{
				pos: position{line: 226, col: 12, offset: 4623},
				run: (*parser).callonSHEBANG1,
				expr: &seqExpr{
					pos: position{line: 226, col: 12, offset: 4623},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 226, col: 12, offset: 4623},
							val:        "#!",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 226, col: 17, offset: 4628},
							label: "a",
							expr: &zeroOrMoreExpr{
								pos: position{line: 226, col: 19, offset: 4630},
								expr: &charClassMatcher{
									pos:        position{line: 226, col: 19, offset: 4630},
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
			pos:  position{line: 230, col: 1, offset: 4676},
			expr: &actionExpr{
				pos: position{line: 230, col: 7, offset: 4682},
				run: (*parser).callonLP1,
				expr: &litMatcher{
					pos:        position{line: 230, col: 7, offset: 4682},
					val:        "(",
					ignoreCase: false,
				},
			},
		},
		{
			name: "RP",
			pos:  position{line: 231, col: 1, offset: 4707},
			expr: &actionExpr{
				pos: position{line: 231, col: 7, offset: 4713},
				run: (*parser).callonRP1,
				expr: &litMatcher{
					pos:        position{line: 231, col: 7, offset: 4713},
					val:        ")",
					ignoreCase: false,
				},
			},
		},
		{
			name: "LB",
			pos:  position{line: 232, col: 1, offset: 4738},
			expr: &actionExpr{
				pos: position{line: 232, col: 7, offset: 4744},
				run: (*parser).callonLB1,
				expr: &litMatcher{
					pos:        position{line: 232, col: 7, offset: 4744},
					val:        "[",
					ignoreCase: false,
				},
			},
		},
		{
			name: "RB",
			pos:  position{line: 233, col: 1, offset: 4769},
			expr: &actionExpr{
				pos: position{line: 233, col: 7, offset: 4775},
				run: (*parser).callonRB1,
				expr: &litMatcher{
					pos:        position{line: 233, col: 7, offset: 4775},
					val:        "]",
					ignoreCase: false,
				},
			},
		},
		{
			name: "LC",
			pos:  position{line: 234, col: 1, offset: 4800},
			expr: &actionExpr{
				pos: position{line: 234, col: 7, offset: 4806},
				run: (*parser).callonLC1,
				expr: &litMatcher{
					pos:        position{line: 234, col: 7, offset: 4806},
					val:        "{",
					ignoreCase: false,
				},
			},
		},
		{
			name: "RC",
			pos:  position{line: 235, col: 1, offset: 4831},
			expr: &actionExpr{
				pos: position{line: 235, col: 7, offset: 4837},
				run: (*parser).callonRC1,
				expr: &litMatcher{
					pos:        position{line: 235, col: 7, offset: 4837},
					val:        "}",
					ignoreCase: false,
				},
			},
		},
		{
			name: "MLEscapeChar",
			pos:  position{line: 237, col: 1, offset: 4863},
			expr: &charClassMatcher{
				pos:        position{line: 237, col: 17, offset: 4879},
				val:        "[\\x00-\\x1f\\\\`]",
				chars:      []rune{'\\', '`'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "MLEscapeSequence",
			pos:  position{line: 238, col: 1, offset: 4894},
			expr: &choiceExpr{
				pos: position{line: 238, col: 21, offset: 4914},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 238, col: 21, offset: 4914},
						name: "MLSingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 238, col: 42, offset: 4935},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "MLSingleCharEscape",
			pos:  position{line: 239, col: 1, offset: 4949},
			expr: &charClassMatcher{
				pos:        position{line: 239, col: 23, offset: 4971},
				val:        "['\\\\/bfnrt]",
				chars:      []rune{'\'', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DQEscapeChar",
			pos:  position{line: 241, col: 1, offset: 4984},
			expr: &charClassMatcher{
				pos:        position{line: 241, col: 17, offset: 5000},
				val:        "[\\x00-\\x1f\\\\\"]",
				chars:      []rune{'\\', '"'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DQEscapeSequence",
			pos:  position{line: 242, col: 1, offset: 5015},
			expr: &choiceExpr{
				pos: position{line: 242, col: 21, offset: 5035},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 242, col: 21, offset: 5035},
						name: "DQSingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 242, col: 42, offset: 5056},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "DQSingleCharEscape",
			pos:  position{line: 243, col: 1, offset: 5070},
			expr: &charClassMatcher{
				pos:        position{line: 243, col: 23, offset: 5092},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "UnicodeEscape",
			pos:  position{line: 245, col: 1, offset: 5105},
			expr: &seqExpr{
				pos: position{line: 245, col: 17, offset: 5123},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 245, col: 17, offset: 5123},
						val:        "u",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 245, col: 21, offset: 5127},
						name: "HexDigit",
					},
					&ruleRefExpr{
						pos:  position{line: 245, col: 30, offset: 5136},
						name: "HexDigit",
					},
					&ruleRefExpr{
						pos:  position{line: 245, col: 39, offset: 5145},
						name: "HexDigit",
					},
					&ruleRefExpr{
						pos:  position{line: 245, col: 48, offset: 5154},
						name: "HexDigit",
					},
				},
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 247, col: 1, offset: 5164},
			expr: &charClassMatcher{
				pos:        position{line: 247, col: 16, offset: 5181},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 249, col: 1, offset: 5188},
			expr: &charClassMatcher{
				pos:        position{line: 249, col: 23, offset: 5212},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 251, col: 1, offset: 5219},
			expr: &charClassMatcher{
				pos:        position{line: 251, col: 12, offset: 5232},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "_",
			displayName: "\"whitespace\"",
			pos:         position{line: 253, col: 1, offset: 5243},
			expr: &actionExpr{
				pos: position{line: 253, col: 18, offset: 5262},
				run: (*parser).callon_1,
				expr: &zeroOrMoreExpr{
					pos: position{line: 253, col: 18, offset: 5262},
					expr: &charClassMatcher{
						pos:        position{line: 253, col: 18, offset: 5262},
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
			pos:  position{line: 257, col: 1, offset: 5298},
			expr: &actionExpr{
				pos: position{line: 257, col: 8, offset: 5305},
				run: (*parser).callonEOF1,
				expr: &notExpr{
					pos: position{line: 257, col: 8, offset: 5305},
					expr: &anyMatcher{
						line: 257, col: 9, offset: 5306,
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

func (c *current) onCOMPARATOR8() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonCOMPARATOR8() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCOMPARATOR8()
}

func (c *current) onOPADD1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonOPADD1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOPADD1()
}

func (c *current) onOPMUL1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonOPMUL1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOPMUL1()
}

func (c *current) onEXPR1(head, tail interface{}) (interface{}, error) {
	return newBinaryExpr(c, head, tail)
}

func (p *parser) callonEXPR1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEXPR1(stack["head"], stack["tail"])
}

func (c *current) onADDITIVE1(head, tail interface{}) (interface{}, error) {
	return newBinaryExpr(c, head, tail)
}

func (p *parser) callonADDITIVE1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onADDITIVE1(stack["head"], stack["tail"])
}

func (c *current) onMULTITIVE1(head, tail interface{}) (interface{}, error) {
	return newBinaryExpr(c, head, tail)
}

func (p *parser) callonMULTITIVE1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMULTITIVE1(stack["head"], stack["tail"])
}

func (c *current) onUNARY2(num interface{}) (interface{}, error) {
	return num, nil
}

func (p *parser) callonUNARY2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUNARY2(stack["num"])
}

func (c *current) onUNARY5(expr interface{}) (interface{}, error) {
	return expr, nil
}

func (p *parser) callonUNARY5() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUNARY5(stack["expr"])
}

func (c *current) onUNARY11(call interface{}) (interface{}, error) {
	return call, nil
}

func (p *parser) callonUNARY11() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUNARY11(stack["call"])
}

func (c *current) onCALLEXPR1(head, tail interface{}) (interface{}, error) {
	return newCallExpr(c, head, tail)
}

func (p *parser) callonCALLEXPR1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCALLEXPR1(stack["head"], stack["tail"])
}

func (c *current) onARGS1(head, tail interface{}) (interface{}, error) {
	return newArglist(c, head, tail)
}

func (p *parser) callonARGS1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onARGS1(stack["head"], stack["tail"])
}

func (c *current) onREF1(name interface{}) (interface{}, error) {
	return newVarRef(c, name)
}

func (p *parser) callonREF1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onREF1(stack["name"])
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
