package packages

import "path"

// FilterStdPkgs
func FilterStdPkgs(pkgs []string) []string {

	var list []string

	for _, pkg := range pkgs {

		// skip CGO and any relative import paths
		if pkg == "C" || pkg[0] == '.' {
			continue
		}

		// if the package is part of the golang standard library, skip it
		name := path.Base(pkg)
		if stds, ok := stdpkgs[name]; ok {
			for _, stdpkg := range stds {
				if pkg == stdpkg.path {
					goto SKIP
				}
			}
		}

		// if the import path doens't exists in pkgs, add it
		list = append(list, pkg)
	SKIP: // skips the appending of packages that are already present
	}

	return list
}
