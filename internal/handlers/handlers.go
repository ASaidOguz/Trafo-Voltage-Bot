package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/CloudyKit/jet/v6"
)

var Trafolar []Trafo
var trafo Trafo
var err error
var info string

type Trafo struct {
	TrafoName     string  `json:"trafo_name"`
	TrafoPow      int     `json:"trafo_pow"`
	TrafoCableL   float64 `json:"trafo_cable"`
	TrafoCabConst float64 `json:"trafo_cconst"`
}

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
)

// Home displays the home page with some sample data
func Home(w http.ResponseWriter, r *http.Request) {

	data := make(jet.VarMap)
	data.Set("user_first_name", "Ahmet Said")
	data.Set("user_last_name", "Oğuz")

	trafo.TrafoName = r.FormValue("trafo_name")
	data.Set("trafo", trafo.TrafoName)

	trafo.TrafoPow, err = strconv.Atoi(r.FormValue("trafo_power"))
	if err != nil {
		data.Set("trafopow", "Please enter valid value, use no point or comma!!")
		data.Set("unit", "")
	} else {
		data.Set("trafopow", trafo.TrafoPow)
		data.Set("unit", "kVA")
	}
	trafo.TrafoCableL, err = strconv.ParseFloat(r.FormValue("trafo_cable_length"), 32)
	if err != nil {
		data.Set("trafo_cable_len", "Please enter valid value, use point for floating values!!")
		data.Set("unit_cable", "")
	} else {
		formattedFloat1 := strconv.FormatFloat(trafo.TrafoCableL, 'f', 3, 64)

		data.Set("trafo_cable_len", formattedFloat1)
		data.Set("unit_cable", "km")
	}
	trafo.TrafoCabConst, err = strconv.ParseFloat(r.FormValue("trafo_cable_const"), 64)
	if err != nil {
		data.Set("trafo_cable_const", "Please enter valid value, use point for floating values!!")

	} else {
		formattedFloat2 := strconv.FormatFloat(trafo.TrafoCabConst, 'f', 3, 64)

		data.Set("trafo_cable_const", formattedFloat2)

	}
	if trafo.TrafoName != "" {
		Trafolar = append(Trafolar, trafo)
		data.Set("info", "")
	} else {
		info = "Please fill the blanks"
		data.Set("info", info)
	}

	dow := []string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
	}
	data.Set("dow", dow)
	err := renderPage(w, "home.jet", data)
	if err != nil {
		_, _ = fmt.Fprint(w, "Error executing template:", err)
	}

}

// SendData displays a second page
func SendData(w http.ResponseWriter, r *http.Request) {

	data := make(jet.VarMap)
	data.Set("trafolar", Trafolar)
	err := renderPage(w, "solution.jet", data)
	Trafolar = []Trafo{}
	if err != nil {
		_, _ = fmt.Fprint(w, "Error executing template:", err)
	}
}

// renderPage renders the page using Jet templates
func renderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {
	view, err := views.GetTemplate(tmpl)
	if err != nil {
		log.Println("Unexpected template err:", err.Error())
		return err
	}

	err = view.Execute(w, data, nil)
	if err != nil {
		log.Println("Error executing template:", err)
		return err
	}
	return nil
}
