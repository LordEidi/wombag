package handler
/*-----------------------------------------------------------------------------
 **
 ** - Wombag -
 **
 ** the alternative, native backend for your Wallabag apps
 **
 ** Copyright 2017-18 by SwordLord - the coding crew - http://www.swordlord.com
 ** and contributing authors
 **
 ** This program is free software; you can redistribute it and/or modify it
 ** under the terms of the GNU Affero General Public License as published by the
 ** Free Software Foundation, either version 3 of the License, or (at your option)
 ** any later version.
 **
 ** This program is distributed in the hope that it will be useful, but WITHOUT
 ** ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
 ** FITNESS FOR A PARTICULAR PURPOSE.  See the GNU Affero General Public License
 ** for more details.
 **
 ** You should have received a copy of the GNU Affero General Public License
 ** along with this program. If not, see <http://www.gnu.org/licenses/>.
 **
 **-----------------------------------------------------------------------------
 **
 ** Original Authors:
 ** LordEidi@swordlord.com
 ** LordLightningBolt@swordlord.com
 **
-----------------------------------------------------------------------------*/
import (
	"fmt"
	"github.com/gorilla/schema"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
	"swordlord.com/wombag/tablemodule"
	"swordlord.com/wombagd/render"
	"swordlord.com/wombagd/respond"
)

// standard query params usually sent with requests
type QueryParams struct {

	Url			string	`schema:"url" validate:"omitempty,url"` 		// Url for the entry.
	Title		string 	`schema:"title" validate:"-"` 	// Optional, we'll get the title from the page.
	Tags 		string 	`schema:"tags" validate:"-"` 		// tag1,tag2,tag3 	a comma-separated list of tags.
	Starred 	int 	`schema:"starred" validate:"omitempty,min=0,max=1"` // 1 or 0 	entry already starred
	Archive 	int 	`schema:"archive" validate:"omitempty,min=0,max=1"` // 1 or 0 	entry already archived
}

// the form params which are sent to get an access token
type AccessTokenReqParams struct {

	ClientID		string	`schema:"client_id" form:"client_id" json:"client_id"` 		// aa
	ClientSecret	string	`schema:"client_secret" form:"client_secret" json:"client_secret"` // aa
	GrantType		string	`schema:"grant_type" form:"grant_type" json:"grant_type"` 		// password
	Password		string	`schema:"password" form:"password" json:"password"` 			// aa
	UserName		string	`schema:"username" form:"username" json:"username"` 			// aa
}

type oAuth2 struct {

	AccessToken 	string 	`json:"access_token"`	//	"..."
	ExpirationDate 	uint 	`json:"expires_in"` 	// 3600,
	RefreshToken 	string 	`json:"refresh_token"` 	// "...",
	Scope 			string 	`json:"scope"` 			// null,
	TokenType 		string 	`json:"token_type"` 	// "bearer"
}

func getNewOAuth2() oAuth2 {

	var oa oAuth2

	oa.ExpirationDate = 3600
	oa.Scope = ""
	oa.TokenType = "bearer"

	return oa
}

func OnRoot(w http.ResponseWriter, req *http.Request){

	respond.WithMessage(w, http.StatusOK, "Welcome to Wombag")
}

func OnRetrieveVersionNumber(w http.ResponseWriter, req *http.Request){

	// TODO randomise the reply text
	respond.WithMessage(w, http.StatusOK, "This is Wombag")
}

func OnOAuth(w http.ResponseWriter, req *http.Request){

	var form AccessTokenReqParams

	err1 := bind(&form, req)

	if err1 != nil {
		fmt.Printf("Error when binding %v\n", err1)
		respond.WithMessage(w, http.StatusBadRequest, "An Error occured: " + err1.Error())
	}

	if form.ClientID == "" || form.ClientSecret == "" {
		log.Printf("Missing authentication credentials. Access denied.\n" )
		respond.WithMessage(w, http.StatusUnauthorized, "Access not authorised")
		return
	}

	device, err := tablemodule.ValidateDeviceInDB(form.ClientID, form.ClientSecret)

	if err != nil {
		log.Printf("Wrong Authentication Request. Access denied.\n" )
		respond.WithMessage(w, http.StatusUnauthorized, "Access not authorised")
		return
	}

	oauth := getNewOAuth2()

	oauth.AccessToken = device.AccessToken
	oauth.RefreshToken = device.AccessToken

	wtext := render.WombagText{}
	wtext.Data = oauth
	respond.Render(w, http.StatusOK, wtext)
}

func bind(obj interface{}, req *http.Request) (e error){

	// parse the URL and form and put result in req.Form and req.PostForm
	err := req.ParseForm()

	if err != nil {
		log.Printf("Binding: Parsing Form returned error: %s.\n", err )
		return err
	}

	// Decoder decodes values from a map[string][]string to a struct.
	// -> Gorilla Schema
	decoder := schema.NewDecoder()

	// req.PostForm is a map of our POST form values
	// req.Form contains both, URL and body values, which is why we decide on .Form
	errDec := decoder.Decode(obj, req.Form)

	if errDec != nil {
		log.Printf("Binding: Decoding Form returned error: %s.\n", err )
		return err
	}

	return
}

func isValid(s interface{}) (bool) {

	isValid := true

	v := validator.New()
	errV := v.Struct(s)

	if errV != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := errV.(*validator.InvalidValidationError); ok {
			fmt.Println(errV)
			return false
		}

		for _, err := range errV.(validator.ValidationErrors) {
			fmt.Printf("Validation error: %s, %s, %s, %s\n", err.Namespace(), err.Field(), err.Tag(), err.Value())

			isValid = false
		}
	}

	return isValid
}