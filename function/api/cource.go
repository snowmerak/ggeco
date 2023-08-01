package api

import (
	"encoding/base64"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/gen/bean"
	"github.com/snowmerak/ggeco/lib/service/courses"
	"net/http"
)

func GetCourse(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		courseIdValue := r.URL.Query().Get("course_id")
		courseId, _ := base64.URLEncoding.DecodeString(courseIdValue)
		authorIdValue := r.URL.Query().Get("author_id")
		authorId, _ := base64.URLEncoding.DecodeString(authorIdValue)
		courseName := r.URL.Query().Get("course_name")

		encoder := json.NewEncoder(wr)

		switch authorId {
		case nil:
			switch courseId {
			case nil:
				wr.WriteHeader(http.StatusBadRequest)
				wr.Write([]byte("Bad Request"))
				return
			default:
				data, err := courses.Get(container, courseId)
				if err != nil {
					wr.WriteHeader(http.StatusInternalServerError)
					wr.Write([]byte(err.Error()))
					return
				}
				if err := encoder.Encode(data); err != nil {
					wr.WriteHeader(http.StatusInternalServerError)
					wr.Write([]byte(err.Error()))
					return
				}
			}
		default:
			switch courseName {
			case "":
				data, err := courses.GetByAuthor(container, authorId)
				if err != nil {
					wr.WriteHeader(http.StatusInternalServerError)
					wr.Write([]byte(err.Error()))
					return
				}
				if err := encoder.Encode(data); err != nil {
					wr.WriteHeader(http.StatusInternalServerError)
					wr.Write([]byte(err.Error()))
					return
				}
			default:
				data, err := courses.GetByAuthorAndName(container, authorId, courseName)
				if err != nil {
					wr.WriteHeader(http.StatusInternalServerError)
					wr.Write([]byte(err.Error()))
					return
				}
				if err := encoder.Encode(data); err != nil {
					wr.WriteHeader(http.StatusInternalServerError)
					wr.Write([]byte(err.Error()))
					return
				}
			}
		}
	}
}
