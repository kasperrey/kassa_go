package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func prijsPagina(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	params := u.Query()
	naam := params.Get("name")
	password := params.Get("password")
	if naam == "" {
		naam = "Bil-Bilbos"
	}

	product := getProductByName(naam)
	producten := getAllProducts()
	bankkaarten := getAllBankkaarts()
	for _, b := range bankkaarten {
		if password == b.code {
			htmlString := fmt.Sprintf("<html>\n"+
				"<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n"+
				"<style>\n"+
				"body {\n"+
				"margin: 0;\n"+
				"font-family: Arial, Helvetica, sans-serif;\n"+
				"}\n"+
				".topnav {\n"+
				"overflow: hidden;\n"+
				"background-color: #333;\n"+
				"}\n"+
				".topnav a {\n"+
				"float: left;\n"+
				"color: #f2f2f2;\n"+
				"text-align: center;\n"+
				"padding: 14px 16px;\n"+
				"text-decoration: none;\n"+
				"font-size: 17px;\n"+
				"}\n"+
				".topnav a:hover {\n"+
				"background-color: #ddd;\n"+
				"color: black;\n"+
				"}\n"+
				".topnav a.active {\n"+
				"background-color: #04AA6D;\n"+
				"color: white;\n"+
				"}\n"+
				"</style>\n"+
				"<div class=\"topnav\">\n"+
				"<a href=\"/?password=%s\">Home</a>\n"+
				"<a href=\"/kassa/producten/maken?password=%s\">Product maken</a>\n"+
				"<a class=\"active\" href=\"/kassa/producten/prijs?password=%s\">Producten prijs</a>\n"+
				"</div>\n", password, password, password)

			htmlString += "<select name='producten' id='producten'>\n"
			for _, p := range producten {
				htmlString += fmt.Sprintf("    <option onclick=\"window.location.assign(%s)\" value='%s'>%s</option>\n",
					"&quot;prijs?name="+p.naam+"&password="+password+"&quot;", p.naam, p.naam)
			}
			htmlString += "</select>\n"
			if product.prijs == 0 {
				htmlString += fmt.Sprintf("<p>Geen product met de naam '%s'</p>", product.naam)
			} else {
				htmlString += fmt.Sprintf("<p>De prijs van %s is %.2f</p>", product.naam, product.prijs)
			}
			htmlString += "<html>\n"
			w.Write([]byte(fmt.Sprintf(htmlString)))
			return
		}
	}
	w.Write([]byte("<head>\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n    <style>\n        body {\n            margin: 0;\n            font-family: Arial, Helvetica, sans-serif;\n        }\n\n        button {\n            background-color: #4CAF50;\n            width: 100%;\n            color: orange;\n            padding: 15px;\n            margin: 10px 0px;\n            border: none;\n            cursor: pointer;\n        }\n\n        form {\n            border: 3px solid #f1f1f1;\n        }\n\n        input[type=password] {\n            width: 100%;\n            margin: 8px 0;\n            padding: 12px 20px;\n            display: inline-block;\n            border: 2px solid green;\n            box-sizing: border-box;\n        }\n\n        button:hover {\n            opacity: 0.7;\n        }\n\n        .container {\n            padding: 25px;\n            background-color: lightblue;\n        }\n    </style>\n</head>\n<body>\n\n<form>\n    <div class=\"container\">\n        <label>Password : </label>\n        <input type=\"password\" placeholder=\"Enter Password\" name=\"password\" required>\n        <button type=\"submit\">Login</button>\n    </div>\n</form>\n</body>"))
}

func homePagina(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	params := u.Query()
	password := params.Get("password")

	bankkaarten := getAllBankkaarts()
	for _, b := range bankkaarten {
		if password == b.code {
			htmlString := fmt.Sprintf("<html>\n"+
				"<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n"+
				"<style>\n"+
				"body {\n"+
				"margin: 0;\n"+
				"font-family: Arial, Helvetica, sans-serif;\n"+
				"}\n"+
				".topnav {\n"+
				"overflow: hidden;\n"+
				"background-color: #333;\n"+
				"}\n"+
				".topnav a {\n"+
				"float: left;\n"+
				"color: #f2f2f2;\n"+
				"text-align: center;\n"+
				"padding: 14px 16px;\n"+
				"text-decoration: none;\n"+
				"font-size: 17px;\n"+
				"}\n"+
				".topnav a:hover {\n"+
				"background-color: #ddd;\n"+
				"color: black;\n"+
				"}\n"+
				".topnav a.active {\n"+
				"background-color: #04AA6D;\n"+
				"color: white;\n"+
				"}\n"+
				"</style>\n"+
				"<div class=\"topnav\">\n"+
				"<a class=\"active\" href=\"/?password=%s\">Home</a>\n"+
				"<a href=\"/kassa/producten/maken?password=%s\">Product maken</a>\n"+
				"<a href=\"/kassa/producten/prijs?password=%s\">Producten prijs</a>\n"+
				"</div>\n"+
				"<html>", password, password, password)
			w.Write([]byte(fmt.Sprintf(htmlString)))
			return
		}
	}
	w.Write([]byte("<head>\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n    <style>\n        body {\n            margin: 0;\n            font-family: Arial, Helvetica, sans-serif;\n        }\n\n        button {\n            background-color: #4CAF50;\n            width: 100%;\n            color: orange;\n            padding: 15px;\n            margin: 10px 0px;\n            border: none;\n            cursor: pointer;\n        }\n\n        form {\n            border: 3px solid #f1f1f1;\n        }\n\n        input[type=password] {\n            width: 100%;\n            margin: 8px 0;\n            padding: 12px 20px;\n            display: inline-block;\n            border: 2px solid green;\n            box-sizing: border-box;\n        }\n\n        button:hover {\n            opacity: 0.7;\n        }\n\n        .container {\n            padding: 25px;\n            background-color: lightblue;\n        }\n    </style>\n</head>\n<body>\n\n<form>\n    <div class=\"container\">\n        <label>Password : </label>\n        <input type=\"password\" placeholder=\"Enter Password\" name=\"password\" required>\n        <button type=\"submit\">Login</button>\n    </div>\n</form>\n</body>"))
}

func makenPagina(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	params := u.Query()
	password := params.Get("password")
	naam := params.Get("naam")
	prijs := params.Get("prijs")
	barcode := params.Get("barcode")

	bankkaarten := getAllBankkaarts()
	for _, b := range bankkaarten {
		if password == b.code {
			htmlString := fmt.Sprintf("<html>\n"+
				"<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n"+
				"<style>\n"+
				"body {\n"+
				"margin: 0;\n"+
				"font-family: Arial, Helvetica, sans-serif;\n"+
				"}\n"+
				".topnav {\n"+
				"overflow: hidden;\n"+
				"background-color: #333;\n"+
				"}\n"+
				".topnav a {\n"+
				"float: left;\n"+
				"color: #f2f2f2;\n"+
				"text-align: center;\n"+
				"padding: 14px 16px;\n"+
				"text-decoration: none;\n"+
				"font-size: 17px;\n"+
				"}\n"+
				".topnav a:hover {\n"+
				"background-color: #ddd;\n"+
				"color: black;\n"+
				"}\n"+
				".topnav a.active {\n"+
				"background-color: #04AA6D;\n"+
				"color: white;\n"+
				"}\n"+
				"</style>\n"+
				"<div class=\"topnav\">\n"+
				"<a href=\"/?password=%s\">Home</a>\n"+
				"<a class=\"active\" href=\"/kassa/producten/maken?password=%s\">Product maken</a>\n"+
				"<a href=\"/kassa/producten/prijs?password=%s\">Producten prijs</a>\n"+
				"</div>\n"+
				"<div>\n"+
				"<input placeholder=\"naam\" id=\"naam\" required>\n"+
				"<input placeholder=\"prijs\" id=\"prijs\" required>\n"+
				"<input placeholder=\"barcode\" id=\"barcode\" required>\n"+
				"<button onclick=\"f()\">maak</button>\n"+
				"</div>"+
				"<script>\n"+
				"function f() {\n"+
				"var n = document.getElementById(\"naam\")\n"+
				"var p = document.getElementById(\"prijs\")\n"+
				"var b = document.getElementById(\"barcode\")\n"+
				"window.location.assign(\"?naam=\"+n.value+\"&prijs=\"+p.value+\"&barcode=\"+b.value+\"&password=%s\")\n"+
				"}\n"+
				"</script>"+
				"<html>", password, password, password, password)
			if (naam != "") && (prijs != "") && (barcode != "") {
				prijsFloat, _ := strconv.ParseFloat(prijs, 64)
				barcodeInt, _ := strconv.ParseInt(barcode, 10, 64)
				createProduct(Product{id: 0, naam: naam, prijs: prijsFloat, barcode: barcodeInt})
			}
			w.Write([]byte(fmt.Sprintf(htmlString)))
			return
		}
	}
	w.Write([]byte("<head>\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n    <style>\n        body {\n            margin: 0;\n            font-family: Arial, Helvetica, sans-serif;\n        }\n\n        button {\n            background-color: #4CAF50;\n            width: 100%;\n            color: orange;\n            padding: 15px;\n            margin: 10px 0px;\n            border: none;\n            cursor: pointer;\n        }\n\n        form {\n            border: 3px solid #f1f1f1;\n        }\n\n        input[type=password] {\n            width: 100%;\n            margin: 8px 0;\n            padding: 12px 20px;\n            display: inline-block;\n            border: 2px solid green;\n            box-sizing: border-box;\n        }\n\n        button:hover {\n            opacity: 0.7;\n        }\n\n        .container {\n            padding: 25px;\n            background-color: lightblue;\n        }\n    </style>\n</head>\n<body>\n\n<form>\n    <div class=\"container\">\n        <label>Password : </label>\n        <input type=\"password\" placeholder=\"Enter Password\" name=\"password\" required>\n        <button type=\"submit\">Login</button>\n    </div>\n</form>\n</body>"))
}

func startWeb() {
	port := "8000"
	mux := http.NewServeMux()

	mux.HandleFunc("/kassa/producten/prijs", prijsPagina)
	mux.HandleFunc("/kassa/producten/maken", makenPagina)
	mux.HandleFunc("/", homePagina)
	http.ListenAndServe(":"+port, mux)
}
