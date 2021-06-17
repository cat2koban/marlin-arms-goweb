func index(w http.ResponseWriter, r *http.Request) {
  threads, err := data.Threads(); if err == nil {
    _, err := session(writer, request)
    generateHTML(writer, threads, "layout", "public.navbar", "index")
  } else {
    generateHTML(writer, threads, "layout", "private.navbar", "index")
  }
}