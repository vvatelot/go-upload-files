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

formUpload.addEventListener("submit", function(e) {
    e.preventDefault();

    var formData = new FormData(formUpload)
    var client = new XMLHttpRequest()

    client.upload.onloadstart = function(e) {
        progressBar.value = e.loaded
        progressBar.max = e.total
        progressBar.classList.remove('is-hidden')
    }
    
    client.upload.onprogress = function (e) {
        progressBar.value = e.loaded
    }
    
    client.onreadystatechange = function() {
        if (this.readyState == 4) {
            var jsonResponse = JSON.parse(this.responseText);
            responseMessage.innerHTML = jsonResponse.message
            modal.classList.add("is-active")
            cleanRenderedList()
            
        }
    };
    
    client.open("POST",'/upload?userid=' + userId)
    client.send(formData)
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