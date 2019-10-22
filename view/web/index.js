function uploadXlsxFile() {
	var elements = 'iput-file-upload';
	var fileLists = document.getElementById(elements).files;
	console.log(document.getElementById(elements).files[0]);
	if (fileLists.length <= 0) {
		console.log('文件为空');
		return;
	}

	var formData = new FormData();

	formData.append("file",fileLists[0]);
	formData.append("id", elements);
	formData.append("type", fileLists[0]['type']);

	console.log(formData.contentType)

	$.ajax({
		url: '/api/v0/file/upload/xlsx',
		type: 'POST',
		processData: false,
		contentType: false,
		data: formData,
		success: function (s) {
			alert(s);
		}
	})
}