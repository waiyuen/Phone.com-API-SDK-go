/* 
 * Phone.com API
 *
 * This is a Phone.com api Swagger definition
 *
 * OpenAPI spec version: 1.0.0
 * Contact: apisupport@phone.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type CreateContactParams struct {

	// First Name
	FirstName string `json:"first_name,omitempty"`

	// Middle Name
	MiddleName string `json:"middle_name,omitempty"`

	// Last Name
	LastName string `json:"last_name,omitempty"`

	// Prefix
	Prefix string `json:"prefix,omitempty"`

	// Phonetic First Name
	PhoneticFirstName string `json:"phonetic_first_name,omitempty"`

	// Phonetic Middle Name
	PhoneticMiddleName string `json:"phonetic_middle_name,omitempty"`

	// Phonetic Last Name
	PhoneticLastName string `json:"phonetic_last_name,omitempty"`

	// Suffix
	Suffix string `json:"suffix,omitempty"`

	// Nickname
	Nickname string `json:"nickname,omitempty"`

	// Company Name
	Company string `json:"company,omitempty"`

	// Department Name
	Department string `json:"department,omitempty"`

	// Job Title
	JobTitle string `json:"job_title,omitempty"`

	// Email Addresses
	Emails []Email `json:"emails,omitempty"`

	// Phone Numbers
	PhoneNumbers []PhoneNumberContact `json:"phone_numbers,omitempty"`

	// Addresses
	Addresses []AddressListContacts `json:"addresses,omitempty"`

	// Contact Group
	Group interface{} `json:"group,omitempty"`
}
