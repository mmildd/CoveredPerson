// Code generated by entc, DO NOT EDIT.

package patient

const (
	// Label holds the string label denoting the patient type in the database.
	Label = "patient"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPatientName holds the string denoting the patient_name field in the database.
	FieldPatientName = "patient_name"
	// FieldPatientAge holds the string denoting the patient_age field in the database.
	FieldPatientAge = "patient_age"
	// FieldPatientWeight holds the string denoting the patient_weight field in the database.
	FieldPatientWeight = "patient_weight"
	// FieldPatientHeight holds the string denoting the patient_height field in the database.
	FieldPatientHeight = "patient_height"
	// FieldPatientPrefix holds the string denoting the patient_prefix field in the database.
	FieldPatientPrefix = "patient_prefix"
	// FieldPatientGender holds the string denoting the patient_gender field in the database.
	FieldPatientGender = "patient_gender"
	// FieldPatientBlood holds the string denoting the patient_blood field in the database.
	FieldPatientBlood = "patient_blood"

	// EdgePatientCoveredPerson holds the string denoting the patient_coveredperson edge name in mutations.
	EdgePatientCoveredPerson = "Patient_CoveredPerson"

	// Table holds the table name of the patient in the database.
	Table = "patients"
	// PatientCoveredPersonTable is the table the holds the Patient_CoveredPerson relation/edge.
	PatientCoveredPersonTable = "covered_persons"
	// PatientCoveredPersonInverseTable is the table name for the CoveredPerson entity.
	// It exists in this package in order to avoid circular dependency with the "coveredperson" package.
	PatientCoveredPersonInverseTable = "covered_persons"
	// PatientCoveredPersonColumn is the table column denoting the Patient_CoveredPerson relation/edge.
	PatientCoveredPersonColumn = "Patient_id"
)

// Columns holds all SQL columns for patient fields.
var Columns = []string{
	FieldID,
	FieldPatientName,
	FieldPatientAge,
	FieldPatientWeight,
	FieldPatientHeight,
	FieldPatientPrefix,
	FieldPatientGender,
	FieldPatientBlood,
}

var (
	// PatientNameValidator is a validator for the "Patient_Name" field. It is called by the builders before save.
	PatientNameValidator func(string) error
	// PatientAgeValidator is a validator for the "Patient_Age" field. It is called by the builders before save.
	PatientAgeValidator func(string) error
	// PatientWeightValidator is a validator for the "Patient_Weight" field. It is called by the builders before save.
	PatientWeightValidator func(string) error
	// PatientHeightValidator is a validator for the "Patient_Height" field. It is called by the builders before save.
	PatientHeightValidator func(string) error
	// PatientPrefixValidator is a validator for the "Patient_Prefix" field. It is called by the builders before save.
	PatientPrefixValidator func(string) error
	// PatientGenderValidator is a validator for the "Patient_Gender" field. It is called by the builders before save.
	PatientGenderValidator func(string) error
	// PatientBloodValidator is a validator for the "Patient_Blood" field. It is called by the builders before save.
	PatientBloodValidator func(string) error
)
