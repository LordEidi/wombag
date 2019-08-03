package wombag

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
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func EnsureTemplateFilesExist() {

	tmplDir := getConfigDefault("templates.dir", "./templates/")

	if tmplDir[len(tmplDir)-1:] != "/" {
		tmplDir += "/"
	}

	if _, err := os.Stat(tmplDir); os.IsNotExist(err) {
		os.MkdirAll(tmplDir, 0700)
	}

	err := ensureSpecificFile(tmplDir+getConfigDefault("templates.auth", "auth.tmpl"), defTmplAuth)
	if err != nil {
		LogError("Ensure template failed.", logrus.Fields{"name": "auth.tmpl", "path": tmplDir, "error": err})
	}

	err = ensureSpecificFile(tmplDir+getConfigDefault("entry.auth", "entry.tmpl"), defTmplEntry)
	if err != nil {
		LogError("Ensure template failed.", logrus.Fields{"name": "entry.tmpl", "path": tmplDir, "error": err})
	}

	err = ensureSpecificFile(tmplDir+getConfigDefault("entries.auth", "entries.tmpl"), defTmplEntries)
	if err != nil {
		LogError("Ensure template failed.", logrus.Fields{"name": "entries.tmpl", "path": tmplDir, "error": err})
	}

	err = ensureSpecificFile(tmplDir+getConfigDefault("templates.tags", "tags.tmpl"), defTmplTags)
	if err != nil {
		LogError("Ensure template failed.", logrus.Fields{"name": "tags.tmpl", "path": tmplDir, "error": err})
	}
}

func getConfigDefault(config string, defaultValue string) string {

	s := GetStringFromConfig(config)

	if len(s) == 0 {
		s = defaultValue
	}

	return s
}

func ensureSpecificFile(pathToTemplate string, byTemplate []byte) error {

	if _, err := os.Stat(pathToTemplate); os.IsNotExist(err) {

		e2 := ioutil.WriteFile(pathToTemplate, byTemplate, 0700)

		return e2

	} else {

		return err
	}
}

var defTmplAuth = []byte(`{
    "access_token": "{{ .AccessToken }}",
    "expires_in": {{ .ExpirationDate }},
    "refresh_token": "{{ .RefreshToken }}",
    "scope": null,
    "token_type": "{{ .TokenType }}"
}`)

var defTmplEntry = []byte(`
    {
    "_links": {
        "self": {
            "href": "/api/entries/{{ .Entry.EntryId }}
        }
    },
    "content": "{{ .Entry.GetContentJSON }}\n",
    "created_at": "{{ .Entry.CrtDat }}",
    "domain_name": "{{ .Entry.Domain }}",
    "id": {{ .Entry.EntryId }},
    "is_archived": {{ .Entry.Archived }},
    "is_starred": {{ .Entry.Starred }},
    "language": "{{ .Entry.Language }}",
    "mimetype": "text/html",
    "preview_picture": "{{ .Entry.PreviewPicture }}",
    "reading_time": 2,
    "tags": [{{ .Entry.GetTags }}],
    "title": "{{ .Entry.GetTitleJSON }}",
    "updated_at": "{{ .Entry.UpdDat }}",
    "url": "{{ .Entry.URL }}",
    "user_email": "",
    "user_id": 1,
    "user_name": "wallabag"
}
`)

var defTmplEntries = []byte(`
    {
    "_embedded": {
        "items": [
        {{range $key,$entry := .Entries}}
        {
            "_links": {
                "self": {
                    "href": "/api/entries/{{ $entry.EntryId }}"
                }
            },
            "annotations": [],
            "content": "{{ $entry.GetContentJSON }}\n",
            "created_at": "{{ .Entry.CrtDat }}",
            "domain_name": "{{ .Entry.Domain }}",
            "id": {{ $entry.EntryId }},
            "is_archived": {{ $entry.Archived }},
            "is_starred": {{ $entry.Starred }},
            "language": "{{ .Entry.Language }}",
            "mimetype": "text/html",
            "preview_picture": "{{ .Entry.PreviewPicture }}",
            "reading_time": 2,
            "tags": [{{ $entry.GetTags }}],
            "title": "{{ $entry.GetTitleJSON }}",
            "updated_at": "{{ .Entry.UpdDat }}",
            "url": "{{ $entry.URL }}",
            "user_email": "",
            "user_id": 1,
            "user_name": "wallabag"
        },
        {{end}}
        ]
    },
    "_links": {
        "first": {
            "href": "http://{{ .Server }}:{{ .Port }}/api/entries?page={{ .Page }}&perPage={{ .Limit }}"
        },
        "last": {
            "href": "http://{{ .Server }}:{{ .Port }}/api/entries?page={{ .Page }}&perPage={{ .Limit }}"
        },
        "self": {
            "href": "http://{{ .Server }}:{{ .Port }}/api/entries?page={{ .Page }}&perPage={{ .Limit }}"
        }
    },
    "limit": {{ .Limit }},
    "page": {{ .Page }},
    "pages": {{ .Page }},
    "total": {{ .Page }}
}
`)

var defTmplTags = []byte(`
[
  {{range $key,$tag := .Tags}}
    {
        "slug":"{{ $tag.Slug }}",
        "label":"{{ $tag.Label }}",
        "id":{{ $tag.TagId }}
    },
  {{end}}
]
`)
