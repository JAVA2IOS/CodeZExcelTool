function uploadXlsxFile() {

	var elements = 'input-xlsx-file-upload';
	var fileLists = document.getElementById(elements).files;
	if (fileLists.length <= 0 || typeof(fileLists) == 'undefined') {
		alert('excel文件为空')
		return;
	}

	var txtElementId = 'input-txt-file-upload';
	var txtFileObjs = document.getElementById(txtElementId).files;

	// if txtFileObjs.length <= 0 || typeof(txtFileObjs) == 'undefined' {
	// 	alert('txt文件为空')
	// 	return;
	// }

	var formData = new FormData();

	formData.append("xlsxFile",fileLists[0]);
	formData.append("txtFile",txtFileObjs[0]);

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