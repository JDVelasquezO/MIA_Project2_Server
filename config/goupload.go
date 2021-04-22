package config

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

//UploadSetup Configuracion Upload
type UploadSetup struct {
	NameInput           string
	MimesType           []string
	PermitirExtenciones bool
	RutaArchivos        string
	LimiteUpload        float32
	Ext                 string
}

//Upload sube un archivo a la carpeta indicada en el servidor
func (u *UploadSetup) Upload(r *http.Request, name string) error {
	file, header, err := r.FormFile(u.NameInput)
	if err != nil {
		return err
	}
	defer file.Close()

	if header.Size > u.getLimiteUpload() {
		return errors.New("El tamaño del archivo supera el limite permitido: ")
	}

	mimeType := header.Header.Get("Content-Type")
	existe := u.buscarMime(mimeType)
	if (u.PermitirExtenciones && !existe) || (!u.PermitirExtenciones && existe) {
		return errors.New("Extencion no permitida: " + mimeType)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	segmentados := strings.Split(header.Filename, ".")
	extencion := segmentados[len(segmentados)-1]
	u.Ext = extencion
	return ioutil.WriteFile(u.RutaArchivos+name+"."+extencion, data, 0666)
}

func (u UploadSetup) buscarMime(mimeType string) bool {
	for _, mime := range u.MimesType {
		if mimeType == mime {
			return true
		}
	}
	return false
}

func (u UploadSetup) getLimiteUpload() int64 {
	return int64(u.LimiteUpload * 1024)
}