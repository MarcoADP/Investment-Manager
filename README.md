# Investment-Manager

- https://go.dev/doc/modules/layout
- https://github.com/golang-standards/project-layout
- https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2
- https://gist.github.com/ayoubzulfiqar/9f1a34049332711fddd4d4b2bfd46096

Packages:

- **/cmd:** The directory name for each application should match the name of the executable you want to have (e.g., /cmd/investment-manager)
- **/internal:** Private application and library code. You use internal directories to make packages private. If you put a package inside an internal directory, then other packages can’t import it unless they share a common ancestor. And it’s the only directory named in Go’s documentation and has special compiler treatment.
- **/pkg:** Library code that's ok to use by external applications (e.g., /pkg/db/repository).
- **/vendor:** Application dependencies (managed manually or by your favorite dependency management tool like the new built-in Go Modules feature). The go mod vendor command will create the /vendor directory for you.

