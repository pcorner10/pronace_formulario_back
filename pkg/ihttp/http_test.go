package ihttp

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupRoutes(t *testing.T) {
	tt := []struct {
		desc       string
		method     string
		uri        string
		payload    []byte
		headers    map[string]string
		StatusCode int
	}{
		{
			desc:   "CreateForm0",
			method: "POST",
			uri:    "http://localhost:8080/api/forms/form0",
			payload: []byte(`{
				"email": "princon0@ucol.mx",
				"zona": {
				  "municipio": "Cuauhtemoc",
				  "localidad": "Cuauhtemoc",
				  "colonia": "8",
				  "manzana": "2",
				  "lote": "200"
				},
				"integrante": [
				  {
					"sexo": "hombre",
					"edad": "27",
					"parentesco": "encuestado",
					"escolaridad": "posgrado",
					"trabaja": "si",
					"horas_dia": "6",
					"dias_semana": "5",
					"ocupacion": "DEV",
					"fumador": "no",
					"tiempo_ciudad": "12"
				  },
				  {
					"sexo": "mujer",
					"edad": "27",
					"parentesco": "padre",
					"escolaridad": "secundaria",
					"trabaja": "si",
					"horas_dia": "5",
					"dias_semana": "6",
					"ocupacion": "SEÑP",
					"fumador": "no",
					"tiempo_ciudad": "13"
				  }
				]
			  }`),
		},
	}

	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			var payload io.Reader
			if tc.payload != nil {
				payload = bytes.NewBuffer(tc.payload)
			}
			req, err := http.NewRequest(tc.method, tc.uri, payload)
			assert.NoError(t, err)

			for k, v := range tc.headers {
				req.Header.Set(k, v)
			}

			resp, err := http.DefaultClient.Do(req)
			assert.NoError(t, err)

			assert.Equal(t, tc.StatusCode, resp.StatusCode)
		})
	}
}
