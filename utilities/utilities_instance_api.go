package utilities

import (
	"encoding/json"
	"log"

	"github.com/jinzhu/gorm"
)

// Response map in this struts after decode
type Response struct {
	Number string
}

// DB object
var db *gorm.DB

// Event is default struct
type Event struct {
	ID          int    `json:"id"`
	EventName   string `json:"event_name"`
	EventSource string `json:"event_source"`
	EventDetail string `json:"event_detail"`
	TicketNo    string `json:"ticket_no"`
	ServiceNow  string `json:"service_now"`
}

// CallInstanceAPI function call third party api
func CallInstanceAPI(utilitieEventString []byte) {
	log.Println("Call instance API")
	var returnResponse string

	event := Event{}

	json.Unmarshal(utilitieEventString, &event)

	log.Println("Environment is : " + Environment)

	if Environment == "production" {

		apiRequest := APIRequest{
			APIURL:  InstanceAPIURL,
			Method:  "POST",
			Headers: "Content-Type:application/json",
			Data:    utilitieEventString,
		}

		returnResponse = CallAPI(apiRequest)
	} else {
		returnResponse = TestResponse()
	}

	response := GetTicketID(returnResponse)

	log.Println("Update ticket event : ")
	UpdateEventTicket(response.Number, returnResponse, event.ID)

}

// TestResponse Return a test response
func TestResponse() string {
	return `{"upon_approval":"proceed","location":"","expected_start":"","reopen_count":"0","close_notes":"","additional_assignee_list":"","impact":"2","urgency":"2","correlation_id":"","sys_tags":"","sys_domain":{"link":"https://instance.service-now.com/api/now/table/sys_user_group/global","value":"global"},"description":"","group_list":"","priority":"3","delivery_plan":"","sys_mod_count":"0","work_notes_list":"","business_service":"","follow_up":"","closed_at":"","sla_due":"","delivery_task":"","sys_updated_on":"2016-01-22 14:28:24","parent":"","work_end":"","number":"INC0010002","closed_by":"","work_start":"","calendar_stc":"","category":"inquiry","business_duration":"","incident_state":"1","activity_due":"","correlation_display":"","company":"","active":"true","due_date":"","assignment_group":{"link":"https://instance.service-now.com/api/now/table/sys_user_group/287ebd7da9fe198100f92cc8d1d2154e","value":"287ebd7da9fe198100f92cc8d1d2154e"},"caller_id":"","knowledge":"false","made_sla":"true","comments_and_work_notes":"","parent_incident":"","state":"1","user_input":"","sys_created_on":"2016-01-22 14:28:24","approval_set":"","reassignment_count":"0","rfc":"","child_incidents":"0","opened_at":"2016-01-22 14:28:24","short_description":"Unable to connect to office wifi","order":"","sys_updated_by":"admin","resolved_by":"","notify":"1","upon_reject":"cancel","approval_history":"","problem_id":"","work_notes":"","calendar_duration":"","close_code":"","sys_id":"c537bae64f411200adf9f8e18110c76e","approval":"not requested","caused_by":"","severity":"3","sys_created_by":"admin","resolved_at":"","assigned_to":"","business_stc":"","wf_activity":"","sys_domain_path":"/","cmdb_ci":"","opened_by":{"link":"https://instance.service-now.com/api/now/table/sys_user/6816f79cc0a8016401c5a33be04be441","value":"6816f79cc0a8016401c5a33be04be441"},"subcategory":"","rejection_goto":"","sys_class_name":"incident","watch_list":"","time_worked":"","contact_type":"phone","escalation":"0","comments":""}`
}

// GetTicketID is extact id from response json
func GetTicketID(returnResponse string) Response {
	response := Response{}
	responseBytes := []byte(returnResponse)
	json.Unmarshal(responseBytes, &response)

	return response
}

// UpdateEventTicket is update event with ticket and service now
func UpdateEventTicket(ticketNo, serviceNow string, ID int) {
	// connect to database
	db = DbConnect(db)

	//close database connection at end of database operation
	defer db.Close()

	event := Event{
		ID: ID,
	}

	db.Model(&event).UpdateColumns(Event{TicketNo: ticketNo, ServiceNow: serviceNow})
}
