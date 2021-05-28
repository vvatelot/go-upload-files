var fileList = [];

inputFiles.addEventListener('change', function () {
    fileList = [];
    for (var i = 0; i < inputFiles.files.length; i++) {
        fileList.push(inputFiles.files[i]);
    }
    renderFileList();
});

formUpload.addEventListener("submit", async (e) => {
    e.preventDefault();

    let response = await fetch('/upload?userid=vvatel22', {
        method: 'POST',
        body: new FormData(formUpload)
    });

    let result = await response.json();

    alert(result.message);
    })


renderFileList = function () {
    fileListDisplay.innerHTML = '';
    submitButton.style.display = "none";

    fileList.forEach(function (file, index) {
        var fileDisplayEl = document.createElement('li');
        fileDisplayEl.innerHTML = (index + 1) + ': ' + file.name;
        fileListDisplay.appendChild(fileDisplayEl);
    });

    if (fileList.length > 0) {
        submitButton.style.display = "block";
    }
};

function handleSubmit(acceptedFiles) {
    const data = new FormData();

    for (const file of acceptedFiles) {
        data.append('files[]', file, file.name);
    }

    console.log(data.getAll("files"))

    return fetch('/upload?userid=vvatel22', {
        method: 'POST',
        body: data,      
    });
}