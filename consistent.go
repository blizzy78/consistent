package consistent

import (
	"go/ast"
	"go/token"
	"slices"
	"strconv"
	"strings"

	"github.com/go-toolsmith/astcast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

type configuration struct {
	params                  enumValue
	returns                 enumValue
	typeParams              enumValue
	funcTypeParams          enumValue
	singleImports           enumValue
	newAllocs               enumValue
	makeAllocs              enumValue
	hexLits                 enumValue
	rangeChecks             enumValue
	andNOTs                 enumValue
	floatLits               enumValue
	lenChecks               enumValue
	switchCases             enumValue
	switchDefaults          enumValue
	emptyIfaces             enumValue
	slogAttrs               enumValue
	labelsRegexp            regexpValue
	fileCommentIgnoreRegexp regexpValue
}

const flagIgnore = "ignore"

// NewAnalyzer returns a new analyzer.
func NewAnalyzer() *analysis.Analyzer {
	cfg := configuration{
		params: enumValue{
			allowed: paramsFlagAllowedValues,
			value:   fieldListExplicit,
		},

		returns: enumValue{
			allowed: returnsFlagAllowedValues,
			value:   fieldListExplicit,
		},

		typeParams: enumValue{
			allowed: typeParamsFlagAllowedValues,
			value:   fieldListExplicit,
		},

		funcTypeParams: enumValue{
			allowed: funcTypeParamsFlagAllowedValues,
			value:   fieldListExplicit,
		},

		singleImports: enumValue{
			allowed: singleImportsFlagAllowedValues,
			value:   singleImportsBare,
		},

		newAllocs: enumValue{
			allowed: newAllocsFlagAllowedValues,
			value:   newAllocsLiteral,
		},

		makeAllocs: enumValue{
			allowed: makeAllocsFlagAllowedValues,
			value:   makeAllocsLiteral,
		},

		hexLits: enumValue{
			allowed: hexLitsFlagAllowedValues,
			value:   hexLitsLower,
		},

		rangeChecks: enumValue{
			allowed: rangeChecksFlagAllowedValues,
			value:   rangeChecksLeft,
		},

		andNOTs: enumValue{
			allowed: andNOTsFlagAllowedValues,
			value:   andNOTsANDNOT,
		},

		floatLits: enumValue{
			allowed: floatLitsFlagAllowedValues,
			value:   floatLitsExplicit,
		},

		lenChecks: enumValue{
			allowed: lenChecksFlagAllowedValues,
			value:   lenChecksEqualZero,
		},

		switchCases: enumValue{
			allowed: switchCasesFlagAllowedValues,
			value:   switchCasesComma,
		},

		switchDefaults: enumValue{
			allowed: switchDefaultsFlagAllowedValues,
			value:   switchDefaultsLast,
		},

		emptyIfaces: enumValue{
			allowed: emptyIfacesFlagAllowedValues,
			value:   emptyIfacesAny,
		},

		slogAttrs: enumValue{
			allowed: slogAttrsFlagAllowedValues,
			value:   slogAttrsAttr,
		},

		labelsRegexp: newRegexpValue("^[a-z][a-zA-Z0-9]*$"),

		fileCommentIgnoreRegexp: newRegexpValue(""),
	}

	ana := analysis.Analyzer{
		Name: "consistent",
		Doc:  "checks that common constructs are used consistently",

		Run: func(pass *analysis.Pass) (any, error) {
			run(pass, &cfg)
			return nil, nil
		},

		Requires: []*analysis.Analyzer{
			inspect.Analyzer,
		},
	}

	ana.Flags.Var(&cfg.params, "params", cfg.params.description("check function/method parameter types"))
	ana.Flags.Var(&cfg.returns, "returns", cfg.returns.description("check function/method return value types"))
	ana.Flags.Var(&cfg.typeParams, "typeParams", cfg.typeParams.description("check type parameter types"))
	ana.Flags.Var(&cfg.funcTypeParams, "funcTypeParams", cfg.funcTypeParams.description("check function type parameter types"))
	ana.Flags.Var(&cfg.singleImports, "singleImports", cfg.singleImports.description("check single import declarations"))
	ana.Flags.Var(&cfg.newAllocs, "newAllocs", cfg.newAllocs.description("check allocations using new"))
	ana.Flags.Var(&cfg.makeAllocs, "makeAllocs", cfg.makeAllocs.description("check allocations using make"))
	ana.Flags.Var(&cfg.hexLits, "hexLits", cfg.hexLits.description("check upper/lowercase in hex literals"))
	ana.Flags.Var(&cfg.rangeChecks, "rangeChecks", cfg.rangeChecks.description("check range checks"))
	ana.Flags.Var(&cfg.andNOTs, "andNOTs", cfg.andNOTs.description("check AND-NOT expressions"))
	ana.Flags.Var(&cfg.floatLits, "floatLits", cfg.floatLits.description("check floating-point literals"))
	ana.Flags.Var(&cfg.lenChecks, "lenChecks", cfg.lenChecks.description("check len/cap checks"))
	ana.Flags.Var(&cfg.switchCases, "switchCases", cfg.switchCases.description("check switch case clauses"))
	ana.Flags.Var(&cfg.switchDefaults, "switchDefaults", cfg.switchDefaults.description("check switch default clauses"))
	ana.Flags.Var(&cfg.emptyIfaces, "emptyIfaces", cfg.emptyIfaces.description("check empty interfaces"))
	ana.Flags.Var(&cfg.slogAttrs, "slogAttrs", cfg.slogAttrs.description("check log/slog argument types"))
	ana.Flags.Var(&cfg.labelsRegexp, "labelsRegexp", "check labels against regexp (\"\" to disable)")
	ana.Flags.Var(&cfg.fileCommentIgnoreRegexp, "fileCommentIgnoreRegexp", "ignore files containing comment matching regexp")

	return &ana
}

func run(pass *analysis.Pass, cfg *configuration) { //nolint:cyclop // it's only basic dispatcher logic
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector) //nolint:forcetypeassert // inspect.Analyzer always returns *inspector.Inspector

	ignoreFiles := filesToIgnore(pass, cfg)

	filter := []ast.Node{
		(*ast.AssignStmt)(nil),
		(*ast.BasicLit)(nil),
		(*ast.BinaryExpr)(nil),
		(*ast.CallExpr)(nil),
		(*ast.CompositeLit)(nil),
		(*ast.Field)(nil),
		(*ast.File)(nil),
		(*ast.FuncDecl)(nil),
		(*ast.FuncLit)(nil),
		(*ast.FuncType)(nil),
		(*ast.InterfaceType)(nil),
		(*ast.LabeledStmt)(nil),
		(*ast.SwitchStmt)(nil),
		(*ast.TypeSpec)(nil),
		(*ast.UnaryExpr)(nil),
		(*ast.ValueSpec)(nil),
	}

	funcDecls := []*ast.FuncDecl{}
	funcLits := []*ast.FuncLit{}
	funcTypes := []*ast.FuncType{}

	ignoreCurrentFile := false

	inspector.Preorder(filter, func(node ast.Node) {
		if ignoreCurrentFile {
			return
		}

		switch node := node.(type) {
		case *ast.AssignStmt:
			checkAndNotAssignStmt(pass, node, cfg.andNOTs.value)

		case *ast.BasicLit:
			checkHexLit(pass, node, cfg.hexLits.value)
			checkFloatLit(pass, node, cfg.floatLits.value)

		case *ast.BinaryExpr:
			checkRangeCheck(pass, node, cfg.rangeChecks.value)
			checkAndNotExpr(pass, node, cfg.andNOTs.value)
			checkLenCheck(pass, node, cfg.lenChecks.value)

		case *ast.CallExpr:
			checkNewAllocNew(pass, node, cfg.newAllocs.value)
			checkMakeAllocMake(pass, node, cfg.makeAllocs.value)
			checkSlogAttrs(pass, node, cfg.slogAttrs.value)

		case *ast.CompositeLit:
			checkMakeAllocLit(pass, node, cfg.makeAllocs.value)

		case *ast.Field:
			checkEmptyIface(pass, node, cfg.emptyIfaces.value)

		case *ast.File:
			ignoreCurrentFile = slices.Contains(ignoreFiles, node)
			if ignoreCurrentFile {
				return
			}

			checkSingleImports(pass, node, cfg.singleImports.value)

		case *ast.FuncDecl:
			funcDecls = append(funcDecls, node)

			if node.Recv == nil {
				checkParamsFunc(pass, node, cfg.params.value)
				checkReturnsFunc(pass, node, cfg.returns.value)
				checkTypeParamsFunc(pass, node, cfg.typeParams.value)

				return
			}

			checkParamsMethod(pass, node, cfg.params.value)
			checkReturnsMethod(pass, node, cfg.returns.value)

		case *ast.FuncLit:
			funcLits = append(funcLits, node)

			checkParamsFuncLit(pass, node, cfg.params.value)
			checkReturnsFuncLit(pass, node, cfg.returns.value)

		case *ast.FuncType:
			funcTypes = append(funcTypes, node)

		case *ast.InterfaceType:
			checkEmptyIface(pass, node, cfg.emptyIfaces.value)

		case *ast.LabeledStmt:
			checkLabel(pass, node, cfg.labelsRegexp)

		case *ast.SwitchStmt:
			checkSwitchCases(pass, node, cfg.switchCases.value)
			checkSwitchDefault(pass, node, cfg.switchDefaults.value)

		case *ast.TypeSpec:
			checkTypeParamsType(pass, node, cfg.typeParams.value)
			checkEmptyIface(pass, node, cfg.emptyIfaces.value)

		case *ast.UnaryExpr:
			checkNewAllocLit(pass, node, cfg.newAllocs.value)

		case *ast.ValueSpec:
			checkEmptyIface(pass, node, cfg.emptyIfaces.value)
		}
	})

	for _, typ := range funcTypes {
		if isFuncDefinition(typ, funcDecls, funcLits) {
			continue
		}

		checkParamsFuncType(pass, typ, cfg.funcTypeParams.value)
	}
}

func filesToIgnore(pass *analysis.Pass, cfg *configuration) []*ast.File {
	if cfg.fileCommentIgnoreRegexp.r == nil {
		return nil
	}

	files := []*ast.File{}

	for _, file := range pass.Files {
		for _, comment := range file.Comments {
			if cfg.fileCommentIgnoreRegexp.r.MatchString(comment.Text()) {
				files = append(files, file)
				break
			}
		}
	}

	return files
}

func reportf(pass *analysis.Pass, pos token.Pos, format string, args ...any) {
	if noLintDirectivePos(pass, pos) {
		return
	}

	pass.Reportf(pos, format, args...)
}

func noLintDirectivePos(pass *analysis.Pass, pos token.Pos) bool {
	line := pass.Fset.Position(pos).Line

	for _, file := range pass.Files {
		if pos < file.FileStart || pos > file.FileEnd {
			continue
		}

		for _, comment := range file.Comments {
			if pass.Fset.Position(comment.Pos()).Line < line {
				continue
			}

			return noLintDirective(comment.List[0].Text)
		}
	}

	return false
}

func noLintDirective(comment string) bool {
	if !strings.HasPrefix(comment, "//nolint:") {
		return false
	}

	comment = comment[9:]

	secondaryCommentPos := strings.Index(comment, "//")
	if secondaryCommentPos >= 0 {
		comment = comment[:secondaryCommentPos]
	}

	for _, p := range strings.Split(comment, ",") {
		if strings.TrimSpace(p) == "consistent" {
			return true
		}
	}

	return false
}

func isFuncDefinition(typ *ast.FuncType, decls []*ast.FuncDecl, lits []*ast.FuncLit) bool {
	for _, decl := range decls {
		if decl.Type == typ {
			return true
		}
	}

	for _, lit := range lits {
		if lit.Type == typ {
			return true
		}
	}

	return false
}

func litInt(expr ast.Expr) (int, bool) {
	lit := astcast.ToBasicLit(expr)

	if lit.Kind != token.INT {
		return 0, false
	}

	val := strings.ToLower(strings.ReplaceAll(lit.Value, "_", ""))
	base := 10

	switch {
	case strings.HasPrefix(val, "0x"):
		val = val[2:]
		base = 16

	case strings.HasPrefix(val, "0b"):
		val = val[2:]
		base = 2

	case strings.HasPrefix(val, "0o"):
		val = val[2:]
		base = 8

	case strings.HasPrefix(val, "0") && strings.TrimLeft(val, "0") == "":
		return 0, true

	case strings.HasPrefix(val, "0"):
		val = val[1:]
		base = 8
	}

	i, err := strconv.ParseInt(val, base, 32)
	if err != nil {
		return 0, false
	}

	return int(i), true
}

func namedFields(fields *ast.FieldList) bool {
	return fields != nil && len(fields.List) != 0 && len(fields.List[0].Names) != 0
}

func unnamedFields(fields *ast.FieldList) bool {
	return fields != nil && len(fields.List) != 0 && len(fields.List[0].Names) == 0
}
