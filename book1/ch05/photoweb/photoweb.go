package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	UPLOAD_DIR = "./uploads"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		io.WriteString(w, "<form method=\"POST\" action=\"/upload\" "+
			"enctype=\"multipart/form-data\">"+
			"Choose an image to upload: <input name=\"image\" type=\"file\" />"+
			"<input type=\"submit\" value=\"Upload\" />"+
			"</form>")

		return
	}

	if r.Method == "POST" {
		//multipart.File
		//*multipart.FileHeader
		f, h, err := r.FromFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		filename := h.Filename
		defer f.Close()

		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer t.Close()

		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)

	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FromValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir("./upload")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var listHtml string
	for _, fileInfo := range fileInfoArr {
		imgid := fileInfo.Name
		listHtml += "<li><a href=\"/view?id=" + imgid + "\">imgid</a></li>"
	}

	io.WriteString(w, "<ol>"+listHtml+"</ol>")
}

func main() {
	http.HandleFunc("/", listHandler)
	http.HandleFunc("/view", viewHandler)
	http.HandleFunc("/upload", uploadHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
