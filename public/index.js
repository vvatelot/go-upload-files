var fileList = []
const urlParams = new URLSearchParams(window.location.search)
const userId = urlParams.get('userid')
var tbody = fileListDisplay.getElementsByTagName('tbody')[0]

inputFiles.addEventListener('change', function () {
    fileList = [];
    for (var i = 0; i < inputFiles.files.length; i++) {
        fileList.push(inputFiles.files[i])
    }
    renderFileList()
});

formUpload.addEventListener("submit", async (e) => {
    e.preventDefault();
    progressBar.classList.remove('is-hidden')

    var formData = new FormData(formUpload)
    var xhr=new XMLHttpRequest()
    xhr.open("POST",'/upload?userid=' + userId)
    xhr.upload.addEventListener('progress', function(e){
        progressBar.value = Math.ceil(e.loaded/e.total)*100
    }, false);
    xhr.onreadystatechange = function() {
        if (this.readyState == 4) {
            var jsonResponse = JSON.parse(this.responseText);
            responseMessage.innerHTML = jsonResponse.message
            modal.classList.add("is-active")
            progressBar.classList.add('is-hidden')
        }
    };

    xhr.send(formData);

    cleanRenderedList()
})


renderFileList = function () {
    cleanRenderedList()
    
    fileList.forEach(function (file, index) {
        var row = tbody.insertRow()
        var cellFile = row.insertCell()
        var cellSize = row.insertCell()
        var cellType = row.insertCell()

        cellFile.innerHTML = file.name
        cellSize.innerHTML = humanFileSize(file.size)
        cellType.innerHTML = file.type
    })

    if (fileList.length > 0) {
        showRenderedList()
    }
}

cleanRenderedList = function () {
    tbody.innerHTML = ""
    submitButton.classList.add('is-hidden')
    fileListDisplay.classList.add('is-hidden')
    progressBar.classList.add('is-hidden')
}

showRenderedList = function() {
    submitButton.classList.remove('is-hidden')
    fileListDisplay.classList.remove('is-hidden')
}

humanFileSize = function (size) {
    var i = Math.floor( Math.log(size) / Math.log(1024) );
    return ( size / Math.pow(1024, i) ).toFixed(2) * 1 + ' ' + ['B', 'kB', 'MB', 'GB', 'TB'][i];
};

closeModal = function() {
    modal.classList.remove("is-active")
}