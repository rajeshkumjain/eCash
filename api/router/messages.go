package router

// ErrorMap : ..
var ErrorMap = map[string]string{
	"E001": "ERROR: (E001) Invalid Header Format : Expecting application/jon in the header ...  ",
	"E002": "ERROR: (E002) Unable to decode the content from request body ... ",
	"E003": "ERROR: (E003) Unable to decode the JSON packet ...  ",
	"E004": "ERROR: (E004) Unable to complete the registration process ... ",
	"E005": "ERROR: (E005) Invalid Activation Key or Corrupted Key ... ",
	"E006": "ERROR: (E006) Record Not found with the matching Activation Key ... ",

	"S001": "SUCCESS: (S001) Thank You - your registration completed successfully ...",
	"S002": "SUCCESS: (S002) Your Account is now Activated Successfully ...",
}

