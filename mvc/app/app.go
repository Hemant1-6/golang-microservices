package app

import "net/http"

func StartApp()  {
    if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}