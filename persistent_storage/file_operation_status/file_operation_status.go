package file_operation_status

const (
	fileOperationStatusHeader = "file_operation_status: "
	DoesNotExist              = fileOperationStatusHeader + "File does not exist"
	NameIsTheSame             = fileOperationStatusHeader + "Name is the same"
	AlreadyExists             = fileOperationStatusHeader + "File already exists"
)
