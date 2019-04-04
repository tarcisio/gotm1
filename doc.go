/*
Package gotm1 provides an idiomatic golang implementation of TM1 rest API.

	tm1 := gotm1.New(
		http.NewSimpleHTTPClient(
			http.WithHost("tm1.example.com"),
			http.WithPort(8091),
			http.WithUsername("c6548654"),
			http.WithPassword("p@$$w0rd"),
			http.WithNamespace("CompanyLDAP"),
			http.WithCognosIntegratedLogin(),
			http.WithTimeoutConnection(5),
		),
	)

	defer tm1.Logout()

	cubes, err := tm1.GetCubes()
	if err != nil {
		log.Fatal(err)
	}

	for _, cube := range cubes {
		fmt.Println(cube.Name)
	}

This package wraps around the oficial API so it doesn't try to be perfect.

*/
package gotm1
