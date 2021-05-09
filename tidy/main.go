//package main
//
//import (
//	"fmt"
//	"sort"
//
//	"golang.org/x/tools/go/callgraph"
//	"golang.org/x/tools/go/loader"
//	"golang.org/x/tools/go/pointer"
//	"golang.org/x/tools/go/ssa"
//	"golang.org/x/tools/go/ssa/ssautil"
//	"golang.org/x/tools/go/packages"
//)
//
//func main() {
//
//	//packages.Load()
//
//	const myprog = `
//package main
//
//import "fmt"
//
//type I interface {
//	f(map[string]int)
//}
//
//type C struct{}
//
//func (C) f(m map[string]int) {
//	fmt.Println("C.f()")
//	fmt.Sprint()
//}
//
//func main() {
//	var i I = C{}
//	x := map[string]int{"one":1}
//	i.f(x) // dynamic method call
//}
//`
//	var conf loader.Config
//
//	// Parse the input file, a string.
//	// (Command-line tools should use conf.FromArgs.)
//	file, err := conf.ParseFile("myprog.go", myprog)
//	if err != nil {
//		fmt.Print(err) // parse error
//		return
//	}
//
//	// Create single-file main package and import its dependencies.
//	conf.CreateFromFiles("main", file)
//
//	iprog, err := conf.Load()
//	if err != nil {
//		fmt.Print(err) // type error in some package
//		return
//	}
//
//	// Create SSA-form program representation.
//	prog := ssautil.CreateProgram(iprog, 0)
//	mainPkg := prog.Package(iprog.Created[0].Pkg)
//
//	// Build SSA code for bodies of all functions in the whole program.
//	prog.Build()
//
//	// Configure the pointer analysis to build a call-graph.
//	config := &pointer.Config{
//		Mains:          []*ssa.Package{mainPkg},
//		BuildCallGraph: true,
//	}
//
//	// Query points-to set of (C).f's parameter m, a map.
//	C := mainPkg.Type("C").Type()
//	Cfm := prog.LookupMethod(C, mainPkg.Pkg, "f").Params[1]
//	config.AddQuery(Cfm)
//
//	// Run the pointer analysis.
//	result, err := pointer.Analyze(config)
//	if err != nil {
//		panic(err) // internal error in pointer analysis
//	}
//
//	// Find edges originating from the main package.
//	// By converting to strings, we de-duplicate nodes
//	// representing the same function due to context sensitivity.
//	var edges []string
//	callgraph.GraphVisitEdges(result.CallGraph, func(edge *callgraph.Edge) error {
//		caller := edge.Caller.Func
//		if caller.Pkg == mainPkg {
//			edges = append(edges, fmt.Sprint(caller, " --> ", edge.Callee.Func))
//		}
//		return nil
//	})
//
//	// Print the edges in sorted order.
//	sort.Strings(edges)
//	for _, edge := range edges {
//		fmt.Println(edge)
//	}
//	fmt.Println()
//
//	// Print the labels of (C).f(m)'s points-to set.
//	fmt.Println("m may point to:")
//	var labels []string
//	for _, l := range result.Queries[Cfm].PointsTo().Labels() {
//		label := fmt.Sprintf("  %s: %s", prog.Fset.Position(l.Pos()), l)
//		labels = append(labels, label)
//	}
//	sort.Strings(labels)
//	for _, label := range labels {
//		fmt.Println(label)
//	}
//
//}
//
//func main1() {
//
//	//packages.Load()
//
//	const myprog = `
//package main
//
//import "fmt"
//
//type I interface {
//	f(map[string]int)
//}
//
//type C struct{}
//
//func (C) f(m map[string]int) {
//	fmt.Println("C.f()")
//	fmt.Sprint()
//}
//
//func main() {
//	var i I = C{}
//	x := map[string]int{"one":1}
//	i.f(x) // dynamic method call
//}
//`
//
//	packages.Load(packages.Config{
//		Mode:       0,
//		Context:    nil,
//		Logf:       nil,
//		Dir:        "",
//		Env:        nil,
//		BuildFlags: nil,
//		Fset:       nil,
//		ParseFile:  nil,
//		Tests:      false,
//		Overlay:    nil,
//	})
//	var conf loader.Config
//
//	// Parse the input file, a string.
//	// (Command-line tools should use conf.FromArgs.)
//	file, err := conf.ParseFile("myprog.go", myprog)
//	if err != nil {
//		fmt.Print(err) // parse error
//		return
//	}
//
//	// Create single-file main package and import its dependencies.
//	conf.CreateFromFiles("main", file)
//
//	iprog, err := conf.Load()
//	if err != nil {
//		fmt.Print(err) // type error in some package
//		return
//	}
//
//	// Create SSA-form program representation.
//	prog := ssautil.CreateProgram(iprog, 0)
//	mainPkg := prog.Package(iprog.Created[0].Pkg)
//
//	// Build SSA code for bodies of all functions in the whole program.
//	prog.Build()
//
//	// Configure the pointer analysis to build a call-graph.
//	config := &pointer.Config{
//		Mains:          []*ssa.Package{mainPkg},
//		BuildCallGraph: true,
//	}
//
//	// Query points-to set of (C).f's parameter m, a map.
//	C := mainPkg.Type("C").Type()
//	Cfm := prog.LookupMethod(C, mainPkg.Pkg, "f").Params[1]
//	config.AddQuery(Cfm)
//
//	// Run the pointer analysis.
//	result, err := pointer.Analyze(config)
//	if err != nil {
//		panic(err) // internal error in pointer analysis
//	}
//
//	// Find edges originating from the main package.
//	// By converting to strings, we de-duplicate nodes
//	// representing the same function due to context sensitivity.
//	var edges []string
//	callgraph.GraphVisitEdges(result.CallGraph, func(edge *callgraph.Edge) error {
//		caller := edge.Caller.Func
//		if caller.Pkg == mainPkg {
//			edges = append(edges, fmt.Sprint(caller, " --> ", edge.Callee.Func))
//		}
//		return nil
//	})
//
//	// Print the edges in sorted order.
//	sort.Strings(edges)
//	for _, edge := range edges {
//		fmt.Println(edge)
//	}
//	fmt.Println()
//
//	// Print the labels of (C).f(m)'s points-to set.
//	fmt.Println("m may point to:")
//	var labels []string
//	for _, l := range result.Queries[Cfm].PointsTo().Labels() {
//		label := fmt.Sprintf("  %s: %s", prog.Fset.Position(l.Pos()), l)
//		labels = append(labels, label)
//	}
//	sort.Strings(labels)
//	for _, label := range labels {
//		fmt.Println(label)
//	}
//
//}

package main

import (
	"flag"
	"fmt"
	"go/ast"
	"os"
	"path"
	"strings"

	"golang.org/x/tools/go/packages"
)

var (
	rootPkg             *packages.Package
	toTidyPkgPathPrefix = "github.com/zhjp0/goes/tidy/tests/internal"

	// map["<package_path>"]map["<ident>"]true
	refMap = map[string]map[string]bool{}
)

func main() {
	flag.Parse()

	// Many tools pass their command-line arguments (after any flags)
	// uninterpreted to packages.Load so that it can interpret them
	// according to the conventions of the underlying build system.
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedSyntax | packages.NeedDeps | packages.NeedModule | packages.NeedImports,
	}
	pkgs, err := packages.Load(cfg, flag.Args()...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "load: %v\n", err)
		os.Exit(1)
	}
	if packages.PrintErrors(pkgs) > 0 {
		os.Exit(1)
	}

	if len(pkgs) < 1 {
		fmt.Fprintf(os.Stderr, "no packages found")
	}

	rootPkg = pkgs[0]
	searchPkg(pkgs[0])

	//for _, pkg := range pkgs {
	//	searchPkg(pkg)
	//}
}

func searchPkg(pkg *packages.Package) {
	if !needSearch(pkg) {
		return
	}
	for _, f := range pkg.Syntax {
		searchFile(f)
	}

	for _, pkg := range pkg.Imports {
		searchPkg(pkg)
	}
}

func searchFile(f *ast.File) {
	if f == nil {
		return
	}

	fmt.Println("\tsearching file:", f.Name)
	//fmt.Printf("name: %s, path: %v\n", f.Imports[0].Name, f.Imports[0].Path)

	var (
		// 包含目标路径的 imports
		importSet = map[string]bool{}
	)
	for _, imp := range f.Imports {
		p := imp.Path.Value
		p = p[1:len(p)-1]
		//fmt.Println(p, len(p))
		if strings.HasPrefix(p, toTidyPkgPathPrefix) {
			importSet[getImportName(imp)] = true
		}
	}
	fmt.Printf("\timportSet: %v\n", importSet)

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.SelectorExpr:
			//fmt.Println(x.X)
			expr := x.X
			if ident, ok := expr.(*ast.Ident); ok {
				fmt.Printf("\t\tX: %s, Selector: %s; ", ident.Name, x.Sel)
				name := ident.Name
				if _, exists := importSet[name]; exists {
					fmt.Printf("referenced.")
				}
				fmt.Println()
			}
		}
		return true
	})
}

func needSearch(pkg *packages.Package) bool {
	if pkg == nil {
		return false
	}
	need := strings.HasPrefix(pkg.PkgPath, rootPkg.PkgPath)
	fmt.Print("pkg: ", pkg.PkgPath, ", ")
	if need {
		fmt.Println("searching...")
	} else {
		fmt.Println("skip...")
	}
	return need
	//return true
}

func getImportName(spec *ast.ImportSpec) string {
	if spec == nil {
		return ""
	}
	if spec.Name != nil {
		return spec.Name.Name
	}
	p := spec.Path.Value
	p = p[1:len(p)-1]
	return path.Base(p)
}

//package main
//
//import (
//	"go/ast"
//	"go/parser"
//	"go/token"
//)
//
//func printAST(code string) {
//	// Create the AST by parsing src.
//	fset := token.NewFileSet() // positions are relative to fset
//	f, err := parser.ParseFile(fset, "", code, 0)
//	if err != nil {
//		panic(err)
//	}
//
//	// Print the AST.
//	ast.Print(fset, f)
//}
//
//func main() {
//	code := `
//package xxx
//
//import (
//	accountCli "clients/account"
//)
//
//func callRPC() {
//	accountCli.Create()
//}
//`
//
//	printAST(code)
//}
