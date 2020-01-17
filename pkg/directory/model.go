package directory

import "encoding/xml"

type include struct {
	XMLName xml.Name `xml:"include"`
	Domain  domain
}

type domain struct {
	XMLName   xml.Name `xml:"domain"`
	Name      string   `xml:"name,attr"`
	Params    params
	Variables variables
	Groups    groups
}

type params struct {
	XMLName xml.Name `xml:"params"`
	Param   []param
}

type variables struct {
	XMLName  xml.Name `xml:"variables"`
	Variable []variable
}

type groups struct {
	XMLName xml.Name `xml:"groups"`
	Group   []group
}

type param struct {
	XMLName xml.Name `xml:"param"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

type variable struct {
	XMLName xml.Name `xml:"variable"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

type group struct {
	XMLName xml.Name `xml:"group"`
	Name    string   `xml:"name,attr"`
	Users   users
}

type users struct {
	XMLName xml.Name `xml:"users"`
	User    []user
}

type user struct {
	XMLName   xml.Name `xml:"user"`
	ID        string   `xml:"id,attr"`
	Params    params
	Variables variables
}
