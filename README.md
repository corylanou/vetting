# vetting

This project illustrates an example of the difference between calling
`go vet` vs directly calling the `vet` tool via `xargs`.

To see the difference, first run:

```sh
go vet ./...
```

You should see no ouptput (everything ran clean).

Now run the following command:

```sh
GOTOOLDIR=$(go env GOTOOLDIR)
git ls-files | grep '.go$' | xargs $GOTOOLDIR/vet 
```

You will get the following result:

```sh
stuff.go:17: foo.Rows composite literal uses unkeyed fields
```
