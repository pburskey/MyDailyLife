package main

import (
	redis_utility "burskey/mydailylife/domain/internal/redis"
	"burskey/mydailylife/domain/internal/utility"
	"burskey/mydailylife/domain/package/dao"
	"burskey/mydailylife/domain/package/dao/redis"
	"burskey/mydailylife/domain/package/domain"
	"burskey/mydailylife/domain/package/service"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/gddo/httputil/header"
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strings"
)

var daoImpl dao.DAO
var taskService *service.TaskService

func main() {

	//arguments := os.Args[1:]
	//config := utility.LoadConfiguration(arguments[0])

	//mysqlconfiguration, err := mysql.Configure(config.MySQL)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer mysqlconfiguration.Close()

	redisConfiguration := utility.RedisConfiguration{
		Password: "letmein",
	}
	redisConnection := redis_utility.Factory(redisConfiguration)
	redisDAOImpl := redis.Factory(redisConnection)
	//daoImpl = mysql.Build(mysqlconfiguration, redisDAOImpl)
	taskService = service.Factory(redisDAOImpl)

	r := mux.NewRouter()
	r.Use(CORS)

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", get).Methods(http.MethodGet)
	api.HandleFunc("", post).Methods(http.MethodPost)
	api.HandleFunc("", put).Methods(http.MethodPut)
	api.HandleFunc("", delete).Methods(http.MethodDelete)

	//api.HandleFunc("/user/{userID}/comment/{commentID}", params).Methods(http.MethodGet)
	//api.HandleFunc("/schools", schools).Methods(http.MethodGet)
	//api.HandleFunc("/categories", categories).Methods(http.MethodGet)
	//api.HandleFunc("/dates", dates).Methods(http.MethodGet)
	//api.HandleFunc("/category/{aCategory}/metrics", metricsByCategory).Methods(http.MethodGet)
	//api.HandleFunc("/school/{aSchool}/metrics", metricsBySchool).Methods(http.MethodGet)
	//api.HandleFunc("/school/{aSchool}/category/{aCategory}/metrics", metricsBySchoolAndCategory).Methods(http.MethodGet)
	//api.HandleFunc("/school/{aSchool}/category/{aCategory}/metricDetails", metricDetailsBySchoolAndCategory).Methods(http.MethodGet)

	//api.HandleFunc("/date/{aDate}/metrics", metricsByDate).Methods(http.MethodGet)
	//api.HandleFunc("/metric/{aMetric}", metric).Methods(http.MethodGet)

	api.HandleFunc("/party", saveParty).Methods(http.MethodPost)
	api.HandleFunc("/party/{partyID}/task", saveTask).Methods(http.MethodPost)
	api.HandleFunc("/party/{partyID}", getParty).Methods(http.MethodGet)
	api.HandleFunc("/task/{taskID}/load", loadTask).Methods(http.MethodPost)
	api.HandleFunc("/task/{taskID}", getTask).Methods(http.MethodGet)
	api.HandleFunc("/task{taskID}/status/{status}", saveTaskInProgressStatus).Methods(http.MethodPost)

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"content-type"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowCredentials(),
	)
	cors(r)
	//log.Fatal(http.ListenAndServe(":8080", r))
	http.ListenAndServe(":8080", setHeaders(r))

}

func getParty(writer http.ResponseWriter, request *http.Request) {
	//var task *domain.Task

	pathParams := mux.Vars(request)
	writer.Header().Set("Content-Type", "application/json")

	id := ""

	if val, ok := pathParams["partyID"]; ok {
		id = val
	} else {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(`{"message": "need a task id"}`))
		return
	}

	party, err := taskService.GetParty(id)
	if err != nil {

	}

	fmt.Fprintf(writer, "Party: %+v", party)
}

func getTask(writer http.ResponseWriter, request *http.Request) {
	//var task *domain.Task

	pathParams := mux.Vars(request)
	writer.Header().Set("Content-Type", "application/json")

	id := ""

	if val, ok := pathParams["taskID"]; ok {
		id = val
	} else {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(`{"message": "need a task id"}`))
		return
	}

	task, err := taskService.GetTask(id)
	if err != nil {

	}

	fmt.Fprintf(writer, "Taswk: %+v", task)
}

func saveTaskInProgressStatus(writer http.ResponseWriter, request *http.Request) {

}

func loadTask(writer http.ResponseWriter, request *http.Request) {

	//var task *domain.Task

	pathParams := mux.Vars(request)
	writer.Header().Set("Content-Type", "application/json")

	taskID := ""

	if val, ok := pathParams["taskID"]; ok {
		taskID = val
	} else {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(`{"message": "need a task id"}`))
		return
	}

	task, err := taskService.GetTask(taskID)
	if err != nil {

	}

	fmt.Fprintf(writer, "Task: %+v", task)

}

//
//func save(aDateAndTime time.Time, dataMap map[string]map[string]*domain.CovidMetric, counter *utility.Counter, daoImpl *dao.DAO) {
//	if dataMap != nil && len(dataMap) > 0 {
//
//		for organization, schools := range dataMap {
//			//log.Println(organization)
//
//			for schoolName, metric := range schools {
//				//log.Println(schoolName)
//				//log.Println(metric)
//				daoImpl.SaveMetric(organization, schoolName, aDateAndTime, metric, counter)
//			}
//
//		}
//	}
//
//}

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world"}`))
}

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete called"}`))
}

func saveTask(w http.ResponseWriter, r *http.Request) {
	var task *domain.Task
	err := decodeJSONBody(w, r, &task)

	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			log.Print(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	partyID := ""

	if val, ok := pathParams["partyID"]; ok {
		partyID = val
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "need a party id"}`))
		return
	}

	task = &domain.Task{
		ID:          uuid.New().String(),
		Name:        task.Name,
		Description: task.Description,
		PartyId:     partyID,
	}
	if err := taskService.SaveTask(task); err != nil {

	}
	fmt.Fprintf(w, "%+v", task.ID)
}

func saveParty(w http.ResponseWriter, r *http.Request) {
	var party *domain.Person
	err := decodeJSONBody(w, r, &party)

	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			log.Print(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}
	party = &domain.Person{
		ID:    uuid.New().String(),
		First: party.First,
		Last:  party.Last,
	}

	if err := taskService.SaveParty(party); err != nil {

	}
	fmt.Fprintf(w, "%+v", party.ID)
}

//	func params(w http.ResponseWriter, r *http.Request) {
//		pathParams := mux.Vars(r)
//		w.Header().Set("Content-Type", "application/json")
//
//		userID := -1
//		var err error
//		if val, ok := pathParams["userID"]; ok {
//			userID, err = strconv.Atoi(val)
//			if err != nil {
//				w.WriteHeader(http.StatusInternalServerError)
//				w.Write([]byte(`{"message": "need a number"}`))
//				return
//			}
//		}
//
//		commentID := -1
//		if val, ok := pathParams["commentID"]; ok {
//			commentID, err = strconv.Atoi(val)
//			if err != nil {
//				w.WriteHeader(http.StatusInternalServerError)
//				w.Write([]byte(`{"message": "need a number"}`))
//				return
//			}
//		}
//
//		query := r.URL.Query()
//		location := query.Get("location")
//
//		w.Write([]byte(fmt.Sprintf(`{"userID": %d, "commentID": %d, "location": "%s" }`, userID, commentID, location)))
//	}
//
// func categories(w http.ResponseWriter, r *http.Request) {
//
//		w.Header().Set("Content-Type", "application/json")
//
//		//for len(aValues) > 0 {
//		//	var aString string
//		//	values, err = redis.Scan(values, &aString)
//		//	if err != nil {
//		//		fmt.Println(err)
//		//		return
//		//	}
//		//
//		//}
//
//		data, err := daoImpl.GetCategories()
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//
//		json, err := json.Marshal(&data)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		w.WriteHeader(http.StatusOK)
//		w.Header().Set("Content-Type", "application/json")
//		w.Write(json)
//	}
//
//	func metricsByCategory(w http.ResponseWriter, r *http.Request) {
//		pathParams := mux.Vars(r)
//		w.Header().Set("Content-Type", "application/json")
//
//		aParam, ok := pathParams["aCategory"]
//		if !ok || len(aParam) == 0 {
//			w.WriteHeader(http.StatusInternalServerError)
//			w.Write([]byte(`{"message": "Category is required"}`))
//			return
//		}
//		categories, _ := daoImpl.GetCategories()
//		metricData, err := daoImpl.GetMetricsByCategory(domain.FindCodeByID(categories, aParam))
//		json, err := json.Marshal(&metricData)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		w.WriteHeader(http.StatusOK)
//		w.Header().Set("Content-Type", "application/json")
//		w.Write(json)
//	}
//
// func schools(w http.ResponseWriter, r *http.Request) {
//
//		w.Header().Set("Content-Type", "application/json")
//
//		data, err := daoImpl.GetSchools()
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		json, err := json.Marshal(&data)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		w.WriteHeader(http.StatusOK)
//		w.Header().Set("Content-Type", "application/json")
//		w.Write(json)
//	}
//
//	func metricsBySchool(w http.ResponseWriter, r *http.Request) {
//		pathParams := mux.Vars(r)
//		w.Header().Set("Content-Type", "application/json")
//
//		//aHeaderValue := r.Header.Get("resolve-metric-detail")
//		//var resolveMetricDetail bool
//		//if len(aHeaderValue) > 0 && aHeaderValue == "1"{
//		//	resolveMetricDetail = true
//		//}
//		aParam, ok := pathParams["aSchool"]
//		if !ok || len(aParam) == 0 {
//			w.WriteHeader(http.StatusInternalServerError)
//			w.Write([]byte(`{"message": "School is required"}`))
//			return
//		}
//
//		codes, _ := daoImpl.GetSchools()
//		code := domain.FindCodeByID(codes, aParam)
//
//		metricData, err := daoImpl.GetMetricsBySchool(code)
//		json, err := json.Marshal(&metricData)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		w.WriteHeader(http.StatusOK)
//		w.Header().Set("Content-Type", "application/json")
//		w.Write(json)
//	}
//
//	func metricsBySchoolAndCategory(w http.ResponseWriter, r *http.Request) {
//		pathParams := mux.Vars(r)
//		w.Header().Set("Content-Type", "application/json")
//
//		aSchool, ok := pathParams["aSchool"]
//
//		if !ok || len(aSchool) == 0 {
//			w.WriteHeader(http.StatusInternalServerError)
//			w.Write([]byte(`{"message": "School is required", "message": "Category is required"}`))
//			return
//		}
//
//		aCategory, ok := pathParams["aCategory"]
//		if !ok || len(aCategory) == 0 {
//			w.WriteHeader(http.StatusInternalServerError)
//			w.Write([]byte(`{"message": "School is required", "message": "Category is required"}`))
//			return
//		}
//
//		schools, _ := daoImpl.GetSchools()
//		school := domain.FindCodeByDescription(schools, aSchool)
//
//		categories, _ := daoImpl.GetCategories()
//		category := domain.FindCodeByDescription(categories, aCategory)
//
//		metricData, err := daoImpl.GetMetricsBySchoolAndCategory(school, category)
//		json, err := json.Marshal(&metricData)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		w.WriteHeader(http.StatusOK)
//		w.Header().Set("Content-Type", "application/json")
//		w.Write(json)
//	}
//
//	func metricDetailsBySchoolAndCategory(w http.ResponseWriter, r *http.Request) {
//		pathParams := mux.Vars(r)
//		w.Header().Set("Content-Type", "application/json")
//
//		aSchool, ok := pathParams["aSchool"]
//
//		if !ok || len(aSchool) == 0 {
//			w.WriteHeader(http.StatusInternalServerError)
//			w.Write([]byte(`{"message": "School is required", "message": "Category is required"}`))
//			return
//		}
//
//		aCategory, ok := pathParams["aCategory"]
//		if !ok || len(aCategory) == 0 {
//			w.WriteHeader(http.StatusInternalServerError)
//			w.Write([]byte(`{"message": "School is required", "message": "Category is required"}`))
//			return
//		}
//
//		schools, _ := daoImpl.GetSchools()
//		school := domain.FindCodeByID(schools, aSchool)
//
//		categories, _ := daoImpl.GetCategories()
//		category := domain.FindCodeByID(categories, aCategory)
//
//		metricData, err := daoImpl.GetMetricsBySchoolAndCategory(school, category)
//
//		sort.Slice(metricData, func(i int, j int) bool {
//			return metricData[i].DateTime.Before(metricData[j].DateTime)
//		})
//
//		json, err := json.Marshal(&metricData)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//
//		w.WriteHeader(http.StatusOK)
//		w.Header().Set("Content-Type", "application/json")
//		w.Write(json)
//	}
//
// //
// //func dates(w http.ResponseWriter, r *http.Request) {
// //
// //	w.Header().Set("Content-Type", "application/json")
// //
// //	data, err := daoImpl.GetDates()
// //	if err != nil {
// //		http.Error(w, err.Error(), http.StatusInternalServerError)
// //		return
// //	}
// //
// //	json, err := json.Marshal(&data)
// //	if err != nil {
// //		http.Error(w, err.Error(), http.StatusInternalServerError)
// //		return
// //	}
// //
// //	w.WriteHeader(http.StatusOK)
// //	w.Header().Set("Content-Type", "application/json")
// //	w.Write(json)
// //}
// //
// //func metricsByDate(w http.ResponseWriter, r *http.Request) {
// //	pathParams := mux.Vars(r)
// //	w.Header().Set("Content-Type", "application/json")
// //
// //	aParam, ok := pathParams["aDate"]
// //	if !ok || len(aParam) == 0 {
// //		w.WriteHeader(http.StatusInternalServerError)
// //		w.Write([]byte(`{"message": "Date as yyyymmddhh24miss is required"}`))
// //		return
// //	}
// //	metricData := daoImpl.GetMetricsByDate(aParam)
// //	json, err := json.Marshal(&metricData)
// //	if err != nil {
// //		http.Error(w, err.Error(), http.StatusInternalServerError)
// //		return
// //	}
// //	w.WriteHeader(http.StatusOK)
// //	w.Header().Set("Content-Type", "application/json")
// //	w.Write(json)
// //}
//
//	func metric(w http.ResponseWriter, r *http.Request) {
//		pathParams := mux.Vars(r)
//		w.Header().Set("Content-Type", "application/json")
//
//		aParam, ok := pathParams["aMetric"]
//		if !ok || len(aParam) == 0 {
//			w.WriteHeader(http.StatusInternalServerError)
//			w.Write([]byte(`{"message": "Date as yyyymmddhh24miss is required"}`))
//			return
//		}
//		skey, err := strconv.Atoi(aParam)
//		metricData, err := daoImpl.GetMetric(skey)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//
//		json, err := json.Marshal(&metricData)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		w.WriteHeader(http.StatusOK)
//		w.Header().Set("Content-Type", "application/json")
//		w.Write(json)
//	}
func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	if (*req).Method == "OPTIONS" {
		return
	}

	// process the request...
}

func handler(w http.ResponseWriter, req *http.Request) {
	// ...
	enableCors(&w)
	// ...
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Set headers
		w.Header().Set("Access-Control-Allow-Headers:", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		fmt.Println("ok")

		// Next
		next.ServeHTTP(w, r)
		return
	})
}

func setHeaders(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//anyone can make a CORS request (not recommended in production)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//only allow GET, POST, and OPTIONS
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		//Since I was building a REST API that returned JSON, I set the content type to JSON here.
		w.Header().Set("Content-Type", "application/json")
		//Allow requests to have the following headers
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, cache-control")
		//if it's just an OPTIONS request, nothing other than the headers in the response is needed.
		//This is essential because you don't need to handle the OPTIONS requests in your handlers now
		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})

}

func decodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			return &malformedRequest{status: http.StatusUnsupportedMediaType, msg: msg}
		}
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := fmt.Sprintf("Request body contains badly-formed JSON")
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			return &malformedRequest{status: http.StatusRequestEntityTooLarge, msg: msg}

		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		msg := "Request body must only contain a single JSON object"
		return &malformedRequest{status: http.StatusBadRequest, msg: msg}
	}

	return nil
}

type malformedRequest struct {
	status int
	msg    string
}

func (mr *malformedRequest) Error() string {
	return mr.msg
}
